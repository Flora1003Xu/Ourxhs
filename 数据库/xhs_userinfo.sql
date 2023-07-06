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
-- Table structure for table `userinfo`
--

DROP TABLE IF EXISTS `userinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userinfo` (
  `userAccount` bigint unsigned NOT NULL AUTO_INCREMENT,
  `userName` varchar(50) NOT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `portrait` varchar(255) NOT NULL COMMENT '用户头像',
  `introduction` varchar(100) DEFAULT NULL COMMENT '用户简介',
  `birthday` date DEFAULT NULL,
  `registTime` datetime NOT NULL COMMENT '注册时间',
  `fansNum` int DEFAULT '0' COMMENT '粉丝数量',
  `noteNum` int DEFAULT '0' COMMENT '笔记数',
  `collectNum` int DEFAULT '0' COMMENT '收藏数量',
  `followNum` int DEFAULT '0' COMMENT '关注数',
  `collectedNum` int DEFAULT '0' COMMENT '被收藏数量',
  `likedNum` int DEFAULT '0' COMMENT '被点赞的数量',
  `phoneNumber` varchar(15) DEFAULT NULL COMMENT '手机号',
  `mail` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`userAccount`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userinfo`
--

LOCK TABLES `userinfo` WRITE;
/*!40000 ALTER TABLE `userinfo` DISABLE KEYS */;
INSERT INTO `userinfo` VALUES (10,'shx','男','head_18055310723_1688602316_AF43867C940F7B5127995A1C0700491F.jpg','shx的简介','2003-02-04','2023-07-06 08:11:56',1,18,0,0,4,5,'18055310723',NULL,'123456'),(11,'Lyx','女','head_13809843393_1688604897_727498b4b453ba158c2350acc8f8e56.jpg','一个低调的人','2004-01-11','2023-07-06 08:54:57',1,7,0,2,1,3,'13809843393',NULL,'lin85305912'),(12,'yzt','女','head_14749948452_1688606781_微信图片_20230706092352.jpg','hahaha','2023-07-04','2023-07-06 09:26:21',2,5,0,1,2,2,'14749948452',NULL,'123456'),(13,'Avenir.','女','head_18922223203_1688607975_微信图片_202307060945114.jpg','早上坏早上坏早上坏','2023-07-03','2023-07-06 09:46:15',0,2,0,1,0,0,'18922223203',NULL,'123456'),(18,'yzt','女','head_14749948457_1688612198_微信图片_20230706092352.jpg','haha ','2023-07-03','2023-07-06 10:56:38',0,0,0,0,0,0,'14749948457',NULL,'123456');
/*!40000 ALTER TABLE `userinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-06 23:18:39
