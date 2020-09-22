CREATE TABLE `media` (
     `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
     `uuid` VARCHAR(36) NOT NULL,
     `url` VARCHAR(255) NULL,
     `title` VARCHAR(150) NULL,
     `alt` VARCHAR(150) NULL,
     `description` TEXT NULL,
     `file_path` VARCHAR(255) NOT NULL,
     `file_size` INT NOT NULL,
     `file_name` VARCHAR(255) NOT NULL,
     `sizes` JSON NULL,
     `type` VARCHAR(150) NOT NULL,
     `user_id` INT NULL,
     `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `created_at` TIMESTAMP
);