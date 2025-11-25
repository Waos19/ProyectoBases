

CREATE TABLE PERSONAL (
    id_personal NUMBER PRIMARY KEY,
    nombre VARCHAR2(100) NOT NULL,
    tipo VARCHAR2(20) NOT NULL,
    CONSTRAINT chk_personal_tipo CHECK (tipo IN ('estudiante','docente','admin'))
);

CREATE TABLE SERVICIO (
    id_servicio NUMBER PRIMARY KEY,
    nombre VARCHAR2(100) NOT NULL,
    categoria VARCHAR2(50) NOT NULL,
    max_integrantes NUMBER NOT NULL
);

CREATE TABLE INSTRUCTOR (
    id_instructor NUMBER PRIMARY KEY,
    nombre VARCHAR2(100) NOT NULL,
    especialidad VARCHAR2(50) NOT NULL
);

CREATE TABLE HORARIO (
    id_horario NUMBER PRIMARY KEY,
    id_servicio NUMBER NOT NULL,
    dia_semana VARCHAR2(20) NOT NULL,
    hora_inicio DATE NOT NULL,
    hora_fin DATE NOT NULL,
    CONSTRAINT fk_horario_servicio FOREIGN KEY (id_servicio)
        REFERENCES SERVICIO(id_servicio)
);

CREATE TABLE ACTIVIDAD (
    id_actividad NUMBER PRIMARY KEY,
    id_servicio NUMBER NOT NULL,
    fecha DATE NOT NULL,
    id_horario NUMBER NOT NULL,
    CONSTRAINT fk_actividad_servicio FOREIGN KEY (id_servicio)
        REFERENCES SERVICIO(id_servicio),
    CONSTRAINT fk_actividad_horario FOREIGN KEY (id_horario)
        REFERENCES HORARIO(id_horario)
);

CREATE TABLE ELEMENTO (
    id_elemento NUMBER PRIMARY KEY,
    nombre VARCHAR2(100) NOT NULL,
    tipo VARCHAR2(50) NOT NULL,
    cantidad_disponible NUMBER NOT NULL
);

CREATE TABLE ACTIVIDAD_ELEMENTO (
    id_actividad NUMBER NOT NULL,
    id_elemento NUMBER NOT NULL,
    cantidad_usada NUMBER NOT NULL,
    CONSTRAINT pk_actividad_elemento PRIMARY KEY (id_actividad, id_elemento),
    CONSTRAINT fk_ae_actividad FOREIGN KEY (id_actividad)
        REFERENCES ACTIVIDAD(id_actividad),
    CONSTRAINT fk_ae_elemento FOREIGN KEY (id_elemento)
        REFERENCES ELEMENTO(id_elemento)
);

CREATE TABLE PARTICIPACION (
    id_personal NUMBER NOT NULL,
    id_actividad NUMBER NOT NULL,
    rol VARCHAR2(20) NOT NULL,
    CONSTRAINT chk_participacion_rol CHECK (rol IN ('participante','instructor')),
    CONSTRAINT pk_participacion PRIMARY KEY (id_personal, id_actividad),
    CONSTRAINT fk_participacion_personal FOREIGN KEY (id_personal)
        REFERENCES PERSONAL(id_personal),
    CONSTRAINT fk_participacion_actividad FOREIGN KEY (id_actividad)
        REFERENCES ACTIVIDAD(id_actividad)
);

CREATE TABLE PAGO (
    id_pago NUMBER PRIMARY KEY,
    id_personal NUMBER NOT NULL,
    id_actividad NUMBER NOT NULL,
    monto NUMBER(10,2) NOT NULL,
    fecha DATE NOT NULL,
    CONSTRAINT fk_pago_personal FOREIGN KEY (id_personal)
        REFERENCES PERSONAL(id_personal),
    CONSTRAINT fk_pago_actividad FOREIGN KEY (id_actividad)
        REFERENCES ACTIVIDAD(id_actividad)
);


-- Datos para PERSONAL
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (1, 'Juan Pérez', 'estudiante');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (2, 'María López', 'docente');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (3, 'Carlos Sánchez', 'admin');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (4, 'Ana Gómez', 'estudiante');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (5, 'Luis Rodríguez', 'docente');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (6, 'Sofía Díaz', 'admin');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (7, 'Pedro Martínez', 'estudiante');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (8, 'Lucía Morales', 'docente');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (9, 'Diego Herrera', 'admin');
INSERT INTO PERSONAL (id_personal, nombre, tipo) VALUES (10, 'Elena Navarro', 'admin');

-- Datos para SERVICIO
INSERT INTO SERVICIO (id_servicio, nombre, categoria, max_integrantes) VALUES (1, 'Fútbol', 'Deporte', 20);
INSERT INTO SERVICIO (id_servicio, nombre, categoria, max_integrantes) VALUES (2, 'Pintura', 'Arte', 15);
INSERT INTO SERVICIO (id_servicio, nombre, categoria, max_integrantes) VALUES (3, 'Guitarra', 'Música', 12);
INSERT INTO SERVICIO (id_servicio, nombre, categoria, max_integrantes) VALUES (4, 'Programación', 'Tecnología', 18);
INSERT INTO SERVICIO (id_servicio, nombre, categoria, max_integrantes) VALUES (5, 'Yoga', 'Deporte', 25);

-- Datos para INSTRUCTOR
INSERT INTO INSTRUCTOR (id_instructor, nombre, especialidad) VALUES (1, 'Roberto Torres', 'Música');
INSERT INTO INSTRUCTOR (id_instructor, nombre, especialidad) VALUES (2, 'Patricia Vega', 'Deporte');
INSERT INTO INSTRUCTOR (id_instructor, nombre, especialidad) VALUES (3, 'Miguel Castillo', 'Tecnología');
INSERT INTO INSTRUCTOR (id_instructor, nombre, especialidad) VALUES (4, 'Laura Rivas', 'Música');
INSERT INTO INSTRUCTOR (id_instructor, nombre, especialidad) VALUES (5, 'Andrés Molina', 'Deporte');

-- Datos para HORARIO
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (1, 3, 'Martes', TO_DATE('2024-03-12 14:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-12 16:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (2, 1, 'Lunes', TO_DATE('2024-03-11 16:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-11 18:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (3, 1, 'Miércoles', TO_DATE('2024-03-13 17:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-13 19:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (4, 1, 'Viernes', TO_DATE('2024-03-15 11:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-15 13:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (5, 1, 'Lunes', TO_DATE('2024-03-11 14:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-11 16:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (6, 4, 'Lunes', TO_DATE('2024-03-11 11:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-11 13:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (7, 1, 'Viernes', TO_DATE('2024-03-15 14:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-15 16:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (8, 1, 'Viernes', TO_DATE('2024-03-15 09:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-15 11:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (9, 2, 'Sábado', TO_DATE('2024-03-16 18:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-16 20:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO HORARIO (id_horario, id_servicio, dia_semana, hora_inicio, hora_fin) VALUES (10, 5, 'Lunes', TO_DATE('2024-03-11 17:00', 'YYYY-MM-DD HH24:MI'), TO_DATE('2024-03-11 19:00', 'YYYY-MM-DD HH24:MI'));

-- Datos para ACTIVIDAD
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (1, 5, TO_DATE('2024-03-11', 'YYYY-MM-DD'), 10);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (2, 1, TO_DATE('2024-03-11', 'YYYY-MM-DD'), 5);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (3, 1, TO_DATE('2024-03-13', 'YYYY-MM-DD'), 3);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (4, 2, TO_DATE('2024-03-16', 'YYYY-MM-DD'), 9);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (5, 4, TO_DATE('2024-03-11', 'YYYY-MM-DD'), 6);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (6, 5, TO_DATE('2024-03-11', 'YYYY-MM-DD'), 10);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (7, 1, TO_DATE('2024-03-15', 'YYYY-MM-DD'), 7);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (8, 1, TO_DATE('2024-03-13', 'YYYY-MM-DD'), 3);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (9, 2, TO_DATE('2024-03-16', 'YYYY-MM-DD'), 9);
INSERT INTO ACTIVIDAD (id_actividad, id_servicio, fecha, id_horario) VALUES (10, 5, TO_DATE('2024-03-11', 'YYYY-MM-DD'), 10);

-- Datos para PARTICIPACION
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (9, 10, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (6, 6, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (6, 10, 'instructor');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (8, 10, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (8, 2, 'instructor');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (2, 5, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (8, 2, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (1, 5, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (10, 8, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (5, 7, 'instructor');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (6, 1, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (8, 6, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (3, 10, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (2, 8, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (1, 4, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (5, 3, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (4, 7, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (7, 8, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (2, 3, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (8, 7, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (9, 5, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (3, 7, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (7, 6, 'participante');
INSERT INTO PARTICIPACION (id_personal, id_actividad, rol) VALUES (7, 4, 'participante');

-- Datos para ELEMENTO
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (1, 'Balón de fútbol', 'Equipo', 11);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (2, 'Pinceles', 'Material', 16);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (3, 'Lienzo', 'Otro', 6);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (4, 'Guitarra acústica', 'Otro', 5);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (5, 'Portátil', 'Balón', 21);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (6, 'Pesas', 'Otro', 10);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (7, 'Colchoneta', 'Material', 10);
INSERT INTO ELEMENTO (id_elemento, nombre, tipo, cantidad_disponible) VALUES (8, 'Proyector', 'Instrumento', 27);

-- Datos para ACTIVIDAD_ELEMENTO
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (4, 2, 6);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (2, 8, 9);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (2, 4, 1);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (10, 3, 1);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (8, 3, 3);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (9, 3, 2);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (6, 1, 4);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (10, 7, 4);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (6, 4, 2);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (4, 5, 3);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (4, 7, 1);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (10, 3, 2);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (9, 5, 8);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (6, 2, 8);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (5, 4, 1);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (2, 6, 5);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (7, 3, 1);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (6, 7, 2);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (8, 6, 4);
INSERT INTO ACTIVIDAD_ELEMENTO (id_actividad, id_elemento, cantidad_usada) VALUES (1, 7, 1);

-- Datos para PAGO
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (1, 3, 2, 22.34, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (2, 4, 1, 43.95, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (3, 3, 5, 29.74, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (4, 7, 9, 35.85, TO_DATE('2024-03-16', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (5, 6, 3, 58.33, TO_DATE('2024-03-13', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (6, 10, 1, 41.97, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (7, 9, 7, 37.86, TO_DATE('2024-03-15', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (8, 2, 8, 54.40, TO_DATE('2024-03-13', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (9, 4, 2, 78.93, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (10, 3, 2, 33.80, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (11, 2, 1, 49.67, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (12, 2, 6, 52.96, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (13, 4, 10, 36.34, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (14, 5, 6, 52.16, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (15, 2, 2, 69.43, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (16, 8, 8, 31.83, TO_DATE('2024-03-13', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (17, 2, 6, 61.82, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (18, 3, 9, 11.62, TO_DATE('2024-03-16', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (19, 9, 6, 52.16, TO_DATE('2024-03-11', 'YYYY-MM-DD'));
INSERT INTO PAGO (id_pago, id_personal, id_actividad, monto, fecha) VALUES (20, 1, 9, 30.87, TO_DATE('2024-03-16', 'YYYY-MM-DD'));

COMMIT;