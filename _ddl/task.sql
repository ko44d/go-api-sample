CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`   VARCHAR(20)  NOT NULL COMMENT 'タスクの状態',
    `created`  VARCHAR(6)   NOT NULL COMMENT 'レコード作成日時',
    `modified` VARCHAR(6)   NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';