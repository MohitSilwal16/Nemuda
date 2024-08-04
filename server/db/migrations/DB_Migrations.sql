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

-- Dumping data for table nemuda.messages_aayush: ~0 rows (approximately)
DELETE FROM `messages_aayush`;
INSERT INTO `messages_aayush` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Aayush', 'Hey Aayush, all good here. What about you?', 'Read', '2024-07-08 10:16:12'),
	('Nimesh', 'Aayush', 'e', 'Sent', '2024-07-13 23:28:18');

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

-- Dumping data for table nemuda.messages_konark: ~8 rows (approximately)
DELETE FROM `messages_konark`;
INSERT INTO `messages_konark` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Konark', 'Yo Kon', 'Read', '2024-07-07 20:32:25'),
	('Nimesh', 'Konark', 'Nothing', 'Read', '2024-07-07 20:38:41'),
	('Palash', 'Konark', 'Hey Konark, todo bien por aquí. ¿Y tú?', 'Read', '2024-07-08 10:19:45'),
	('Nimesh', 'Konark', 'Wbu ?', 'Read', '2024-07-08 19:10:25'),
	('Nimesh', 'Konark', 'Hello', 'Read', '2024-07-15 16:53:37'),
	('Nimesh', 'Konark', 'Hello', 'Sent', '2024-07-16 21:25:54'),
	('Nimesh', 'Konark', 'Yo', 'Sent', '2024-07-16 21:25:55'),
	('Nimesh', 'Konark', 'Wyd ?', 'Sent', '2024-07-16 21:25:58');

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
	('Nimesh', 'Krish', 'Hello', 'Sent', '2024-07-13 19:42:54'),
	('Nimesh', 'Krish', 'Yo', 'Sent', '2024-07-14 16:01:42');

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

-- Dumping data for table nemuda.messages_nimesh: ~38 rows (approximately)
DELETE FROM `messages_nimesh`;
INSERT INTO `messages_nimesh` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Konark', 'Nimesh', 'Yep Nim', 'Read', '2024-07-07 20:32:32'),
	('Konark', 'Nimesh', 'Wyd Nimesh', 'Read', '2024-07-07 20:37:18'),
	('Konark', 'Nimesh', 'Nimesh ??', 'Read', '2024-07-07 20:37:24'),
	('Aayush', 'Nimesh', 'Hello Nimesh! How\'s it going?', 'Read', '2024-07-08 10:15:00'),
	('Prachin', 'Nimesh', 'Hi Nimesh, what\'s up?', 'Read', '2024-07-08 10:21:10'),
	('Konark', 'Nimesh', 'Nothing', 'Read', '2024-07-08 19:10:42'),
	('Konark', 'Nimesh', 'Wyd ?', 'Read', '2024-07-15 16:53:50'),
	('Parth', 'Nimesh', 'Msg 2 P', 'Read', '2024-07-16 21:36:46'),
	('Parth', 'Nimesh', 'Msg 3 P', 'Read', '2024-07-16 21:36:51'),
	('Parth', 'Nimesh', 'Msg 4 P', 'Read', '2024-07-16 21:36:54'),
	('Parth', 'Nimesh', 'Msg 7 P', 'Read', '2024-07-16 21:37:03'),
	('Parth', 'Nimesh', 'Msg 9 P', 'Read', '2024-07-16 21:37:11'),
	('Parth', 'Nimesh', 'Msg 11 P', 'Read', '2024-07-16 21:37:34'),
	('Parth', 'Nimesh', 'Msg 12 P', 'Read', '2024-07-16 21:37:36'),
	('Parth', 'Nimesh', 'Msg 17 N', 'Read', '2024-07-16 21:39:29'),
	('Parth', 'Nimesh', 'Msg 18 P', 'Read', '2024-07-16 21:39:39'),
	('Parth', 'Nimesh', 'Msg 19 P', 'Read', '2024-07-16 21:39:42'),
	('Parth', 'Nimesh', 'Msg 20 P', 'Read', '2024-07-16 21:39:45'),
	('Parth', 'Nimesh', 'Msg 21 P', 'Read', '2024-07-16 21:39:49'),
	('Parth', 'Nimesh', 'Msg 23 P', 'Read', '2024-07-16 21:54:11'),
	('Parth', 'Nimesh', 'Msg 25 P', 'Read', '2024-07-16 21:54:20'),
	('Parth', 'Nimesh', 'Msg 27 P', 'Read', '2024-07-16 21:54:36'),
	('Parth', 'Nimesh', 'Msg 31 P', 'Read', '2024-07-16 21:54:57'),
	('Parth', 'Nimesh', 'Msg 32 P', 'Read', '2024-07-16 21:55:34'),
	('Parth', 'Nimesh', 'Msg 35 P', 'Read', '2024-07-16 22:12:57');

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

-- Dumping data for table nemuda.messages_palash: ~0 rows (approximately)
DELETE FROM `messages_palash`;
INSERT INTO `messages_palash` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Konark', 'Palash', 'Hola Palash! Qué tal?', 'Read', '2024-07-08 10:18:30'),
	('Nimesh', 'Palash', 'Aiyo', 'Sent', '2024-07-13 20:29:04');

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

-- Dumping data for table nemuda.messages_parth: ~17 rows (approximately)
DELETE FROM `messages_parth`;
INSERT INTO `messages_parth` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Parth', 'Msg 1', 'Read', '2024-07-16 21:36:34'),
	('Nimesh', 'Parth', 'Msg 5 N', 'Read', '2024-07-16 21:36:57'),
	('Nimesh', 'Parth', 'Msg 6 N', 'Read', '2024-07-16 21:36:59'),
	('Nimesh', 'Parth', 'Msg 8 N', 'Read', '2024-07-16 21:37:07'),
	('Nimesh', 'Parth', 'Msg 10N', 'Read', '2024-07-16 21:37:15'),
	('Nimesh', 'Parth', 'Msg 13 N', 'Read', '2024-07-16 21:39:14'),
	('Nimesh', 'Parth', 'Msg 14 N', 'Read', '2024-07-16 21:39:17'),
	('Nimesh', 'Parth', 'Msg 15 N', 'Read', '2024-07-16 21:39:19'),
	('Nimesh', 'Parth', 'Msg 16 N', 'Read', '2024-07-16 21:39:21'),
	('Nimesh', 'Parth', 'Msg 22 N', 'Read', '2024-07-16 21:54:06'),
	('Nimesh', 'Parth', 'Msg 24 N', 'Read', '2024-07-16 21:54:15'),
	('Nimesh', 'Parth', 'Msg 26 N', 'Read', '2024-07-16 21:54:31'),
	('Nimesh', 'Parth', 'Msg 28 N', 'Read', '2024-07-16 21:54:40'),
	('Nimesh', 'Parth', 'Msg 29 N', 'Read', '2024-07-16 21:54:44'),
	('Nimesh', 'Parth', 'Msg 30 N', 'Read', '2024-07-16 21:54:47'),
	('Nimesh', 'Parth', 'Msg 33 N', 'Read', '2024-07-16 22:04:26'),
	('Nimesh', 'Parth', 'Msg 34 N', 'Read', '2024-07-16 22:06:42'),
	('Nimesh', 'Parth', 'Msg 36 N', 'Read', '2024-07-16 22:23:44');

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

-- Dumping structure for table nemuda.messages_rudra
CREATE TABLE IF NOT EXISTS `messages_rudra` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` = 'Rudra'),
  `MessageContent` varchar(100) NOT NULL CHECK (`MessageContent` <> '' and char_length(`MessageContent`) <= 100),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_rudra_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_rudra_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.messages_rudra: ~0 rows (approximately)
DELETE FROM `messages_rudra`;

-- Dumping structure for table nemuda.users
CREATE TABLE IF NOT EXISTS `users` (
  `Username` varchar(20) NOT NULL CHECK (`Username` <> ''),
  `Password` varchar(20) NOT NULL CHECK (`Password` <> ''),
  `Token` char(8) NOT NULL,
  PRIMARY KEY (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table nemuda.users: ~9 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`Username`, `Password`, `Token`) VALUES
	('Aayush', 'Aayush1@', 'eedaea78'),
	('Konark', 'Konark1@', ''),
	('Krish', 'Krish12@', ''),
	('Nimesh', 'Nimesh1@', '041a1766'),
	('Palash', 'Palash1@', '8639d8c8'),
	('Parth', 'Parth12@', ''),
	('Prachin', 'Prachin1@', ''),
	('Robin', 'Robin12@', '9361b56a'),
	('Rudra', 'Rudra12@', '5526d9ff');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
