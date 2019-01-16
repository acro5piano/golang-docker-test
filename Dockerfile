FROM alpine:3.8

RUN mkdir /app
WORKDIR /app
ADD ./golang-docker-test /app/

EXPOSE 8000
CMD ["/app/golang-docker-test"]
