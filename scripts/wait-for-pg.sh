#!/bin/bash
# wait-for-postgres.sh

set -e

host="$1"
shift
port="$1"
shift
db="$1"
shift
user="$1"
shift
password="$1"
shift
cmd="$@"

until PGPASSWORD="$password" psql -h "$host" -p "$port" -U "$user" "$db" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
bash -c "$cmd"