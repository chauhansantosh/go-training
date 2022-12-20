CREATE DATABASE IF NOT EXISTS bankdb;

use bankdb;

CREATE TABLE IF NOT EXISTS `customer` (
  `customer_id` int unsigned NOT NULL,
  `customer_name` text NOT NULL,
  `customer_type` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `customer_pan` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `pan_UNIQUE` (`customer_pan`),
  CONSTRAINT `chk_acc_type` CHECK ((`customer_type` in ('INDIVIDUAL','COMPANY')))
);

CREATE TABLE IF NOT EXISTS `bank_account` (
 `account_id` int unsigned NOT NULL,
  `customer_id` int unsigned NOT NULL,
  `balance` float unsigned NOT NULL,
  `account_type` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `account_pan` varchar(10) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `is_locked` tinyint(1) DEFAULT '0',
  `lock_period_fd` int unsigned DEFAULT NULL,
  `penalty_fd` decimal(10,3) DEFAULT NULL,
  `locked_until` datetime DEFAULT NULL,
  PRIMARY KEY (`account_id`),
  KEY `fk_customer` (`customer_id`),
  CONSTRAINT `chk_account_type` CHECK ((`account_type` in ('SAVINGS','FIXED','CURRENT')))
);

#--------------------------------------------------- BEFORE INSERT --------------------------------------------------------

DROP TRIGGER IF EXISTS `bankdb`.`bank_account_BEFORE_INSERT`;

DELIMITER $$
USE `bankdb`$$
CREATE DEFINER=`root`@`localhost` TRIGGER `bank_account_BEFORE_INSERT` BEFORE INSERT ON `bank_account` FOR EACH ROW BEGIN
	IF EXISTS (SELECT 1 FROM bank_account WHERE customer_id = NEW.customer_id AND NEW.account_type in ('SAVINGS','CURRENT')) THEN 
        SELECT CONCAT(NEW.customer_id, ' customer already have a ', NEW.account_type, ' account. Can not have more than one.') INTO @error_text; 
		SIGNAL SQLSTATE '45000' SET message_text = @error_text; 
	END IF;
END;$$
DELIMITER ;

DROP TRIGGER IF EXISTS `bankdb`.`account_limit_before_insert`;

DELIMITER $$
USE `bankdb`$$
CREATE DEFINER=`root`@`localhost` TRIGGER `account_limit_before_insert` BEFORE INSERT ON `bank_account` FOR EACH ROW BEGIN
	if NEW.account_type = 'SAVINGS' and NEW.balance > 10000000 then
SELECT CONCAT('Balance limit reached. Can not have more than 10 millions in SAVINGS account.') INTO @error_text; 
		SIGNAL SQLSTATE '48000' SET message_text = @error_text; 
end if;
END;$$
DELIMITER ;

DROP TRIGGER IF EXISTS `bankdb`.`account_customer_must_exists`;

DELIMITER $$
USE `bankdb`$$
CREATE DEFINER=`root`@`localhost` TRIGGER `account_customer_must_exists` BEFORE INSERT ON `bank_account` FOR EACH ROW BEGIN
	IF NOT EXISTS (SELECT 1 FROM customer WHERE customer_id = NEW.customer_id) THEN 
    SELECT CONCAT('Customer ', NEW.customer_id, ' does not exist. Create customer before creating account.') INTO @error_text; 
		SIGNAL SQLSTATE '47000' SET message_text = @error_text; 
	END IF;
END;$$
DELIMITER ;

#--------------------------------------------------------- BEFORE UPDATE--------------------------------------------------------

DROP TRIGGER IF EXISTS `bankdb`.`bank_account_BEFORE_UPDATE`;

DELIMITER $$
USE `bankdb`$$
CREATE DEFINER=`root`@`localhost` TRIGGER `bank_account_BEFORE_UPDATE` BEFORE UPDATE ON `bank_account` FOR EACH ROW BEGIN
IF (NEW.balance - OLD.balance > 50000 AND OLD.account_type='SAVINGS' AND (NEW.account_pan IS NULL OR NEW.account_pan = '')) THEN
     SELECT CONCAT('Pan number is mandatory for depositing more than 50000.') INTO @error_text; 
	 SIGNAL SQLSTATE '46000' SET message_text = @error_text; 
elseif (NEW.balance - OLD.balance > 250000 AND OLD.account_type='CURRENT' AND (NEW.account_pan IS NULL OR NEW.account_pan = '')) THEN
     SELECT CONCAT('Pan number is mandatory for depositing more than 250000.') INTO @error_text; 
	 SIGNAL SQLSTATE '46000' SET message_text = @error_text; 
END IF;
END$$
DELIMITER ;


DROP TRIGGER IF EXISTS `bankdb`.`account_limit_before_update`;

DELIMITER $$
USE `bankdb`$$
CREATE DEFINER=`root`@`localhost` TRIGGER `account_limit_before_update` BEFORE UPDATE ON `bank_account` FOR EACH ROW BEGIN
if OLD.account_type = 'SAVINGS' and NEW.balance > 10000000 then
SELECT CONCAT('Balance limit reached. Can not have more than 10 millions in SAVINGS account.') INTO @error_text; 
		SIGNAL SQLSTATE '48000' SET message_text = @error_text; 
end if;
END;$$
DELIMITER ;

#-------------------------------------------------------------------------------


CREATE TABLE IF NOT EXISTS `transaction` (
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
