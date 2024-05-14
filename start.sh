#!/bin/sh


#exit in error
set -e

source /app/app.env

echo "Waiting for 10 seconds before running the database migration..."
sleep 10

echo "run db migration"

/app/migrate -path /app/migration -database "$DB_SOURCE" --verbose up

echo "start the app"
exec "$@"
