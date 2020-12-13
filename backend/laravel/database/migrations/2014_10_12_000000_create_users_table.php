<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateUsersTable extends Migration
{
    // Create table
    public function up()
    {
        Schema::create('users', function (Blueprint $table) {
            $table->increments('id');
            $table->string('username')->unique();
            $table->string('email')->unique();
            $table->string('password');
            $table->string('bio')->nullable();
            $table->string('image')->nullable();
            $table->string('type');
        });
    }

    // Drop table
    public function down()
    {
        Schema::dropIfExists('users');
    }
}
