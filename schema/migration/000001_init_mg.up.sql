CREATE TABLE `adminReq` (
  `reqId` int(11) NOT NULL AUTO_INCREMENT,
  `userId` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 0,
)

CREATE TABLE `books_record` (
  `bookId` int(11) NOT NULL AUTO_INCREMENT,
  `bookName` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `copies` int(11) NOT NULL DEFAULT 0,
)

CREATE TABLE `cookie` (
  `sessionId` varchar(255) NOT NULL,
  `userId` int(11) NOT NULL DEFAULT 0,
)

CREATE TABLE `requests` (
  `reqId` int(11) NOT NULL AUTO_INCREMENT,
  `bookId` int(11) NOT NULL,
  `userId` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 0,
)

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `admin` tinyint(1) NOT NULL DEFAULT 0,
) 
