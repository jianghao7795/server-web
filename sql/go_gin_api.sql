mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: localhost    Database: go_gin_api
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
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','f78382de80cf583cf854bbac0b6e796fbde36fe2739ca4ae072637010f179cb0','管理员','13888888888',1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34','');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_menu`
--

DROP TABLE IF EXISTS `admin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  PRIMARY KEY (`id`),
  KEY `idx_admin_id` (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 COMMENT='管理员菜单栏表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_menu`
--

LOCK TABLES `admin_menu` WRITE;
/*!40000 ALTER TABLE `admin_menu` DISABLE KEYS */;
INSERT INTO `admin_menu` VALUES (1,1,16,'2022-07-24 04:32:34','init'),(2,1,21,'2022-07-24 04:32:34','init'),(3,1,20,'2022-07-24 04:32:34','init'),(4,1,19,'2022-07-24 04:32:34','init'),(5,1,18,'2022-07-24 04:32:34','init'),(6,1,17,'2022-07-24 04:32:34','init'),(7,1,13,'2022-07-24 04:32:34','init'),(8,1,15,'2022-07-24 04:32:34','init'),(9,1,14,'2022-07-24 04:32:34','init'),(10,1,10,'2022-07-24 04:32:34','init'),(11,1,12,'2022-07-24 04:32:34','init'),(12,1,11,'2022-07-24 04:32:34','init'),(13,1,7,'2022-07-24 04:32:34','init'),(14,1,9,'2022-07-24 04:32:34','init'),(15,1,8,'2022-07-24 04:32:34','init'),(16,1,4,'2022-07-24 04:32:34','init'),(17,1,6,'2022-07-24 04:32:34','init'),(18,1,5,'2022-07-24 04:32:34','init'),(19,1,1,'2022-07-24 04:32:34','init'),(20,1,3,'2022-07-24 04:32:34','init'),(21,1,2,'2022-07-24 04:32:34','init'),(22,1,22,'2022-07-24 04:32:34','init'),(23,1,23,'2022-07-24 04:32:34','init'),(24,1,24,'2022-07-24 04:32:34','init'),(25,1,25,'2022-07-24 04:32:34','init');
/*!40000 ALTER TABLE `admin_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authorized`
--

DROP TABLE IF EXISTS `authorized`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authorized` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `business_key` varchar(32) NOT NULL DEFAULT '' COMMENT '调用方key',
  `business_secret` varchar(60) NOT NULL DEFAULT '' COMMENT '调用方secret',
  `business_developer` varchar(60) NOT NULL DEFAULT '' COMMENT '调用方对接人',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_business_key` (`business_key`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='已授权的调用方表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authorized`
--

LOCK TABLES `authorized` WRITE;
/*!40000 ALTER TABLE `authorized` DISABLE KEYS */;
INSERT INTO `authorized` VALUES (1,'admin','12878dd962115106db6d','管理员','管理面板调用',1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34','');
/*!40000 ALTER TABLE `authorized` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authorized_api`
--

DROP TABLE IF EXISTS `authorized_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authorized_api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `business_key` varchar(32) NOT NULL DEFAULT '' COMMENT '调用方key',
  `method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',
  `api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='已授权接口地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authorized_api`
--

LOCK TABLES `authorized_api` WRITE;
/*!40000 ALTER TABLE `authorized_api` DISABLE KEYS */;
INSERT INTO `authorized_api` VALUES (1,'admin','GET','/api/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(2,'admin','POST','/api/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(3,'admin','PUT','/api/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(4,'admin','DELETE','/api/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(5,'admin','PATCH','/api/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34','');
/*!40000 ALTER TABLE `authorized_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cron_task`
--

DROP TABLE IF EXISTS `cron_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cron_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `spec` varchar(64) NOT NULL DEFAULT '' COMMENT 'crontab 表达式',
  `command` varchar(255) NOT NULL DEFAULT '' COMMENT '执行命令',
  `protocol` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '执行方式 1:shell 2:http',
  `http_method` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'http 请求方式 1:get 2:post',
  `timeout` int(11) unsigned NOT NULL DEFAULT '60' COMMENT '超时时间(单位:秒)',
  `retry_times` tinyint(1) NOT NULL DEFAULT '3' COMMENT '重试次数',
  `retry_interval` int(11) NOT NULL DEFAULT '60' COMMENT '重试间隔(单位:秒)',
  `notify_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知',
  `notify_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '通知类型 1:邮件 2:webhook',
  `notify_receiver_email` varchar(255) NOT NULL DEFAULT '' COMMENT '通知者邮箱地址(多个用,分割)',
  `notify_keyword` varchar(255) NOT NULL DEFAULT '' COMMENT '通知匹配关键字(多个用,分割)',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台任务表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cron_task`
--

LOCK TABLES `cron_task` WRITE;
/*!40000 ALTER TABLE `cron_task` DISABLE KEYS */;
/*!40000 ALTER TABLE `cron_task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `link` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',
  `icon` varchar(60) NOT NULL DEFAULT '' COMMENT '图标',
  `level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型 1:一级菜单 2:二级菜单',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是 -1:否',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 COMMENT='左侧菜单栏表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,0,'配置信息','','mdi-settings-box',1,10,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(2,1,'告警邮箱','/config/email','',2,101,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(3,1,'错误码','/config/code','',2,102,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(4,0,'代码生成器','','mdi-code-not-equal-variant',1,20,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(5,4,'生成数据表 CURD','/generator/gorm','',2,201,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(6,4,'生成控制器方法','/generator/handler','',2,202,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(7,0,'授权调用方','','mdi-playlist-check',1,30,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(8,7,'调用方','/authorized/list','',2,301,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(9,7,'使用说明','/authorized/demo','',2,302,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(10,0,'系统管理员','','mdi-account',1,50,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(11,10,'管理员','/admin/list','',2,501,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(12,10,'菜单管理','/admin/menu','',2,502,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(13,0,'查询小助手','','mdi-database-search',1,60,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(14,13,'查询缓存','/tool/cache','',2,601,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(15,13,'查询数据','/tool/data','',2,602,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(16,0,'实用工具箱','','mdi-tools',1,70,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(17,16,'Hashids','/tool/hashids','',2,702,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(18,16,'调用日志','/tool/logs','',2,703,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(19,16,'接口文档','/swagger/index.html','',2,704,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(20,16,'GraphQL','/graphql','',2,705,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(21,16,'接口指标','/metrics','',2,706,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(22,16,'服务升级','/upgrade','',2,701,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(23,0,'后台任务','','mdi-av-timer',1,40,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(24,23,'任务列表','/cron/list','',2,401,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(25,16,'WebSocket','/tool/websocket','',2,707,1,-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34','');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_action`
--

DROP TABLE IF EXISTS `menu_action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu_action` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',
  `method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',
  `api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
  PRIMARY KEY (`id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8 COMMENT='功能权限表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_action`
--

LOCK TABLES `menu_action` WRITE;
/*!40000 ALTER TABLE `menu_action` DISABLE KEYS */;
INSERT INTO `menu_action` VALUES (1,17,'GET','/api/tool/hashids/**',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(2,14,'POST','/api/tool/cache/search',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(3,14,'PATCH','/api/tool/cache/clear',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(4,15,'GET','/api/tool/data/dbs',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(5,15,'POST','/api/tool/data/mysql',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(6,15,'POST','/api/tool/data/tables',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(7,2,'PATCH','/api/config/email',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(8,5,'POST','/generator/gorm/execute',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(9,6,'POST','/generator/handler/execute',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(10,8,'GET','/authorized/add',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(11,8,'GET','/authorized/api/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(12,8,'GET','/api/authorized',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(13,8,'PATCH','/api/authorized/used',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(14,8,'DELETE','/api/authorized/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(15,8,'POST','/api/authorized',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(16,8,'GET','/api/authorized_api',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(17,8,'POST','/api/authorized_api',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(18,8,'DELETE','/api/authorized_api/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(19,11,'GET','/admin/add',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(20,11,'POST','/api/admin',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(21,11,'GET','/api/admin',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(22,11,'PATCH','/api/admin/used',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(23,11,'PATCH','/api/admin/reset_password/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(24,11,'DELETE','/api/admin/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(25,11,'GET','/admin/action/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(26,11,'GET','/api/admin/menu/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(27,11,'POST','/api/admin/menu',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(28,12,'GET','/admin/menu_action/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(29,12,'GET','/api/menu',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(30,12,'DELETE','/api/menu/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(31,12,'GET','/api/menu/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(32,12,'PATCH','/api/menu/used',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(33,12,'POST','/api/menu',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(34,12,'GET','/api/menu_action',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(35,12,'POST','/api/menu_action',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(36,12,'DELETE','/api/menu_action/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(37,22,'POST','/upgrade/execute',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(38,11,'PATCH','/api/admin/offline',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(39,12,'PATCH','/api/menu/sort',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(40,24,'GET','/cron/add',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(41,24,'GET','/cron/edit/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(42,24,'POST','/api/cron',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(43,24,'POST','/api/cron/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(44,24,'GET','/api/cron',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(45,24,'GET','/api/cron/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(46,24,'PATCH','/api/cron/used',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(47,24,'PATCH','/api/cron/exec/*',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34',''),(48,25,'POST','/api/tool/send_message',-1,'2022-07-24 04:32:34','init','2022-07-24 04:32:34','');
/*!40000 ALTER TABLE `menu_action` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-14  7:46:21
