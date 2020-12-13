<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use App\Category;

class Meal extends Model
{
    protected $table = 'meals';
    protected $fillable = ['name', 'description', 'price'];
}
