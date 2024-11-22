# Utiliza una imagen base de Go
FROM golang:1.23-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Compila el binario
RUN go build -o frontend-microservice .

# Define el puerto en el que la aplicación escuchará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./frontend-microservice"]
