import { Test, TestingModule } from '@nestjs/testing';

import { ConflictException } from '@nestjs/common';
import { DRIZZLE } from '../drizzle/drizzle.module';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { TeamsService } from './teams.service';

describe('TeamsService', () => {
  let service: TeamsService;
  let mockDrizzle: Partial<DrizzleDB>;

  beforeEach(async () => {
    mockDrizzle = {
      select: jest.fn().mockReturnValue({
        from: jest.fn().mockReturnValue({
          where: jest.fn().mockReturnValue({
            limit: jest.fn().mockResolvedValue(null),
          }),
        }),
      }),
      insert: jest.fn().mockReturnValue({
        values: jest.fn().mockResolvedValue({
          returning: jest.fn().mockResolvedValue([{ id: 1 }]),
        }),
      }),
    };
    const module: TestingModule = await Test.createTestingModule({
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
    (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
      from: jest.fn().mockReturnValue({
        where: jest.fn().mockReturnValue({
          limit: jest.fn().mockResolvedValue([null]),
        }),
      }),
    });
    (mockDrizzle.insert as jest.Mock)
      .mockReturnValueOnce({
        values: jest.fn().mockReturnValue({
          returning: jest.fn().mockResolvedValue([{ id: 1 }]),
        }),
      })
      .mockReturnValueOnce({
        values: jest.fn().mockReturnValue({
          returning: jest
            .fn()
            .mockResolvedValue([{ teamId: 1, userId: 1, isLeader: true }]),
        }),
      });
    const team = await service.create({ name: 'Test Team' }, 1);
    expect(team).toEqual({ id: 1 });
  });

  it('should throw an error if the team name is already taken', async () => {
    (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
      from: jest.fn().mockReturnValue({
        where: jest.fn().mockReturnValue({
          limit: jest.fn().mockResolvedValue([{ id: 1, name: 'Test Team' }]),
        }),
      }),
    });
    (mockDrizzle.insert as jest.Mock).mockReturnValueOnce({
      values: jest.fn().mockReturnValue({
        returning: jest.fn().mockResolvedValue([{ id: 1 }]),
      }),
    });
    await expect(service.create({ name: 'Test Team' }, 1)).rejects.toThrow(
      ConflictException,
    );
  });
});
