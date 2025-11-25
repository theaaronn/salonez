-- Script completo de inicialización para Salonez
-- Incluye schema y datos de demo

-- ============================================================================
-- SCHEMA
-- ============================================================================

-- Tabla de usuarios (sirve para User, Usuario, Propietario, Admin)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) NOT NULL UNIQUE,
    tipo_usuario VARCHAR(20) DEFAULT 'usuario' CHECK (tipo_usuario IN ('usuario', 'propietario', 'admin')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de ubicaciones
CREATE TABLE IF NOT EXISTS locations (
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
CREATE TABLE IF NOT EXISTS halls (
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
CREATE TABLE IF NOT EXISTS reservations (
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
CREATE TABLE IF NOT EXISTS cancelations (
    id SERIAL PRIMARY KEY,
    reservation_id INTEGER NOT NULL REFERENCES reservations(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    time TIME NOT NULL,
    motivo TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Índices para mejorar el rendimiento
CREATE INDEX IF NOT EXISTS idx_users_correo ON users(correo);
CREATE INDEX IF NOT EXISTS idx_users_tipo ON users(tipo_usuario);
CREATE INDEX IF NOT EXISTS idx_halls_propietario ON halls(propietario_id);
CREATE INDEX IF NOT EXISTS idx_halls_nombre ON halls(nombre);
CREATE INDEX IF NOT EXISTS idx_halls_imgs_path ON halls USING GIN(imgs_path);
CREATE INDEX IF NOT EXISTS idx_reservations_hall ON reservations(hall_id);
CREATE INDEX IF NOT EXISTS idx_reservations_user ON reservations(user_id);
CREATE INDEX IF NOT EXISTS idx_reservations_date ON reservations(date);
CREATE INDEX IF NOT EXISTS idx_reservations_cancelled ON reservations(cancelled);

-- Función para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggers para actualizar updated_at
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_locations_updated_at ON locations;
CREATE TRIGGER update_locations_updated_at BEFORE UPDATE ON locations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_halls_updated_at ON halls;
CREATE TRIGGER update_halls_updated_at BEFORE UPDATE ON halls
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_reservations_updated_at ON reservations;
CREATE TRIGGER update_reservations_updated_at BEFORE UPDATE ON reservations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================================
-- DATOS DE DEMO
-- ============================================================================

-- Limpiar datos existentes (solo para demo)
TRUNCATE users, locations, halls, reservations, cancelations CASCADE;

-- Insertar usuarios de demo
INSERT INTO users (nombre, correo, tipo_usuario) VALUES
('Admin User', 'admin', 'admin'),
('Owner User', 'owner', 'propietario'),
('Regular User', 'user', 'usuario');

-- Insertar algunas ubicaciones
INSERT INTO locations (calle, numero, colonia, cp, ciudad, estado) VALUES
('Avenida Insurgentes', '1234', 'Centro', 06000, 'Ciudad de México', 'CDMX'),
('Calle Reforma', '567', 'Polanco', 11560, 'Ciudad de México', 'CDMX'),
('Boulevard Juárez', '890', 'Del Valle', 03100, 'Ciudad de México', 'CDMX');

-- Insertar salones de demo (asumiendo que el owner tiene id=2)
INSERT INTO halls (imgs_path, nombre, direccion, capacidad, numero_telefono, precio, propietario_id, location_id) VALUES
(ARRAY['/static/hallImgs/hall1-1.jpg', '/static/hallImgs/hall1-2.jpg'], 
 'Salón Principal', 
 'Avenida Insurgentes 1234, Centro, 06000, Ciudad de México, CDMX', 
 '200 personas', 
 '555-1234-5678', 
 5000.00, 
 2, 
 1),

(ARRAY['/static/hallImgs/hall2-1.jpg'], 
 'Salón Ejecutivo', 
 'Calle Reforma 567, Polanco, 11560, Ciudad de México, CDMX', 
 '100 personas', 
 '555-2345-6789', 
 3500.00, 
 2, 
 2),

(ARRAY['/static/hallImgs/hall3-1.jpg', '/static/hallImgs/hall3-2.jpg', '/static/hallImgs/hall3-3.jpg'], 
 'Salón Jardín', 
 'Boulevard Juárez 890, Del Valle, 03100, Ciudad de México, CDMX', 
 '150 personas', 
 '555-3456-7890', 
 4200.00, 
 2, 
 3);

-- Insertar reservaciones de demo (asumiendo user id=3, halls ids 1,2,3)
INSERT INTO reservations (hall_id, user_id, date, time, percentage_paid, cancelled) VALUES
(1, 3, '2025-12-15', '18:00', 50.00, false),
(2, 3, '2025-12-20', '14:00', 100.00, false),
(3, 3, '2025-11-25', '16:00', 25.00, true);

-- Insertar cancelación para la reservación cancelada
INSERT INTO cancelations (reservation_id, date, time, motivo) VALUES
(3, '2025-11-24', '10:30', 'Cambio de planes del usuario');

-- Mensaje de confirmación
SELECT 'Database initialized successfully!' as status;
SELECT COUNT(*) as total_users FROM users;
SELECT COUNT(*) as total_halls FROM halls;
SELECT COUNT(*) as total_reservations FROM reservations;
