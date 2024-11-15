# API for Cosmetics Shop

Esta es una API desarrollada para gestionar una tienda de cosméticos. Utiliza el framework web **Gin** y el ORM **GORM** para manejar la base de datos. Proporciona funcionalidad para gestionar las siguientes entidades:

- **Buyer** (Compradores)
- **Customer** (Clientes)
- **Product** (Productos)
- **Invoice** (Facturas)

Los endpoints están documentados en una colección de **Postman** adjunta.

---

## Tecnologías utilizadas

- **Go**: Lenguaje de programación principal.
- **Gin**: Framework web utilizado para manejar las rutas y las solicitudes HTTP.
- **GORM**: ORM utilizado para gestionar la persistencia de datos en la base de datos.

---

## Requisitos previos

1. Tener instalado Go en tu sistema (versión 1.18 o superior).
2. Instalar las dependencias necesarias mediante `go mod`.

---

## Configuración del proyecto

### Clonar el repositorio
```bash
git clone <URL_DEL_REPOSITORIO>
cd cosmetics-shop-api
```

### Instalar dependencias
```bash
go mod tidy
```

### Configurar la base de datos

El proyecto está configurado para usar SQLite por defecto, funcionando en memoria de momento. Puedes cambiar la configuración de la base de datos editando el archivo correspondiente en el código fuente (función `InitDB`).

### Ejecutar la aplicación
```bash
go run main.go
```

La API estará disponible en [http://localhost:8080](http://localhost:8080).

---

## Endpoints disponibles

Las rutas están agrupadas bajo el prefijo `/api/v1`. A continuación, algunos ejemplos de los endpoints:

### Buyers (Compradores)
- **GET** `/api/v1/buyers`: Obtener todos los compradores.
- **POST** `/api/v1/buyers`: Crear un nuevo comprador.
- **PUT** `/api/v1/buyers/:id`: Actualizar un comprador por ID.
- **DELETE** `/api/v1/buyers/:id`: Eliminar un comprador por ID.

### Customers (Clientes)
- **GET** `/api/v1/customers`: Obtener todos los clientes.
- **POST** `/api/v1/customers`: Crear un nuevo cliente.

### Products (Productos)
- **GET** `/api/v1/products`: Obtener todos los productos.
- **POST** `/api/v1/products`: Crear un nuevo producto.

### Invoices (Facturas)
- **GET** `/api/v1/invoices`: Obtener todas las facturas.
- **POST** `/api/v1/invoices`: Crear una nueva factura.

---

## Colección de Postman

La colección de Postman para probar los endpoints está incluida en este repositorio. Puedes importarla en Postman para realizar pruebas rápidamente.

1. Abre Postman.
2. Haz clic en **Importar**.
3. Selecciona el archivo `postman_collection.json` incluido en este repositorio.

---

## Contribuciones

Las contribuciones son bienvenidas. Si deseas colaborar, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una rama para tu funcionalidad o corrección de errores.
3. Envía un pull request.

---

## Autor

Camilo Alejandro Marin Bustamante

