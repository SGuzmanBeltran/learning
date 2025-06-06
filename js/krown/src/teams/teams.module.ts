import { DrizzleModule } from '../drizzle/drizzle.module';
import { Module } from '@nestjs/common';
import { TeamsController } from './teams.controller';
import { TeamsService } from './teams.service';

@Module({
  imports: [DrizzleModule],
  controllers: [TeamsController],
  providers: [TeamsService],
})
export class TeamsModule {}
