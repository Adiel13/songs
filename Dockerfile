# Utiliza la imagen oficial de golang como base
FROM golang:latest

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum al directorio de trabajo dentro del contenedor
COPY go.mod go.sum ./

# Instala las dependencias del módulo
RUN go mod download

# Copia el resto de los archivos de la aplicación
COPY . .

# Establece las directivas de construcción necesarias y compila la aplicación
RUN GO111MODULE=on && GOOS=linux GOARCH=amd64
RUN go build -o main .

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta la aplicación al iniciar el contenedor
CMD ["/app/main"]
