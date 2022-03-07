FROM postgres:12-alpine as db
WORKDIR /app
COPY ./init.sh /docker-entrypoint-initdb.d
COPY ./grey.sql ./db/grey.sql