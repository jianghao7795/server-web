mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: localhost    Database: hlayout
-- ------------------------------------------------------
-- Server version	5.7.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(100) NOT NULL,
  `cid` bigint(20) unsigned NOT NULL,
  `desc` varchar(200) DEFAULT NULL,
  `content` longtext,
  `img` varchar(100) DEFAULT NULL,
  `comment_count` bigint(20) NOT NULL DEFAULT '0',
  `read_count` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_article_deleted_at` (`deleted_at`) USING BTREE,
  KEY `fk_article_category` (`cid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `content` varchar(500) NOT NULL,
  `status` tinyint(4) DEFAULT '2',
  `article_title` longtext,
  `username` longtext,
  `title` longtext,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `profile` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `desc` varchar(200) DEFAULT NULL,
  `qqchat` varchar(200) DEFAULT NULL,
  `wechat` varchar(100) DEFAULT NULL,
  `weibo` varchar(200) DEFAULT NULL,
  `bili` varchar(200) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `img` varchar(200) DEFAULT NULL,
  `avatar` varchar(200) DEFAULT NULL,
  `icp_record` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profile`
--

LOCK TABLES `profile` WRITE;
/*!40000 ALTER TABLE `profile` DISABLE KEYS */;
/*!40000 ALTER TABLE `profile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(500) NOT NULL,
  `role` bigint(20) DEFAULT '2',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (3,'2022-02-15 15:55:36.299','2022-02-15 15:55:36.299','2022-02-20 14:47:00.215','55555','11111',2),(4,'2022-02-20 15:03:14.295','2022-02-20 15:03:14.295',NULL,'bobo11','shsdfgsdfgwagsdfhgjghfghfj',2),(5,'2022-02-20 15:03:24.439','2022-02-20 15:03:24.439',NULL,'bobo113','shsdfgsdfgwadfhgjghfghfj',2),(6,'2022-02-20 15:03:29.519','2022-02-20 15:03:29.519',NULL,'bobo114','shsdfgsdfgwadfhgjghfghfj',2),(7,'2022-02-20 15:03:33.890','2022-02-20 15:03:33.890',NULL,'bobo134','shsdfgsdfgwadfhgjghfghfj',2),(8,'2022-02-20 15:03:38.213','2022-02-20 15:03:38.213',NULL,'bobo154','shsdfgsdfgwadfhgjghfghfj',2),(9,'2022-02-20 15:03:42.802','2022-02-20 15:03:42.802',NULL,'bobo254','shsdfgsdfgwadfhgjghfghfj',2),(10,'2022-02-20 15:03:50.481','2022-02-20 15:03:50.481',NULL,'bobo354','shsdfgsdfgwadfhgjghfghfj',2),(11,'2022-02-20 15:07:34.737','2022-02-20 15:07:34.737',NULL,'bobo355','shsdfgsdfgwadfhgjghfghfj',2),(12,'2022-02-20 15:07:39.384','2022-02-20 15:07:39.384',NULL,'bobaa55','shsdfgsdfgwadfhgjghfghfj',2),(13,'2022-02-20 15:07:45.196','2022-02-20 15:07:45.196',NULL,'bobaa455','shsdfgsdfgwadfhgjghfghfj',2),(14,'2022-02-20 15:07:51.102','2022-02-20 15:07:51.102',NULL,'bobaa465','shsdfgsdfgwadfhgjghfghfj',2),(15,'2022-02-20 15:07:55.600','2022-02-20 15:07:55.600',NULL,'bobaa485','shsdfgsdfgwadfhgjghfghfj',2),(16,'2022-02-20 15:08:02.149','2022-02-20 15:08:02.149',NULL,'bobaa785','shsdfgsdfgwadfhgjghfghfj',2),(17,'2022-02-20 15:08:08.911','2022-02-20 15:08:08.911',NULL,'bobaa885','shsdfgsdfgwadfhgjghfghfj',2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-14  7:48:36
