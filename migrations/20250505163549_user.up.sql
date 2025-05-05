CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name TEXT NOT NULL,
                       second_name TEXT NOT NULL,
                       birthdate DATE NOT NULL,
                       gender TEXT NOT NULL,
                       biography TEXT,
                       city TEXT,
                       password TEXT NOT NULL
);
