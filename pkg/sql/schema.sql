CREATE TABLE `user` (
  `id` BINARY(16) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `role` text NOT NULL,
  `is_active` boolean NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime
);

CREATE TABLE `protocol` (
  `id` BINARY(16) PRIMARY KEY AUTO_INCREMENT,
  `diagnosis` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL
);

CREATE TABLE `action` (
  `id` BINARY(16) PRIMARY KEY AUTO_INCREMENT,
  `protocol_id` BINARY(16),
  `drug` varchar(255) NOT NULL,
  `dosage` text,
  `time` text,
  `duration` text,
  `description` text
);

ALTER TABLE `action` ADD FOREIGN KEY (`protocol_id`) REFERENCES `protocol` (`id`);
