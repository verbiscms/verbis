##############################################
# Create Files Table
##############################################

CREATE TABLE `files`
(
    `id`          int           NOT NULL AUTO_INCREMENT,
    `uuid`        varchar(36)   NOT NULL,
    `url`         varchar(500)  NOT NULL,
    `name`        varchar(255)  NOT NULL,
    `bucket_id`   text          NOT NULL,
    `mime`        varchar(150)  NOT NULL,
    `source_type` varchar(255)  NOT NULL,
    `provider`    varchar(255)  NOT NULL,
    `region`      varchar(255)  NOT NULL,
    `bucket`      varchar(255)  NOT NULL,
    `file_size`    int           NOT NULL,
    `private`     bit DEFAULT 0 NOT NULL,
    PRIMARY KEY (`id`),
    KEY `files_url_index` (`url`),
    CONSTRAINT files_unique UNIQUE (url)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

##############################################
# Create Media Sizes Table
##############################################

CREATE TABLE `media_sizes` (
   `id`        int           NOT NULL AUTO_INCREMENT,
   `uuid`      varchar(36)   NOT NULL,
   `file_id`    int           NOT NULL,
   `media_id`  int           NOT NULL,
   `size_key`  varchar(255)  NOT NULL,
   `size_name` varchar(255)  NOT NULL,
   `width`     int           NOT NULL,
   `height`    int           NOT NULL,
   `crop`      bit DEFAULT 0 NOT NULL,
   PRIMARY KEY (`id`),
   KEY `media_size_file_index` (`file_id`),
   KEY `media_size_media_index` (`media_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

##############################################
# Inserts for Storage
##############################################

# Delete Duplicate URLS
DELETE FROM media m WHERE id IN (
    WITH DeleteURL as (SELECT `id`,
      ROW_NUMBER() OVER(PARTITION BY url ORDER BY ID DESC) row_num
FROM media) SELECT id FROM DeleteURL WHERE row_num > 1
);

# Insert main files (no media sizes)
INSERT INTO files (`uuid`, `url`, `file_size`, `name`, `mime`, `bucket_id`, `region`, `provider`, `bucket`, `source_type`)
SELECT `uuid`, `url`, `file_size`, `file_name`, `mime`, CONCAT('uploads/', media.file_path, '/', media.uuid, '.', SUBSTRING_INDEX(media.file_name, '.', -1)) AS bucket_id, '' AS region, 'local' AS provider,  '' AS bucket, 'media'
FROM media;

# Insert media sizes
INSERT INTO files (`url`, `uuid`, `file_size`, `name`, `mime`, `bucket_id`, `region`, `provider`, `bucket`, `source_type`)
    (SELECT sizes.url, sizes.uuid,  sizes.file_size, sizes.name, media.mime, CONCAT('uploads/', media.file_path, '/', sizes.uuid, '.', SUBSTRING_INDEX(media.file_name, '.', -1)) AS bucket_id, '' AS region, 'local' AS provider, '' AS bucket, 'media'
     FROM media,
          JSON_TABLE(sizes, '$.*'
                     COLUMNS (
                         uuid VARCHAR(255) PATH '$**.uuid',
                         url VARCHAR(255) PATH '$**.url',
                         file_size VARCHAR(255) PATH '$**.file_size',
                         name VARCHAR(255) PATH '$**.name'
                         )
              ) AS sizes);

##############################################
# Inserts for Media Sizes
##############################################

INSERT INTO media_sizes (`file_id`, `media_id`, `uuid`, `width`, `height`, `crop`, `size_name`, `size_key`)
    (SELECT 0, media.id, sizes.uuid, sizes.width, sizes.height, sizes.crop, sizes.size_name,
            CASE
                WHEN sizes.size_name = 'HD Size' THEN 'hd'
                WHEN sizes.size_name = 'Large Size' THEN 'large'
                WHEN sizes.size_name = 'Medium Size' THEN 'medium'
                WHEN sizes.size_name = 'Thumbnail Size' THEN 'thumbnail'
                ELSE 'Default Size' END AS keyname
     FROM media,
          JSON_TABLE(sizes, '$.*'
                     COLUMNS (
                         width INT PATH '$**.width',
                         height INT PATH '$**.height',
                         size_name VARCHAR(255) PATH '$**.size_name',
                         crop INT PATH '$**.crop',
                         uuid VARCHAR(255) PATH '$**.uuid'
                         )
              ) AS sizes);

##############################################
# MediaTable Alter
##############################################

ALTER TABLE media ADD COLUMN (`file_id` INT NOT NULL);

##############################################
# File ID Inserts
##############################################

UPDATE `media` INNER JOIN `files` ON media.uuid = files.uuid SET media.file_id = files.id WHERE media.uuid = files.uuid;
UPDATE `media_sizes` INNER JOIN `files` ON media_sizes.uuid = files.uuid SET media_sizes.file_id = files.id WHERE media_sizes.uuid = files.uuid;

##############################################
# Drop Columns
##############################################

ALTER TABLE media
    DROP COLUMN uuid,
    DROP COLUMN url,
    DROP COLUMN file_path,
    DROP COLUMN file_size,
    DROP COLUMN file_name,
    DROP COLUMN sizes,
    DROP COLUMN mime;

ALTER TABLE media_sizes
    DROP COLUMN uuid;

##############################################
# Add To Options
##############################################

INSERT INTO options(`option_name`, `option_value`) VALUES('storage_provider', '\"local\"');
INSERT INTO options(`option_name`, `option_value`) VALUES('storage_bucket', '\"\"')

