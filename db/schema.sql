SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema gorani
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Table `word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `word` (
  `word_id` INT NOT NULL AUTO_INCREMENT,
  `word` VARCHAR(255) NOT NULL,
  `word_pronunciation` VARCHAR(255) NULL,
  PRIMARY KEY (`word_id`),
  UNIQUE INDEX `uq_word_idx` (`word` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `relevant_word_type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `relevant_word_type` (
  `relevant_word_type_code` INT NOT NULL AUTO_INCREMENT,
  `relevant_word_type_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`relevant_word_type_code`),
  UNIQUE INDEX `relevant_word_type_name_UNIQUE` (`relevant_word_type_name` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `relevant_word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `relevant_word` (
  `word_id` INT NOT NULL,
  `target_word_id` INT NOT NULL,
  `relevant_word_type_code` INT NOT NULL,
  `relevant_word_score` INT NOT NULL,
  `relevant_word_vote_sum` INT NOT NULL,
  PRIMARY KEY (`word_id`, `target_word_id`),
  INDEX `fk_relevant_word2_idx` (`target_word_id` ASC),
  INDEX `fk_relevant_word3_idx` (`relevant_word_type_code` ASC),
  CONSTRAINT `fk_relevant_word`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relevant_word2`
    FOREIGN KEY (`target_word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relevant_word3`
    FOREIGN KEY (`relevant_word_type_code`)
    REFERENCES `relevant_word_type` (`relevant_word_type_code`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `book`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `book` (
  `book_isbn` VARCHAR(13) NOT NULL,
  `book_name` VARCHAR(255) NOT NULL,
  `book_author` VARCHAR(255) NULL,
  `book_cover_image` VARCHAR(255) NULL,
  PRIMARY KEY (`book_isbn`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sentence`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sentence` (
  `sentence_id` INT NOT NULL AUTO_INCREMENT,
  `sentence` VARCHAR(1024) NOT NULL,
  `book_isbn` VARCHAR(13) NULL,
  PRIMARY KEY (`sentence_id`),
  INDEX `fk_sentence_idx` (`book_isbn` ASC),
  CONSTRAINT `fk_sentence`
    FOREIGN KEY (`book_isbn`)
    REFERENCES `book` (`book_isbn`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `word_sentence`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `word_sentence` (
  `sentence_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `word_position` INT NOT NULL,
  PRIMARY KEY (`sentence_id`, `word_id`, `word_position`),
  CONSTRAINT `fk_word_sentence`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_word_sentence2`
    FOREIGN KEY (`sentence_id`)
    REFERENCES `sentence` (`sentence_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `definition`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `definition` (
  `definition_id` INT NOT NULL AUTO_INCREMENT,
  `word_id` INT NOT NULL,
  `definition` VARCHAR(1024) NOT NULL,
  `definition_pos` ENUM('verb', 'aux', 'tverb', 'iverb', 'noun', 'adj', 'adv', 'abr', 'prep', 'symbol', 'pronoun', 'conj', 'suffix', 'prefix', 'det') NULL,
  PRIMARY KEY (`definition_id`),
  CONSTRAINT `fk_definition`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `known_word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `known_word` (
  `user_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `known_word_number` INT NOT NULL,
  PRIMARY KEY (`user_id`, `word_id`),
  INDEX `fk_known_word2_idx` (`word_id` ASC),
  CONSTRAINT `fk_known_word`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_known_word2`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_detail`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_detail` (
  `user_id` INT NOT NULL,
  `user_profile_image` VARCHAR(256) NOT NULL,
  `user_added_date` DATETIME NOT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_detail`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `oauth_service`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `oauth_service` (
  `oauth_service_code` INT NOT NULL,
  `oauth_service_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`oauth_service_code`),
  UNIQUE INDEX `oauth_service_name_UNIQUE` (`oauth_service_name` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `oauth_passport`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `oauth_passport` (
  `user_id` INT NOT NULL,
  `oauth_service_code` INT NOT NULL,
  `oauth_user_id` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`user_id`, `oauth_service_code`),
  INDEX `fk_oauth_passport2_idx` (`oauth_service_code` ASC),
  CONSTRAINT `fk_oauth_passport`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_oauth_passport2`
    FOREIGN KEY (`oauth_service_code`)
    REFERENCES `oauth_service` (`oauth_service_code`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `book_rating`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `book_rating` (
  `book_isbn` VARCHAR(13) NOT NULL,
  `book_rating_provider` VARCHAR(255) NOT NULL,
  `rating` DOUBLE NOT NULL,
  PRIMARY KEY (`book_isbn`, `book_rating_provider`),
  CONSTRAINT `fk_book_rating`
    FOREIGN KEY (`book_isbn`)
    REFERENCES `book` (`book_isbn`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `unknown_word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `unknown_word` (
  `user_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `unknown_word_added_date` DATETIME NOT NULL,
  `unknown_word_memory_sentence` VARCHAR(1024) NULL,
  INDEX `fk_unknown_word_idx` (`user_id` ASC),
  INDEX `fk_unknown_word2_idx` (`word_id` ASC),
  PRIMARY KEY (`user_id`, `word_id`),
  CONSTRAINT `fk_unknown_word`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_unknown_word2`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `wordbook_entries_update_date`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `wordbook_entries_update_date` (
  `wordbook_uuid` BINARY(16) NOT NULL,
  `wordbook_entry_update_date` DATETIME NOT NULL,
  PRIMARY KEY (`wordbook_uuid`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `example`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `example` (
  `definition_id` INT NOT NULL,
  `foreign` VARCHAR(255) NOT NULL,
  `native` VARCHAR(255) NULL,
  INDEX `fk_example_idx` (`definition_id` ASC),
  CONSTRAINT `fk_example`
    FOREIGN KEY (`definition_id`)
    REFERENCES `definition` (`definition_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `relevant_word_vote`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `relevant_word_vote` (
  `word_id` INT NOT NULL,
  `target_word_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `relevant_word_type_code` INT NOT NULL,
  PRIMARY KEY (`word_id`, `target_word_id`),
  INDEX `fk_relevant_word_vote2_idx` (`target_word_id` ASC),
  INDEX `fk_relevant_word_vote3_idx` (`relevant_word_type_code` ASC),
  CONSTRAINT `fk_relevant_word_vote`
    FOREIGN KEY (`word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relevant_word_vote2`
    FOREIGN KEY (`target_word_id`)
    REFERENCES `word` (`word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relevant_word_vote3`
    FOREIGN KEY (`relevant_word_type_code`)
    REFERENCES `relevant_word_type` (`relevant_word_type_code`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `genre`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `genre` (
  `genre_code` INT NOT NULL AUTO_INCREMENT,
  `genre_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`genre_code`),
  UNIQUE INDEX `genre_name_UNIQUE` (`genre_name` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `book_genre`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `book_genre` (
  `book_isbn` VARCHAR(13) NOT NULL,
  `genre_code` INT NOT NULL,
  PRIMARY KEY (`book_isbn`, `genre_code`),
  INDEX `fk_book_genre2_idx` (`genre_code` ASC),
  CONSTRAINT `fk_book_genre`
    FOREIGN KEY (`book_isbn`)
    REFERENCES `book` (`book_isbn`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_book_genre2`
    FOREIGN KEY (`genre_code`)
    REFERENCES `genre` (`genre_code`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `book_review`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `book_review` (
  `book_isbn` VARCHAR(13) NOT NULL,
  `book_review_provider` VARCHAR(45) NOT NULL,
  `review` TEXT NOT NULL,
  INDEX `fk_book_review_idx` (`book_isbn` ASC),
  CONSTRAINT `fk_book_review`
    FOREIGN KEY (`book_isbn`)
    REFERENCES `book` (`book_isbn`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `quiz_unknown_word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `quiz_unknown_word` (
  `user_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `quiz_unknown_word_grade` INT NOT NULL,
  `quiz_unknown_word_date` DATETIME NOT NULL,
  PRIMARY KEY (`user_id`, `word_id`),
  CONSTRAINT `fk_unknown_word_quiz`
    FOREIGN KEY (`user_id` , `word_id`)
    REFERENCES `unknown_word` (`user_id` , `word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `suggested_quiz_unknown_word`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `suggested_quiz_unknown_word` (
  `user_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `suggested_quiz_type_code` INT NOT NULL,
  `suggested_quiz_start_date` DATETIME NOT NULL,
  `suggested_quiz_ef` DOUBLE NOT NULL,
  `suggested_quiz_recent_interval` INT NOT NULL,
  `suggested_quiz_repetition_number` INT NOT NULL,
  PRIMARY KEY (`user_id`, `word_id`, `suggested_quiz_type_code`),
  CONSTRAINT `fk_suggested_quiz_unknown_word`
    FOREIGN KEY (`user_id` , `word_id`)
    REFERENCES `unknown_word` (`user_id` , `word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `unknown_word_source`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `unknown_word_source` (
  `user_id` INT NOT NULL,
  `word_id` INT NOT NULL,
  `definition_id` INT NOT NULL,
  `unknown_word_source_book` VARCHAR(45) NULL,
  `unknown_word_source_sentence` VARCHAR(1024) NULL,
  `unknown_word_source_word_index` INT NULL,
  INDEX `fk_unknown_word_source_idx` (`user_id` ASC, `word_id` ASC),
  INDEX `fk_unknown_word_source2_idx` (`definition_id` ASC),
  CONSTRAINT `fk_unknown_word_source`
    FOREIGN KEY (`user_id` , `word_id`)
    REFERENCES `unknown_word` (`user_id` , `word_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_unknown_word_source2`
    FOREIGN KEY (`definition_id`)
    REFERENCES `definition` (`definition_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_prefer_genre`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_prefer_genre` (
  `user_id` INT NOT NULL,
  `genre_code` INT NOT NULL,
  PRIMARY KEY (`user_id`, `genre_code`),
  INDEX `fk_user_prefer_genre2_idx` (`genre_code` ASC),
  CONSTRAINT `fk_user_prefer_genre`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_prefer_genre2`
    FOREIGN KEY (`genre_code`)
    REFERENCES `genre` (`genre_code`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
