import { BadRequestException, UnauthorizedException } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';

import { AuthService } from './auth.service';
import { DRIZZLE } from '../drizzle/drizzle.module';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { JwtService } from '@nestjs/jwt';

jest.mock('bcryptjs', () => ({
  hash: jest.fn().mockResolvedValue('hashedPassword'),
  compare: jest.fn().mockResolvedValue(true),
}));

describe('AuthService', () => {
  let service: AuthService;
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
      providers: [
        AuthService,
        {
          provide: JwtService,
          useValue: {
            signAsync: jest.fn().mockResolvedValue('mock.jwt.token'),
          },
        },
        {
          provide: DRIZZLE,
          useValue: mockDrizzle,
        },
      ],
    }).compile();

    service = module.get<AuthService>(AuthService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('register', () => {
    it('should throw BadRequestException when passwords do not match', async () => {
      const registerDto = {
        email: 'test@example.com',
        password: 'password',
        passwordConfirmation: 'different',
        username: 'testuser',
        cellphone: '1234567890',
      };

      await expect(service.register(registerDto)).rejects.toThrow(
        BadRequestException,
      );
    });

    it('should register a new user successfully', async () => {
      (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
        from: jest.fn().mockReturnValue({
          where: jest.fn().mockReturnValue({
            limit: jest.fn().mockResolvedValue(null),
          }),
        }),
      });

      const registerDto = {
        email: 'test@example.com',
        password: 'password',
        passwordConfirmation: 'password',
        username: 'testuser',
        cellphone: '1234567890',
      };

      const result = await service.register(registerDto);
      expect(result).toEqual({ message: 'User registered successfully' });
    });
  });

  describe('login', () => {
    it('should throw UnauthorizedException for invalid credentials', async () => {
      (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
        from: jest.fn().mockReturnValue({
          where: jest.fn().mockReturnValue({
            limit: jest.fn().mockResolvedValue([]),
          }),
        }),
      });

      const loginDto = {
        email: 'test@example.com',
        password: 'password',
      };

      await expect(service.login(loginDto)).rejects.toThrow(
        UnauthorizedException,
      );
    });

    it('should return access token for valid credentials', async () => {
      (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
        from: jest.fn().mockReturnValue({
          where: jest.fn().mockReturnValue({
            limit: jest.fn().mockResolvedValue([
              {
                id: 1,
                email: 'test@example.com',
                password: 'hashedPassword',
              },
            ]),
          }),
        }),
      });

      const loginDto = {
        email: 'test@example.com',
        password: 'password',
      };

      const result = await service.login(loginDto);
      expect(result).toEqual({
        message: 'Login successful',
        accessToken: 'mock.jwt.token',
      });
    });
  });
});
