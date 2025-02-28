FROM postgres:17.4 AS db

WORKDIR /app

COPY ./seed.sql /docker-entrypoint-initdb.d/
