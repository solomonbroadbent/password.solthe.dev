# Use the official NGINX base image
FROM nginx:1.25.3

# Copy your NGINX configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]

