# Image Recognition API

## About
This will be a go application that will receive images and identify objects in the image using it's underlying model

## How to run
This application will require a google cloud storage bucket, with accompanying service account credentials
All environment variables can be seen in `.env.sample`

Run using `go run main.go`. This will use create a default Gin that will listen on `localhost:8080`

## Sample Request & Response
Send an HTTP POST request using a curl command:

```
curl -F "file=@<filepath/image-name>" "localhost:8080/label-image"
```

Response:
```
{
  "labels": ["Water","Photograph","Sky","Building","Boat","Rectangle","Art","Urban design","Travel","Magenta"]
}

```