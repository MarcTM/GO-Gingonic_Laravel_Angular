<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class AuthController extends Controller
{

    public function _construct()
    {
        $this->middleware('auth:api', ['except' => ['login']]);
    }


    // Login
    public function login()
    {
        $credentials = request(['email', 'password']);

        if (! $token = auth()->attempt($credentials)) {
            return response()->json(['error' => 'Unathorized'], 401);
        }

        return $this->respondWithToken($token);
    }


    // Returns current user
    public function me(){
        return response()->json(auth()->user());
    }


    // Logout
    public function logout()
    {
        auth()->logout();

        return response()->json(['message' => 'Successfully logged out']);
    }


    // Refresh token
    public function refresh()
    {
        return $this->respondWithToken(auth()->refresh());
    }


    // Get the token array structure
    protected function respondWithToken($token)
    {
        return response()->json([
            'access_token' => $token,
            'token_type' => 'bearer',
            'expires_in' => auth()->factory()->getTTL() * 60
        ]);
    }

}
