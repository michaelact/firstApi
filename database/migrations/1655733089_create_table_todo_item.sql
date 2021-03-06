CREATE TABLE Todo (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        activity_group_id INT(6) UNSIGNED NOT NULL,
        title VARCHAR(30) NOT NULL,
        is_active TINYINT(1) UNSIGNED DEFAULT '1' NOT NULL,
        priority VARCHAR(30) DEFAULT 'very-high' NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL, 
        deleted_at TIMESTAMP, 
        FOREIGN KEY (activity_group_id) REFERENCES Activity(id)
);
