-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `age` bigint NOT NULL, `name` varchar(255) NOT NULL DEFAULT 'unknown', PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "ymirs" table
CREATE TABLE `ymirs` (`id` bigint NOT NULL AUTO_INCREMENT, `version` varchar(255) NOT NULL DEFAULT 'alpha-test-dev1', PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "cars" table
CREATE TABLE `cars` (`id` bigint NOT NULL AUTO_INCREMENT, `model` varchar(255) NOT NULL, `registered_at` timestamp NULL, `user_cars` bigint NULL, PRIMARY KEY (`id`), INDEX `cars_users_cars` (`user_cars`), CONSTRAINT `cars_users_cars` FOREIGN KEY (`user_cars`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "groups" table
CREATE TABLE `groups` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "group_users" table
CREATE TABLE `group_users` (`group_id` bigint NOT NULL, `user_id` bigint NOT NULL, PRIMARY KEY (`group_id`, `user_id`), INDEX `group_users_user_id` (`user_id`), CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
