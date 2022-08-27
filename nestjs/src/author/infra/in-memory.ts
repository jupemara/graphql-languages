import { Injectable } from '@nestjs/common';
import { IAuthorRepository } from '../domain/repository';
import { Author } from '../models/author.model';

export const SYMBOL = 'InMemoryRepository';

@Injectable()
export class InMemoryRepository implements IAuthorRepository {
  private vs = [
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
  public findById(id: number): Author {
    return this.vs.find((v) => {
      return v.id === id;
    });
  }
}
