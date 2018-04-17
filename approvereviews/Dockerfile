# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/validator.v2
RUN go get encoding/json
RUN go get fmt
RUN go get log
RUN go get net/http

# copy the local package files to the container workspace
ADD . /go/src/github.com/mismaeel/moviesapp/approvereviews

# Setting up working directory
WORKDIR /go/src/github.com/mismaeel/moviesapp/approvereviews

# Build the movies command inside the container.
RUN go install github.com/mismaeel/moviesapp/approvereviews

# Run the movies microservice when the container starts.
ENTRYPOINT /go/bin/approvereviews

# Service listens on port 8080.
EXPOSE 8080
