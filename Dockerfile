FROM centos:7

COPY GO-redis /app/

RUN chmod 777 /app/GO-redis

WORKDIR /app

ENTRYPOINT ["./GO-redis"]