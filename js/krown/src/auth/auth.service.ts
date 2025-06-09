import {
  BadRequestException,
  Inject,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { LoginDto } from './dto/login.dto';
import { RegisterDto } from './dto/register.dto';
import { DRIZZLE } from '../drizzle/drizzle.module';
import { users } from '../drizzle/schemas/users.schema';
import { eq, or } from 'drizzle-orm';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcryptjs';

export interface LoginResponse {
  message: string;
  accessToken: string;
}

@Injectable()
export class AuthService {
  constructor(
    @Inject(DRIZZLE) private drizzle: DrizzleDB,
    private readonly jwtService: JwtService,
  ) {}

  //todo: catch unique constraint error
  async register(registerDto: RegisterDto): Promise<{ message: string }> {
    const { email, password, passwordConfirmation, username, cellphone } =
      registerDto;

    if (password !== passwordConfirmation) {
      throw new BadRequestException('Passwords do not match');
    }

    const [user] = await this.drizzle
      .select()
      .from(users)
      .where(or(eq(users.email, email), eq(users.username, username)))
      .limit(1);

    if (user) {
      throw new BadRequestException('User already exists');
    }

    const hashedPassword = await bcrypt.hash(password, 10);
    await this.drizzle.insert(users).values({
      email: email,
      password: hashedPassword,
      username: username,
      cellphone: cellphone,
    });

    return {
      message: 'User registered successfully',
    };
  }

  async login(loginDto: LoginDto): Promise<LoginResponse> {
    const { email, password } = loginDto;
    const [user] = await this.drizzle
      .select()
      .from(users)
      .where(eq(users.email, email))
      .limit(1);

    if (!user) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const isPasswordValid = await bcrypt.compare(password, user.password);

    if (!isPasswordValid) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const accessToken = await this.jwtService.signAsync({
      userId: user.id,
      email: user.email,
    });

    return {
      message: 'Login successful',
      accessToken,
    };
  }
}
