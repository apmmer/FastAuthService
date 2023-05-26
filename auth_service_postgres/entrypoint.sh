#!/bin/bash
# set -e

# Append SSL config to the postgresql.conf file
# echo "ssl = on" >> "$PGDATA/postgresql.conf"
# echo "ssl_cert_file = '/var/lib/postgresql/server.crt'" >> "$PGDATA/postgresql.conf"
# echo "ssl_key_file = '/var/lib/postgresql/server.key'" >> "$PGDATA/postgresql.conf"


# exec docker-entrypoint.sh postgres

set -e
# Копирование сертификата и ключа из директории certificates
echo "ENTRYPOINT: Copying certificates..."
cp /var/lib/postgresql/certificates/server.crt /var/lib/postgresql/server.crt
cp /var/lib/postgresql/certificates/server.key /var/lib/postgresql/server.key

# Установите права доступа для файла ключа сертификата
echo "ENTRYPOINT: Setting permissions..."
chmod 600 /var/lib/postgresql/server.key
chmod 644 /var/lib/postgresql/server.crt
echo "ENTRYPOINT: Setting user-postgres..."
chown postgres:postgres /var/lib/postgresql/server.key
chown postgres:postgres /var/lib/postgresql/server.crt
echo "ENTRYPOINT: Executing command..."
exec docker-entrypoint.sh postgres
