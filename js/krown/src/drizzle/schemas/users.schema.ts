import { integer, sqliteTable, text } from 'drizzle-orm/sqlite-core';

import { sql } from 'drizzle-orm';

export const users = sqliteTable('users', {
  id: integer('id').primaryKey(),
  email: text('email').unique().notNull(),
  username: text('username').unique().notNull(),
  cellphone: text('cellphone').unique().notNull(),
  password: text('password').notNull(),
  createdAt: integer('created_at').default(sql`CURRENT_TIMESTAMP`),
  updatedAt: integer('updated_at').default(sql`CURRENT_TIMESTAMP`),
});
