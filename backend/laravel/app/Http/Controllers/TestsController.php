<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Test;


class TestsController extends Controller
{

    public function create(Request $request)
    {
        $tests = new Test();
        
        error_log($tests);
        error_log($request);

    	$tests->username = $request->username;
        $tests->email = $request->email;
        $tests->password = $request->password;
        $tests->bio = $request->bio;
        $tests->image = $request->image;

        $tests->save();

    	return response()->json($tests);
    }

    public function show()
    {
        $tests = Test::all();
    	return response()->json($tests);
    }

}
