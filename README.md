# Proyecto en Go - API Extractor de stock con inventario disponible para almacenamiento de datos.

## Descripción

Este proyecto es una API desarrollada en Go que guarda el token de vercel y permite la gestión de productos mediante una conexión a MongoDB. Proporciona los siguientes endpoints:

- **Autenticación**: Recibe un token de autenticación desde Vercel.
- **Gestión de Productos**: Obtiene una lista de productos almacenados en MongoDB.
- **Proxy**: Redirecciona solicitudes a una API externa alojada en Vercel.

## Tecnologías utilizadas

- **Go**
- **Fiber** (Framework web)
- **MongoDB** (Base de datos NoSQL)
- **Vercel** (Autenticación y servicios externos)

## Instalación y Configuración

### Prerrequisitos

1. Tener instalado Go en su versión más reciente.
2. Tener acceso a una instancia de MongoDB.
3. Contar con una cuenta en Vercel y un token de autenticación válido.

### Pasos

1. Clonar el repositorio:

    ```bash
    git clone https://github.com/usuario/proyecto-go-api.git
    cd proyecto-go-api
    ```

2. Configurar variables de entorno:

    Crear un archivo `.env` con las siguientes variables:

    ```env
    DB_CONEXION=mongodb+srv://usuario:password@cluster.mongodb.net/nombreBD
    DB_NAME=WMS
    COLLECCION=stock_available
    URL_EXTERNA=https://sil-recibo-despacho-nodejs-git-desarrollo-johancuervos-projects.vercel.app
    ```

3. Instalar dependencias:

    ```bash
    go mod tidy
    ```

4. Ejecutar el servidor:

    ```bash
    go run main.go
    ```

## Endpoints

### Autenticación

**POST** `/auth/vercel`  
Autentica la API con Vercel mediante un token.

- **Request**:

  ```json
  {
     "token": "tu_token_vercel"
  }
  ```

- **Response**:

  ```json
  {
     "message": "Autenticación exitosa",
     "status": 200
  }
  ```

### Productos

**GET** `/products`  
Obtiene la lista de productos almacenados en MongoDB.

- **Response**:

  ```json
  [
     {
        "id": "1",
        "nombre": "Producto A",
        "stock": 100
     },
     {
        "id": "2",
        "nombre": "Producto B",
        "stock": 50
     }
  ]
  ```

### Proxy

**GET** `/proxy/vercel/api/Productos/Stock`  
Redirecciona una solicitud a la API externa de Vercel para obtener información sobre stock de productos.

- **Response**:

  ```json
  {
     "id": "1",
     "nombre": "Producto A",
     "stock": 100
  }
  ```

## Despliegue

Para desplegar en un servidor o en Vercel, seguir la [documentación oficial de despliegue de aplicaciones en Go](https://vercel.com/docs).

## Desplegado

https://apiservicego-production.up.railway.app/swagger/index.html


## Contribuciones

Si deseas contribuir, crea un fork del repositorio y realiza un pull request con tus cambios.

## Licencia

Este proyecto se distribuye bajo la licencia [MIT](LICENSE).