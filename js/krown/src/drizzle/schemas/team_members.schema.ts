import { integer, primaryKey, sqliteTable } from 'drizzle-orm/sqlite-core';
import { relations, sql } from 'drizzle-orm';

import { teams } from './teams.schema';
import { users } from './users.schema';

export type TeamMember = typeof teamMembers.$inferSelect;

export const teamMembers = sqliteTable(
  'team_members',
  {
    teamId: integer('team_id')
      .notNull()
      .references(() => teams.id),
    userId: integer('user_id')
      .notNull()
      .references(() => users.id),
    isLeader: integer('is_leader', { mode: 'boolean' })
      .notNull()
      .default(false),
    createdAt: integer('created_at').default(sql`CURRENT_TIMESTAMP`),
    updatedAt: integer('updated_at').default(sql`CURRENT_TIMESTAMP`),
  },
  (t) => [primaryKey({ columns: [t.teamId, t.userId] })],
);

export const teamMembersRelations = relations(teamMembers, ({ one }) => ({
  team: one(teams, {
    fields: [teamMembers.teamId],
    references: [teams.id],
  }),
  user: one(users, {
    fields: [teamMembers.userId],
    references: [users.id],
  }),
}));
