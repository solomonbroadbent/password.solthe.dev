#!/bin/sh

# CAN RUN THIS COMMAND on machine to use without pulling repo:
# 	sh -c "$(curl -fsSL https://raw.githubusercontent.com/solomonbroadbent/password.solthe.dev/production/build-scripts/debian-install-docker.sh)"

# Add Docker GPG key
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list

# Update the package list again
apt update

# Install Docker
apt install -y docker-ce docker-ce-cli containerd.io

# Add your user to the docker group to run Docker without sudo
usermod -aG docker $USER

echo "Docker has been successfully installed. Please log out and log back in for group changes to take effect."
