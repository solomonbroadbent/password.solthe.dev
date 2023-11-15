#!/bin/sh

# can run script without dl:
# 	sh -c "$(curl -fsSL https://raw.githubusercontent.com/solomonbroadbent/password.solthe.dev/production/build-scripts/debian-start-image.sh)"
docker pull solomonbroadbent/password.solthe.dev:api-latest

docker container create --name password.solthe.dev-api -p 80:80 solomonbroadbent/password.solthe.dev:api-latest
docker container start password.solthe.dev-api
