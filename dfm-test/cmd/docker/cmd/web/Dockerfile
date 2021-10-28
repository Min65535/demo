FROM nginx:1.15.2-alpine

RUN mkdir -p /var/www/html

COPY html /var/www/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]

