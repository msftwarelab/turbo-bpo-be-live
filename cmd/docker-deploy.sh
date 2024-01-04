#!/bin/bash
docker stop  turbo-bpo-backend 
docker rmi  turbo-bpo-backend 
docker rm turbo-bpo-backend
docker container prune -f
docker system prune --volumes -f


#echo -e "creating build"
#cd /cmd
#go build -o main .
#cd ..
#echo -e "creating build success"
env GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ./main
echo -e "Building docker image.."
docker build -t pamisa56/turbo-bpo-backend .
echo -e "Docker image was created successfully"
docker run -p 6969:6969 --name turbo-bpo-backend -d pamisa56/turbo-bpo-backend
echo -e "Docker turbo-bpo-backend container is running..."