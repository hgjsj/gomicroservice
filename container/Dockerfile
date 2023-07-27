FROM docker.io/library/golang:1.20.3-alpine
RUN mkdir -p /opt/go-microservice && mkdir -p /var/go-microservice/data
COPY go-microservice go-service.yml /opt/go-microservice/
WORKDIR /opt/go-microservice
ENTRYPOINT ["/opt/go-microservice/go-microservice"]
CMD [ "service", "cloudservice" ]
