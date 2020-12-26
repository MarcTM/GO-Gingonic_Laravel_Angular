<?php

namespace App\Http\Controllers\Api\Auth;

use Illuminate\Http\Request;
use App\Http\Controllers\Api\Controller;
use Illuminate\Support\Facades\Hash;
use App\User;
use App\Meal;


class AuthMealsController extends Controller
{

    // Create meal
    public function createMeal(Request $request)
    {
        $user = $this->authUser();
        if($user->type != 'admin') {
            return response()->json("You do not have perimssions to do this");
        }

        $meals = new Meal();

    	$meals->name = $request->name;
    	$meals->description = $request->description;
        $meals->price = $request->price;

    	$meals->save();
    }


    // Update meal
    public function updateMeal(Request $request, $id)
    {
        $user = $this->authUser();
        if($user->type != 'admin') {
            return response()->json("You do not have perimssions to do this");
        }

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
        $user = $this->authUser();
        if($user->type != 'admin') {
            return response()->json("You do not have perimssions to do this");
        }
        
        $meal = Meal::find($id);
        $meal->delete();
        return response()->json($meal);
    }

}
