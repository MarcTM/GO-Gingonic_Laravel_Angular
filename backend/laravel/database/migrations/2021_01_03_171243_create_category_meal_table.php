<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateCategoryMealTable extends Migration
{
    public function up()
    {
        Schema::create('category_meal', function (Blueprint $table) {
            $table->increments('id');
            $table->unsignedInteger('category_id');
            $table->unsignedInteger('meal_id');
            $table->timestamps();
        });
    }

    
    public function down()
    {
        Schema::dropIfExists('category_meal');
    }
}
