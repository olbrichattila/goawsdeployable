CREATE TABLE status_updates (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `instruction_id` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `message` MEDIUMTEXT
);
