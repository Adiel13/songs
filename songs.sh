#!/bin/bash

# Clonar el repositorio
git clone https://github.com/Adiel13/songs

# Moverse al directorio clonado
cd songs
export GO111MODULE=on
GOOS=linux GOARCH=amd64 go build -o main .

# Crear red bridge para los contenedores
docker network create songs_network

# Crear contenedor de MySQL
docker run -d \
  --name db_songs \
  --network songs_network \
  -e MYSQL_ROOT_PASSWORD=songs \
  -p 3306:3306 \
  mysql:latest

# Esperar unos segundos para que MySQL se inicie completamente
sleep 10

cd ..

# Copiar el script de creación de la base de datos y las tablas al contenedor MySQL
docker cp script.sql db_songs:/script.sql

# Ejecutar el script SQL dentro del contenedor MySQL
docker exec -i db_songs mysql -uroot -psongs < script.sql

cd songs

# Ejecutar la aplicación Go en un contenedor
docker run -d \
  --name go_app \
  --network songs_network \
  -p 8080:8080 \
  -v $(pwd):/app \
  golang:latest \
  /app/main

cd ..
# Configurar Nginx como proxy inverso
docker run -d \
  --name nginx_container \
  --network songs_network \
  -p 80:80 \
  -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf \
  nginx:latest

echo $(pwd)
