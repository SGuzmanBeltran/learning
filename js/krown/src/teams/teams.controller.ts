import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  UseGuards,
  Req,
} from '@nestjs/common';
import { TeamsService } from './teams.service';
import { CreateTeamDto } from './dto/create-team.dto';
import { UpdateTeamDto } from './dto/update-team.dto';
import { ApiCreatedResponse } from '@nestjs/swagger';
import { Team } from '../drizzle/schemas/teams.schema';
import { AuthGuard, AuthenticatedRequest } from '../auth/auth.guard';

@UseGuards(AuthGuard)
@Controller('teams')
export class TeamsController {
  constructor(private readonly teamsService: TeamsService) {}

  @Get()
  findAll(@Req() req: AuthenticatedRequest) {
    return this.teamsService.findAll(req.user!.userId);
  }

  @Post()
  @ApiCreatedResponse({
    description: 'The team has been successfully created.',
    type: Number,
  })
  async create(
    @Body() createTeamDto: CreateTeamDto,
    @Req() req: AuthenticatedRequest,
  ): Promise<Team> {
    return await this.teamsService.create(createTeamDto, req.user!.userId);
  }

  @Get(':teamID')
  async findOne(@Param('teamID') teamID: string) {
    return await this.teamsService.findOne(+teamID);
  }

  @Patch(':teamID')
  async update(
    @Param('teamID') teamID: string,
    @Body() updateTeamDto: UpdateTeamDto,
  ) {
    return await this.teamsService.update(+teamID, updateTeamDto);
  }

  @Delete(':teamID')
  async remove(
    @Param('teamID') teamID: string,
    @Req() req: AuthenticatedRequest,
  ): Promise<{
    message: string;
    deleted: boolean;
  }> {
    const user = req.user!;
    return await this.teamsService.remove(+teamID, user.userId);
  }

  @Post(':teamID/members')
  async addMember(
    @Param('teamID') teamID: string,
    @Req() req: AuthenticatedRequest,
  ) {
    return await this.teamsService.addMember(+teamID, req.user!.userId);
  }

  @Delete(':teamID/members/:userID')
  async removeMember(
    @Param('teamID') teamID: string,
    @Param('userID') userID: string,
  ) {
    return await this.teamsService.removeMember(+teamID, +userID);
  }
}
