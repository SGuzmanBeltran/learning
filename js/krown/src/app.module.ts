import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AuthModule } from './auth/auth.module';
import { DrizzleModule } from './drizzle/drizzle.module';
import { TeamsModule } from './teams/teams.module';

@Module({
  imports: [AuthModule, DrizzleModule, TeamsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
