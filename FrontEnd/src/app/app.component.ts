import { Component } from '@angular/core';
import { WeatherDataService } from './weather-data.service';
import { WeatherData } from './app.model';
import {faEarthAmericas, faThermometerHalf, faDroplet, faGaugeMed} from '@fortawesome/free-solid-svg-icons';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { animate, state, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  animations: [
    trigger('opacityDrop',[
      state('visible', style({
        'opacity':1,
        'display':"inline-block"
      })),
      state('hidden', style({
        'opacity':0,
        'display':"none"
      })),
      transition('visible => hidden', [
        animate(1000)
      ])
    ]),
    trigger('opacityGain',[
      state('hidden', style({
        'opacity':0,
        'display':"none"
      })),
      state('visible', style({
        'opacity':1,
        'display':"inline-block"
      })),
      transition('hidden => visible', [
        animate(1000)
      ])
    ]),
  ],
})
export class AppComponent {
  title = 'frontEnd';
  city:any
  state = 'visible'
  stateBehind = 'hidden'
  

  webResponse:WeatherData = {};
  faEarthAmericas = faEarthAmericas
  faThermometerHalf =  faThermometerHalf
  faDroplet = faDroplet
  faGaugeMed = faGaugeMed

  constructor(
    private callAPI:WeatherDataService
  ){
    // document.documentElement.style.setProperty('--frontDisplay','inline-block')
    // document.documentElement.style.setProperty('--rearDisplay','none')
  }


  getCountry() {
    this.city = document.querySelector('input')?.value
    console.log(this.city)
    let wDataURL:string = "http://localhost:8000/wData?city=" + this.city
    this.callAPI.getWdata(wDataURL).subscribe (
      res => {
        this.webResponse = res;
      }
    )
    this.state == 'visible' ?
    this.state = 'hidden' : this.state = 'hidden';

    this.stateBehind == 'hidden' ?
    this.stateBehind = 'visible' : this.stateBehind = 'visible'
    // document.documentElement.style.setProperty('--frontDisplay','none')
    // document.documentElement.style.setProperty('--rearDisplay','inline-block')

  }
  
}
