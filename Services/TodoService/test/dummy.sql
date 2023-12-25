/*
    This file contains list of SQL queries to run in order to get some dummy data in database.
    Ensure the database has been created and tables migrated before running these commands
*/

/* Add records into user table */
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `firstname`, `lastname`)
    VALUES (1, '2023-12-24 11:50:54.688', '2023-12-24 11:50:54.688', null, 'johndoe', 'John', 'Doe');
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `firstname`, `lastname`)
    VALUES (2, '2023-12-25 11:50:54.688', '2023-12-25 11:50:54.688', null, 'ammylee', 'Ammy', 'Lee');

/* Add records into todos table */
INSERT INTO `todos`(`id`, `created_at`, `updated_at`, `deleted_at`, `priority`, `title`, `description`, `status`, `user_id`)
    VALUES (1, '2023-12-24 11:50:54.688', '2023-12-24 11:50:54.688', null, 1, 'First Todo', 'This is the first todo', 'open', 1);
INSERT INTO `todos`(`id`, `created_at`, `updated_at`, `deleted_at`, `priority`, `title`, `description`, `status`, `user_id`)
    VALUES (2, '2023-12-20 11:50:54.688', '2023-12-24 11:50:54.688', null, 2, 'Second Todo', 'This is the second todo', 'in_progress', 1);
INSERT INTO `todos`(`id`, `created_at`, `updated_at`, `deleted_at`, `priority`, `title`, `description`, `status`, `user_id`)
    VALUES (3, '2023-12-22 11:50:54.688', '2023-12-24 11:50:54.688', null, 3, 'Third Todo', 'This is the third todo', 'done', 1);
INSERT INTO `todos`(`id`, `created_at`, `updated_at`, `deleted_at`, `priority`, `title`, `description`, `status`, `user_id`)
    VALUES (4, '2023-12-21 11:50:54.688', '2023-12-24 11:50:54.688', null, 4, 'Fourth Todo', 'This is the fourth todo', 'cancelled', 2);
INSERT INTO `todos`(`id`, `created_at`, `updated_at`, `deleted_at`, `priority`, `title`, `description`, `status`, `user_id`)
    VALUES (5, '2023-12-20 11:50:54.688', '2023-12-20 11:50:54.688', null, 5, 'Fifth Todo', 'This is the fifth todo', 'open', 2);
