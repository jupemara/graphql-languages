import { Author } from '../models/author.model';

export interface IAuthorRepository {
  findById(id: number): Author;
}
