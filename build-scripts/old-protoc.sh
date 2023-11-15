#!/bin/sh


mkdir -p generated/password-generator

# go_opt=module stuff is to ensure it's generated in the folder i want.
# 	if this gets fucked up affects importing the module after pushing to github.
# 	idk wtf am doowin tho
#
protoc --go_out=generated/password-generator \
	--go_opt=module=github.com/solomonbroadbent/password.solthe.dev/project/generated/password-generator \
	--go-grpc_out=generated/password-generator \
	--go-grpc_opt=module=github.com/solomonbroadbent/password.solthe.dev/project/generated/password-generator \
	protos/password-generator/password-generator.proto
