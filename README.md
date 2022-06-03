# HomeVision Backend Take Home


## Description:

Take home interview for HomeVision.
Focusing on writing a clean coded go program that consumes and download images from an API that can fail randomly and is slow.

Full description at:
https://www.notion.so/HomeVision-Backend-Take-Home-94ec9ce5da0341fb9d3dc547a5bc2e15


## How to build & run:
```
go build .            // Generates the executable for your OS
go run main.go        // Runs the generated executable
```
Images are downloaded on "housesImages" folder


## Dev Insights:

### Concurrency:
There are two places where go concurrency comes handy:
1. Fetching all pages at once.
2. Downloading all houses images from a page at once.

When the program starts, we run a go routine for each page that must be fetched.
Each time we successfully fetch a new page, a go routine starts for each house to download its image.

### Backoff:
As the API could fail randomly, I implemented a simple exponentialBackoff limited with max attempts.
Each request attempt will wait double the time of the last request

### HttpMockClient:
I needed a simple way to mock the http requests to test the houseService.
I tried many ways of doing It like implementing interfaces, httptest package to generate a custom mock Server, and some external libraries. But ended up creating a simple mocking system similar to testify library.
```.Simulate(url, statusCode, bodyResponse)``` method from httpMockClient will only return the mock request if you match the exact url

### DependencyInjection
To be able to use my httpMockClient, the houseService needs to be injected with an httpClient.
So I can use httpMockClient on tests, and the httpRetryClient on the running app.

### Log
Huge emphasis on a clean log with wrapped errors for a clear stacktrace
