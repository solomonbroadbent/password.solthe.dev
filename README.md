- copy and paste button
- word list
- remove similarly spelt words
- ax certain characters
- choose length

- add to bitwarden
- setup server
- grpc client



# COMMANDS
test the api...
grpcurl -plaintext localhost:50051 password_generator.PasswordGenerator/GeneratePassword

DOCKER
docker image build --tag "solomonbroadbent/password.solthe.dev:api-latest" .
docker container create --name api.password.solthe.dev -p 80:80 "solomonbroadbent/password.solthe.dev:api-latest"
docker container start api.password.solthe.dev

docker container run -it --entrypoint /bin/sh "solomonbroadbent/password.solthe.dev:api-latest"
