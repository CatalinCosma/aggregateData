# Aggregate data app
It's a small service that uses the UpFluence stream api for aggregating data 

## Prerequisites
* go 1.20.7
* gin

## Instalation

* clone or download repository
* in the project root run `go mod tidy` and `go mod vendor`

## How to use

To run the application, in the project root run

`go run build && ./UpChallenge` or `go run main.go`

This will start a web server on port 8080, with one route defined:
* `GET/http://localhost:8080/analysis?duration=6s&dimension=likes` 

The `duration` parameter and `dimension` parameter can be set as:

`{3s,5s,10s,30s}` seconds for `duration` the time
that the app si blocking and getting data from the stream

`{likes or comments}` for dimesion

Based on this query parameters the service will offer different responses:

* for likes:
{
  "total_posts": 20
  "minimum_timestamp": X,
  "maximum_timestamp": Y,
  "avg_likes": 50
}
* for comments:
{
  "total_posts":10
  "minimum_timestamp": X,
  "maximum_timestamp": Y,
  "avg_comments": 20
}
