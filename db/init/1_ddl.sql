SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;

SET
    @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS,
    FOREIGN_KEY_CHECKS = 0;

SET
    @OLD_SQL_MODE = @@SQL_MODE,
    SQL_MODE = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `yubisuma_api` DEFAULT CHARACTER SET utf8mb4;

USE `yubisuma_api`;

SET CHARSET utf8mb4;

-- こちらが新しく作成するanimalsテーブルの定義

CREATE TABLE
    IF NOT EXISTS `yubisuma_api`.`animals` (
        `id` VARCHAR(128) NOT NULL COMMENT '動物ID (UNK)',
        `name` VARCHAR(64) NOT NULL COMMENT '動物名',
        `numFinger` INT NOT NULL COMMENT '指の数',
        `skill` VARCHAR(64) NOT NULL COMMENT 'スキル',
        `skillDesc` VARCHAR(512) NOT NULL COMMENT 'スキル', 
        `handUrl` VARCHAR(256) NOT NULL COMMENT '手の画像URL',
        `upFingerUrl` VARCHAR(256) NOT NULL COMMENT '上向きの指の画像URL',
        `downFingerUrl` VARCHAR(256) NOT NULL COMMENT '下向きの指の画像URL',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB COMMENT = '動物';



SET SQL_MODE=@OLD_SQL_MODE;

SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;

SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;