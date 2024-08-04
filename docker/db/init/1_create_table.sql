CREATE TABLE users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  email text NOT NULL,
  password text NOT NULL
);

CREATE TABLE persons (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID  integer,
  name text,
  gender  text,
  brith integer,
  stations integer[],
  areas integer[],
  gyms integer[],
  times text[],
  bp integer,
  sq integer,
  dl integer
);

CREATE TABLE lines (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE stations (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  lineID integer,
  name text NOT NULL
);

CREATE TABLE areas (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE gyms (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE muscles (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  part text NOT NULL,
  name text NOT NULL
);

CREATE TABLE menus (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  image text NOT NULL,
  type text NOT NULL,
  target integer,
  muscles integer[]
);

CREATE TABLE matches (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  requester integer,
  approver integer,
  status  text
);

CREATE TABLE chats (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  member integer[]
);

CREATE TABLE messages (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  chatID integer,
  sender integer,
  receiver integer,
  content text NOT NULL,
  date timestamp
);

CREATE TABLE notices (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID  integer,
  chatID integer,
  type  text
);