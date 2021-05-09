DROP DATABASE IF EXISTS prac;
CREATE DATABASE IF NOT EXISTS prac;
use prac;

DROP TABLE IF EXISTS `member`;


CREATE TABLE IF NOT EXISTS `member`
(
 `id`               INT(20) AUTO_INCREMENT,
 `last_name`        VARCHAR(20) NOT NULL,
 `first_name`        VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO member (last_name,first_name) VALUES ('tanaka','tarou');
INSERT INTO member (last_name,first_name) VALUES ('satou','kazuo');
INSERT INTO member (last_name,first_name) VALUES ('itou','kenta');