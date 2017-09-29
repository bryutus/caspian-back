# Dockerfile_Go
FROM golang:1.9

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y cron && \
    apt-get install -y supervisor && \
    apt-get install -y vim

RUN go get -u github.com/labstack/echo && \
    go get -u github.com/dgrijalva/jwt-go && \
    go get -u github.com/BurntSushi/toml && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/go-sql-driver/mysql

EXPOSE 1323
