import * as DrizzleSpec from '../common/test/drizzle.mock';

import {
  ConflictException,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';

import { AuthModule } from '../auth/auth.module';
import { ConfigModule } from '@nestjs/config';
import { DRIZZLE } from '../drizzle/drizzle.module';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { TeamsService } from './teams.service';

const setCreateMockTransaction = (
  mockDrizzle: Partial<DrizzleDB>,
  insertReturnValues: any[],
) => {
  (mockDrizzle.transaction as jest.Mock).mockImplementationOnce(
    async <T>(callback: (tx: Partial<DrizzleDB>) => Promise<T>): Promise<T> => {
      const tx: Partial<DrizzleDB> = {
        insert: jest
          .fn()
          .mockReturnValueOnce({
            values: jest.fn().mockReturnValue({
              returning: jest.fn().mockResolvedValue([insertReturnValues[0]]),
            }),
          })
          .mockReturnValueOnce({
            values: jest.fn().mockReturnValue({
              returning: jest.fn().mockResolvedValue([insertReturnValues[1]]),
            }),
          }),
      };
      return await callback(tx);
    },
  );
};

describe('TeamsService', () => {
  let service: TeamsService;
  let mockDrizzle: Partial<DrizzleDB>;

  beforeEach(async () => {
    mockDrizzle = DrizzleSpec.setupMockDrizzle();

    const module: TestingModule = await Test.createTestingModule({
      imports: [
        ConfigModule.forRoot({
          isGlobal: true,
        }),
        AuthModule,
      ],
      providers: [
        TeamsService,
        {
          provide: DRIZZLE,
          useValue: mockDrizzle,
        },
      ],
    }).compile();

    service = module.get<TeamsService>(TeamsService);
  });

  describe('create team', () => {
    it('should be defined', () => {
      expect(service).toBeDefined();
    });

    it('should create a team', async () => {
      DrizzleSpec.setMockSelect(mockDrizzle, [null]);
      setCreateMockTransaction(mockDrizzle, [
        { id: 1, name: 'Test Team' },
        { teamId: 1, userId: 1, isLeader: true },
      ]);
      DrizzleSpec.setMockSelect(mockDrizzle, [{ id: 1, name: 'Test Team' }]);

      const team = await service.create({ name: 'Test Team' }, 1);
      expect(team).toEqual({ id: 1, name: 'Test Team' });
      expect(mockDrizzle.transaction).toHaveBeenCalled();
      expect(mockDrizzle.select).toHaveBeenCalledTimes(2);
    });

    it('should throw an error if the team name is already taken', async () => {
      DrizzleSpec.setMockSelect(mockDrizzle, [{ id: 1, name: 'Test Team' }]);

      await expect(service.create({ name: 'Test Team' }, 1)).rejects.toThrow(
        ConflictException,
      );
    });

    it.each([
      {
        scenario: 'team creation fails',
        insertReturnValues: [null],
      },
      {
        scenario: 'team member creation fails',
        insertReturnValues: [null, null],
      },
    ])(
      'should throw an error if the team creation fails with $scenario',
      async ({ insertReturnValues }) => {
        DrizzleSpec.setMockSelect(mockDrizzle, [null]);
        setCreateMockTransaction(mockDrizzle, insertReturnValues);

        await expect(service.create({ name: 'Test Team' }, 1)).rejects.toThrow(
          InternalServerErrorException,
        );

        expect(mockDrizzle.transaction).toHaveBeenCalled();
      },
    );
  });

  describe('findAll teams', () => {
    it('should find all teams', async () => {
      DrizzleSpec.setMockSelect(mockDrizzle, [
        {
          teams: { id: 1, name: 'Test Team' },
          teamMembers: { teamId: 1, userId: 1, isLeader: true },
        },
      ]);

      const teams = await service.findAll(1);
      expect(teams).toEqual([{ id: 1, name: 'Test Team' }]);
    });
  });

  describe('findOne team', () => {
    it('should find a team', async () => {
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [
        {
          teams: { id: 1, name: 'Test Team' },
          team_members: [{ teamId: 1, userId: 1, isLeader: true }],
        },
      ]);

      const team = await service.findOne(1);
      expect(team).toEqual({
        team: { id: 1, name: 'Test Team' },
        members: [{ teamId: 1, userId: 1, isLeader: true }],
      });
    });

    it('should throw an error if the team is not found', async () => {
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [null]);

      await expect(service.findOne(1)).rejects.toThrow(NotFoundException);
    });
  });

  describe('update team', () => {
    it('should update a team', async () => {
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [
        {
          teams: { id: 1, name: 'Old Team' },
          team_members: [{ teamId: 1, userId: 1, isLeader: true }],
        },
      ]);
      DrizzleSpec.setMockUpdate(mockDrizzle, [{ id: 1, name: 'Updated Team' }]);

      const team = await service.update(1, { name: 'Updated Team' });
      expect(team).toEqual({ id: 1, name: 'Updated Team' });
      expect(mockDrizzle.update).toHaveBeenCalled();
      expect(mockDrizzle.select).toHaveBeenCalled();
    });

    it('should throw an error if the team update fails', async () => {
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [
        {
          teams: { id: 1, name: 'Old Team' },
          team_members: [{ teamId: 1, userId: 1, isLeader: true }],
        },
      ]);

      DrizzleSpec.setMockUpdateFails(mockDrizzle);

      await expect(service.update(1, { name: 'Updated Team' })).rejects.toThrow(
        InternalServerErrorException,
      );
    });
  });

  describe('remove team', () => {
    it('should remove a team', async () => {
      const teamLeader = { teamId: 1, userId: 1, isLeader: true };
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [
        {
          teams: { id: 1, name: 'Test Team' },
          team_members: [teamLeader],
        },
      ]);
      DrizzleSpec.setMockDelete(mockDrizzle);

      const result = await service.remove(1, teamLeader.userId);

      expect(result).toEqual({
        message: 'Team 1 deleted successfully',
        deleted: true,
      });
      expect(mockDrizzle.delete).toHaveBeenCalled();
      expect(mockDrizzle.select).toHaveBeenCalled();
    });

    it('should throw an error if the team deletion fails', async () => {
      const teamLeader = { teamId: 1, userId: 1, isLeader: true };
      DrizzleSpec.setMockSelectWithInnerJoinLimit(mockDrizzle, [
        {
          teams: { id: 1, name: 'Test Team' },
          team_members: [teamLeader],
        },
      ]);
      DrizzleSpec.setMockDeleteFails(mockDrizzle);

      await expect(service.remove(1, teamLeader.userId)).rejects.toThrow(
        InternalServerErrorException,
      );

      expect(mockDrizzle.delete).toHaveBeenCalled();
      expect(mockDrizzle.select).toHaveBeenCalled();
    });
  });
});
