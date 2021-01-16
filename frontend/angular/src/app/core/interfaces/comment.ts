import { User } from './user';

export interface Comment {
    id: number;
    body: string;
    recipe_id: number;
    User: User;
}