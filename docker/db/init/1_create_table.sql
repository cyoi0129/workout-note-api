CREATE TABLE users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL
);

CREATE TABLE types (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE muscles (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  part text NOT NULL,
  name text NOT NULL
);

CREATE TABLE masters (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID integer,
  name text NOT NULL,
  image text NOT NULL,
  type integer,
  target integer,
  muscles integer[]
);