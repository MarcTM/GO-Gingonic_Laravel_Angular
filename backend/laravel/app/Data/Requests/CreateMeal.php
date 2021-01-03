<?php

namespace App\Data\Requests;

use Illuminate\Http\Request;

class CreateMeal
{
    public function validate(Request $request)
    {
        return $request->validate([
            'meal.name' => 'required|string',
            'meal.description' => 'required|string',
            'meal.price' => 'required|integer',
            'meal.category' => 'required|integer'
        ]);
    }
}