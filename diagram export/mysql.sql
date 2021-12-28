CREATE TABLE `accounts` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `owner` varchar(199) NOT NULL,
  `balance` decimal NOT NULL,
  `currency` varchar(5) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `entries` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `account_id` bigint NOT NULL,
  `amount` decimal NOT NULL COMMENT 'can be negative or positive',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `transfers` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` bigint NOT NULL,
  `to_account_id` bigint NOT NULL,
  `amount` decimal NOT NULL COMMENT 'must be positive',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `users` (
  `username` varchar(199) PRIMARY KEY,
  `hashed_password` varchar(199) NOT NULL,
  `full_name` varchar(199) NOT NULL,
  `email` varchar(50) UNIQUE NOT NULL,
  `password_changed_at` timestamp DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE UNIQUE INDEX `accounts_index_1` ON `accounts` (`owner`, `currency`);

CREATE INDEX `entries_index_2` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_5` ON `transfers` (`from_account_id`, `to_account_id`);
