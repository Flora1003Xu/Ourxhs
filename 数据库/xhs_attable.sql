-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: xhs
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `attable`
--

DROP TABLE IF EXISTS `attable`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `attable` (
  `atId` int NOT NULL AUTO_INCREMENT,
  `userAct` varchar(50) NOT NULL COMMENT '用户帐号',
  `noteId` int DEFAULT NULL,
  `atUserName` varchar(20) DEFAULT NULL COMMENT '@的用户的名字',
  `atLocation` int DEFAULT NULL COMMENT '@符号在正文中的位置',
  `atState` int DEFAULT '1' COMMENT '@在评论还是正文，1表示正文',
  PRIMARY KEY (`atId`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attable`
--

LOCK TABLES `attable` WRITE;
/*!40000 ALTER TABLE `attable` DISABLE KEYS */;
INSERT INTO `attable` VALUES (12,'10',48,'shx',8,1),(16,'0',6,'',0,0),(17,'0',7,'',0,0),(18,'0',8,'Lyx',3,0),(19,'0',9,'Lyx',3,0),(20,'11',55,'Lyx',5,1),(21,'11',55,'Lyx',5,1),(22,'11',55,'Lyx',6,1),(23,'0',10,'Lyx',3,0),(24,'11',56,'Lyx',1,1),(25,'0',11,'Lyx',2,0),(26,'11',57,'Lyx',8,1),(27,'0',12,'',0,0),(28,'12',58,'yzt',10,1);
/*!40000 ALTER TABLE `attable` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-06 23:18:38
