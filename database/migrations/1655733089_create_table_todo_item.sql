CREATE TABLE Todo (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        activity_group_id INT(6) UNSIGNED NOT NULL,
        is_active INT(6) NOT NULL,
        priority VARCHAR(30) NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
        deleted_at TIMESTAMP, 
        FOREIGN KEY (activity_group_id) REFERENCES Activity(id)
);
