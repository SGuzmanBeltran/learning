import * as bcrypt from 'bcryptjs';

import { BadRequestException, UnauthorizedException } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';

import { AuthService } from './auth.service';
import { DRIZZLE } from '../drizzle/drizzle.module';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { JwtService } from '@nestjs/jwt';
import { users } from '@src/drizzle/schemas/users.schema';

jest.mock('bcryptjs', () => ({
  hash: jest.fn().mockImplementation((password) => {
    return Promise.resolve(`hashed${password}`);
  }),
  compare: jest.fn().mockImplementation((password, hashedPassword) => {
    return Promise.resolve(`hashed${password}` === hashedPassword);
  }),
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

    it.each([
      {
        scenario: 'existing email',
        registerTestDto: {
          email: 'test@example.com',
          username: 'testuser',
        },
        existingUser: {
          id: 1,
          email: 'test@example.com',
          username: 'different',
        },
      },
      {
        scenario: 'existing username',
        registerTestDto: {
          email: 'test@example.com',
          username: 'testuser',
        },
        existingUser: {
          id: 2,
          email: 'different@example.com',
          username: 'testuser',
        },
      },
      {
        scenario: 'existing email and username',
        registerTestDto: {
          email: 'test@example.com',
          username: 'testuser',
        },
        existingUser: {
          id: 3,
          email: 'test@example.com',
          username: 'testuser',
        },
      },
    ])(
      'should throw BadRequestException when $scenario',
      async ({ registerTestDto, existingUser }) => {
        const registerDto = {
          email: registerTestDto.email,
          password: 'password',
          passwordConfirmation: 'password',
          username: registerTestDto.username,
          cellphone: '1234567890',
        };

        (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
          from: jest.fn().mockReturnValue({
            where: jest.fn().mockReturnValue({
              limit: jest.fn().mockResolvedValue([existingUser]),
            }),
          }),
        });

        await expect(service.register(registerDto)).rejects.toThrow(
          BadRequestException,
        );
      },
    );

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

    it('should hash password', async () => {
      const registerDto = {
        email: 'test@example.com',
        password: 'password',
        passwordConfirmation: 'password',
        username: 'testuser',
        cellphone: '1234567890',
      };

      const result = await service.register(registerDto);

      expect(bcrypt.hash).toHaveBeenCalledWith('password', 10);
      expect(result).toEqual({
        message: 'User registered successfully',
      });
    });
  });

  describe('login', () => {
    it.each([
      {
        scenario: 'invalid email',
        loginTestDto: {
          email: 'invalid@example.com',
          password: 'password',
        },
        mockUser: [],
      },
      {
        scenario: 'invalid password',
        loginTestDto: {
          email: 'test@example.com',
          password: 'invalid',
        },
        mockUser: [
          {
            id: 1,
            email: 'test@example.com',
            password: 'hashedPassword',
          },
        ],
      },
    ])(
      'should throw UnauthorizedException for $scenario',
      async ({ loginTestDto, mockUser }) => {
        (mockDrizzle.select as jest.Mock).mockReturnValueOnce({
          from: jest.fn().mockReturnValue({
            where: jest.fn().mockReturnValue({
              limit: jest.fn().mockResolvedValue(mockUser),
            }),
          }),
        });

        await expect(service.login(loginTestDto)).rejects.toThrow(
          UnauthorizedException,
        );
      },
    );

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
        password: 'Password',
      };

      const result = await service.login(loginDto);
      expect(result).toEqual({
        message: 'Login successful',
        accessToken: 'mock.jwt.token',
      });
    });
  });
});
