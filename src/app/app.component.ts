import { Component } from '@angular/core';
import { Http, Response } from '@angular/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  //model
  product: string;

  categories: string[];

  constructor(private http: Http) {
    this.categories = [];
  }

  doScrap() {
    console.log("asking for categories...");
    this.http.get('http://localhost:3001/go/api/scraping?keyword='+this.product).subscribe(res => {
      console.log('The response: ', res);
      this.categories = res.json()
    });
  } 
}
