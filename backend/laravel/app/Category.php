<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use App\Meal;

class Category extends Model
{
    protected $table = 'categories';
    protected $fillable = ['name', 'image'];

    
    public function meals()
    {
        return $this->belongsToMany(Meal::class);
    }
}