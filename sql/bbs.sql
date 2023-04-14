mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: localhost    Database: bbs
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
-- Table structure for table `blog_article`
--

DROP TABLE IF EXISTS `blog_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` mediumtext,
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_by` varchar(255) DEFAULT NULL COMMENT '修改人',
  `state` tinyint(3) unsigned DEFAULT NULL,
  `created_on` int(11) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT NULL,
  `deleted_on` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_article`
--

LOCK TABLES `blog_article` WRITE;
/*!40000 ALTER TABLE `blog_article` DISABLE KEYS */;
INSERT INTO `blog_article` VALUES (2,1,'test211','test-desc111','test-content111','upload/images/20210625/29b7fccf67f3cf0074eff01d3fb860bf.jpg','test-created','',1,1627483780,0,0),(3,2,'test554445','test113332244desc','我的blogee','/nocw/sdfasdf/fasdf/asdfasdf.png','admin','admin',1,0,0,0),(4,3,'test555','test112244desc','sdfdsdhsdfasf','/nocw/sdfasdf/asdfasdf.png','admin','',1,1648083887,0,0),(5,2,'test57755','test112244desc','sdfdsdhsdfasf','/nocw/sdfasdf/asdfasdf.png','admin','',1,1649240654,0,0);
/*!40000 ALTER TABLE `blog_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_auth`
--

DROP TABLE IF EXISTS `blog_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(255) DEFAULT '' COMMENT '密码',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'create time',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modified time',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'delete time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_auth`
--

LOCK TABLES `blog_auth` WRITE;
/*!40000 ALTER TABLE `blog_auth` DISABLE KEYS */;
INSERT INTO `blog_auth` VALUES (1,'test','test123456',0,0,0),(2,'admin','$2a$10$REjgUlSQP0zZ8kVlDILmyOFXSKUs93SzNPjbDJwxDKhHGbj9bBsO6',1627483508,1627483508,0);
/*!40000 ALTER TABLE `blog_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_comment`
--

DROP TABLE IF EXISTS `blog_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `auth_id` int(10) unsigned NOT NULL,
  `article_id` int(10) unsigned NOT NULL,
  `content` mediumtext,
  `created_on` int(11) DEFAULT NULL,
  `modified_on` int(11) DEFAULT NULL,
  `deleted_on` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_comment`
--

LOCK TABLES `blog_comment` WRITE;
/*!40000 ALTER TABLE `blog_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_dog`
--

DROP TABLE IF EXISTS `blog_dog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_dog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `gril_god_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_dog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_dog`
--

LOCK TABLES `blog_dog` WRITE;
/*!40000 ALTER TABLE `blog_dog` DISABLE KEYS */;
INSERT INTO `blog_dog` VALUES (1,'2022-01-25 17:07:26','2022-01-25 17:07:26',NULL,'bibi',1);
/*!40000 ALTER TABLE `blog_dog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_gril_god`
--

DROP TABLE IF EXISTS `blog_gril_god`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_gril_god` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_gril_god_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_gril_god`
--

LOCK TABLES `blog_gril_god` WRITE;
/*!40000 ALTER TABLE `blog_gril_god` DISABLE KEYS */;
INSERT INTO `blog_gril_god` VALUES (1,'2022-01-25 17:07:26','2022-01-25 17:07:26',NULL,'didid');
/*!40000 ALTER TABLE `blog_gril_god` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_log`
--

DROP TABLE IF EXISTS `blog_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(100) NOT NULL DEFAULT '' COMMENT 'code',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'user-id',
  `user_name` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `method` char(10) NOT NULL DEFAULT '' COMMENT '请求方法',
  `ip_address` varchar(100) NOT NULL DEFAULT '' COMMENT 'IP ip_address',
  `created_on` int(10) unsigned DEFAULT '0',
  `modified_on` int(10) unsigned DEFAULT '0',
  `deleted_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1007 DEFAULT CHARSET=utf8mb4 COMMENT='log表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_log`
--

LOCK TABLES `blog_log` WRITE;
/*!40000 ALTER TABLE `blog_log` DISABLE KEYS */;
INSERT INTO `blog_log` VALUES (1005,'200',2,'admin','POST','127.0.0.1',1627568953,1627568953,0),(1006,'200',1,'test','GET','127.0.0.1',1627568987,1627568987,0);
/*!40000 ALTER TABLE `blog_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_menu`
--

DROP TABLE IF EXISTS `blog_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'Menus名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT 'URL',
  `created_on` int(11) DEFAULT '0',
  `updated_on` int(11) DEFAULT '0',
  `deleted_on` int(11) DEFAULT '0',
  `icon` varchar(50) DEFAULT '' COMMENT 'icon',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Menus表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_menu`
--

LOCK TABLES `blog_menu` WRITE;
/*!40000 ALTER TABLE `blog_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_tag`
--

DROP TABLE IF EXISTS `blog_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='文章标签管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_tag`
--

LOCK TABLES `blog_tag` WRITE;
/*!40000 ALTER TABLE `blog_tag` DISABLE KEYS */;
INSERT INTO `blog_tag` VALUES (1,'Go',1627483726,'test2',1627483726,'',0,0),(2,'PHP',1627483742,'test2',1627483742,'',0,1),(3,'JavaScript',1627483754,'test2',1627483754,'',0,1);
/*!40000 ALTER TABLE `blog_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_user`
--

DROP TABLE IF EXISTS `blog_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `age` tinyint(3) unsigned DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `member_number` varchar(191) DEFAULT NULL,
  `activated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_user`
--

LOCK TABLES `blog_user` WRITE;
/*!40000 ALTER TABLE `blog_user` DISABLE KEYS */;
INSERT INTO `blog_user` VALUES (1,'2022-01-25 14:42:09','2022-01-25 14:42:09',NULL,'jj',NULL,13,NULL,NULL,NULL),(2,'2022-01-25 14:46:24','2022-01-25 14:46:24',NULL,'jj',NULL,13,NULL,NULL,NULL),(3,'2022-01-25 15:02:12','2022-01-25 15:02:12',NULL,'jj',NULL,13,NULL,NULL,NULL),(4,'2022-01-25 16:09:07','2022-01-25 16:09:07',NULL,'jj',NULL,13,NULL,NULL,NULL);
/*!40000 ALTER TABLE `blog_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-14  7:43:20
