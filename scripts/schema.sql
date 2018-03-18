CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    CONSTRAINT uc_username UNIQUE (username),
    CONSTRAINT uc_email UNIQUE (email)
);

CREATE TABLE words (
    id INT PRIMARY KEY AUTO_INCREMENT,
    word VARCHAR(255) NOT NULL,
    pron VARCHAR(255) DEFAULT NULL,
    source VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL,
    CONSTRAINT uc_source UNIQUE (source)
);

CREATE TABLE defs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    word_id INT NOT NULL,
    part VARCHAR(255) DEFAULT NULL,
    def VARCHAR(255) NOT NULL,
    FOREIGN KEY (word_id)
        REFERENCES words(id)
        ON DELETE CASCADE
);

CREATE TABLE examples (
    id INT PRIMARY KEY AUTO_INCREMENT,
    def_id INT NOT NULL,
    kor VARCHAR(255) DEFAULT NULL,
    eng VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (def_id)
        REFERENCES defs(id)
        ON DELETE CASCADE
);

CREATE TABLE wordbooks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    seen_at timestamp NOT NULL DEFAULT current_timestamp,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    name VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT uc_user_name UNIQUE (user_id, name)
);

CREATE TABLE wordbook_entries (
    wordbook_id INT NOT NULL,
    sr_no INT DEFAULT NULL,
    def_id INT NOT NULL,
    star BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (def_id)
        REFERENCES defs(id)
        ON DELETE CASCADE,
    FOREIGN KEY (wordbook_id)
        REFERENCES wordbooks(id)
        ON DELETE CASCADE,
    CONSTRAINT uc_wordbook_def UNIQUE (wordbook_id, def_id)
);

delimiter $$
CREATE TRIGGER wordbooks_composite_auto BEFORE INSERT ON wordbook_entries
FOR EACH ROW BEGIN
    SET NEW.sr_no = (
       SELECT IFNULL(MAX(sr_no), 0) + 1
       FROM wordbook_entries
       WHERE wordbook_id  = NEW.wordbook_id
    );
END;$$
delimiter ;

CREATE VIEW defs_of_wordbooks AS 
SELECT wordbook_id, sr_no, defs.id as def_id, star, def, part FROM wordbook_entries
INNER JOIN defs ON defs.id = wordbook_entries.def_id;