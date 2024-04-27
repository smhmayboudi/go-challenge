#!/bin/sh

set -Eefuvx

POSTGRES_HOST=${1:-127.0.0.1}
POSTGRES_PORT=${2:-5432}

POSTGRES_DB=${3:-gorm}
POSTGRES_PASSWORD=${4:-gorm}
POSTGRES_USER=${5:-gorm}

DB_NAME=${6:-server-arium}
VERSION=${7:-development}

DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}"
SOURCE=file://db/${DB_NAME}/${VERSION}

migrate -database=${DSN} -source=${SOURCE} -verbose up
