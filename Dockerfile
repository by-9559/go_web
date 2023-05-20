# FROM  golang
FROM  ubuntu:latest
WORKDIR /go_web
COPY ./main .
COPY ./templates ./templates
RUN apt update && apt install -y ca-certificates
RUN chmod 777 main
EXPOSE 8081
EXPOSE 8082
ENTRYPOINT [ "./main" ]