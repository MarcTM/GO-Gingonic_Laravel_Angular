<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Meal;
use App\Category;
use App\Data\Transformers\MealsTransformer;


class MealsController extends Controller
{

    // Show all meals
    public function show()
    {
        $meals = Meal::all();

        $transformer = new MealsTransformer;
        $transform = $transformer->transform($meals);
        
        return response()->json($transform);
    }


    // Show one meal
    public function showMeal($id)
    {
        $meal = Meal::find($id);
        return response()->json($meal);
    }

}
