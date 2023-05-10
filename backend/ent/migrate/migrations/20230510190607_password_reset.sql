-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `email_confirmed` bool NOT NULL DEFAULT true;
-- Create "tokens" table
CREATE TABLE `tokens` (`id` bigint NOT NULL AUTO_INCREMENT, `token` varchar(255) NOT NULL, `action` varchar(255) NOT NULL, `valid_until` timestamp NULL, `send` bool NOT NULL DEFAULT false, `created_at` timestamp NULL, `data` json NOT NULL, `user_tokens` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `tokens_users_tokens` FOREIGN KEY (`user_tokens`) REFERENCES `users` (`id`) ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
