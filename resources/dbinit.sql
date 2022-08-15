CREATE TABLE goChallenge.transaction(
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            amount FLOAT NOT NULL ,
                            currency STRING NOT NULL,
                            createdAt TIMESTAMP NOT NULL
);


CREATE DATABASE goChallenge;

DROP TABLE goChallenge.transaction;

INSERT INTO goChallenge.transaction(amount, currency, createdAt) VALUES (9348.3,'CLP',now());
SELECT * FROM goChallenge.transaction;


