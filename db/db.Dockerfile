FROM mysql

ENV MYSQL_DATABASE test 

COPY ./db/mysql-scripts /docker-entrypoint-initdb.d/

EXPOSE 3306