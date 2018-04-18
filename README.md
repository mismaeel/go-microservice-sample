# go-microservice-sample

The main goal of this task is to create a simple application for managing movies. Application should consist of 2 microservices responsible for: 
(A) Movies management
(B) Approving reviews 

The sample is devloped in go and mongodb 

### Refrences

* [microservices-an-example-with-docker-go-and-mongod](https://dzone.com/articles/microservices-an-example-with-docker-go-and-mongod)
* [request-response-pattern-using-go-channles](http://hassansin.github.io/request-response-pattern-using-go-channles)
* [Making rest api go](https://thenewstack.io/make-a-restful-json-api-go/)
* [go-blog-series](http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part1/)

**Virtual domains** has been defined in `docker-compose.yml` file and configured in `/etc/hosts` file. Add the following line in your `/etc/hosts` file:

127.0.0.1   movies.local approvereviews.local

Command for Starting services
=============================

docker-compose up -d

Command for Stoping services
============================
docker-compose stop


Documentation
============================
(A) Movies management services

   **Create movie:**
   ```
   curl -H "Content-Type: application/json" -X POST -d '{"data":{"title":"myawsomemovie","director":"Ken Block","actors":["John Snow","ChristinaJake"],"rating": 3.55,"createdAt": "2018-04-15T08:02:06.029Z"}}' http://movies.local/movies
   ```
   ```
   curl -H "Content-Type: application/json" -X POST -d '{"data":{"title":"myawsomemovie2","director":"Ken Block","actors":["John Snow","ChristinaJake"],"rating": 8.55,"createdAt": "2018-04-15T08:02:06.029Z"}}' http://movies.local/movies
  ```
   ```
   curl -H "Content-Type: application/json" -X POST -d '{"data":{"title":"myawsomemovie3","director":"Ken Block","actors":["John Snow","ChristinaJake"],"rating": 9.55,"createdAt": "2018-04-15T08:02:06.029Z"}}' http://movies.local/movies
   ```
   
  **Get movie:**
  ALL
  ```
  curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://movies.local/movies
  ```
  By ID
  ```
  curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://movies.local/movies/5ad68743ad8f9c00075dab57
  ```
  
  **Delete movie:**
  ```
  curl -X DELETE http://movies.local/movies/5ad6a6641c419b0009c5ff38
  ```
  
**Add review to movie:**
```
curl -X PUT -H "Content-Type: application/json" -d '{"data":{"title":"myawsomemovie","director":"Ken Block","actors":["John Snow2","ChristinaJake"],"rating": 3.55,"createdAt": "2018-04-15T08:02:06.029Z","review" : "super cool movie"}}' http://movies.local/movies/addreview/5ad6a86d1c419b0009c5ff39
```


**Simple review apprvoal service:**
```
curl -H "Content-Type: application/json" -X POST -d '{"id": "5ad68743ad8f9c00075dab57","title": "MyAwesomeMovie22","rating": 2.55,"director": "Ken Block","actors": ["John Snow2","Christina Jake"],"createdAt": "2018-04-17T18:05:33.631Z","review" : "super cool review2"}' http://approvereviews.local/approvereview
```


