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

        $bearer = auth()->attempt($credentials);
        return response()->json(['token' => $bearer]);
    }


    // // Returns current user
    // public function me(){
    //     return response()->json(auth()->user());
    // }


    // // Logout
    // public function logout()
    // {
    //     auth()->logout();

    //     return response()->json(['message' => 'Successfully logged out']);
    // }


    // // Refresh token
    // public function refresh()
    // {
    //     return $this->respondWithToken(auth()->refresh());
    // }

}
