import { eq } from 'drizzle-orm';
import { CreateTeamDto } from './dto/create-team.dto';
import { DrizzleDB } from '../drizzle/types/drizzle';
import {
  ConflictException,
  Inject,
  Injectable,
  InternalServerErrorException,
} from '@nestjs/common';
import { UpdateTeamDto } from './dto/update-team.dto';
import { Team, teams } from '../drizzle/schemas/teams.schema';
import { DRIZZLE } from '../drizzle/drizzle.module';
import {
  TeamMember,
  teamMembers,
} from '../drizzle/schemas/team_members.schema';

@Injectable()
export class TeamsService {
  constructor(@Inject(DRIZZLE) private drizzle: DrizzleDB) {}

  async create(createTeamDto: CreateTeamDto, userID: number): Promise<Team> {
    const [existingTeam] = await this.drizzle
      .select()
      .from(teams)
      .where(eq(teams.name, createTeamDto.name))
      .limit(1);

    if (existingTeam) {
      throw new ConflictException('Team name already exists');
    }

    const [team]: Team[] = await this.drizzle
      .insert(teams)
      .values({
        name: createTeamDto.name,
      })
      .returning();

    const [teamMember]: TeamMember[] = await this.drizzle
      .insert(teamMembers)
      .values({
        teamId: team.id,
        userId: userID,
        isLeader: true,
      })
      .returning();

    if (!teamMember) {
      throw new InternalServerErrorException('Failed to add team member');
    }

    return team;
  }

  findAll() {
    return `This action returns all teams`;
  }

  findOne(id: number) {
    return `This action returns a #${id} team`;
  }

  update(id: number, updateTeamDto: UpdateTeamDto) {
    return `This action updates a #${id} team`;
  }

  remove(id: number) {
    return `This action removes a #${id} team`;
  }

  addMember(teamID: number) {
    return `This action adds a member to a #${teamID} team`;
  }

  removeMember(teamID: number, userID: number) {
    return `This action removes a member from a #${teamID} team`;
  }
}
