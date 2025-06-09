/* eslint-disable @typescript-eslint/no-unsafe-return */
import { eq } from 'drizzle-orm';
import { CreateTeamDto } from './dto/create-team.dto';
import {
  ConflictException,
  ForbiddenException,
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

    const transactionTeam: Team = await this.drizzle.transaction(async (tx) => {
      const [team]: Team[] = await tx
        .insert(teams)
        .values({
          name: createTeamDto.name,
          playersCount: 1,
        })
        .returning();

      if (!team) {
        throw new InternalServerErrorException('Failed to create team');
      }

      try {
        await tx.insert(teamMembers).values({
          teamId: team.id,
          userId: userID,
          isLeader: true,
        });
      } catch {
        throw new InternalServerErrorException('Failed to add team member');
      }

      return team;
    });

    const [team] = await this.drizzle
      .select()
      .from(teams)
      .where(eq(teams.id, transactionTeam.id))
      .limit(1);

    return team;
  }

  async findAll(userId: number): Promise<Team[]> {
    const userTeams = await this.drizzle
      .select({ teamId: teamMembers.teamId })
      .from(teamMembers)
      .where(eq(teamMembers.userId, userId));

    const teamIds = userTeams.map((t) => t.teamId);

    return this.drizzle.query.teams.findMany({
      where: (teams, { inArray }) => inArray(teams.id, teamIds),
      with: {
        teamMembers: true,
      },
    });
    // return userTeams.map((team) => {
    //   return {
    //     ...team.teams,
    //     members: team.team_members as unknown as TeamMember[],
    //   };
    // });
  }

  async findOne(id: number): Promise<{
    members: TeamMember[];
  }> {
    const rows = await this.drizzle
      .select()
      .from(teams)
      .leftJoin(teamMembers, eq(teams.id, teamMembers.teamId))
      .where(eq(teams.id, id));

    if (!rows.length) {
      throw new NotFoundException(`Team with id ${id} not found`);
    }

    const members: TeamMember[] = [];
    for (const row of rows) {
      if (row.team_members) {
        members.push(row.team_members as unknown as TeamMember);
      }
    }

    const team = rows[0].teams as Team;

    return {
      ...team,
      members,
    };
  }

  async update(id: number, updateTeamDto: UpdateTeamDto): Promise<Team> {
    await this.findOne(id);
    try {
      const [updatedTeam] = await this.drizzle
        .update(teams)
        .set(updateTeamDto)
        .where(eq(teams.id, id))
        .returning();

      return updatedTeam;
    } catch {
      throw new InternalServerErrorException('Failed to update team');
    }
  }

  async remove(
    id: number,
    userId: number,
  ): Promise<{ message: string; deleted: boolean }> {
    const team = await this.findOne(id);
    const isTeamMember = this.isTeamMember(team, userId);

    if (!isTeamMember) {
      throw new ForbiddenException(
        'You are not authorized to delete this team',
      );
    }

    try {
      await this.drizzle.delete(teams).where(eq(teams.id, id));
    } catch {
      throw new InternalServerErrorException('Failed to delete team');
    }

    return { message: `Team ${id} deleted successfully`, deleted: true };
  }

  async addMember(teamID: number, userID: number) {
    const team = await this.findOne(teamID);
    const isTeamMember = this.isTeamMember(team, userID);

    if (isTeamMember) {
      throw new ConflictException('User is already a member of this team');
    }

    try {
      await this.drizzle.insert(teamMembers).values({
        teamId: teamID,
        userId: userID,
      });

      return { message: `Member ${userID} added to team ${teamID}` };
    } catch {
      throw new InternalServerErrorException('Failed to add team member');
    }
  }

  async removeMember(teamID: number, userID: number) {
    return `This action removes a member from a #${teamID} team`;
  }

  isTeamMember(team: { members: TeamMember[] }, userID: number): boolean {
    return team.members.some((member) => member.userId === userID);
  }
}
