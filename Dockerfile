FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/sfreiberg/gotwilio

WORKDIR /go/src/github.com/heaptracetechnology/microservice-twilio

ADD . /go/src/github.com/heaptracetechnology/microservice-twilio

RUN go install github.com/heaptracetechnology/microservice-twilio

ENTRYPOINT microservice-twilio

EXPOSE 3000