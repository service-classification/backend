-- migrations/initial.sql

CREATE TABLE services
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    parameters  TEXT[] NOT NULL
);

CREATE TABLE groups
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE classified_services
(
    service_id INT REFERENCES services (id),
    group_id   INT REFERENCES groups (id),
    PRIMARY KEY (service_id, group_id)
);
