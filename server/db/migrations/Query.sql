USE Nemuda;

CREATE TABLE Users(
	Username VARCHAR(20) PRIMARY KEY CHECK (Username <> ""),
	Password VARCHAR(20) NOT NULL CHECK (Password <> ""),
	Token CHAR(8) UNIQUE NOT NULL
);

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

SELECT *
FROM users;

SELECT *
FROM Messages;

SELECT *
FROM Messages
WHERE
	(Sender = "Nimesh" AND Receiver = "Konark" )
OR
	(Sender = "Konark" AND Receiver = "Nimesh" );
	
SELECT *
FROM Messages
WHERE
	(Sender = "Palash" AND Receiver = "Konark" )
OR
	(Sender = "Konark" AND Receiver = "Palash" );


INSERT INTO Messages (Sender, Receiver, MessageContent, Status, DateTime)
VALUES 
    ('Nimesh', 'Konark', 'Yo Kon', 'Sent', '2024-07-07 20:32:25'),
    ('Konark', 'Nimesh', 'Yep Nim', 'Sent', '2024-07-07 20:32:32'),
    ('Konark', 'Nimesh', 'Wyd Nimesh', 'Sent', '2024-07-07 20:37:18'),
    ('Konark', 'Nimesh', 'Nimesh ??', 'Sent', '2024-07-07 20:37:24'),
    ('Nimesh', 'Konark', 'Nothing', 'Sent', '2024-07-07 20:38:41');


INSERT INTO Messages (Sender, Receiver, MessageContent, Status, DateTime)
VALUES 
    ('Aayush', 'Nimesh', 'Hello Nimesh! How\'s it going?', 'Sent', '2024-07-08 10:15:00'),
    ('Nimesh', 'Aayush', 'Hey Aayush, all good here. What about you?', 'Sent', '2024-07-08 10:16:12'),
    ('Konark', 'Palash', 'Hola Palash! Qué tal?', 'Sent', '2024-07-08 10:18:30'),
    ('Palash', 'Konark', 'Hey Konark, todo bien por aquí. ¿Y tú?', 'Sent', '2024-07-08 10:19:45'),
    ('Prachin', 'Nimesh', 'Hi Nimesh, what\'s up?', 'Sent', '2024-07-08 10:21:10'),
    ('Nimesh', 'Prachin', 'Hey Prachin, just chilling. You?', 'Sent', '2024-07-08 10:22:27');
    

SHOW TABLES;

DESCRIBE messages;

DELETE
FROM messages
WHERE DATETIME LIKE '2024-07-09%';

SELECT *
FROM messages_konark;

SELECT *
FROM messages_nimesh;

DELETE
FROM messages_nimesh
WHERE DATETIME LIKE '2024-07-14%';

SELECT *
FROM messages;

UPDATE messages
SET STATUS = "Read";

ALTER TABLE `messages`
MODIFY `MessageContent` varchar(100) NOT NULL;

-- Step 2: Add a new constraint to ensure message length is less than or equal to 100
ALTER TABLE `messages`
ADD CONSTRAINT `message_length_check` CHECK (CHAR_LENGTH(`MessageContent`) <= 100);

DESCRIBE messages;

SELECT *
FROM Messages_Nimesh
WHERE Sender = 'Konark'
UNION
SELECT *
FROM Messages_Konark
WHERE Sender = 'Nimesh'
ORDER BY DATETIME;

DELIMITER $$

DROP PROCEDURE IF EXISTS Create_Messages_Table $$

CREATE PROCEDURE Create_Messages_Table()
BEGIN
    DECLARE done INT DEFAULT 1;
    DECLARE user VARCHAR(20);
    DECLARE tableName VARCHAR(29);
    DECLARE createTableSQL VARCHAR(1000);
    
    DECLARE myCursor CURSOR FOR SELECT Username FROM Users;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 0;
    
    OPEN myCursor;
    	WHILE done DO
        FETCH myCursor INTO user;
        
        SET @tableName = CONCAT('Messages_', user);
        
        SELECT @tableName;
			
			SET @sql = CONCAT('CREATE TABLE IF NOT EXISTS ', @tableName, ' (
			    Sender VARCHAR(20) NOT NULL CHECK (Sender <> ''''),
			    Receiver VARCHAR(20) NOT NULL CHECK (Receiver = ''', user, '''),
			    MessageContent VARCHAR(100) NOT NULL CHECK (MessageContent <> '''' AND CHAR_LENGTH(MessageContent) <= 100),
			    Status VARCHAR(10) NOT NULL CHECK (Status IN (''Sent'', ''Delivered'', ''Read'')),
			    DateTime DATETIME NOT NULL,
			    FOREIGN KEY (Sender) REFERENCES users (Username),
			    FOREIGN KEY (Receiver) REFERENCES users (Username)
			)');
			
			PREPARE stmt FROM @sql;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;        
    END WHILE;
    CLOSE myCursor;
END $$

DELIMITER ;

CALL Create_Messages_Table;

DROP PROCEDURE Create_Messages_Table;

SHOW TABLES;

DELIMITER $$

DROP PROCEDURE IF EXISTS Data_Migration $$

CREATE PROCEDURE Data_Migration()
DETERMINISTIC
BEGIN
    DECLARE done INT DEFAULT 1;
    DECLARE sender_1 VARCHAR(20);
    DECLARE receiver_1 VARCHAR(20);
    DECLARE messageContent_1 VARCHAR(100);
    DECLARE status_1 VARCHAR(10);
    DECLARE date_time_1 DATETIME;
    DECLARE tableName VARCHAR(29);
    DECLARE createTableSQL VARCHAR(1000);
    DECLARE insertSQL VARCHAR(1000);
    
    DECLARE myCursor CURSOR FOR SELECT Sender, Receiver, MessageContent, Status, DateTime FROM messages;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 0;
    
    OPEN myCursor;
    	WHILE done DO
        FETCH myCursor INTO sender_1, receiver_1, messageContent_1, status_1, date_time_1;
        
        SET tableName = CONCAT('Messages_', receiver_1);
        
        SELECT sender_1, receiver_1, messageContent_1, status_1, date_time_1, tableName;
        
        SET @tableName = CONCAT('Messages_', receiver_1);
			
			SET @sql = CONCAT('CREATE TABLE IF NOT EXISTS ', @tableName, ' (
			    Sender VARCHAR(20) NOT NULL CHECK (Sender <> ''''),
			    Receiver VARCHAR(20) NOT NULL CHECK (Receiver = ''', receiver_1, '''),
			    MessageContent VARCHAR(100) NOT NULL CHECK (MessageContent <> '''' AND CHAR_LENGTH(MessageContent) <= 100),
			    Status VARCHAR(10) NOT NULL CHECK (Status IN (''Sent'', ''Delivered'', ''Read'')),
			    DateTime DATETIME NOT NULL,
			    FOREIGN KEY (Sender) REFERENCES users (Username),
			    FOREIGN KEY (Receiver) REFERENCES users (Username)
			)');
			
			PREPARE stmt FROM @sql;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;

        -- Prepare the dynamic SQL to insert the data
		   SET @insertSQL = CONCAT(
		   'INSERT INTO ', @tableName, ' (Sender, Receiver, MessageContent, Status, DateTime) ',
		   'VALUES ("', sender_1, '", "', receiver_1, '", "', messageContent_1, '", "', status_1, '", "', date_time_1, '")'
			);

        
        -- Execute the dynamic SQL to insert the data
        PREPARE stmt FROM @insertSQL;
        EXECUTE stmt;
        DEALLOCATE PREPARE stmt;
        
    END WHILE;
    CLOSE myCursor;
END $$

DELIMITER ;

CALL Data_Migration;

SHOW TABLES;
