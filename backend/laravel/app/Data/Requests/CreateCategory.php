<?php

namespace App\Data\Requests;

use Illuminate\Http\Request;

class CreateCategory
{
    public function validate(Request $request)
    {
        return $request->validate([
            'category.name' => 'required|string',
            'category.image' => 'required|string',
        ]);
    }
}