# Create Storage Table
# TODO add no duplicates for url and name
# Index for url for faster lookups
CREATE TABLE `files`
(
    `id`          int           NOT NULL AUTO_INCREMENT,
    `bucket_id`   text          NOT NULL,
    `uuid`        varchar(36)   NOT NULL,
    `url`         text          NOT NULL,
    `name`        varchar(255)  NOT NULL,
    `path`        text          NOT NULL,
    `mime`        varchar(150)  NOT NULL,
    `source_type` varchar(255)  NOT NULL,
    `provider`    varchar(255)  NOT NULL,
    `region`      varchar(255)  NOT NULL,
    `bucket`      varchar(255)  NOT NULL,
    `file_size`   int           NOT NULL,
    `private`     bit DEFAULT 0 NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

# Create Media Sizes Table
CREATE TABLE `media_sizes`
(
    `id`        int           NOT NULL AUTO_INCREMENT,
    `file_id`   int           NOT NULL,
    `media_id`  int           NOT NULL,
    `size_key`  varchar(255)  NOT NULL,
    `size_name` varchar(255)  NOT NULL,
    `width`     int           NOT NULL,
    `height`    int           NOT NULL,
    `crop`      bit DEFAULT 0 NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

# Inserts for Storage
INSERT INTO storage (`uuid`, `url`, `file_size`, `name`, `mime`, `path`, `region`, `provider`, `bucket`, `source_type`)
SELECT `uuid`,
       `url`,
       `file_size`,
       `file_name`,
       `mime`,
       `file_path`,
       '''',
       '' local'', '''',
       '' media ''
FROM media;

INSERT INTO storage (`uuid`, `url`, `file_size`, `name`, `mime`, `path`, `region`, `provider`, `bucket`, `source_type`)
    (SELECT sizes.uuid,
            sizes.url,
            sizes.file_size,
            sizes.name,
            media.mime,
            media.file_path,
            '''',
            '' local'', '''',
            '' media ''
     FROM media,
          json_table(sizes, '' $.* ''
                     COLUMNS(
                             uuid VARCHAR(255) PATH ''$**.uuid'',
                             url VARCHAR(255) PATH ''$**.url'',
                             file_size VARCHAR(255) PATH ''$**.file_size'',
                             name VARCHAR(255) PATH ''$**.name''
                         )
              ) AS sizes);

# Inserts for Media
ALTER table media
    ADD COLUMN `storage_id` int NOT NULL DEFAULT 0;

UPDATE media m
    JOIN storage s on m.uuid = s.uuid
SET m.file_id = s.id
WHERE m.uuid = s.uuid;

# Inserts for Media Sizes
# INSERT INTO media_sizes (`file_id`, `media_id`, `name`, `width`, `height`, `crop`)
#     (SELECT storage.id, m.id, sizes.width, sizes.height, sizes.sizename, crop
#     FROM media AS m,
#          json_table(sizes, ''$.*''
#                 COLUMNS (
#                     width VARCHAR(255) PATH ''$**.width'',
#                     height VARCHAR(255) PATH ''$**.height'',
#                     sizename VARCHAR(255) PATH ''$**.size_name'',
#                     crop VARCHAR(255) PATH ''$**.crop''
#                 )
#              ) AS sizes);

SELECT media.id,
       sizes.width,
       sizes.height,
       sizes.size_name,
       sizes.crop,
       CASE
           WHEN sizes.size_name = ''HD Size'' THEN ''hd''
           WHEN sizes.size_name = ''Large Size'' THEN ''large''
           WHEN sizes.size_name = ''Medium Size'' THEN ''medium''
           WHEN sizes.size_name = ''Thumbnail Size'' THEN ''thumbnail''
           ELSE '''' END AS keyname
FROM media AS media,
     JSON_TABLE(sizes, '' $.* ''
                COLUMNS(
                        width INT PATH ''$**.width'',
                        height INT PATH ''$**.height'',
                        size_name VARCHAR(255) PATH ''$**.size_name'',
                        crop INT PATH ''$**.crop''
                    )
         ) AS sizes;



ALTER TABLE media
    DROP COLUMN uuid,
    DROP COLUMN url,
    DROP COLUMN file_path,
    DROP COLUMN file_size,
    DROP COLUMN file_name,
    DROP COLUMN sizes,
    DROP COLUMN mime;

ALTER TABLE media
    DROP COLUMN uuid;
