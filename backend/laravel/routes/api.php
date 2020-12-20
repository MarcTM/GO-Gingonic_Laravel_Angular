<?php

use Illuminate\Http\Request;


Route::group(['prefix'=>'auth'], function ($router) {
    
    // Authentication
    Route::post('register', 'Api\Auth\AuthController@register');
    Route::post('login', 'Api\Auth\AuthController@login');
    Route::post('refresh', 'Api\Auth\AuthController@refresh');
    Route::post('validate', 'Api\Auth\AuthController@isAdmin');
    // Route::post('create', 'Api\Auth\AuthController@create');
});


Route::group(['prefix'=>'laravel'], function ($router) {
    
    // Meals
    Route::post('/meal', 'MealsController@create');
    Route::get('/meals', 'MealsController@show');
    Route::get('/meal/{id}', 'MealsController@showMeal');
    Route::put('/meal/{id}', 'MealsController@updateMeal');
    Route::delete('/meal/{id}', 'MealsController@deleteMeal');
});