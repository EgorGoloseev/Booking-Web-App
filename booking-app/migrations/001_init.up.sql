CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       role TEXT NOT NULL
);

CREATE TABLE rooms (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       location TEXT,
                       capacity INT
);

CREATE TABLE bookings (
                          id SERIAL PRIMARY KEY,
                          room_id INT REFERENCES rooms(id),
                          user_id INT REFERENCES users(id),
                          start_time TIMESTAMP NOT NULL,
                          end_time TIMESTAMP NOT NULL,
                          purpose TEXT,
                          created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_room_time
    ON bookings(room_id, start_time, end_time);

CREATE INDEX idx_user_time
    ON bookings(user_id, start_time);