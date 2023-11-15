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

doctl compute firewall create --name password.solthe.dev \
	--inbound-rules "protocol:tcp,ports:80,address:0.0.0.0/0"

doctl compute firewall add-droplets a7603bd2-1f23-44ce-a29f-2b14d5b6d9fa --droplet-ids 385124377

doctl compute domain records create solthe.dev --record-type A --record-name password --record-data 170.64.158.163 --record-ttl 30

