FROM golang:alpine

ARG MIGRATION_DIR_ARG
ARG MIGRATION_TARGET_DIR_ARG="/opt/migrations"

ENV MIGRATION_DIR=$MIGRATION_DIR_ARG
ENV MIGRATION_TARGET_DIR=$MIGRATION_TARGET_DIR_ARG

WORKDIR /opt

RUN apk add --no-cache git gcc g++ postgresql-client bash pkgconfig && \
    go get -v github.com/rubenv/sql-migrate/...

COPY ${MIGRATION_DIR} ${MIGRATION_TARGET_DIR}

# ADD dbconfig.prod.yml /opt/dbconfig.yml
