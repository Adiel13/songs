#!/bin/bash

# Clone del repo
git clone https://github.com/Adiel13/songs

# activar directivas necesarias
cd songs
export GO111MODULE=on
GOOS=linux GOARCH=amd64 go build -o main .

# red de contenedores 
docker network create songs_network

# base de atos
docker run -d \
  --name db_songs \
  --network songs_network \
  -e MYSQL_ROOT_PASSWORD=songs \
  -p 3306:3306 \
  mysql:latest

sleep 10

# cargar cript de base de datos
cd ..
docker cp script.sql db_songs:/script.sql

docker exec -i db_songs mysql -uroot -psongs < script.sql

cd songs

# contenedor de songs go
docker run -d \
  --name go_app \
  --network songs_network \
  -p 8080:8080 \
  -v $(pwd):/app \
  golang:latest \
  /app/main

#contenedor de nginx
cd ..
docker run -d \
  --name nginx_container \
  --network songs_network \
  -p 80:80 \
  -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf \
  nginx:latest
