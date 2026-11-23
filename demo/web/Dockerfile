FROM nginx:1.31.9-alpine3.24
# where /api requests are forwarded. The API is published on the host's port 8080, so: localhost:8080
ENV API_URL=http://localhost:8080
# the official image fills ${API_URL} in this template when the container starts
COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY index.html og.svg /usr/share/nginx/html/
