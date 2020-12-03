<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use App\Meal;

class Category extends Model
{
    protected $table = 'categories';
    protected $fillable = ['name'];

    public function meals(){
        return $this->hasMany(Meal::class)->latest();
    }
}
