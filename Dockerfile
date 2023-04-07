FROM golang
WORKDIR /go_web
COPY . .
RUN go build main.go
RUN chmod 777 main
ENTRYPOINT [ "./main" ]