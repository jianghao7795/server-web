mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: localhost    Database: SQLbooks
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
-- Table structure for table `Customers`
--

DROP TABLE IF EXISTS `Customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Customers` (
  `cust_id` char(10) NOT NULL,
  `cust_name` char(50) NOT NULL,
  `cust_address` char(50) DEFAULT NULL,
  `cust_city` char(50) DEFAULT NULL,
  `cust_state` char(5) DEFAULT NULL,
  `cust_zip` char(10) DEFAULT NULL,
  `cust_country` char(50) DEFAULT NULL,
  `cust_contact` char(50) DEFAULT NULL,
  `cust_email` char(255) DEFAULT NULL,
  PRIMARY KEY (`cust_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Customers`
--

LOCK TABLES `Customers` WRITE;
/*!40000 ALTER TABLE `Customers` DISABLE KEYS */;
INSERT INTO `Customers` VALUES ('1000000001','Village Toys','200 Maple Lane','Detroit','MI','44444','USA','John Smith','sales@villagetoys.com'),('1000000002','Kids Place','333 South Lake Drive','Columbus','OH','43333','USA','Michelle Green',NULL),('1000000003','Fun4All','1 Sunny Place','Muncie','IN','42222','USA','Jim Jones','jjones@fun4all.com'),('1000000004','Fun4All','829 Riverside Drive','Phoenix','AZ','88888','USA','Denise L. Stephens','dstephens@fun4all.com'),('1000000005','The Toy Store','4545 53rd Street','Chicago','IL','54545','USA','Kim Howard',NULL);
/*!40000 ALTER TABLE `Customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OrderItems`
--

DROP TABLE IF EXISTS `OrderItems`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OrderItems` (
  `order_num` int(11) NOT NULL,
  `order_item` int(11) NOT NULL,
  `prod_id` char(10) NOT NULL,
  `quantity` int(11) NOT NULL,
  `item_price` decimal(8,2) NOT NULL,
  PRIMARY KEY (`order_num`,`order_item`),
  KEY `FK_OrderItems_Products` (`prod_id`),
  CONSTRAINT `FK_OrderItems_Orders` FOREIGN KEY (`order_num`) REFERENCES `Orders` (`order_num`),
  CONSTRAINT `FK_OrderItems_Products` FOREIGN KEY (`prod_id`) REFERENCES `Products` (`prod_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderItems`
--

LOCK TABLES `OrderItems` WRITE;
/*!40000 ALTER TABLE `OrderItems` DISABLE KEYS */;
INSERT INTO `OrderItems` VALUES (20005,1,'BR01',100,5.49),(20005,2,'BR03',100,10.99),(20006,1,'BR01',20,5.99),(20006,2,'BR02',10,8.99),(20006,3,'BR03',10,11.99),(20007,1,'BR03',50,11.49),(20007,2,'BNBG01',100,2.99),(20007,3,'BNBG02',100,2.99),(20007,4,'BNBG03',100,2.99),(20007,5,'RGAN01',50,4.49),(20008,1,'RGAN01',5,4.99),(20008,2,'BR03',5,11.99),(20008,3,'BNBG01',10,3.49),(20008,4,'BNBG02',10,3.49),(20008,5,'BNBG03',10,3.49),(20009,1,'BNBG01',250,2.49),(20009,2,'BNBG02',250,2.49),(20009,3,'BNBG03',250,2.49);
/*!40000 ALTER TABLE `OrderItems` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Orders`
--

DROP TABLE IF EXISTS `Orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Orders` (
  `order_num` int(11) NOT NULL,
  `order_date` datetime NOT NULL,
  `cust_id` char(10) NOT NULL,
  PRIMARY KEY (`order_num`),
  KEY `FK_Orders_Customers` (`cust_id`),
  CONSTRAINT `FK_Orders_Customers` FOREIGN KEY (`cust_id`) REFERENCES `Customers` (`cust_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Orders`
--

LOCK TABLES `Orders` WRITE;
/*!40000 ALTER TABLE `Orders` DISABLE KEYS */;
INSERT INTO `Orders` VALUES (20005,'2020-05-01 00:00:00','1000000001'),(20006,'2020-01-12 00:00:00','1000000003'),(20007,'2020-01-30 00:00:00','1000000004'),(20008,'2020-02-03 00:00:00','1000000005'),(20009,'2020-02-08 00:00:00','1000000001');
/*!40000 ALTER TABLE `Orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Products`
--

DROP TABLE IF EXISTS `Products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Products` (
  `prod_id` char(10) NOT NULL,
  `vend_id` char(10) NOT NULL,
  `prod_name` char(255) NOT NULL,
  `prod_price` decimal(8,2) NOT NULL,
  `prod_desc` text,
  PRIMARY KEY (`prod_id`),
  KEY `FK_Products_Vendors` (`vend_id`),
  CONSTRAINT `FK_Products_Vendors` FOREIGN KEY (`vend_id`) REFERENCES `Vendors` (`vend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Products`
--

LOCK TABLES `Products` WRITE;
/*!40000 ALTER TABLE `Products` DISABLE KEYS */;
INSERT INTO `Products` VALUES ('BNBG01','DLL01','Fish bean bag toy',3.49,'Fish bean bag toy, complete with bean bag worms with which to feed it'),('BNBG02','DLL01','Bird bean bag toy',3.49,'Bird bean bag toy, eggs are not included'),('BNBG03','DLL01','Rabbit bean bag toy',3.49,'Rabbit bean bag toy, comes with bean bag carrots'),('BR01','BRS01','8 inch teddy bear',5.99,'8 inch teddy bear, comes with cap and jacket'),('BR02','BRS01','12 inch teddy bear',8.99,'12 inch teddy bear, comes with cap and jacket'),('BR03','BRS01','18 inch teddy bear',11.99,'18 inch teddy bear, comes with cap and jacket'),('RGAN01','DLL01','Raggedy Ann',4.99,'18 inch Raggedy Ann doll'),('RYL01','FNG01','King doll',9.49,'12 inch king doll with royal garments and crown'),('RYL02','FNG01','Queen doll',9.49,'12 inch queen doll with royal garments and crown');
/*!40000 ALTER TABLE `Products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vendors`
--

DROP TABLE IF EXISTS `Vendors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Vendors` (
  `vend_id` char(10) NOT NULL,
  `vend_name` char(50) NOT NULL,
  `vend_address` char(50) DEFAULT NULL,
  `vend_city` char(50) DEFAULT NULL,
  `vend_state` char(5) DEFAULT NULL,
  `vend_zip` char(10) DEFAULT NULL,
  `vend_country` char(50) DEFAULT NULL,
  PRIMARY KEY (`vend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vendors`
--

LOCK TABLES `Vendors` WRITE;
/*!40000 ALTER TABLE `Vendors` DISABLE KEYS */;
INSERT INTO `Vendors` VALUES ('BRE02','Bear Emporium','500 Park Street','Anytown','OH','44333','USA'),('BRS01','Bears R Us','123 Main Street','Bear Town','MI','44444','USA'),('DLL01','Doll House Inc.','555 High Street','Dollsville','CA','99999','USA'),('FNG01','Fun and Games','42 Galaxy Road','London',NULL,'N16 6PS','England'),('FRB01','Furball Inc.','1000 5th Avenue','New York','NY','11111','USA'),('JTS01','Jouets et ours','1 Rue Amusement','Paris',NULL,'45678','France');
/*!40000 ALTER TABLE `Vendors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `course` (
  `c_no` varchar(20) NOT NULL COMMENT '课程号',
  `c_name` varchar(20) NOT NULL COMMENT '课程名称',
  `t_no` varchar(20) NOT NULL COMMENT '教师编号 外键关联teacher表',
  PRIMARY KEY (`c_no`),
  KEY `t_no` (`t_no`),
  CONSTRAINT `course_ibfk_1` FOREIGN KEY (`t_no`) REFERENCES `teacher` (`t_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES ('3-105','计算机导论','825'),('3-245','操作系统','804'),('6-166','数字电路','856'),('9-888','高等数学','831');
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `score`
--

DROP TABLE IF EXISTS `score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `score` (
  `s_no` varchar(20) NOT NULL COMMENT '成绩表的编号 依赖学生学号',
  `c_no` varchar(20) NOT NULL COMMENT '课程号 依赖于课程表中的c_id',
  `sc_degree` decimal(10,0) DEFAULT NULL,
  PRIMARY KEY (`s_no`,`c_no`),
  KEY `c_no` (`c_no`),
  CONSTRAINT `score_ibfk_1` FOREIGN KEY (`s_no`) REFERENCES `student` (`s_no`),
  CONSTRAINT `score_ibfk_2` FOREIGN KEY (`c_no`) REFERENCES `course` (`c_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `score`
--

LOCK TABLES `score` WRITE;
/*!40000 ALTER TABLE `score` DISABLE KEYS */;
INSERT INTO `score` VALUES ('103','3-105',92),('103','3-245',86),('103','6-166',85),('105','3-105',88),('105','3-245',75),('105','6-166',79),('109','3-105',76),('109','3-245',68),('109','6-166',81);
/*!40000 ALTER TABLE `score` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student` (
  `s_no` varchar(20) NOT NULL COMMENT '学生学号',
  `s_name` varchar(20) NOT NULL COMMENT '学生姓名 不能为空',
  `s_sex` varchar(10) NOT NULL COMMENT '学生性别',
  `s_birthday` datetime DEFAULT NULL COMMENT '学生生日',
  `s_class` varchar(20) DEFAULT NULL COMMENT '学生所在的班级',
  PRIMARY KEY (`s_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES ('101','曾华','男','1977-09-01 00:00:00','95033'),('102','匡明','男','1975-10-02 00:00:00','95031'),('103','王丽','女','1976-01-23 00:00:00','95033'),('104','李军','男','1976-02-20 00:00:00','95033'),('105','王芳','女','1975-02-10 00:00:00','95031'),('106','陆军','男','1974-06-03 00:00:00','95031'),('107','王尼玛','男','1976-02-20 00:00:00','95033'),('108','张全蛋','男','1975-02-10 00:00:00','95031'),('109','赵铁柱','男','1974-06-03 00:00:00','95031');
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teacher`
--

DROP TABLE IF EXISTS `teacher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teacher` (
  `t_no` varchar(20) NOT NULL COMMENT '教师编号',
  `t_name` varchar(20) NOT NULL COMMENT '教师姓名',
  `t_sex` varchar(20) NOT NULL COMMENT '教师性别',
  `t_birthday` datetime DEFAULT NULL COMMENT '教师生日',
  `t_rof` varchar(20) NOT NULL COMMENT '教师职称',
  `t_depart` varchar(20) NOT NULL COMMENT '教师所在的部门',
  PRIMARY KEY (`t_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher`
--

LOCK TABLES `teacher` WRITE;
/*!40000 ALTER TABLE `teacher` DISABLE KEYS */;
INSERT INTO `teacher` VALUES ('804','李诚','男','1958-12-02 00:00:00','副教授','计算机系'),('825','王萍','女','1972-05-05 00:00:00','助教','计算机系'),('831','刘冰','女','1977-08-14 00:00:00','助教','电子工程系'),('856','张旭','男','1969-03-12 00:00:00','讲师','电子工程系');
/*!40000 ALTER TABLE `teacher` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user5`
--

DROP TABLE IF EXISTS `user5`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user5` (
  `id` int(11) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  UNIQUE KEY `name_2` (`name`),
  UNIQUE KEY `name` (`name`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user5`
--

LOCK TABLES `user5` WRITE;
/*!40000 ALTER TABLE `user5` DISABLE KEYS */;
INSERT INTO `user5` VALUES (1,'lisi'),(2,'zhangsan');
/*!40000 ALTER TABLE `user5` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-14  7:55:58
