
(SELECT *
FROM messages_parth
WHERE Sender = "Nimesh"
UNION
SELECT *
FROM messages_nimesh
WHERE Sender = "Parth"
ORDER BY DATETIME DESC
LIMIT 4 OFFSET 0
)
ORDER BY DATETIME;

SELECT *
FROM messages_konark
WHERE Sender = "Konark"

DELETE
FROM messages_nimesh
WHERE Sender = "Nimesh"

SELECT *
FROM messages_parth
WHERE Sender = "Nimesh"
UNION
SELECT *
FROM messages_nimesh
WHERE Sender = "Parth"
ORDER BY DATETIME DESC;


DELETE
FROM messages_nimesh
WHERE Sender <> "Parth";


DESCRIBE users;

SELECT *
FROM messages_aayush
WHERE Sender = "Nimesh"
UNION
SELECT *
FROM messages_nimesh
WHERE Sender = "Aayush"
ORDER BY DATETIME DESC;

DELETE
FROM messages_aayush
WHERE Sender = "Nimesh";

DELETE
FROM messages_nimesh
WHERE Sender = "Aayush";

SHOW TABLES;

SELECT *
FROM messages_aayush;

SELECT *
FROM messages_konark;

SELECT *
FROM messages_krish;

SELECT *
FROM messages_legend;

SELECT *
FROM messages_nimesh;

SELECT *
FROM messages_palash;

SELECT *
FROM messages_parth;

SELECT *
FROM messages_prachin;

SELECT *
FROM messages_robin;

SELECT *
FROM messages_rudra;
