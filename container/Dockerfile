FROM docker.io/library/golang:1.20.3-alpine
RUN mkdir -p /opt/go-microservice && mkdir -p /var/go-microservice/data $$ mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /opt/go-microservice
ENTRYPOINT ["/opt/go-microservice/go-microservice"]
CMD [ "service", "cloudservice" ]
