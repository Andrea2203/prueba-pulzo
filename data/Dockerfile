# Usar una imagen base de Go
FROM golang:1.21-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente al contenedor
COPY . .

# Descargar las dependencias
RUN go mod tidy

# Compilar la aplicación
RUN go build -o app .

# Exponer el puerto que usará la aplicación
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./app"]
