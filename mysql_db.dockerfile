FROM mysql:8.0.32

LABEL author="Bruno"
LABEL description="A database for this demo app"
LABEL version="1.0"

COPY db/mysql/*.sql /docker-entrypoint-initdb.d/
