FROM postgres:alpine
ADD ./database_schema.sql /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*