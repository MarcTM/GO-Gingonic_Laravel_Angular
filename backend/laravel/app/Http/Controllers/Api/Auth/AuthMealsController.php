<?php

namespace App\Http\Controllers\Api\Auth;

use Illuminate\Http\Request;
use App\Http\Controllers\Api\Controller;
use Illuminate\Support\Facades\Hash;
use App\User;
use App\Meal;
use App\Category;
use App\Data\Requests\CreateMeal;


class AuthMealsController extends Controller
{

    // Create meal
    public function createMeal(Request $request)
    {
        $user = $this->authUser();
        if($user->type != 'admin') {
            return response()->json("You do not have perimssions to do this");
        }

        $validation = new CreateMeal;
        $validated = $validation->validate($request);

        $meals = new Meal();
    	$meals->name = $validated['meal']['name'];
    	$meals->description = $validated['meal']['description'];
        $meals->price = $validated['meal']['price'];

        $meals->save();

        $cat_id = $validated['meal']['category'];
        if($cat_id > 0) {
            $meal = Meal::latest('created_at')->first();
            $category = Category::where('id', $cat_id)->first();
            if (!$meal->categories->contains($category->id)) {
                $meal->categories()->attach($category);
            }
        }

        return response()->json($meals);
    }

    // // Update meal
    // public function updateMeal(Request $request, $id)
    // {
    //     $user = $this->authUser();
    //     if($user->type != 'admin') {
    //         return response()->json("You do not have perimssions to do this");
    //     }

    //     $meal = Meal::find($id);

    //     // error_log($meal);
    //     $meal->name = $request->name;
    //     $meal->description = $request->description;

    //     $meal->save();

    //     return response()->json($meal);
    // }

    // // Delete meal
    // public function deleteMeal($id)
    // {
    //     $user = $this->authUser();
    //     if($user->type != 'admin') {
    //         return response()->json("You do not have perimssions to do this");
    //     }
        
    //     $meal = Meal::find($id);
    //     $meal->delete();
    //     return response()->json($meal);
    // }

}
