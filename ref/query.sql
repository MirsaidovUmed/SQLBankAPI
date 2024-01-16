CREATE TABLE account(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(250) NOT NULL,
    balance NUMERIC NOT NULL DEFAULT 0,
    phone_number VARCHAR(9) NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE transfer(
    id SERIAL PRIMARY KEY,
    sender_id SERIAL REFERENCES account NOT NULL,
    recipient_id SERIAL REFERENCES account NOT NULL,
    amount NUMERIC NOT NULL
);