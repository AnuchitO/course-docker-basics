FROM nginx:1.31.9-alpine3.24
# where /api requests are forwarded. "api" is the network alias from docker-compose.yml, 8080 is the port inside the API container
ENV API_URL=http://api:8080
# the official image fills ${API_URL} in this template when the container starts
COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY index.html og.svg /usr/share/nginx/html/
