CREATE TABLE `seo_meta_options` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `page_id` INT NOT NULL,
    `seo` JSON,
    `meta` JSON
);