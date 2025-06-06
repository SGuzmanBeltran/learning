import { Test, TestingModule } from '@nestjs/testing';

import { DRIZZLE } from '../drizzle/drizzle.module';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { TeamsController } from './teams.controller';
import { TeamsService } from './teams.service';

describe('TeamsController', () => {
  let controller: TeamsController;
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
        values: jest.fn().mockResolvedValue(undefined),
      }),
    };
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TeamsController],
      providers: [
        TeamsService,
        {
          provide: DRIZZLE,
          useValue: mockDrizzle,
        },
      ],
    }).compile();

    controller = module.get<TeamsController>(TeamsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
