import { Controller, Get } from '@nestjs/common';

@Controller('people')
export class PeopleController {
  @Get()
  fn(): string {
    return 'people controller';
  }
}
