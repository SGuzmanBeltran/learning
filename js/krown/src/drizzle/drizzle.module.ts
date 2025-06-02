import * as schema from './schemas/schema';

import { ConfigService } from '@nestjs/config';
import { Module } from '@nestjs/common';
import { createClient } from '@libsql/client';
import { drizzle } from 'drizzle-orm/libsql';

export const DRIZZLE = Symbol('DRIZZLE');
@Module({
  providers: [
    {
      provide: DRIZZLE,
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => {
        const dbName = configService.get<string>('DB_FILE_NAME');
        const client = createClient({
          url: `file:${dbName}`,
        });
        return drizzle(client, { schema });
      },
    },
  ],
  exports: [DRIZZLE],
})
export class DrizzleModule {}
