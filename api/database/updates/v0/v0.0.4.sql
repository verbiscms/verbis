CREATE TABLE `storage`
(
    `id`          int           NOT NULL AUTO_INCREMENT,
    `uuid`        varchar(36)   NOT NULL,
    `provider`    varchar(255)  NOT NULL,
    `region`      varchar(255)  NOT NULL,
    `bucket`      varchar(255)  NOT NULL,
    `url`         text          NOT NULL,
    `name`        varchar(255)  NOT NULL,
    `path`        text          NOT NULL,
    `private`     bit DEFAULT 0 NOT NULL,
    `file_size`   int           NOT NULL,
    `source_type` varchar(255)  NOT NULL,
    `mime`        varchar(150)  NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO storage (`uuid`, `url`, `file_size`, `name`, `mime`, `path`, `region`, `provider`, `bucket`, `source_type`)
SELECT `uuid`, `url`, `file_size`, `file_name`, `mime`, `file_path`, '', 'local',  '', 'media'
FROM media;

INSERT INTO storage (`uuid`, `url`, `file_size`, `name`, `mime`, `path`, `region`, `provider`, `bucket`, `source_type`)
    (SELECT sizes.uuid, sizes.url, sizes.file_size, sizes.name, media.mime, media.file_path, '', 'local', '', 'media'
     FROM media,
          json_table(sizes, '$.*'
                     COLUMNS (
                         uuid VARCHAR(255) PATH '$**.uuid',
                         url VARCHAR(255) PATH '$**.url',
                         file_size VARCHAR(255) PATH '$**.file_size',
                         name VARCHAR(255) PATH '$**.name'
                         )
              ) AS sizes)
