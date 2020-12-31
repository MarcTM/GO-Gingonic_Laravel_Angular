<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class HomeSliderTable extends Migration
{
    // Create table
    public function up()
    {
        Schema::create('slider', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name');
            $table->string('image');
            $table->timestamps();
        });
    }

    // Drop table
    public function down()
    {
        Schema::dropIfExists('slider');
    }
}
