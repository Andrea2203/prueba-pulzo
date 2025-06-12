# Prueba Pulzo

Este es un proyecto de consta de un servicio en Go que permita la autenticación y el consumo de datos externos.

## Estructura del Proyecto

```
prueba-pulzo/
├── auth/              
│   ├── handlers/           
│   │   ├── token.go
│   ├── models/           
│   │   ├── token.go
│   ├── main.go
│   ├── Dockerfile
├── data/
│   ├── handlers/           
│   │   ├── data.go
│   ├── main.go
│   ├── Dockerfile
├── .env.example       
├── .gitignore         
```

## Tecnologías Utilizadas

- **Go** - Lenguaje de programación principal
- **Docker** - Containerización
- **Git** - Control de versiones

## Instalación y Configuración

### Prerrequisitos

- Go 1.21 o superior
- Docker y Docker Compose
- Git

### Instalación Local

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/Andrea2203/prueba-pulzo.git
   cd prueba-pulzo
   ```

2. **Configurar variables de entorno**
   ```bash
   cp .env.example .env
   # Editar el archivo .env con tus configuraciones
   ```
### Ejecución local de AUTH
1. **Ingresar a la carpeta**
   ```bash
   cd auth
   ```
2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```
   
3. **Instalar dependencias**
   ```bash
   go mod tidy
   ```
   ### Ejecución local de Data
1. **Ingresar a la carpeta**
   ```bash
   cd data
   ```
2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```
   
### Instalación con Docker
1. **Ejecutar en la raiz del proyecto**
   ```bash
   docker compose up --build
   ```

## API Endpoints

### Autenticación
**POST** `/create-token` - Crear Token
#### Respuesta correcta:
```json
{
    "value": "314c7f14-9c0c-486b-9afd-9a73dcc9d874",
    "uses": 5,
    "is_valid": true
}
```
![image](https://github.com/user-attachments/assets/06c864a4-278d-4428-9e73-b07e6dc29971)

### Traer información
**GET** `/use-api?token={{token}}&api={{api}}` - Obtener data con validación de Token
```bach
api: 
-characters
-locations
-episodes

```
#### Respuesta correcta:
```json
{
    "data": {
        "info": {
            "count": 51,
            "pages": 3,
            "next": "https://rickandmortyapi.com/api/episode?page=2",
            "prev": null
        },
        "results": [
            {
                "id": 1,
                "name": "Pilot",
                "air_date": "December 2, 2013",
                "episode": "S01E01",
                "characters": [
                    "https://rickandmortyapi.com/api/character/1",
                ],
                "url": "https://rickandmortyapi.com/api/episode/1",
                "created": "2017-11-10T12:56:33.798Z"
            },
        ],
    },
    "token": {
        "value": "158317e5-c372-4ea2-9c2f-245124e13555",
        "uses": 3,
        "is_valid": true
    }
}
```
![image](https://github.com/user-attachments/assets/45db4a16-8869-4b10-b959-5abde44b6c84)

**GET** `/get-data?api={{api}}` - Obtener data con validación de Token
```bach
api: 
-characters
-locations
-episodes

```
#### Respuesta correcta:
```json
{
  "info": {
      "count": 51,
      "pages": 3,
      "next": "https://rickandmortyapi.com/api/episode?page=2",
      "prev": null
  },
  "results": [
      {
          "id": 1,
          "name": "Pilot",
          "air_date": "December 2, 2013",
          "episode": "S01E01",
          "characters": [
              "https://rickandmortyapi.com/api/character/1",
          ],
          "url": "https://rickandmortyapi.com/api/episode/1",
          "created": "2017-11-10T12:56:33.798Z"
      },
  ],
}
```
![image](https://github.com/user-attachments/assets/5b4e9a6f-2811-4f19-be8c-e56ec43d5f65)

Puedes importar la colección de Postman desde [Pulzo-prueba.postman_collection.json](Pulzo-prueba.postman_collection.json).


