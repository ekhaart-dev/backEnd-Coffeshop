FROM golang:1.20.6-alpine

WORKDIR /goapp

COPY . .

RUN go mod download

RUN go build -v -o /goapp/backEnd_Coffeshop ./cmd/main.go

EXPOSE 8081

ENTRYPOINT [ "/goapp/backEnd_Coffeshop" ]