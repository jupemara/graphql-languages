import { Inject, Injectable } from '@nestjs/common';
import { IAuthorRepository } from './domain/repository';
import { SYMBOL } from './infra/in-memory';
import { Author } from './models/author.model';

@Injectable()
export class AuthorService {
  constructor(
    @Inject(SYMBOL) private readonly authorRepository: IAuthorRepository,
  ) {}
  private vs: Author[] = [
    {
      id: 1,
      firstName: 'hoge',
      lastName: 'piyo',
    },
    {
      id: 2,
      firstName: 'chim',
      lastName: 'taro',
    },
    {
      id: 3,
      firstName: 'fuga',
      lastName: 'soso',
    },
  ];
  findById(id: number): Author {
    return this.authorRepository.findById(id);
  }
}
