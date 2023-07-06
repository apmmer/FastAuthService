set -e

cp /var/lib/postgresql/certificates/server.crt /var/lib/postgresql/server.crt
cp /var/lib/postgresql/certificates/server.key /var/lib/postgresql/server.key

chmod 600 /var/lib/postgresql/server.key
chmod 644 /var/lib/postgresql/server.crt

chown postgres:postgres /var/lib/postgresql/server.key
chown postgres:postgres /var/lib/postgresql/server.crt

exec docker-entrypoint.sh postgres
