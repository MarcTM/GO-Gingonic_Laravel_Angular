<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Meal;


class MealsController extends Controller
{
    // Meals

    // Create meal
    public function create(Request $request)
    {
        $meals = new Meal();

        // error_log($request);
    	$meals->name = $request->name;
    	$meals->description = $request->description;
        $meals->price = $request->price;
        $meals->category_id = $request->category_id;

    	$meals->save();

    	return response()->json($meals);
    }


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


    // Update meal
    public function updateMeal(Request $request, $id)
    {
        $meal = Meal::find($id);

        // error_log($meal);
        $meal->name = $request->name;
        $meal->description = $request->description;

        $meal->save();

        return response()->json($meal);
    }


    // Delete meal
    public function deleteMeal($id)
    {
        $meal = Meal::find($id);
        $meal->delete();
        return response()->json($meal);
    }

}
