import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { TeamsService } from './teams.service';
import { CreateTeamDto } from './dto/create-team.dto';
import { UpdateTeamDto } from './dto/update-team.dto';

@Controller('teams')
export class TeamsController {
  constructor(private readonly teamsService: TeamsService) {}

  @Get()
  findAll() {
    return this.teamsService.findAll();
  }

  @Post()
  create(@Body() createTeamDto: CreateTeamDto) {
    return this.teamsService.create(createTeamDto);
  }

  @Get(':teamID')
  findOne(@Param('teamID') teamID: string) {
    return this.teamsService.findOne(+teamID);
  }

  @Patch(':teamID')
  update(
    @Param('teamID') teamID: string,
    @Body() updateTeamDto: UpdateTeamDto,
  ) {
    return this.teamsService.update(+teamID, updateTeamDto);
  }

  @Delete(':teamID')
  remove(@Param('teamID') teamID: string) {
    return this.teamsService.remove(+teamID);
  }

  @Post(':teamID/members')
  addMember(@Param('teamID') teamID: string) {
    return this.teamsService.addMember(+teamID);
  }

  @Delete(':teamID/members/:userID')
  removeMember(
    @Param('teamID') teamID: string,
    @Param('userID') userID: string,
  ) {
    return this.teamsService.removeMember(+teamID, +userID);
  }
}
