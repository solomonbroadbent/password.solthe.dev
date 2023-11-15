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
docker image build --tag password-generator:latest .
docker container create --name password-generator -p 80:80 password-generator:latest
docker container start password-generator

docker container run -it --entrypoint /bin/sh password-generator

