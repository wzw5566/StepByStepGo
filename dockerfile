FROM golang:latest

MAINTAINER vincent 

WORKDIR $GOPATH/src/hellodocker
ADD . $GOPATH//users/vincentwen/MyCode/go/StepByStepGo/src/dockerfile/hellodocker
RUN go build src/dockerfile/hellodocker/hellodocker.go

EXPOSE 8080

ENTRYPOINT ["./hellodocker"]
