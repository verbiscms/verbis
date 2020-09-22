CREATE TABLE `categories` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid` VARCHAR(36) NOT NULL,
    `slug` VARCHAR(150) NOT NULL UNIQUE,
    `name` VARCHAR(150) NOT NULL,
    `description` TEXT DEFAULT NULL,
    `hidden` TINYINT(1) NOT NULL DEFAULT '0',
    `parent_id` INT(11) DEFAULT NULL,
    `page_template` VARCHAR(150) NULL,
    `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL
);