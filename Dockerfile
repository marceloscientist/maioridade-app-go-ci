FROM golang:1.23.2

WORKDIR /app
 
COPY . .
 
RUN go build -o main

CMD ["./main"]