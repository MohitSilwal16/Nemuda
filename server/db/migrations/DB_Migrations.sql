-- Dumping database structure for nemuda
CREATE DATABASE IF NOT EXISTS `nemuda` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `nemuda`;

-- Dumping structure for table nemuda.messages
CREATE TABLE IF NOT EXISTS `messages` (
  `Sender` varchar(20) NOT NULL CHECK (`Sender` <> ''),
  `Receiver` varchar(20) NOT NULL CHECK (`Receiver` <> ''),
  `MessageContent` varchar(50) NOT NULL CHECK (`MessageContent` <> ''),
  `Status` varchar(10) NOT NULL CHECK (`Status` in ('Sent','Delivered','Read')),
  `DateTime` datetime NOT NULL,
  KEY `Sender` (`Sender`),
  KEY `Receiver` (`Receiver`),
  CONSTRAINT `messages_ibfk_1` FOREIGN KEY (`Sender`) REFERENCES `users` (`Username`),
  CONSTRAINT `messages_ibfk_2` FOREIGN KEY (`Receiver`) REFERENCES `users` (`Username`)
);

-- Dumping data for table nemuda.messages: ~11 rows (approximately)
DELETE FROM `messages`;
INSERT INTO `messages` (`Sender`, `Receiver`, `MessageContent`, `Status`, `DateTime`) VALUES
	('Nimesh', 'Konark', 'Yo Kon', 'Sent', '2024-07-07 20:32:25'),
	('Konark', 'Nimesh', 'Yep Nim', 'Sent', '2024-07-07 20:32:32'),
	('Konark', 'Nimesh', 'Wyd Nimesh', 'Sent', '2024-07-07 20:37:18'),
	('Konark', 'Nimesh', 'Nimesh ??', 'Sent', '2024-07-07 20:37:24'),
	('Nimesh', 'Konark', 'Nothing', 'Sent', '2024-07-07 20:38:41'),
	('Aayush', 'Nimesh', 'Hello Nimesh! How\'s it going?', 'Sent', '2024-07-08 10:15:00'),
	('Nimesh', 'Aayush', 'Hey Aayush, all good here. What about you?', 'Sent', '2024-07-08 10:16:12'),
	('Konark', 'Palash', 'Hola Palash! Qué tal?', 'Sent', '2024-07-08 10:18:30'),
	('Palash', 'Konark', 'Hey Konark, todo bien por aquí. ¿Y tú?', 'Sent', '2024-07-08 10:19:45'),
	('Prachin', 'Nimesh', 'Hi Nimesh, what\'s up?', 'Sent', '2024-07-08 10:21:10'),
	('Nimesh', 'Prachin', 'Hey Prachin, just chilling. You?', 'Sent', '2024-07-08 10:22:27');

-- Dumping structure for table nemuda.users
CREATE TABLE IF NOT EXISTS `users` (
  `Username` varchar(20) NOT NULL CHECK (`Username` <> ''),
  `Password` varchar(20) NOT NULL CHECK (`Password` <> ''),
  `Token` char(8) NOT NULL,
  PRIMARY KEY (`Username`)
) ;

-- Dumping data for table nemuda.users: ~7 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`Username`, `Password`, `Token`) VALUES
	('Aayush', 'Aayush1@', 'd9ac1d95'),
	('Konark', 'Konark1@', ''),
	('Nimesh', 'Nimesh1@', 'e4b86ffd'),
	('Palash', 'Palash1@', 'a5ec9dc1'),
	('Parth', 'Parth12@', '702d736d'),
	('Prachin', 'Prachin1@', '4a97f180'),
	('Robin', 'Robin12@', 'a8f68aeb');
