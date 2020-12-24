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