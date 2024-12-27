FROM golang:1.19

WORKDIR /app
 
COPY . .

RUN go mod init maioridade-app 

RUN go mod tidy 

RUN go build -o main

CMD ["./main"]