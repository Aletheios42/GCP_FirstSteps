#FROM adminer:latest
#EXPOSE 8080
FROM wordpress:latest
ENV WORDPRESS_DB_HOST "10.73.0.3"
ENV WORDPRESS_DB_USER "user"
ENV WORDPRESS_DB_PASSWORD "password"
ENV WORDPRESS_DB_NAME "my-database"
EXPOSE 80