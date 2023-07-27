FROM docker.io/library/golang:1.20.3-alpine
RUN mkdir -p /opt/go-microservice 
RUN mkdir -p /var/go-microservice/data
WORKDIR /opt/go-microservice
COPY --chmod=755 go-microservice /opt/go-microservice/
COPY --chmod=644 go-microservice.yml /opt/go-microservice/
ENTRYPOINT ["/opt/go-microservice/go-microservice"]
CMD [ "service", "cloudservice" ]
