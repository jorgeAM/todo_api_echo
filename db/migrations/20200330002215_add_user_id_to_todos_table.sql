-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE todos ADD userId INT UNSIGNED AFTER `description`;
ALTER TABLE todos ADD CONSTRAINT FK_user FOREIGN KEY (userId) REFERENCES users(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE todos DROP FOREIGN KEY FK_user;
ALTER TABLE todos DROP COLUMN userId;
