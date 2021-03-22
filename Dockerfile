FROM golang:alpine

ADD . /api
WORKDIR /api

RUN go get
RUN go build

CMD go get github.com/pilu/fresh && \
	fresh;
	
EXPOSE 8080

