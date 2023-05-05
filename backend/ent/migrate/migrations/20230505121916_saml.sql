-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `token`, ADD COLUMN `sso_identifier` varchar(255) NULL, ADD COLUMN `last_login` timestamp NULL;
