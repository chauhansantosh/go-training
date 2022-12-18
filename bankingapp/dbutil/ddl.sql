CREATE DATABASE IF NOT EXISTS bankdb;

CREATE TABLE `customer` (
  `customer_id` int unsigned NOT NULL,
  `customer_name` text NOT NULL,
  `customer_type` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `customer_pan` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `pan_UNIQUE` (`customer_pan`)
);

CREATE TABLE `bank_account` (
  `account_id` int unsigned NOT NULL,
  `customer_id` int unsigned NOT NULL,
  `balance` float unsigned NOT NULL,
  `account_type` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `account_pan` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`account_id`),
  KEY `fk_customer` (`customer_id`)
);

CREATE DEFINER=`root`@`localhost` TRIGGER `bank_account_BEFORE_INSERT` BEFORE INSERT ON `bank_account` FOR EACH ROW BEGIN
	IF EXISTS (SELECT 1 FROM bank_account WHERE customer_id = NEW.customer_id AND NEW.account_type in ('SAVINGS','CURRENT')) THEN 
        SELECT CONCAT(NEW.customer_id, ' customer already have a ', NEW.account_type, ' account. Can not have more than one.') INTO @error_text; 
		SIGNAL SQLSTATE '45000' SET message_text = @error_text; 
	END IF;
END

CREATE DEFINER=`root`@`localhost` TRIGGER `bank_account_BEFORE_UPDATE` BEFORE UPDATE ON `bank_account` FOR EACH ROW BEGIN
IF (NEW.balance - OLD.balance > 50000 AND OLD.account_type='SAVINGS' AND (NEW.account_pan IS NULL OR NEW.account_pan = '')) THEN
     SELECT CONCAT('Pan number is mandatory for depositing more than 50000.') INTO @error_text; 
	 SIGNAL SQLSTATE '46000' SET message_text = @error_text; 
elseif (NEW.balance - OLD.balance > 250000 AND OLD.account_type='CURRENT' AND (NEW.account_pan IS NULL OR NEW.account_pan = '')) THEN
     SELECT CONCAT('Pan number is mandatory for depositing more than 250000.') INTO @error_text; 
	 SIGNAL SQLSTATE '46000' SET message_text = @error_text; 
END IF;
END

CREATE TABLE `transaction` (
  `transaction_id` int unsigned NOT NULL AUTO_INCREMENT,
  `account_id` int unsigned DEFAULT NULL,
  `customer_id` int unsigned DEFAULT NULL,
  `opening_balance` float DEFAULT NULL,
  `amount` float DEFAULT NULL,
  `new_balance` float DEFAULT NULL,
  `transaction_type` char(6) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  KEY `fk_customer` (`customer_id`),
  KEY `fk_account` (`account_id`),
  CONSTRAINT `fk_account` FOREIGN KEY (`account_id`) REFERENCES `bank_account` (`account_id`),
  CONSTRAINT `fk_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`)
);