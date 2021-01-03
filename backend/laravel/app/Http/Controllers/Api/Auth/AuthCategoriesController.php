<?php

namespace App\Http\Controllers\Api\Auth;

use Illuminate\Http\Request;
use App\Http\Controllers\Api\Controller;
use Illuminate\Support\Facades\Hash;
use App\User;
use App\Category;
use App\Data\Requests\CreateCategory;


class AuthCategoriesController extends Controller
{
    
    // Create category
    public function createCategory(Request $request)
    {
        $user = $this->authUser();
        if($user->type != 'admin') {
            return response()->json("You do not have perimssions to do this");
        }

        $validation = new CreateCategory;
        $validated = $validation->validate($request);

        $categories = new Category();
    	$categories->name = $validated['category']['name'];
    	$categories->image = $validated['category']['image'];

    	$categories->save();
    }

}
