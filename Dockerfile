FROM golang:1.19.3-alpine

WORKDIR /usr/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v

CMD ["app"]
