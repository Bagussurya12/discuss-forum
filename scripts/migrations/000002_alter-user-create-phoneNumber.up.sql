ALTER TABLE users
ADD phone_number VARCHAR(18) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT unique_phone_number UNIQUE (phone_number);