CREATE TABLE city(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    region varchar(255) NOT NULL
);

CREATE TABLE account(
    id SERIAL PRIMARY KEY,
    city_id SERIAL REFERENCES city NOT NULL,
    name VARCHAR(255) NOT NULL,
    balance NUMERIC NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);


CREATE TABLE transfer(
    id SERIAL PRIMARY KEY,
    sender_id SERIAL REFERENCES account NOT NULL,
    recipient_id SERIAL REFERENCES account NOT NULL,
    amount NUMERIC NOT NULL
);