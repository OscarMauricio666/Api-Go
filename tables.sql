CREATE TABLE tickets(
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    user VARCHAR(66) NOT NULL,
    status VARCHAR(66) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);