import { Time } from '@angular/common';
import { Timestamp } from 'rxjs';

export interface Meal {
    id: number;
    name: string;
    description: string;
    price: number;
    // category_id: number;
    created_at: string;
    updated_at: string;
}