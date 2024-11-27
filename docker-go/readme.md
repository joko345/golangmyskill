go mod init github.com/joko345/golangmyskill
go mod tidy
docker build -t "docker-go" .
docker run -p 8080:8081 -it docker-go

FROM golang:alpine //versi golang dan provider gambar
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]
//docker menjalankan perintah layaknya cmd ubuntu
docker build -t "nama folder project" .
