import { Args, Int, Query, Resolver } from '@nestjs/graphql';
import { AuthorService } from './author.service';
import { Author } from './models/author.model';

@Resolver((of) => Author)
export class AuthorResolver {
  constructor(private authorService: AuthorService) {}

  @Query((returns) => Author)
  async author(@Args('id', { type: () => Int }) id: number) {
    return this.authorService.findById(id);
  }
}
