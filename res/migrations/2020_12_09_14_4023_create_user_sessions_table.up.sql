CREATE TABLE `user_sessions` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `session_key` VARCHAR(38) NOT NULL,
    `login_time` TIMESTAMP NOT NULL,
    `last_seen_time` TIMESTAMP NOT NULL
);