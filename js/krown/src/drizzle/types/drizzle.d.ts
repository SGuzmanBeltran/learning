import { LibSQLDatabase } from 'drizzle-orm/libsql';
import * as schema from '../schemas/schema';

export type DrizzleDB = LibSQLDatabase<typeof schema>;
