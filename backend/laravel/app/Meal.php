<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use App\Category;

class Meal extends Model
{
    protected $table = 'meals';
    protected $fillable = ['name', 'description', 'price', 'category_id'];

    public function categories(){
        return $this->hasMany(Category::class)->latest();
    }
}
