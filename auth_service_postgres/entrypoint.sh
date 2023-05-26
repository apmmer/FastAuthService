#!/bin/bash
set -e
# Copying the certificate and key from the certificate directory
echo "ENTRYPOINT: Copying certificates..."
cp /var/lib/postgresql/certificates/server.crt /var/lib/postgresql/server.crt
cp /var/lib/postgresql/certificates/server.key /var/lib/postgresql/server.key

# Setting the access rights for the certificate and key file
echo "ENTRYPOINT: Setting permissions..."
chmod 600 /var/lib/postgresql/server.key
chmod 644 /var/lib/postgresql/server.crt
echo "ENTRYPOINT: Setting user-postgres..."
chown postgres:postgres /var/lib/postgresql/server.key
chown postgres:postgres /var/lib/postgresql/server.crt
echo "ENTRYPOINT: Executing command..."
exec docker-entrypoint.sh postgres
