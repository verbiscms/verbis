CREATE TABLE `users` (
   `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   `uuid` VARCHAR(36) NOT NULL,
   `first_name` VARCHAR(150) NOT NULL,
   `last_name` VARCHAR(150) NOT NULL,
   `email` VARCHAR(255) NOT NULL UNIQUE,
   `password` VARCHAR(60) NOT NULL,
   `website` TEXT NULL,
   `facebook` TEXT NULL,
   `twitter` TEXT NULL,
   `linked_in` TEXT NULL,
   `instagram` TEXT NULL,
   `token` VARCHAR(38) NOT NULL,
   `email_verified_at` TIMESTAMP NULL,
   `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
   `created_at` TIMESTAMP NOT NULL
);