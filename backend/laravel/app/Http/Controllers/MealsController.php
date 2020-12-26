<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Meal;


class MealsController extends Controller
{
    // Meals

    // Show all meals
    public function show()
    {
        $meals = Meal::all();
        return response()->json($meals);
    }


    // Show one meal
    public function showMeal($id)
    {
        $meal = Meal::find($id);
        return response()->json($meal);
    }

}
