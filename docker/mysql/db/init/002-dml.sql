DELETE FROM `rooms`;
INSERT INTO `rooms` (`id`, `title`) VALUES ("00000000000000000000000001", "ルームA");
INSERT INTO `rooms` (`id`, `title`) VALUES ("00000000000000000000000002", "ルームB");
INSERT INTO `rooms` (`id`, `title`) VALUES ("00000000000000000000000003", "ルームC");

DELETE FROM `messages`;
INSERT INTO `messages` (`id`, `room_id`, `body`) VALUES ("00000000000000000000000001", "00000000000000000000000001", "ルームAのメッセージ1");
INSERT INTO `messages` (`id`, `room_id`, `body`) VALUES ("00000000000000000000000002", "00000000000000000000000001", "ルームAのメッセージ2");
INSERT INTO `messages` (`id`, `room_id`, `body`) VALUES ("00000000000000000000000003", "00000000000000000000000001", "ルームAのメッセージ3");
