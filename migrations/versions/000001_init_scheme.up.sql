CREATE TABLE IF NOT EXISTS "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS role_name_idx ON role (name);

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    pinged_at TIMESTAMP NOT NULL,
    tg_username VARCHAR(32) DEFAULT NULL,
    tg_id BIGINT UNIQUE NOT NULL,
    phone_number VARCHAR(20) DEFAULT NULL,
    name VARCHAR(50) DEFAULT NULL,
    surname VARCHAR(50) DEFAULT NULL,
    username VARCHAR(50) DEFAULT NULL UNIQUE,
    access_token VARCHAR(128) DEFAULT NULL UNIQUE,
    refresh_token VARCHAR(128) DEFAULT NULL UNIQUE,
    description VARCHAR(200) DEFAULT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS user_tg_id_idx ON "user" (tg_id);
CREATE UNIQUE INDEX IF NOT EXISTS user_username_idx ON "user" (username);

CREATE TABLE IF NOT EXISTS "user_role" (
    user_id INTEGER NOT NULL REFERENCES "user"(id),
    role_id INTEGER NOT NULL REFERENCES role(id),
    expired_on TIMESTAMP DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS "route" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user"(id),
    start_address VARCHAR(128) NOT NULL,
    start_latitude DECIMAL(10, 8),
    start_longitude DECIMAL(11, 8),
    stop_address VARCHAR(128) NOT NULL,
    stop_latitude DECIMAL(10, 8),
    stop_longitude DECIMAL(11, 8)
);

CREATE TABLE IF NOT EXISTS vehicle (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user"(id),
    brand VARCHAR(64) NOT NULL,
    license_plate VARCHAR(20) NOT NULL UNIQUE
);
CREATE UNIQUE INDEX IF NOT EXISTS vehicle_license_plate_idx ON vehicle (license_plate);

CREATE TABLE IF NOT EXISTS trip (
    id SERIAL PRIMARY KEY,
    driver_id INTEGER NOT NULL REFERENCES "user"(id),
    route_id INTEGER NOT NULL REFERENCES route(id),
    vehicle_id INTEGER NOT NULL REFERENCES vehicle(id),
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    seats_count INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    is_completed BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS booking (
    trip_id INTEGER NOT NULL REFERENCES trip(id),
    user_id INTEGER NOT NULL REFERENCES "user"(id),
    booking_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (trip_id, user_id)
);
