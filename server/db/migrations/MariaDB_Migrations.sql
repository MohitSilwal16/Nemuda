-- Dumping database structure for nemuda
CREATE DATABASE IF NOT EXISTS `nemuda` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `nemuda`;

-- Dumping structure for table nemuda.users
CREATE TABLE IF NOT EXISTS `users` (
  `Username` varchar(20) NOT NULL CHECK (`Username` <> ''),
  `Password` varchar(20) NOT NULL CHECK (`Password` <> ''),
  `Token` char(8) NOT NULL,
  PRIMARY KEY (`Username`)
);

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