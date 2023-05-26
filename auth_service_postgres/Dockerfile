FROM postgres:latest

# # Copy certificates
COPY ./certificates /var/lib/postgresql/certificates
# # Copy the entrypoint script
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
