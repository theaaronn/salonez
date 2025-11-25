# Salonez - Sistema de Gestión de Salones

## Instalación y Configuración

### 1. Clonar e instalar dependencias

```bash
git clone https://github.com/TheAaronn/Salonez.git
cd Salonez
go mod tidy
```

### 2. Configurar base de datos PostgreSQL

Crear archivo `.env` en la raíz:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=salonez
```

### 3. Crear base de datos y cargar schema

```bash
# Crear base de datos
psql -U postgres -c "CREATE DATABASE salonez;"

# Cargar schema
psql -U postgres -d salonez -f db/schema.sql

# Cargar datos de demo
psql -U postgres -d salonez -f db/demo_data.sql
```

### 4. Generar archivos Templ

```bash
templ generate
```

### 5. Ejecutar la aplicación

```bash
# Modo desarrollo con Air (hot reload)
air

# O directamente con Go
go run cmd/launchServer.go
```

La aplicación estará disponible en `http://127.0.0.1:6969`

## Usuarios de Demo

Para probar la aplicación, usa estas credenciales:

- **Admin**: email: `admin`, password: cualquiera
- **Propietario**: email: `owner`, password: cualquiera  
- **Usuario**: email: `user`, password: cualquiera

## Funcionalidades por Rol

### Admin (admin)
- Ver todas las reservaciones del sistema
- Estado de cada reservación (activa/cancelada)

### Propietario (owner)
- Crear nuevos salones
- Ver sus propios salones
- Ver reservaciones de sus salones

### Usuario (user)
- Ver salones disponibles
- Crear reservaciones
- Ver sus reservaciones
- Cancelar reservaciones

## Estructura del Proyecto

```
salonez/
├── cmd/              # Punto de entrada
├── db/               # Capa de base de datos
├── handlers/         # Controladores HTTP
├── models/           # Modelos de datos
├── static/           # Archivos estáticos
│   └── views/       # Plantillas Templ
└── utils/           # Utilidades
```

## Comandos útiles

```bash
# Tests
go test ./models -v

# Compilar CSS con Tailwind
tailwindcss -i "./static/views/css/app.css" -o "./public/styles.css"

# Compilar binario
go build -o salonez cmd/launchServer.go
```

## Tecnologías

- **Backend**: Go 1.23
- **Framework**: Echo v4
- **Templates**: Templ
- **Base de datos**: PostgreSQL
- **Estilos**: Tailwind CSS
