-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               11.3.2-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for nemuda
CREATE DATABASE IF NOT EXISTS `nemuda` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `nemuda`;

-- Dumping structure for table nemuda.messages_aayush
CREATE TABLE IF NOT EXISTS `messages_aayush` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Aayush'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_aayush_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_aayush_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_aayush: ~1 rows (approximately)
DELETE FROM `messages_aayush`;
INSERT INTO `messages_aayush` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Aayush', 'Hey Aayush, all good here. What about you?', 'Read', '2024-07-08 10:16:12');

-- Dumping structure for table nemuda.messages_konark
CREATE TABLE IF NOT EXISTS `messages_konark` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Konark'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_konark_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_konark_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_konark: ~63 rows (approximately)
DELETE FROM `messages_konark`;
INSERT INTO `messages_konark` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Konark', 'Yo Kon', 'Read', '2024-07-07 20:32:25'),
	('Nimesh', 'Konark', 'Nothing', 'Read', '2024-07-07 20:38:41'),
	('Palash', 'Konark', 'Hey Konark, todo bien por aquí. ¿Y tú?', 'Read', '2024-07-08 10:19:45'),
	('Nimesh', 'Konark', 'Wbu ?', 'Read', '2024-07-08 19:10:25');

-- Dumping structure for table nemuda.messages_krish
CREATE TABLE IF NOT EXISTS `messages_krish` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Krish'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_krish_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_krish_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_krish: ~2 rows (approximately)
DELETE FROM `messages_krish`;
INSERT INTO `messages_krish` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Krish', 'Yo', 'Sent', '2024-07-13 18:18:20'),
	('Nimesh', 'Krish', 'Hello', 'Sent', '2024-07-13 19:42:54');

-- Dumping structure for table nemuda.messages_nimesh
CREATE TABLE IF NOT EXISTS `messages_nimesh` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Nimesh'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_nimesh_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_nimesh_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_nimesh: ~34 rows (approximately)
DELETE FROM `messages_nimesh`;
INSERT INTO `messages_nimesh` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Konark', 'Nimesh', 'Yep Nim', 'Read', '2024-07-07 20:32:32'),
	('Konark', 'Nimesh', 'Wyd Nimesh', 'Read', '2024-07-07 20:37:18'),
	('Konark', 'Nimesh', 'Nimesh ??', 'Read', '2024-07-07 20:37:24'),
	('Aayush', 'Nimesh', 'Hello Nimesh! How\'s it going?', 'Read', '2024-07-08 10:15:00'),
	('Prachin', 'Nimesh', 'Hi Nimesh, what\'s up?', 'Read', '2024-07-08 10:21:10');

-- Dumping structure for table nemuda.messages_palash
CREATE TABLE IF NOT EXISTS `messages_palash` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Palash'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_palash_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_palash_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_palash: ~1 rows (approximately)
DELETE FROM `messages_palash`;
INSERT INTO `messages_palash` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Konark', 'Palash', 'Hola Palash! Qué tal?', 'Read', '2024-07-08 10:18:30');

-- Dumping structure for table nemuda.messages_parth
CREATE TABLE IF NOT EXISTS `messages_parth` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Parth'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_parth_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_parth_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_parth: ~0 rows (approximately)
DELETE FROM `messages_parth`;

-- Dumping structure for table nemuda.messages_prachin
CREATE TABLE IF NOT EXISTS `messages_prachin` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Prachin'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_prachin_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_prachin_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_prachin: ~1 rows (approximately)
DELETE FROM `messages_prachin`;
INSERT INTO `messages_prachin` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Prachin', 'Hey Prachin, just chilling. You?', 'Read', '2024-07-08 10:22:27');

-- Dumping structure for table nemuda.messages_robin
CREATE TABLE IF NOT EXISTS `messages_robin` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Robin'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_robin_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_robin_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_robin: ~0 rows (approximately)
DELETE FROM `messages_robin`;

-- Dumping structure for table nemuda.users
CREATE TABLE IF NOT EXISTS `users` (
  `Username` varchar(20) NOT NULL CHECK (`Username` <> ''),
  `Password` varchar(20) NOT NULL CHECK (`Password` <> ''),
  `Token` char(8) NOT NULL,
  PRIMARY KEY (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.users: ~8 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`Username`, `Password`, `Token`) VALUES
	('Aayush', 'Aayush1@', '7152f812'),
	('Konark', 'Konark1@', '3ca42477'),
	('Krish', 'Krish12@', '117c49b6'),
	('Nimesh', 'Nimesh1@', '6df8eb2d'),
	('Palash', 'Palash1@', 'a5ec9dc1'),
	('Parth', 'Parth12@', '47e92144'),
	('Prachin', 'Prachin1@', '4a97f180'),
	('Robin', 'Robin12@', 'a8f68aeb');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
