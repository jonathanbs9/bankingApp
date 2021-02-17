CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
	`customer_id` int(11) NOT NULL AUTO_INCREMENT,
	`first_name` varchar(20) NOT NULL,
	`last_name` varchar(30) NOT NULL, 
 	`date_birth` date NOT NULL,
	`city` varchar(100) NOT NULL,
	`zipcode` varchar(6) NOT NULL, 
	`status` tinyint(1) NOT NULL DEFAULT '1',
	PRIMARY KEY (`customer_id`)

 ) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
	(9001, 'Jonathan', 'Brull Schroeder', '1986-06-23', 'Mar del Plata', '7600', 1),
	(9002, 'Rafael', 'Santos Borre', '1990-05-23', 'Cali', '12345', 1),
	(9003, 'Marcelo', 'Gallardo', '1978-10-10', 'Buenos Aires', '7100', 1),
	(9004, 'Leonardo', 'Ponzio', '1980-09-01','Santa Fe', '8967', 1);


DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts`(
	`account_id` int(11) NOT NULL AUTO_INCREMENT,
	`customer_id` int(11) NOT NULL, 
	`opening_date`datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`account_type` varchar(10) NOT NULL,
	`pin` varchar(10) NOT NULL,
	`status` tinyint(4) NOT NULL DEFAULT '1',
	PRIMARY KEY (`account_id`),
	KEY `accounts_FK` (`customer_id`),
	CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES customers(`customer_id`) ) ENGINE=InnoDB AUTO_INCREMENT=95476 DEFAULT CHARSET=latin1;


INSERT INTO `accounts` VALUES 
	(95472, 9004, '2021-02-10', 'Checking', '1234', 1),
	(95473, 9001, '2021-02-09', 'Saving', '0000', 1),
	(95474, 9002, '2021-01-01', 'Checking', '9876', 1),
	(95475, 9003, '2002-05-25', 'Saving', '1111', 0);

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
	`transaction_id` int(11) NOT NULL AUTO_INCREMENT,
	`account_id` int(11) NOT NULL,
	`amount` int(11) NOT NULL,
	`transaction_type` varchar(10) NOT NULL,
	`transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`transaction_id`),
	KEY `transactions_FK` (`account_id`),
	CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=INNODB DEFAULT CHARSET=latin1;