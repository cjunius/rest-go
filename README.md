# rest-go
Test REST-API written in Go

Provides a REST endpoint for any resource accepting a JSON body
| HTTP METHOD | ENDPOINT | Description |
| ----- | ----- | ----- |
| POST | localhost:8080/{resource} | Creates a resource and adds/overwrites an id attribute |
| GET | localhost:8080/{resource} | Retrieves a list of resources |
| GET | localhost:8080/{resource}/{id} | Retrieves a specific resource |
| PUT | localhost:8080/{resource}/{id} | Replaces a specific resource |
| PATCH | localhost:8080/{resource}/{id} | Updates a specific resource with a new/updated attribute |
| DELETE | localhost:8080/{resource}/{id} | Removes a specific resource |

Usage:
- git clone https://github.com/cjunius/rest-go.git
- cd rest-go
- docker build -t rest-go .
- docker run -d -p 8080:8080 rest-go:latest

Alternative Usage:
- docker pull cjunius/rest-go:main
- docker run --name rest-go -d -p 8080:8080 cjunius/rest-go:main




