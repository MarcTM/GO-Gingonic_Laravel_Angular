<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Category;
use App\Data\Transformers\CategoriesTransformer;
use App\Data\Transformers\MealsTransformer;


class CategoriesController extends Controller
{

    // Show all categories
    public function show()
    {
        $categories = Category::all();

        $transformer = new CategoriesTransformer;
        $transform = $transformer->transform($categories);
        
        return response()->json($transform);
    }

    // Get category meals
    public function getMeals($id)
    {
        $category = Category::find($id);

        $meals = $category->meals;
        $transformer = new MealsTransformer;
        $transform = $transformer->transform($meals);
        
        return response()->json($category);
    }

}
