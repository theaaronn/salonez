-- Schema PostgreSQL para Salonez
-- Este script crea todas las tablas necesarias para la aplicación

-- Tabla de usuarios (sirve para User, Usuario, Propietario, Admin)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) NOT NULL UNIQUE,
    tipo_usuario VARCHAR(20) DEFAULT 'usuario' CHECK (tipo_usuario IN ('usuario', 'propietario', 'admin')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de ubicaciones
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

-- Tabla de salones (halls) con array nativo para ImgsPath
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

-- Tabla de reservaciones
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

-- Tabla de cancelaciones
CREATE TABLE cancelations (
    id SERIAL PRIMARY KEY,
    reservation_id INTEGER NOT NULL REFERENCES reservations(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    time TIME NOT NULL,
    motivo TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Índices para mejorar el rendimiento
CREATE INDEX idx_users_correo ON users(correo);
CREATE INDEX idx_users_tipo ON users(tipo_usuario);
CREATE INDEX idx_halls_propietario ON halls(propietario_id);
CREATE INDEX idx_halls_nombre ON halls(nombre);
CREATE INDEX idx_halls_imgs_path ON halls USING GIN(imgs_path);
CREATE INDEX idx_reservations_hall ON reservations(hall_id);
CREATE INDEX idx_reservations_user ON reservations(user_id);
CREATE INDEX idx_reservations_date ON reservations(date);
CREATE INDEX idx_reservations_cancelled ON reservations(cancelled);

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

-- Comentarios de tablas
COMMENT ON TABLE users IS 'Almacena información de usuarios, propietarios y administradores';
COMMENT ON TABLE locations IS 'Almacena ubicaciones geográficas detalladas';
COMMENT ON TABLE halls IS 'Almacena información de salones de eventos';
COMMENT ON TABLE reservations IS 'Almacena reservaciones de salones';
COMMENT ON TABLE cancelations IS 'Almacena registro de cancelaciones de reservaciones';

-- Comentarios de columnas importantes
COMMENT ON COLUMN halls.imgs_path IS 'Array de rutas a imágenes del salón';
COMMENT ON COLUMN reservations.percentage_paid IS 'Porcentaje pagado de la reservación (0-100)';
COMMENT ON COLUMN reservations.cancelled IS 'Indica si la reservación fue cancelada';
