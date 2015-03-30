CREATE TABLE IF NOT EXISTS `config` (
    `id` INTEGER PRIMARY KEY NOT NULL,
    `key` TEXT NOT NULL,
    `value` TEXT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS `config_ix1` ON `config`(`key`);

CREATE TABLE IF NOT EXISTS `feed` (
    `id` INTEGER PRIMARY KEY NOT NULL,
    `title` TEXT NOT NULL,
    `url` TEXT NOT NULL,
    `feed` TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `post` (
    `id` INTEGER PRIMARY KEY NOT NULL,
    `feed_id` INTEGER NOT NULL,
    `title` TEXT NOT NULL,
    `url` TEXT NOT NULL,
    `published` INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS `post_ix1` ON `post`(`published`);

