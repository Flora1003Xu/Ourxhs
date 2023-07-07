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
-- Table structure for table `noteinfo`
--

DROP TABLE IF EXISTS `noteinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `noteinfo` (
  `noteId` int unsigned NOT NULL AUTO_INCREMENT,
  `creatorAccount` bigint DEFAULT NULL,
  `title` varchar(100) NOT NULL,
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '正文',
  `numOfPic` int DEFAULT NULL COMMENT '图片数量',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面图片',
  `createTime` datetime NOT NULL COMMENT '创建时间',
  `updateTime` datetime NOT NULL COMMENT '最近更新',
  `tag` varchar(20) DEFAULT NULL,
  `tag1` varchar(20) DEFAULT NULL,
  `tag2` varchar(20) DEFAULT NULL,
  `tag3` varchar(20) DEFAULT NULL,
  `tag4` varchar(20) DEFAULT NULL,
  `tag5` varchar(20) DEFAULT NULL,
  `tag6` varchar(20) DEFAULT NULL,
  `tag7` varchar(20) DEFAULT NULL,
  `tag8` varchar(20) DEFAULT NULL,
  `tag9` varchar(20) DEFAULT NULL,
  `tag10` varchar(20) DEFAULT NULL,
  `location` varchar(100) DEFAULT NULL COMMENT '地理位置',
  `atUserId` varchar(50) DEFAULT NULL COMMENT '好像没用',
  `likeNum` int DEFAULT '0',
  `collectNum` int DEFAULT '0',
  `commentNum` int DEFAULT '0',
  PRIMARY KEY (`noteId`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `noteinfo`
--

LOCK TABLES `noteinfo` WRITE;
/*!40000 ALTER TABLE `noteinfo` DISABLE KEYS */;
INSERT INTO `noteinfo` VALUES (30,10,'情感小作文','??',2,'10_30_1688602433_2F68C16B2CACB3416B06A45CC85C5704.jpg','2023-07-06 08:13:53','2023-07-06 08:13:53','','','','','','情感','','游戏','','','','暗影岛',NULL,0,0,0),(42,10,'吃吃吃','过年?~',3,'10_42_1688604442_IMG_20230118_150155.jpg','2023-07-06 08:47:22','2023-07-06 08:47:22','美食','','','','','','','','','','','',NULL,0,1,0),(43,10,'传统服饰','公选课',2,'10_43_1688604512_IMG_20221129_185858.jpg','2023-07-06 08:48:32','2023-07-06 08:48:32','','','穿搭','','','','','','','','','海琴',NULL,2,1,1),(44,10,'厦门','?',5,'10_44_1688604599_IMG_20210624_174449.jpg','2023-07-06 08:49:59','2023-07-06 08:49:59','','','','','','','','','旅游','','','厦门',NULL,1,1,1),(45,10,'锻炼身体','长跑月……',3,'10_45_1688604665_Screenshot_2023-05-11-21-09-25-275_com.gotokeep.keep.jpg','2023-07-06 08:51:05','2023-07-06 08:51:05','','','','','','','','','','健身','','珠海',NULL,0,0,0),(46,10,'命运石之门','很好看的喵',3,'10_46_1688604721_IMG_20230424_235315.jpg','2023-07-06 08:52:01','2023-07-06 08:52:01','','','','影视','','','','','','','','',NULL,1,0,0),(47,11,'救命！潮州广济桥也太好看了吧！','虽然人很多但是氛围感拉满了！?',3,'11_47_1688605115_c7bbb0d98759663874480867a968658.jpg','2023-07-06 08:58:35','2023-07-06 08:58:35','','','','','','','','','旅游','','','潮州广济桥',NULL,0,0,0),(48,11,'牛肉火锅yyds','牛肉火锅yyds @shx ???',1,'11_48_1688605194_36246419801ab19726bcb86bbd76223.jpg','2023-07-06 08:59:54','2023-07-06 08:59:54','美食','','','','','','','','','','','家里',NULL,3,1,2),(49,10,'今天上大分','?',2,'10_49_1688606156_IMG_20230527_232210.jpg','2023-07-06 09:15:56','2023-07-06 09:15:56','','','','','','','','游戏','','','','无',NULL,1,1,1),(51,12,'美丽广州塔','喀喀喀',1,'12_51_1688607234_微信图片_20230706092331.jpg','2023-07-06 09:33:54','2023-07-06 09:33:54','','','','','','','','','旅游','健身','','广州',NULL,1,1,1),(52,12,'宿舍泡面','嘎嘎香',1,'12_52_1688607808_微信图片_20230706092345.jpg','2023-07-06 09:43:28','2023-07-06 09:43:28','美食','','','','','','','','','','','珠海',NULL,1,1,2),(53,13,'今天你瑞了吗','一杯更比一杯提神醒脑的冰美??',2,'13_53_1688608093_微信图片_202307060945111.jpg','2023-07-06 09:48:13','2023-07-06 09:48:13','美食','','','','','','','','','','','瑞幸',NULL,0,0,0),(54,13,'今天运动啦','今天运动啦！好累！！！！！！！！',2,'13_54_1688608154_微信图片_202307060945115.jpg','2023-07-06 09:49:14','2023-07-06 09:49:14','','','','','','','','','','健身','','珠海',NULL,0,0,0),(55,12,'美丽','何尝不是的 @Lyx  ',1,'12_55_1688609450_9f9481d49fb7a9bc2ae1d662a1afd27.jpg','2023-07-06 10:10:50','2023-07-06 10:10:50','美食','彩妆','','','','','','','','','','抓hi',NULL,0,0,0),(56,12,'爬山','上 @Lyx ',1,'12_56_1688609908_微信图片_20230706092311.jpg','2023-07-06 10:18:28','2023-07-06 10:18:28','美食','彩妆','穿搭','','','','','','','','','抓hi',NULL,0,0,0),(57,12,'卡卡','差不多时间从☝️ @Lyx ',1,'12_57_1688612346_微信图片_20230706092331.jpg','2023-07-06 10:59:06','2023-07-06 10:59:06','','彩妆','穿搭','','','','','','','','','珠海',NULL,0,0,0),(58,11,'来看鸭鸭！','快来看像玥玥的鸭鸭！ @yzt ',1,'11_58_1688622278_c39366c408944dde53124fdb6e59907.jpg','2023-07-06 13:44:38','2023-07-06 13:44:38','','','','','','情感','','','旅游','','','榕园7栋',NULL,0,0,0);
/*!40000 ALTER TABLE `noteinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-06 23:18:37
