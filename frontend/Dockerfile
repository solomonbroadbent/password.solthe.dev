FROM busybox:1.36.1

WORKDIR /app

COPY . .

EXPOSE 80

CMD ["httpd", "-f", "-p", "80", "-h", "/app"]
