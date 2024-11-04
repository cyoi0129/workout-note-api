CREATE TABLE workout_users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  email text NOT NULL,
  password text NOT NULL
);

CREATE TABLE workout_persons (
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

CREATE TABLE workout_lines (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE workout_stations (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  lineID integer,
  name text NOT NULL
);

CREATE TABLE workout_areas (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE workout_gyms (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE workout_muscles (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  part text NOT NULL,
  name text NOT NULL
);

CREATE TABLE workout_menus (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  image text NOT NULL,
  type text NOT NULL,
  target integer,
  muscles integer[]
);

CREATE TABLE workout_matches (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  requester integer,
  approver integer,
  status  text
);

CREATE TABLE workout_chats (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  member integer[]
);

CREATE TABLE workout_messages (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  chatID integer,
  sender integer,
  receiver integer,
  content text NOT NULL,
  date timestamp
);

CREATE TABLE workout_notices (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID  integer,
  chatID integer,
  type  text
);