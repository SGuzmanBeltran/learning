import * as DrizzleSpec from '../common/test/drizzle.mock';

import { Test, TestingModule } from '@nestjs/testing';

import { AuthModule } from '../auth/auth.module';
import { ConfigModule } from '@nestjs/config';
import { ConflictException } from '@nestjs/common';
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

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it('should create a team', async () => {
    DrizzleSpec.setMockSelectLimit(mockDrizzle, [null]);
    setCreateMockTransaction(mockDrizzle, [
      { id: 1, name: 'Test Team' },
      { teamId: 1, userId: 1, isLeader: true },
    ]);

    const team = await service.create({ name: 'Test Team' }, 1);
    expect(team).toEqual({ id: 1, name: 'Test Team' });
    expect(mockDrizzle.transaction).toHaveBeenCalled();
  });

  it('should throw an error if the team name is already taken', async () => {
    DrizzleSpec.setMockSelectLimit(mockDrizzle, [{ id: 1, name: 'Test Team' }]);

    await expect(service.create({ name: 'Test Team' }, 1)).rejects.toThrow(
      ConflictException,
    );
  });
});
