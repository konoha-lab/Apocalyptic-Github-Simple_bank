CREATE TABLE `accounts` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `owner` varchar(255) NOT NULL,
  `balance` decimal NOT NULL,
  `currency` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL 
);

CREATE TABLE `entries` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `account_id` bigint NOT NULL,
  `amount` decimal NOT NULL COMMENT 'can be negative or positive',
  `created_at` timestamp NOT NULL 
);

CREATE TABLE `transfers` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` bigint NOT NULL,
  `to_account_id` bigint NOT NULL,
  `amount` decimal NOT NULL COMMENT 'must be positive',
  `created_at` timestamp NOT NULL 
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `hashed_password` varchar(191),
  `full_name` varchar(191),
  `email` varchar(191) NOT NULL,
  `password_changed_at` timestamp DEFAULT NULL,
  `created_at` timestamp NOT NULL 
);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);
