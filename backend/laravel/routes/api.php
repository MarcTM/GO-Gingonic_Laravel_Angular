<?php

// use Illuminate\Http\Request;



// Route::middleware('auth:api')->get('/user', function (Request $request) {
//     return $request->user();
// });


Route::group([

    // 'middleware' => 'api',
    'prefix' => 'auth'

], function ($router) {
    
    Route::post('login', 'AuthController@login');
    Route::post('logout', 'AuthController@logout');
    Route::post('refresh', 'AuthController@refresh');
    Route::post('me', 'AuthController@me');

});



// Students

Route::post('/student', 'ApiController@create');

Route::get('/students', 'ApiController@show');

Route::get('/student/{id}', 'ApiController@showStudent');

Route::put('/student/{id}', 'ApiController@updateStudent');

Route::delete('/student/{id}', 'ApiController@deleteStudent');


// Meals

Route::post('/meal', 'MealsController@create');

Route::get('/meals', 'MealsController@show');

Route::get('/meal/{id}', 'MealsController@showMeal');

Route::put('/meal/{id}', 'MealsController@updateMeal');

Route::delete('/meal/{id}', 'MealsController@deleteMeal');


// Categories

Route::post('/category', 'CategoriesController@create');

Route::get('/categories', 'CategoriesController@show');

Route::get('/category/{id}', 'CategoriesController@showCategory');

Route::put('/category/{id}', 'CategoriesController@updateCategory');

Route::delete('/category/{id}', 'CategoriesController@deleteCategory');


// Test 

Route::post('/tests', 'TestsController@create');

Route::get('/tests', 'TestsController@show');