import { Module } from '@nestjs/common';
import { AuthorResolver } from './author.resolver';
import { AuthorService } from './author.service';
import { InMemoryRepository, SYMBOL } from './infra/in-memory';

@Module({
  providers: [
    {
      provide: SYMBOL,
      useClass: InMemoryRepository,
    },
    AuthorService,
    AuthorResolver,
  ],
})
export class AuthorModule {}
