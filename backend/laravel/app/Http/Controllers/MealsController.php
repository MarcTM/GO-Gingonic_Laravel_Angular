<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Meal;
use App\Category;


class MealsController extends Controller
{
    // Meals

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

    public function show()
    {
        $meals = Meal::select('meals.id', 'meals.name AS meal_name', 'meals.category_id', 'categories.name AS category_name')->leftJoin('categories', 'categories.id', '=', 'meals.category_id')->get();
    	return response()->json($meals);
    }

    public function showMeal($id)
    {
        $meal = Meal::find($id);
        return response()->json($meal);
    }

    public function updateMeal(Request $request, $id)
    {
        $meal = Meal::find($id);

        // error_log($meal);
        $meal->name = $request->name;
        $meal->description = $request->description;

        $meal->save();

        return response()->json($meal);
    }

    public function deleteMeal($id)
    {
        $meal = Meal::find($id);
        $meal->delete();
        return response()->json($meal);
    }

}
