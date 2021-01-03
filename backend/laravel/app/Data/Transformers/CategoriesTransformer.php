<?php

namespace App\Data\Transformers;

class CategoriesTransformer
{
    public function transform($data)
    {
        $categories = [];
        foreach ($data as &$category) {
            $result = (object)array(
                "id"=>$category->id,
                "name"=>$category->name,
                "image"=>$category->image
            );
            array_push($categories, $result);
        }
        return $categories;
    }
}