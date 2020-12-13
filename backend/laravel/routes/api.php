<?php

use Illuminate\Http\Request;


Route::group([

    'middleware' => ['cors'],
    'prefix' => 'auth'

], function ($router) {
    
    // Register and login
    Route::post('register', 'Api\Auth\AuthController@register');
    Route::post('login', 'Api\Auth\AuthController@login');
    // Route::post('logout', 'Api\Auth\AuthController@logout');
    // Route::post('refresh', 'Api\Auth\AuthController@refresh');
    // Route::post('me', 'Api\Auth\AuthController@me');

});


Route::group([

    'middleware' => ['cors'],
    'prefix' => 'laravel'

], function ($router) {
    
    // Meals
    Route::post('/meal', 'MealsController@create');
    Route::get('/meals', 'MealsController@show');
    Route::get('/meal/{id}', 'MealsController@showMeal');
    Route::put('/meal/{id}', 'MealsController@updateMeal');
    Route::delete('/meal/{id}', 'MealsController@deleteMeal');
});