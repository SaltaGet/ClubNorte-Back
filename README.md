# Backend del Club Norte

Este es el repositorio del backend para la aplicación de gestión del Club Norte. La API está construida con Go y el framework Fiber, proporcionando una solución de alto rendimiento para gestionar las operaciones del club.

## ✨ Características Principales

- **Gestión de Usuarios y Roles**: Sistema completo para manejar usuarios, permisos y roles.
- **Autenticación Segura**: Implementación de autenticación basada en JWT (JSON Web Tokens) con cookies.
- **Gestión Financiera**: Endpoints para registrar ingresos, egresos, depósitos y movimientos de caja.
- **Control de Inventario**: API para gestionar productos, stock, categorías y movimientos de stock.
- **Punto de Venta (PDV)**: Lógica de negocio para operaciones en puntos de venta.
- **Gestión de Recursos**: Administración de canchas deportivas y otros recursos del club.
- **Documentación de API**: Documentación autogenerada y disponible con Swagger y Scalar.

## 🚀 Tecnologías Utilizadas

- **Lenguaje**: [Go](https://golang.org/) (Versión 1.21+)
- **Framework**: [Fiber](https://gofiber.io/) v2
- **Base de Datos**: [MariaDB](https://mariadb.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Contenerización**: [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/)
- **Documentación**: [Swagger](https://swagger.io/)
- **Variables de Entorno**: [godotenv](https://github.com/joho/godotenv)

## 📋 Requisitos Previos

- [Go](https://golang.org/doc/install) 1.21 o superior.
- [Docker](https://docs.docker.com/get-docker/) y [Docker Compose](https://docs.docker.com/compose/install/).
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

## ⚙️ Instalación y Ejecución

Se recomienda utilizar Docker para un despliegue rápido y consistente.

### 1. Con Docker (Método Recomendado)

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/tu-usuario/club-norte-back.git
    cd club-norte-back
    ```

2.  **Crea el archivo de variables de entorno:**
    Crea un archivo llamado `.env` en la raíz del proyecto. Puedes copiar el siguiente ejemplo y modificar los valores según tu configuración.

    ```env
    # Archivo .env de ejemplo
    
    # Base de Datos
    DB_HOST=db
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=root_password
    DB_NAME=club_norte
    MYSQL_ROOT_PASSWORD=root_password
    MYSQL_DATABASE=club_norte

    # Configuración de la API
    API_PORT=3000
    JWT_SECRET=tu_super_secreto_jwt
    
    # CORS
    ORIGIN=http://localhost:5173
    METHODS=GET,POST,PUT,DELETE
    HEADERS=Origin,Content-Type,Accept
    CREDENTIALS=true
    MAXAGE=300

    # New Relic (Opcional)
    NEW_RELIC_APP_NAME=club-norte-api
    NEW_RELIC_LICENSE_KEY=tu_licencia_de_new_relic
    ```

3.  **Levanta los servicios con Docker Compose:**
    Este comando construirá la imagen de la API (si no existe) y levantará los contenedores de la base de datos y la aplicación.

    ```bash
    docker-compose up --build
    ```

4.  **¡Listo!** La API estará corriendo en `http://localhost:3000`.

### 2. De forma local (Sin Docker)

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/tu-usuario/club-norte-back.git
    cd club-norte-back
    ```

2.  **Inicia una base de datos MariaDB:**
    Asegúrate de tener una instancia de MariaDB corriendo y accesible desde tu máquina.

3.  **Configura las variables de entorno:**
    Crea un archivo `.env` como en el paso de Docker, pero ajusta las variables de la base de datos (`DB_HOST`, `DB_PORT`, etc.) para que apunten a tu instancia local.

4.  **Instala las dependencias de Go:**
    ```bash
    go mod tidy
    ```

5.  **Ejecuta la aplicación:**
    ```bash
    go run cmd/api/main.go
    ```

6.  La API estará disponible en `http://localhost:3000`.

## 📚 Documentación de la API

Una vez que la aplicación está en funcionamiento, puedes acceder a la documentación de la API a través de las siguientes rutas:

-   **Swagger UI**: `http://localhost:3000/api/swagger/index.html`
-   **Scalar API Reference**: `http://localhost:3000/reference`

## 📂 Estructura del Proyecto

```
/
├── cmd/api/             # Punto de entrada de la API, configuración y rutas.
│   ├── controllers/     # Lógica de manejo de peticiones HTTP.
│   ├── middleware/      # Middlewares de Fiber.
│   └── routes/          # Definición de las rutas de la API.
├── internal/            # Lógica de negocio principal y acceso a datos.
│   ├── database/        # Configuración y conexión a la base de datos.
│   ├── models/          # Estructuras de datos (entidades de la BD).
│   ├── repositories/    # Lógica de acceso a datos (CRUD).
│   ├── schemas/         # Esquemas de validación y DTOs.
│   └── services/        # Lógica de negocio.
├── .dockerignore        # Archivos a ignorar por Docker.
├── .gitignore           # Archivos a ignorar por Git.
├── docker-compose.yml   # Orquestación de contenedores.
├── Dockerfile           # Definición de la imagen Docker de la API.
├── go.mod               # Dependencias del proyecto.
└── README.md            # Este archivo.
```
