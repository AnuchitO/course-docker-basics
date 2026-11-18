FROM nginx:1.27-alpine
# where /api requests are forwarded; override with: docker run -e API_URL=http://other:8080
ENV API_URL=http://api:8080
# the official image fills ${API_URL} in this template when the container starts
COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY index.html app.js og.svg /usr/share/nginx/html/
