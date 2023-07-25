FROM golang:1.20.3-alpine
RUN makedir -p /opt/go-microservice && makedir -p /var/go-microservice/data
COPY ../go-microservice ../go-service.yml /opt/go-microservice/
WORKDIR /opt/go-microservice
ENTRYPOINT ["go-microservice"]
CMD [ "service", "cloudservice" ]
