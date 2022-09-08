import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import { WeatherData } from './app.model';

@Injectable({
  providedIn: 'root'
})
export class WeatherDataService {

  constructor(
    private http: HttpClient
  ) { }

  getWdata(PwDataURL:string): Observable<WeatherData> {
    return this.http.get<WeatherData>(PwDataURL, {})
  }
}
