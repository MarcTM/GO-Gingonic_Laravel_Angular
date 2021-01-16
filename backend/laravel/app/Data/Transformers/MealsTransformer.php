<?php

namespace App\Data\Transformers;

class MealsTransformer
{
    public function transform($data)
    {
        $meals = [];
        foreach ($data as &$meal) {
            $result = (object)array(
                "id" => $meal->id,
                "name" => $meal->name,
                "description" => $meal->description,
                "image" => $meal->image,
                "price" => $meal->price
            );
            array_push($meals, $result);
        }
        return $meals;
    }
}