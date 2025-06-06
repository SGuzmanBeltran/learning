import { integer, sqliteTable, text } from 'drizzle-orm/sqlite-core';
import { relations, sql } from 'drizzle-orm';

import { teamMembers } from './team_members.schema';

export type Team = typeof teams.$inferSelect;

export const teams = sqliteTable('teams', {
  id: integer('id').primaryKey(),
  name: text('name').notNull(),
  playersCount: integer('players_count').notNull().default(0),
  tournamentsPlayed: integer('tournaments_played').notNull().default(0),
  tournamentsWon: integer('tournaments_won').notNull().default(0),
  createdAt: integer('created_at').default(sql`CURRENT_TIMESTAMP`),
  updatedAt: integer('updated_at').default(sql`CURRENT_TIMESTAMP`),
});

export const teamsRelations = relations(teams, ({ many }) => ({
  teamMembers: many(teamMembers),
}));
