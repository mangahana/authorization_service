CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  username    VARCHAR(20) NOT NULL UNIQUE,
  phone       VARCHAR(10) UNIQUE NOT NULL,
  password    VARCHAR(256) NOT NULL,
  description TEXT DEFAULT '',
  photo       TEXT,
  role_id     INT DEFAULT 1,
  is_banned   BOOLEAN NOT NULL DEFAULT FALSE,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE phone_block_list (
  phone       VARCHAR(10) UNIQUE NOT NULL,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE confirmation_codes (
  code        VARCHAR(6) UNIQUE NOT NULL,
  phone       VARCHAR(10) NOT NULL,
  ip          VARCHAR(32) NOT NULL,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE sessions (
  user_id    SERIAL NOT NULL,
  token      VARCHAR(128) NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE roles (
  id          SERIAL PRIMARY KEY,
  name        TEXT NOT NULL,
  permissions INTEGER[] NOT NULL DEFAULT '{}'
);
INSERT INTO roles (name) VALUES ('Қолданушы');


CREATE TABLE permissions (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
INSERT INTO permissions (name) VALUES ('approve_team');