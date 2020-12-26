import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent {

  constructor() {}

  
  customOptions: any = {
    loop: true,
    items: 1,
    autoplay:true,
    dots: true
  }

  slides = [
    {
      image: "../../assets/img/meal1.jpeg"
    },
    {
      image: "../../assets/img/meal2.jpeg"
    },
  ]

}
