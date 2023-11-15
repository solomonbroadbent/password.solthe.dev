- copy and paste button
- word list
- remove similarly spelt words
- ax certain characters
- choose length

- add to bitwarden
- setup server
- grpc client



# COMMANDS
protoc -I=protos --go_out=paths=source_relative:project/generated/password-generator --go-grpc_out=paths=source_relative:project/generated/password-generator project/protos/password-generator.proto
protoc --go_out=../../generated/password-generator --go_opt=paths=source_relative --go-grpc_out=../../generated/password-generator --go-grpc_opt=paths=source_relative protos/password-generator/password-generator.proto
