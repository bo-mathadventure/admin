#!/bin/sh

echo "Running database migration"

if [ -f ".env" ]; then
    set -o allexport
    ENVFILE=$(cat .env | tr -d '\r')
    echo $ENVFILE > ./.env.tmp
    source ./.env.tmp
    rm ./.env.tmp
    set +o allexport
fi

if ! [ -x "$(command -v nc)" ]; then
    echo "Error: nc is not installed. Skipping db connection test."
else
    until nc -z -v -w30 ${DB_HOST} ${DB_PORT:-3306}
    do
        echo "Waiting for database connection..."
        sleep 5
    done
fi


MIGRATION_PATH="ent/migrate/migrations"
if [ -n "$IS_DEPLOYMENT" ]; then
    MIGRATION_PATH="/app/migrations"
fi

atlas migrate apply --dir "file://${MIGRATION_PATH}" --url ${DB_TYPE:-"mysql"}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT:-3306}/${DB_NAME}
echo "Database migrated"