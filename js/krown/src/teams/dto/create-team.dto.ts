import { IsNotEmpty, IsString } from 'class-validator';

import { ApiProperty } from '@nestjs/swagger';

export class CreateTeamDto {
  @ApiProperty({
    description: 'The name of the team',
    example: 'Team 1',
  })
  @IsString()
  @IsNotEmpty()
  name: string;
}
