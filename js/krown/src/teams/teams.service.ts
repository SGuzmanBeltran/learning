import { eq } from 'drizzle-orm';
import { CreateTeamDto } from './dto/create-team.dto';
import {
  ConflictException,
  Inject,
  Injectable,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { UpdateTeamDto } from './dto/update-team.dto';
import { Team, teams } from '../drizzle/schemas/teams.schema';
import { DRIZZLE } from '../drizzle/drizzle.module';
import {
  TeamMember,
  teamMembers,
} from '../drizzle/schemas/team_members.schema';
import { DrizzleDB } from 'drizzle/types/drizzle';

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

    const team: Team = await this.drizzle.transaction(async (tx) => {
      const [team]: Team[] = await tx
        .insert(teams)
        .values({
          name: createTeamDto.name,
        })
        .returning();

      if (!team) {
        throw new InternalServerErrorException('Failed to create team');
      }

      const [teamMember]: TeamMember[] = await tx
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
    });
    return team;
  }

  async findAll(userId: number): Promise<Team[]> {
    const userTeams = await this.drizzle
      .select()
      .from(teams)
      .innerJoin(teamMembers, eq(teams.id, teamMembers.teamId))
      .where(eq(teamMembers.userId, userId));

    return userTeams.map((team) => team.teams);
  }

  async findOne(id: number): Promise<Team> {
    const [team] = await this.drizzle
      .select()
      .from(teams)
      .where(eq(teams.id, id))
      .limit(1);

    if (!team) {
      throw new NotFoundException(`Team with id ${id} not found`);
    }

    return team;
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
