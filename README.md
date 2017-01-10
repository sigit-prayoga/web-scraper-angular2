# Web Scraper using Angular2 and Go

This project is intended to share a basic knowledge about Angular 2 and Go as a RESTful API.
Simple use case here that user can input any product in input field and it's going to show 2 top categories related to the product.

Thanks to `http://bukalapak.com/` to provide a good categories tree that we can scrap :)

## Run the Go server
```sh
$ cd web-scraper-angular2/src/server

$ go run main.go
```

## Run the Angular2 client in development environment
```sh
$ cd web-scraper-angular2

$ ng serve
```

Note: Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.
