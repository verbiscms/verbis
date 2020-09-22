CREATE TABLE `password_resets` (
    `email` VARCHAR(255) NOT NULL PRIMARY KEY,
    `token` VARCHAR(255),
    `created_at` TIMESTAMP
);