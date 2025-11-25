-- Solo inserta datos de demo en Salonez
-- Las tablas ya deben existir

-- Limpiar datos existentes de demo (opcional)
DELETE FROM cancelations;
DELETE FROM reservations;
DELETE FROM halls;
DELETE FROM locations;
DELETE FROM users;

-- Reiniciar secuencias
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE locations_id_seq RESTART WITH 1;
ALTER SEQUENCE halls_id_seq RESTART WITH 1;
ALTER SEQUENCE reservations_id_seq RESTART WITH 1;
ALTER SEQUENCE cancelations_id_seq RESTART WITH 1;

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

-- Insertar salones de demo (propietario id=2)
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

-- Insertar reservaciones de demo (usuario id=3, halls ids 1,2,3)
INSERT INTO reservations (hall_id, user_id, date, time, percentage_paid, cancelled) VALUES
(1, 3, '2025-12-15', '18:00', 50.00, false),
(2, 3, '2025-12-20', '14:00', 100.00, false),
(3, 3, '2025-11-25', '16:00', 25.00, true);

-- Insertar cancelación para la reservación cancelada
INSERT INTO cancelations (reservation_id, date, time, motivo) VALUES
(3, '2025-11-24', '10:30', 'Cambio de planes del usuario');

-- Verificar datos insertados
SELECT 'Datos insertados exitosamente!' as status;
SELECT COUNT(*) as total_users FROM users;
SELECT COUNT(*) as total_halls FROM halls;
SELECT COUNT(*) as total_reservations FROM reservations;
