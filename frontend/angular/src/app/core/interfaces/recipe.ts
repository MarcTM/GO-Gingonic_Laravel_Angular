import { RecipeAuthor } from './user';

export interface Recipe {
    id: number;
    name: string;
    description: string;
    image: string;
    author: RecipeAuthor;
}