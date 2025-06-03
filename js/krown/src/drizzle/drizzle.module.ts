import * as schema from './schemas/schema';

import { ConfigModule, ConfigService } from '@nestjs/config';
import { Global, Module } from '@nestjs/common';
import { LibSQLDatabase, drizzle } from 'drizzle-orm/libsql';

import { createClient } from '@libsql/client';

export const DRIZZLE = Symbol('DRIZZLE');

@Global()
@Module({
  imports: [ConfigModule],
  providers: [
    {
      provide: DRIZZLE,
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => {
        const dbName = configService.get<string>('DB_FILE_NAME');
        const client = createClient({
          url: `file:${dbName}`,
        });
        return drizzle(client, { schema }) as LibSQLDatabase<typeof schema>;
      },
    },
  ],
  exports: [DRIZZLE],
})
export class DrizzleModule {}
