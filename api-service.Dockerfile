FROM alpine:3.21 AS api-service

WORKDIR /service

COPY ./bin/api-service ./

EXPOSE 8081

CMD ["/service/api-service"]