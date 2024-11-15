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
-- Table structure for table `commentinfo`
--

DROP TABLE IF EXISTS `commentinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commentinfo` (
  `commentId` int NOT NULL AUTO_INCREMENT,
  `noteID` int DEFAULT NULL,
  `commentatorId` bigint DEFAULT NULL COMMENT '评论者账号',
  `content` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '评论内容',
  `commentTime` datetime NOT NULL COMMENT '评论时间',
  `state` int DEFAULT '0' COMMENT '是否已读',
  PRIMARY KEY (`commentId`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commentinfo`
--

LOCK TABLES `commentinfo` WRITE;
/*!40000 ALTER TABLE `commentinfo` DISABLE KEYS */;
INSERT INTO `commentinfo` VALUES (5,43,11,'????????','2023-07-06 08:56:14',0),(6,48,12,'确实好吃','2023-07-06 09:31:16',0),(7,51,13,'好米好米','2023-07-06 09:49:38',0),(8,52,12,'嘎嘎想 @Lyx ','2023-07-06 09:58:20',0),(9,52,12,'嘎嘎想 @Lyx ','2023-07-06 10:08:36',0),(10,44,12,'嘎嘎帮 @Lyx ','2023-07-06 10:17:18',0),(11,49,12,'卡卡 @Lyx ','2023-07-06 10:57:49',0),(12,48,11,'呼呼??','2023-07-06 13:06:15',0);
/*!40000 ALTER TABLE `commentinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-06 23:18:40
