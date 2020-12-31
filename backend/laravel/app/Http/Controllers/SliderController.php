<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Slider;

class SliderController extends Controller
{
    // Show slider images
    public function show()
    {
        $slider = Slider::all();
        return response()->json($slider);
    }
}