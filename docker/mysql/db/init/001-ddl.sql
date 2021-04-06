DROP TABLE IF EXISTS `messages`;
DROP TABLE IF EXISTS `rooms`;

CREATE TABLE IF NOT EXISTS `rooms` (
  `id`    CHAR(26)    NOT NULL        COMMENT 'トークルームULID',
  `title` VARCHAR(64) NOT NULL UNIQUE COMMENT 'トークルーム名',
  PRIMARY KEY (`id`))
ENGINE=InnoDB
COMMENT='トークルーム'
DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `messages` (
  `id`      CHAR(26)    NOT NULL        COMMENT 'メッセージULID',
  `room_id` CHAR(26)    NOT NULL        COMMENT 'トークルームULID',
  `body`    VARCHAR(64) NOT NULL        COMMENT 'メッセージ本文',
  PRIMARY KEY (`id`),
  CONSTRAINT fk_room_id
    FOREIGN KEY (`room_id`)
    REFERENCES rooms (`id`)
    ON DELETE CASCADE)
ENGINE=InnoDB
COMMENT='メッセージ'
DEFAULT CHARSET=utf8mb4;
