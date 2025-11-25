# Documentación Completa - Salonez

**Versión:** 1.0.0  
**Fecha:** 24 de noviembre de 2025  
**Framework:** Go + Echo + Templ  
**Base de datos:** PostgreSQL

---

## Tabla de Contenidos

1. [Descripción General](#descripción-general)
2. [Arquitectura del Proyecto](#arquitectura-del-proyecto)
3. [Instalación y Configuración](#instalación-y-configuración)
4. [Estructura de Directorios](#estructura-de-directorios)
5. [Modelos de Datos](#modelos-de-datos)
6. [Base de Datos](#base-de-datos)
7. [Handlers y Rutas](#handlers-y-rutas)
8. [Vistas (Templ)](#vistas-templ)
9. [Utilidades](#utilidades)
10. [Testing](#testing)
11. [Despliegue](#despliegue)
12. [API Reference](#api-reference)

---

## Descripción General

**Salonez** es una aplicación web desarrollada en Go para la gestión y reservación de salones de eventos. Permite a los usuarios buscar y reservar salones, mientras que los propietarios pueden administrar sus espacios y los administradores tienen control total del sistema.

### Características principales

- ✅ Búsqueda y listado de salones disponibles
- ✅ Sistema de reservaciones con gestión de pagos parciales
- ✅ Panel de administración para propietarios
- ✅ Gestión de usuarios (Usuario, Propietario, Administrador)
- ✅ Galería de imágenes por salón
- ✅ Sistema de autenticación y autorización
- ✅ Gestión de cancelaciones

### Tecnologías utilizadas

- **Backend:** Go 1.23.0
- **Framework Web:** Echo v4.13.4
- **Motor de plantillas:** Templ v0.3.865
- **Base de datos:** PostgreSQL (con soporte para arrays nativos)
- **ORM/Driver:** go-sql-driver/mysql v1.9.2
- **Configuración:** godotenv v1.5.1
- **Hot reload:** Air (desarrollo)

---

## Arquitectura del Proyecto

### Patrón de diseño

El proyecto sigue una arquitectura MVC (Model-View-Controller) adaptada para Go:

```
┌─────────────┐
│   Cliente   │
└──────┬──────┘
       │ HTTP Request
       ▼
┌─────────────┐
│   Router    │ (Echo Framework)
│  (main.go)  │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Handlers   │ (Controllers)
└──────┬──────┘
       │
       ├──────────┐
       ▼          ▼
  ┌────────┐  ┌──────┐
  │ Models │  │ Views│ (Templ)
  └────┬───┘  └──────┘
       │
       ▼
  ┌────────┐
  │   DB   │ (PostgreSQL)
  └────────┘
```

### Flujo de datos

1. Cliente envía solicitud HTTP
2. Router (Echo) dirige a handler correspondiente
3. Handler procesa lógica de negocio usando models
4. Models interactúan con la base de datos
5. Handler renderiza vista con datos
6. Respuesta HTML enviada al cliente

---

## Instalación y Configuración

### Requisitos previos

- Go 1.23.0 o superior
- PostgreSQL 14+ (recomendado)
- Git
- Air (opcional, para desarrollo con hot reload)

### Instalación paso a paso

```bash
# 1. Clonar el repositorio
git clone https://github.com/TheAaronn/Salonez.git

# 2. Navegar al directorio
cd Salonez

# 3. Instalar dependencias
go mod tidy

# 4. Configurar variables de entorno
cp .env.example .env
# Editar .env con tus credenciales de base de datos

# 5. Crear base de datos
psql -U postgres -f db/schema.sql

# 6. Ejecutar la aplicación
go run cmd/launchServer.go
```

### Variables de entorno

Crea un archivo `.env` en la raíz del proyecto:

```env
# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=salonez

# Servidor
SERVER_HOST=127.0.0.1
SERVER_PORT=6969

# Sesiones
SESSION_SECRET=tu_secret_key_aqui
```

### Desarrollo con Air (Hot reload)

```bash
# Instalar Air
go install github.com/air-verse/air@latest

# Ejecutar con hot reload
air
```

---

## Estructura de Directorios

```
salonez/
├── cmd/                      # Puntos de entrada de la aplicación
│   └── launchServer.go      # Servidor principal
├── db/                       # Capa de base de datos
│   ├── getDb.go             # Conexión a BD
│   ├── hallList.go          # Queries de salones
│   ├── login.go             # Queries de autenticación
│   ├── schema.sql           # Schema de PostgreSQL
│   ├── functions/           # Funciones de BD
│   ├── sqlite/              # (Deprecado)
│   ├── tables/              # Definiciones de tablas
│   ├── triggers/            # Triggers de BD
│   └── views/               # Vistas de BD
├── docs/                     # Documentación
│   └── doc.md               # Este archivo
├── handlers/                 # Controladores HTTP
│   ├── hallList.go          # Handler de listado de salones
│   ├── index.go             # Handler de página principal
│   ├── login.go             # Handler de login
│   └── signup.go            # Handler de registro
├── models/                   # Modelos de datos
│   ├── dataModel.go         # Definiciones de estructuras
│   ├── dataModel_test.go    # Tests de modelos
│   └── modelFuncs.go        # Funciones de modelos
├── public/                   # Archivos públicos
│   └── styles.css           # Estilos CSS públicos
├── static/                   # Recursos estáticos
│   ├── assets/              # Recursos multimedia
│   │   ├── pallette         # Paleta de colores
│   │   ├── script.js        # JavaScript
│   │   └── styles.css       # Estilos CSS
│   ├── hallImgs/            # Imágenes de salones
│   └── views/               # Plantillas Templ
│       ├── layout.templ     # Layout base
│       ├── index.templ      # Página principal
│       ├── hallList.templ   # Listado de salones
│       ├── hall.templ       # Detalle de salón
│       ├── login-signup.templ  # Autenticación
│       ├── userHome.templ   # Dashboard usuario
│       ├── propietaryHome.templ  # Dashboard propietario
│       ├── adminHome.templ  # Dashboard admin
│       ├── 402.templ        # Error 402
│       ├── 500.templ        # Error 500
│       ├── loginErr.templ   # Error de login
│       └── css/
│           └── app.css      # Estilos de vistas
├── utils/                    # Utilidades
│   ├── imgPath.go           # Manejo de rutas de imágenes
│   └── render.go            # Renderizado de vistas
├── tmp/                      # Archivos temporales (ignorado)
├── .air.toml                # Configuración de Air
├── .gitignore               # Archivos ignorados por Git
├── go.mod                   # Dependencias de Go
├── go.sum                   # Checksums de dependencias
└── README.md                # Readme básico
```

---

## Modelos de Datos

### User

Representa un usuario en el sistema. Puede ser un usuario regular, propietario de salón, o administrador.

```go
type User struct {
    Nombre string  // Nombre completo del usuario
    Correo string  // Dirección de correo electrónico (debe ser única)
}
```

**Alias de tipos:**

- `Usuario`: Alias de User para usuarios regulares
- `Propietario`: Alias de User para propietarios de salones
- `Admin`: Alias de User para administradores del sistema

**Uso:**

```go
user := models.User{
    Nombre: "Juan Pérez",
    Correo: "juan@example.com",
}

propietario := models.Propietario{
    Nombre: "María García",
    Correo: "maria@salonexample.com",
}
```

### Hall

Representa un salón de eventos disponible para reservación.

```go
type Hall struct {
    Id             int32      // Identificador único del salón
    ImgsPath       []string   // Rutas a las imágenes del salón
    Nombre         string     // Nombre del salón
    Direccion      string     // Dirección completa del salón
    Capacidad      string     // Capacidad del salón (ej: "200 personas")
    NumeroTelefono string     // Número de contacto del salón
    Precio         float64    // Precio base de renta del salón
}
```

**Notas:**

- `ImgsPath` puede contener múltiples rutas de imágenes
- `Precio` debe ser un valor positivo o cero
- `Capacidad` se almacena como string para flexibilidad (ej: "100-150 personas")

**Uso:**

```go
hall := models.Hall{
    Id:             1,
    ImgsPath:       []string{"/uploads/hall1_1.jpg", "/uploads/hall1_2.jpg"},
    Nombre:         "Salón Principal",
    Direccion:      "Av. Insurgentes 123, Col. Centro",
    Capacidad:      "200 personas",
    NumeroTelefono: "555-1234-5678",
    Precio:         5000.00,
}
```

### Location

Representa una ubicación geográfica detallada.

```go
type Location struct {
    Calle   string  // Nombre de la calle
    Numero  string  // Número exterior/interior
    Colonia string  // Colonia o barrio
    CP      int     // Código postal
    Ciudad  string  // Ciudad
    Estado  string  // Estado o provincia
}
```

**Validaciones recomendadas:**

- `CP` debe ser un número positivo
- Todos los campos string no deben estar vacíos
- `CP` en México debe tener 5 dígitos

**Uso:**

```go
location := models.Location{
    Calle:   "Avenida Insurgentes",
    Numero:  "1234",
    Colonia: "Centro",
    CP:      06000,
    Ciudad:  "Ciudad de México",
    Estado:  "CDMX",
}
```

### Reservation

Representa una reservación de un salón.

```go
type Reservation struct {
    Date           string   // Fecha de la reservación (formato: YYYY-MM-DD)
    Time           string   // Hora de la reservación (formato: HH:MM)
    PercentagePaid float32  // Porcentaje pagado (0-100)
    Cancelled      bool     // Estado de cancelación
}
```

**Validaciones recomendadas:**

- `Date` debe ser una fecha válida en formato ISO (YYYY-MM-DD)
- `Time` debe ser una hora válida (HH:MM en formato 24h)
- `PercentagePaid` debe estar entre 0 y 100
- Las reservaciones canceladas (`Cancelled = true`) no deben ser modificables

**Uso:**

```go
reservation := models.Reservation{
    Date:           "2025-12-25",
    Time:           "18:00",
    PercentagePaid: 50.0,
    Cancelled:      false,
}
```

### Cancelation

Representa el registro de una cancelación de reservación.

```go
type Cancelation struct {
    Date string  // Fecha de la cancelación (formato: YYYY-MM-DD)
    Time string  // Hora de la cancelación (formato: HH:MM)
}
```

**Notas:**

- Se crea automáticamente cuando se cancela una reservación
- `Date` y `Time` registran cuándo se realizó la cancelación, no la fecha original de la reservación

**Uso:**

```go
cancelation := models.Cancelation{
    Date: "2025-11-24",
    Time: "14:30",
}
```

### Relaciones entre modelos

```
User (Propietario)
  └─── Hall (1:N)
         ├─── Location (1:1)
         └─── Reservation (1:N)
                ├─── User (N:1)
                └─── Cancelation (1:1 opcional)
```

### Consideraciones de diseño

1. **Tipos de usuario**: Se usan alias de tipos para mayor claridad semántica, pero todos comparten la misma estructura base `User`.

2. **Imágenes**: Se almacenan como rutas en lugar de datos binarios para mejor rendimiento y facilidad de gestión.

3. **Precios y porcentajes**: Se usan tipos numéricos de punto flotante. En producción, considera usar tipos decimales exactos para evitar problemas de precisión.

4. **Fechas y horas**: Se almacenan como strings. Considera usar `time.Time` en la lógica de negocio para mejor manejo de zonas horarias y validaciones.

---

## Base de Datos

### Schema PostgreSQL

El proyecto utiliza PostgreSQL con características avanzadas como arrays nativos, triggers y funciones.

#### Tablas principales

**users**

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) NOT NULL UNIQUE,
    tipo_usuario VARCHAR(20) DEFAULT 'usuario' CHECK (tipo_usuario IN ('usuario', 'propietario', 'admin')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**locations**

```sql
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    calle VARCHAR(255) NOT NULL,
    numero VARCHAR(50) NOT NULL,
    colonia VARCHAR(255) NOT NULL,
    cp INTEGER NOT NULL,
    ciudad VARCHAR(255) NOT NULL,
    estado VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**halls**

```sql
CREATE TABLE halls (
    id SERIAL PRIMARY KEY,
    imgs_path TEXT[] DEFAULT '{}',
    nombre VARCHAR(255) NOT NULL,
    direccion VARCHAR(500) NOT NULL,
    location_id INTEGER REFERENCES locations(id) ON DELETE SET NULL,
    capacidad VARCHAR(100) NOT NULL,
    numero_telefono VARCHAR(20),
    precio NUMERIC(10, 2) NOT NULL,
    propietario_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**reservations**

```sql
CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    hall_id INTEGER NOT NULL REFERENCES halls(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    time TIME NOT NULL,
    percentage_paid NUMERIC(5, 2) DEFAULT 0.00 CHECK (percentage_paid >= 0 AND percentage_paid <= 100),
    cancelled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**cancelations**

```sql
CREATE TABLE cancelations (
    id SERIAL PRIMARY KEY,
    reservation_id INTEGER NOT NULL REFERENCES reservations(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    time TIME NOT NULL,
    motivo TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Índices

```sql
-- Índices para mejorar rendimiento
CREATE INDEX idx_users_correo ON users(correo);
CREATE INDEX idx_users_tipo ON users(tipo_usuario);
CREATE INDEX idx_halls_propietario ON halls(propietario_id);
CREATE INDEX idx_halls_nombre ON halls(nombre);
CREATE INDEX idx_halls_imgs_path ON halls USING GIN(imgs_path);
CREATE INDEX idx_reservations_hall ON reservations(hall_id);
CREATE INDEX idx_reservations_user ON reservations(user_id);
CREATE INDEX idx_reservations_date ON reservations(date);
CREATE INDEX idx_reservations_cancelled ON reservations(cancelled);
```

### Triggers

```sql
-- Función para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggers para actualizar updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_locations_updated_at BEFORE UPDATE ON locations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_halls_updated_at BEFORE UPDATE ON halls
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_reservations_updated_at BEFORE UPDATE ON reservations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Operaciones con arrays en PostgreSQL

```sql
-- Insertar hall con imágenes
INSERT INTO halls (imgs_path, nombre, direccion, capacidad, precio)
VALUES (ARRAY['/img/hall1.jpg', '/img/hall2.jpg'], 'Salón Principal', 'Calle 123', '200 personas', 5000.00);

-- Agregar una imagen al array
UPDATE halls SET imgs_path = array_append(imgs_path, '/img/hall3.jpg') WHERE id = 1;

-- Eliminar una imagen del array
UPDATE halls SET imgs_path = array_remove(imgs_path, '/img/hall2.jpg') WHERE id = 1;

-- Buscar halls que contengan una imagen específica
SELECT * FROM halls WHERE '/img/hall1.jpg' = ANY(imgs_path);
```

---

## Handlers y Rutas

### Servidor principal (cmd/launchServer.go)

```go
package main

import (
    "Salonez/handlers"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.Static("/", "/static")

    // Rutas públicas
    e.GET("/", handlers.IndexHandler)
    e.GET("/halls", handlers.HallList)

    // Autenticación
    e.GET("/login", handlers.ShowLoginHandler)
    e.POST("/login", handlers.ValidateLogin)
    e.POST("/signup", handlers.SignUpHandler)

    e.Logger.Fatal(e.Start("127.0.0.1:6969"))
}
```

### Handlers disponibles

#### IndexHandler (handlers/index.go)

Renderiza la página principal con el listado de salones.

```go
func IndexHandler(c echo.Context) error {
    halls, err := db.GetAllHalls()
    if err != nil {
        return utils.Render(c, 505, views.InternalServerError())
    }
    content := views.HallList(halls)
    return utils.Render(c, 200, views.Layout("Salonez", content, c))
}
```

#### HallList (handlers/hallList.go)

Maneja el listado y filtrado de salones.

**Características:**

- Listado completo de salones
- Filtrado por capacidad, precio, ubicación
- Paginación
- Ordenamiento

#### Login/Signup Handlers

**ShowLoginHandler**: Muestra el formulario de login  
**ValidateLogin**: Valida credenciales de usuario  
**SignUpHandler**: Registra nuevos usuarios

### Rutas del sistema

| Método | Ruta               | Handler          | Descripción              |
| ------ | ------------------ | ---------------- | ------------------------ |
| GET    | `/`                | IndexHandler     | Página principal         |
| GET    | `/halls`           | HallList         | Listado de salones       |
| GET    | `/halls/:id`       | HallDetail       | Detalle de salón         |
| GET    | `/login`           | ShowLoginHandler | Formulario de login      |
| POST   | `/login`           | ValidateLogin    | Validar credenciales     |
| POST   | `/signup`          | SignUpHandler    | Registro de usuario      |
| GET    | `/dashboard/user`  | UserDashboard    | Dashboard de usuario     |
| GET    | `/dashboard/owner` | OwnerDashboard   | Dashboard de propietario |
| GET    | `/dashboard/admin` | AdminDashboard   | Dashboard de admin       |

---

## Vistas (Templ)

### Sistema de plantillas

El proyecto utiliza [Templ](https://templ.guide/) para generar HTML de manera type-safe desde Go.

### Estructura de vistas

#### Layout base (layout.templ)

```go
templ Layout(title string, content templ.Component, c echo.Context) {
    <!DOCTYPE html>
    <html lang="es">
    <head>
        <meta charset="UTF-8"/>
        <title>{title}</title>
        <link rel="stylesheet" href="/static/assets/styles.css"/>
    </head>
    <body>
        @content
        <script src="/static/assets/script.js"></script>
    </body>
    </html>
}
```

### Vistas principales

**index.templ**: Página de inicio  
**hallList.templ**: Listado de salones con cards  
**hall.templ**: Detalle de un salón con galería  
**login-signup.templ**: Formulario de autenticación  
**userHome.templ**: Dashboard de usuario  
**propietaryHome.templ**: Dashboard de propietario  
**adminHome.templ**: Dashboard de administrador

### Componentes reutilizables

- Cards de salones
- Formularios de búsqueda
- Galerías de imágenes
- Navegación
- Footers

### Compilación de templates

```bash
# Generar archivos Go desde templates
templ generate

# O usar Air que lo hace automáticamente
air
```

---

## Utilidades

### render.go

Utilidad para renderizar componentes Templ con Echo.

```go
func Render(c echo.Context, statusCode int, t templ.Component) error {
    buf := templ.GetBuffer()
    defer templ.ReleaseBuffer(buf)

    if err := t.Render(c.Request().Context(), buf); err != nil {
        return err
    }

    return c.HTMLBlob(statusCode, buf.Bytes())
}
```

### imgPath.go

Manejo de rutas de imágenes y uploads.

**Funciones:**

- Validación de tipos de archivo
- Generación de nombres únicos
- Almacenamiento seguro
- Redimensionamiento de imágenes

---

## Testing

### Ejecutar tests

```bash
# Todos los tests
go test ./... -v

# Tests de un paquete específico
go test ./models -v

# Con cobertura
go test ./models -cover

# Generar reporte HTML de cobertura
go test ./models -coverprofile=coverage.out
go tool cover -html=coverage.out

# Ejecutar benchmarks
go test ./models -bench=.
```

### Tests de modelos (models/dataModel_test.go)

El proyecto incluye tests exhaustivos para todos los modelos:

- **TestUserCreation**: Creación de usuarios
- **TestUserTypeAliases**: Verificación de alias de tipos
- **TestHallCreation**: Creación de salones
- **TestHallEmptyImages**: Manejo de arrays vacíos
- **TestLocationCreation**: Creación de ubicaciones
- **TestReservationCreation**: Creación de reservaciones
- **TestReservationCancellation**: Cancelación de reservaciones
- **TestReservationPercentageValidation**: Validación de porcentajes
- **TestCancelationCreation**: Creación de cancelaciones
- **TestHallPriceValidation**: Validación de precios
- **TestLocationPostalCode**: Validación de códigos postales

### Benchmarks

- **BenchmarkHallCreation**: Rendimiento de creación de Hall
- **BenchmarkReservationCreation**: Rendimiento de creación de Reservation

---

## Despliegue

### Desarrollo local

```bash
# Con Air (hot reload)
air

# Sin Air
go run cmd/launchServer.go
```

### Compilación para producción

```bash
# Generar templates
templ generate

# Compilar binario
go build -o salonez cmd/launchServer.go

# Ejecutar binario
./salonez
```

### Docker

```dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o salonez cmd/launchServer.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/salonez .
COPY --from=builder /app/static ./static

EXPOSE 6969
CMD ["./salonez"]
```

### Despliegue en producción

#### Variables de entorno necesarias

```env
DB_HOST=tu-db-host
DB_PORT=5432
DB_USER=tu-usuario
DB_PASSWORD=tu-password
DB_NAME=salonez
SERVER_HOST=0.0.0.0
SERVER_PORT=6969
SESSION_SECRET=secret-largo-y-seguro
```

#### Consideraciones

1. **Base de datos**: Usa un servicio administrado (AWS RDS, Google Cloud SQL, etc.)
2. **Archivos estáticos**: Considera usar CDN para imágenes
3. **HTTPS**: Configura certificados SSL/TLS
4. **Logs**: Implementa logging estructurado
5. **Monitoreo**: Usa herramientas como Prometheus + Grafana
6. **Backups**: Configura backups automáticos de la base de datos

---

## API Reference

### Endpoints públicos

#### GET /

Página principal con listado de salones destacados.

**Response:**

- `200 OK`: HTML con salones

#### GET /halls

Listado completo de salones.

**Query Parameters:**

- `capacidad`: Filtrar por capacidad mínima
- `precio_max`: Precio máximo
- `ciudad`: Ciudad
- `page`: Número de página (paginación)

**Response:**

- `200 OK`: HTML con listado de salones

#### GET /halls/:id

Detalle de un salón específico.

**Parameters:**

- `id`: ID del salón

**Response:**

- `200 OK`: HTML con detalle del salón
- `404 Not Found`: Salón no encontrado

### Endpoints de autenticación

#### GET /login

Muestra formulario de login.

**Response:**

- `200 OK`: HTML con formulario

#### POST /login

Valida credenciales de usuario.

**Body:**

```json
{
  "correo": "usuario@example.com",
  "password": "password123"
}
```

**Response:**

- `200 OK`: Redirección al dashboard
- `401 Unauthorized`: Credenciales inválidas

#### POST /signup

Registra un nuevo usuario.

**Body:**

```json
{
  "nombre": "Juan Pérez",
  "correo": "juan@example.com",
  "password": "password123",
  "tipo_usuario": "usuario"
}
```

**Response:**

- `201 Created`: Usuario creado exitosamente
- `400 Bad Request`: Datos inválidos
- `409 Conflict`: Usuario ya existe

### Endpoints protegidos (requieren autenticación)

#### GET /dashboard/user

Dashboard de usuario regular.

**Response:**

- `200 OK`: HTML con dashboard
- `401 Unauthorized`: No autenticado

#### GET /dashboard/owner

Dashboard de propietario de salones.

**Response:**

- `200 OK`: HTML con dashboard
- `401 Unauthorized`: No autenticado
- `403 Forbidden`: No es propietario

#### GET /dashboard/admin

Dashboard de administrador.

**Response:**

- `200 OK`: HTML con dashboard
- `401 Unauthorized`: No autenticado
- `403 Forbidden`: No es administrador

---

## Mejoras futuras

### Corto plazo

- [ ] Implementar sistema de autenticación completo (JWT/Sessions)
- [ ] Agregar validación de campos en formularios
- [ ] Implementar paginación en listados
- [ ] Agregar sistema de búsqueda avanzada
- [ ] Implementar carga de imágenes

### Mediano plazo

- [ ] Sistema de notificaciones (email/SMS)
- [ ] Integración con pasarelas de pago
- [ ] Dashboard con gráficas y estadísticas
- [ ] Sistema de reviews y calificaciones
- [ ] API REST completa
- [ ] Aplicación móvil

### Largo plazo

- [ ] Sistema de recomendaciones con ML
- [ ] Chat en tiempo real entre usuarios y propietarios
- [ ] Calendario de disponibilidad en tiempo real
- [ ] Integración con redes sociales
- [ ] Sistema de promociones y descuentos

---

## Contribución

### Flujo de trabajo

1. Fork del repositorio
2. Crear branch de feature (`git checkout -b feature/AmazingFeature`)
3. Commit de cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push al branch (`git push origin feature/AmazingFeature`)
5. Abrir Pull Request

### Estándares de código

- Seguir las convenciones de Go (gofmt, golint)
- Tests para nuevas funcionalidades
- Documentación de funciones públicas
- Commits descriptivos en español

---

## Licencia

Este proyecto está bajo la licencia MIT. Ver archivo `LICENSE` para más detalles.

---

## Contacto

**Proyecto:** Salonez  
**Repositorio:** https://github.com/TheAaronn/Salonez  
**Autor:** TheAaronn

---

## Changelog

### [1.0.0] - 2025-11-24

#### Añadido

- Estructura inicial del proyecto
- Sistema de modelos de datos
- Base de datos PostgreSQL con schema completo
- Handlers básicos (index, hallList, login, signup)
- Sistema de vistas con Templ
- Tests unitarios para modelos
- Documentación completa
- Configuración de Air para desarrollo

#### Por implementar

- Sistema de autenticación completo
- CRUD completo para todas las entidades
- Sistema de reservaciones
- Panel de administración
- Integración de pagos

---

**Última actualización:** 24 de noviembre de 2025
