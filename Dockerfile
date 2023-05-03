FROM golang:1.18-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git
RUN go mod download
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]