#!/bin/bash
yum update -y
hostnamectl set-hostname gallery-dev-server
amazon-linux-extras install docker -y
systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user
curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" \
-o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
yum install git -y
yum install wget -y
cd /home/ec2-user
wget -c https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
tar -C /usr/local/ -xzf go1.22.0.linux-amd64.tar.gz
echo 'export GOROOT=/usr/local/go' >> /home/ec2-user/.bashrc
echo 'export PATH=$GOROOT/bin:$PATH' >> /home/ec2-user/.bashrc
git clone https://github.com/muskrt/gallery-microservice.git
cd gallery-microservice
git checkout dev
git config --global --add safe.directory /home/ec2-user/gallery-microservice
