FROM golang:latest as builder

RUN mkdir -p /home/go/app/bin

WORKDIR /home/go/app

COPY . .
RUN go env -w GOPROXY="https://goproxy.cn" && \
    go get 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app

FROM alpine:latest

EXPOSE 80

RUN mkdir -p /home/go/app

WORKDIR /home/go/app

COPY --from=builder /home/go/app/bin /home/go/app

# ENV ADDR defualt 
# ENV MONGO defualt1 
ENTRYPOINT ["./app"]
CMD ["-l ${ARRD:-test}", "-m ${MONGO:-test2}"]