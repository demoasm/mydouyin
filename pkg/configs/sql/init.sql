
CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password', 
    `follow_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User follow count' ,
    `follower_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User follower count' ,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author`      bigint unsigned NOT NULL COMMENT 'author id',
    `play_url`    varchar(128) NOT NULL DEFAULT '' COMMENT 'video play url',
    `cover_url`   varchar(128) NOT NULL DEFAULT '' COMMENT 'vidoe cover url',
    `title`       varchar(128) NOT NULL DEFAULT '' COMMENT 'video title',
    `favorite_count` int unsigned Not NULL DEFAULT 0 COMMENT 'video favorite count' ,
    `comment_count` int unsigned Not NULL DEFAULT 0 COMMENT 'video comment count' ,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'vidoe update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_author_id` (`author`) COMMENT 'Author id index',
    CONSTRAINT   `author_id` FOREIGN KEY (`author`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video information table';

CREATE TABLE `favorite`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`   bigint unsigned NOT NULL COMMENT 'user id',
    `video_id`   bigint unsigned NOT NULL COMMENT 'video id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'vidoe update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id_video_id` (`user_id`, `video_id`) COMMENT 'User id and Video id index',
    CONSTRAINT   `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT   `video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite information table';