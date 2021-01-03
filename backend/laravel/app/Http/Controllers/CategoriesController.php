<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Category;
use App\Data\Transformers\CategoriesTransformer;


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

}
