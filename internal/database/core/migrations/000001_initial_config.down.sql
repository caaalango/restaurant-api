BEGIN;

DROP TABLE IF EXISTS comments;

DROP TABLE IF EXISTS dishes;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS restaurants;

DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;
