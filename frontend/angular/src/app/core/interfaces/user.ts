import { Recipe } from './recipe';

export interface User {
    id: number;
    username: string;
    email: string;
    password: string;
    bio: string;
    image: string;
    type: string;
}

export interface RecipeAuthor {
    id: number;
    username: string;
    image: string;
}

export interface Profile {
    id: number;
    username: string;
    image: string;
    bio: string;
    recipes: Recipe[];
}