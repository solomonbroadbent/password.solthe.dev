#!/bin/sh

doctl compute ssh-key import \
	--public-key-file 


doctl compute droplet create \
	--image debian-12-x64 \
	--region syd1 \
	--size s-1vcpu-512mb-10gb \
	--ssh-keys 39984891 \
	password.solthe.dev


doctl compute ssh password.solthe.dev
