ALTER TABLE `users` ADD `color` INT NOT NULL DEFAULT '0' AFTER `stories_max_id`, ADD `color_background_emoji_id` BIGINT NOT NULL DEFAULT '0' AFTER `color`, ADD `profile_color` INT NOT NULL DEFAULT '0' AFTER `color_background_emoji_id`, ADD `profile_color_background_emoji_id` BIGINT NOT NULL DEFAULT '0' AFTER `profile_color`;
