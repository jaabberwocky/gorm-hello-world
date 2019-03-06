FROM golang

WORKDIR /go/src/gorm-hello-world
COPY . .

RUN go install

ENTRYPOINT /go/bin/gorm-hello-world

EXPOSE 4531
