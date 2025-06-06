import { eq } from 'drizzle-orm';
import { CreateTeamDto } from './dto/create-team.dto';
import { DrizzleDB } from '../drizzle/types/drizzle';
import { ConflictException, Inject, Injectable } from '@nestjs/common';
import { UpdateTeamDto } from './dto/update-team.dto';
import { teams } from '../drizzle/schemas/teams.schema';
import { DRIZZLE } from '../drizzle/drizzle.module';

@Injectable()
export class TeamsService {
  constructor(@Inject(DRIZZLE) private drizzle: DrizzleDB) {}

  async create(createTeamDto: CreateTeamDto): Promise<number> {
    const [existingTeam] = await this.drizzle
      .select()
      .from(teams)
      .where(eq(teams.name, createTeamDto.name))
      .limit(1);

    if (existingTeam) {
      throw new ConflictException('Team name already exists');
    }

    const [team] = await this.drizzle
      .insert(teams)
      .values({
        name: createTeamDto.name,
      })
      .returning();
    return team.id;
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
