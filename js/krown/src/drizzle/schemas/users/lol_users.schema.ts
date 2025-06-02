import { integer, sqliteTable, text } from 'drizzle-orm/sqlite-core';

import { sql } from 'drizzle-orm';
import { users } from '../users.schema';

export const lolUsers = sqliteTable('lol_users', {
  id: integer('id').primaryKey(),
  userId: integer('user_id').references(() => users.id),
  summonerName: text('summoner_name').notNull(),
  summonerId: text('summoner_id').notNull(),
  rank: text('rank').notNull(),
  region: text('region').notNull(),
  createdAt: integer('created_at').default(sql`CURRENT_TIMESTAMP`),
  updatedAt: integer('updated_at').default(sql`CURRENT_TIMESTAMP`),
});
