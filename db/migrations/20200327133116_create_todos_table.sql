-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE todos (
  id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  createdAt DATETIME,
  updatedAt DATETIME
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE todos;