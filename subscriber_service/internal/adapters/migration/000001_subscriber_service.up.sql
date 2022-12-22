CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE averages
(
    average_id UUID PRIMARY KEY         NOT NULL DEFAULT uuid_generate_v4(),
    average    NUMERIC(5, 2)            NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE sensors
(
    sensor_id  UUID PRIMARY KEY         NOT NULL DEFAULT uuid_generate_v4(),
    average_id UUID                     NOT NULL,
    name       VARCHAR(128)             NOT NULL,
    timestamp  TIMESTAMP                NOT NULL,
    value      NUMERIC(5, 2)            NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE "sensors"
    ADD FOREIGN KEY ("average_id") REFERENCES "averages" ("average_id");
