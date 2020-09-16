FROM golang:1.14-buster

WORKDIR /go/src/benkyoukai-go

RUN apt-get update \
  && apt-get -y upgrade \
  && apt-get -y install memcached netcat

RUN curl -sL https://deb.nodesource.com/setup_12.x | bash - \
  && apt-get install -y nodejs \
  && npm i -g knex

RUN go get -u github.com/cosmtrek/air

# COPY ./ .

CMD go build main.go && ./main
