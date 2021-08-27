
-- +migrate Up
CREATE TABLE movies_log_detail
(
    id INT NOT NULL AUTO_INCREMENT,
    deleted_at DATETIME DEFAULT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    payload TEXT DEFAULT NULL,
    PRIMARY KEY (id)
);


-- +migrate Down
DROP TABLE movies_log_detail;
