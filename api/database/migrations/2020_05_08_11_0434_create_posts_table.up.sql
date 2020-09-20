CREATE TABLE `posts` (
   `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   `uuid` VARCHAR(36) NOT NULL,
   `slug` VARCHAR(150) NOT NULL UNIQUE,
   `title` VARCHAR(255) NULL,
   `status` VARCHAR(150) NOT NULL DEFAULT 'draft',
   `resource` VARCHAR(150) NULL,
   `page_template` VARCHAR(150) NOT NULL,
   `layout` VARCHAR(150) NOT NULL,
   `fields` JSON NULL,
   `codeinjection_head` LONGTEXT NULL,
   `codeinjection_foot` LONGTEXT NULL,
   `user_id` INT NOT NULL,
   `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
   `created_at` TIMESTAMP NOT NULL
);