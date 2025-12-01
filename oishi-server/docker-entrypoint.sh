#!/bin/sh
set -e

echo "Waiting for MySQL to be ready..."

# Wait for MySQL to be ready
max_retries=30
retry_count=0

until nc -z mysqldb 3306 || [ $retry_count -eq $max_retries ]; do
  retry_count=$((retry_count + 1))
  echo "MySQL is unavailable - sleeping (attempt $retry_count/$max_retries)"
  sleep 2
done

if [ $retry_count -eq $max_retries ]; then
  echo "Failed to connect to MySQL after $max_retries attempts"
  exit 1
fi

echo "MySQL is up - starting application"
exec "$@"
