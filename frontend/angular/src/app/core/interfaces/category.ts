import { Meal } from './meal';

export interface Category {
    id: number;
    name: string;
    image: string;
    created_at: string;
    updated_at: string;
}

export interface CategoryMeals {
    id: number;
    name: string;
    image: string;
    created_at: string;
    updated_at: string;
    meals: Meal[];
}