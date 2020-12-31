<?php

use Illuminate\Http\Request;


Route::group(['prefix'=>'auth'], function ($router) {
    
    // Authentication
    Route::post('register', 'Api\Auth\AuthController@register');
    Route::post('login', 'Api\Auth\AuthController@login');
    Route::post('refresh', 'Api\Auth\AuthController@refresh');
    Route::post('validate', 'Api\Auth\AuthController@isAdmin');
});


Route::group(['prefix'=>'meals'], function ($router) {
    
    // Meals authed
    Route::post('/', 'Api\Auth\AuthMealsController@createMeal');
    Route::put('/{id}', 'Api\Auth\AuthMealsController@updateMeal');
    Route::delete('/{id}', 'Api\Auth\AuthMealsController@deleteMeal');

    // Meals
    Route::get('/', 'MealsController@show');
    Route::get('/{id}', 'MealsController@showMeal');
});


Route::group(['prefix'=>'slider'], function ($router) {
    
    // Slider
    Route::get('/', 'SliderController@show');
});