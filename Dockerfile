FROM golang
WORKDIR /go_web
COPY . .
RUN go build main.go
RUN chmod 777 main
EXPOSE 8081
ENTRYPOINT [ "./main" ]