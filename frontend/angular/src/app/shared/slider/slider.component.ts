import { Component, OnInit } from '@angular/core';

import { SliderService } from '../../core/services/slider.service';
import { Slider } from '../../core/interfaces/slider';

@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent implements OnInit {

  constructor(
    private sliderService: SliderService
  ) {}


  homeSlides = [];

  customOptions: any = {
    loop: true,
    items: 1,
    autoplay:true,
    dots: true
  }

  getSlides(): void {
    this.sliderService.getSlides()
    .subscribe(slides => {
      slides.map(slide => {
        slide.image = '../../assets/img/' + slide.image;
        this.homeSlides.push(slide);
      })
    });
  }

  ngOnInit(): void {
    this.getSlides();
  }

}
