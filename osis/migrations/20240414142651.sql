-- Create "orgs" table
CREATE TABLE `orgs` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL
);
-- Create index "idx_orgs_deleted_at" to table: "orgs"
CREATE INDEX `idx_orgs_deleted_at` ON `orgs` (`deleted_at`);
-- Create "users" table
CREATE TABLE `users` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `name` text NULL,
  `email` text NULL,
  `phone` text NULL,
  `org_id` integer NULL,
  CONSTRAINT `fk_users_org` FOREIGN KEY (`org_id`) REFERENCES `orgs` (`id`) ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX `idx_users_deleted_at` ON `users` (`deleted_at`);
-- Create "accounts" table
CREATE TABLE `accounts` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  CONSTRAINT `fk_accounts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX `idx_accounts_deleted_at` ON `accounts` (`deleted_at`);
-- Create "sessions" table
CREATE TABLE `sessions` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `account_id` integer NULL,
  CONSTRAINT `fk_sessions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `fk_sessions_account` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_sessions_deleted_at" to table: "sessions"
CREATE INDEX `idx_sessions_deleted_at` ON `sessions` (`deleted_at`);
-- Create "tokens" table
CREATE TABLE `tokens` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `session_id` integer NULL,
  `access_token` text NULL,
  `refresh_token` text NULL,
  `expires_at` datetime NULL,
  CONSTRAINT `fk_tokens_session` FOREIGN KEY (`session_id`) REFERENCES `sessions` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_tokens_deleted_at" to table: "tokens"
CREATE INDEX `idx_tokens_deleted_at` ON `tokens` (`deleted_at`);
