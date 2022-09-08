export interface WeatherData {
    country?: string,
    temp?: string,
    pressure?: string,
    humidity?: string
  }
  
  export class WeatherDataModel implements WeatherData{
    constructor(
      public country?: string,
      public temp?: string,
      public pressure?: string,
      public humidity?: string
    ) {
    }
  }