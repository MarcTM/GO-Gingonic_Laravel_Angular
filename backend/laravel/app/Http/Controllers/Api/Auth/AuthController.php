<?php

namespace App\Http\Controllers\Api\Auth;

use Illuminate\Http\Request;
use App\Http\Controllers\Api\Controller;
use Illuminate\Support\Facades\Hash;
use App\User;


class AuthController extends Controller
{

    // Register
    public function register(Request $request)
    {
        $user = new User();
        $user->username = $request->username;
    	$user->email = $request->email;
        $user->password = Hash::make($request->password);
        $user->type = "user";

        $user->save();

        return response()->json($user);
    }


    // Login
    public function login(Request $request)
    {
        $credentials = $request->only(['email','password']);

        if(!$bearer = auth()->attempt($credentials)) {
            return response()->json(['error' => 'Incorrect email or password'], 401);
        }

        return response()->json(['token' => $bearer]);
    }


    // Refresh token
    public function refresh()
    {
        try {
            $newToken = auth()->refresh();
        } catch(\Tymon\JWTAuth\Exceptions\TokenInvalidException $e) {
            return response()->json(['error' => $e->getMessage()], 422);
        }

        return response()->json(['bearer' => $newToken]);
    }


    // Check if user is admin
    public function isAdmin()
    {
        $isAdmin = auth()->user();

        if(!$isAdmin){
            return response('error', 404);
        } else {
            if($isAdmin->type!="admin"){
                return response('error', 404);
            } else {
                return response()->json("ok");
            }
        }
    }


    // public function create(Request $request)
    // {
    //     $user = $this->authUser();

    //     return response()->json($user);
    // }

}
