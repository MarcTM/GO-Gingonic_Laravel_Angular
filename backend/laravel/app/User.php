<?php

namespace App;

use Illuminate\Notifications\Notifiable;
use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Tymon\JWTAuth\Contracts\JWTSubject;

// class User extends Authenticatable
class User extends Authenticatable implements JWTSubject
{
    use Notifiable;

    protected $fillable = ['username', 'email', 'password', 'bio', 'image'];

    protected $hidden = ['password', 'type'];


    public function getJWTIdentifier()
    {
        return $this->getKey();
    }


    public function getJWTCustomClaims()
    {
        return [];
    }
}
