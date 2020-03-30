-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  firstName VARCHAR(100) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(100) NOT NULL,
  createdAt DATETIME,
  updatedAt DATETIME
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;