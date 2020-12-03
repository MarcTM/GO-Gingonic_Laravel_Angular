<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Category;
use App\Meal;


class CategoriesController extends Controller
{
    // Categories

    public function create(Request $request)
    {
    	$categories = new Category();

    	$categories->name = $request->name;

    	$categories->save();

    	return response()->json($categories);
    }

    public function show()
    {
        // $categories = Category::all()->meals();
        $categories = Category::all();
    	return response()->json($categories);
    }

    public function showCategory($id)
    {
        $category = Category::find($id);
        return response()->json($category);
    }

    public function updateCategory(Request $request, $id)
    {
        $category = Category::find($id);

        // error_log($category);
        $category->name = $request->name;
        $category->description = $request->description;

        $category->save();

        return response()->json($category);
    }

    public function deleteCategory($id)
    {
        $category = Category::find($id);
        $category->delete();
        return response()->json($category);
    }
}
