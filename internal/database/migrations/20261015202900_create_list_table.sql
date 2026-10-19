-- +goose Up
create table lists (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  status VARCHAR(100) NOT NULL DEFAULT 'created'
);

-- +goose Down
DROP TABLE lists;
