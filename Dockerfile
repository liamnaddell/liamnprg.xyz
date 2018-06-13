FROM golang:1.10.3-alpine3.7


WORKDIR /go/src/app
COPY . .

RUN apk update && apk upgrade && apk add sassc git
RUN sh -c 'PREFIX=/go/src/app/ ./sass.sh'
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80
EXPOSE 443 

CMD ["/go/bin/app"]
