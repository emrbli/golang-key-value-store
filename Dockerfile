FROM golang:1.17

WORKDIR /app

COPY go.mod ./
 #COPY go.sum ./

RUN go mod download

COPY . . 
 
RUN go build -o /key-value-store

EXPOSE 8080

CMD ["/key-value-store"]

