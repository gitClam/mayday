-- MySQL dump 10.13  Distrib 5.7.31, for linux-glibc2.12 (x86_64)
--
-- Host: localhost    Database: jietong
-- ------------------------------------------------------
-- Server version	5.7.31

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
-- Table structure for table `sd_application`
--

DROP TABLE IF EXISTS `sd_application`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_application` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workspace_id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_application_sd_workspace_1` (`workspace_id`),
  CONSTRAINT `fk_sd_application_sd_workspace_1` FOREIGN KEY (`workspace_id`) REFERENCES `sd_workspace` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_application`
--

LOCK TABLES `sd_application` WRITE;
/*!40000 ALTER TABLE `sd_application` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_application` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_department`
--

DROP TABLE IF EXISTS `sd_department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_department` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workspace_id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `remark` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_department_sd_workspace_1` (`workspace_id`),
  CONSTRAINT `fk_sd_department_sd_workspace_1` FOREIGN KEY (`workspace_id`) REFERENCES `sd_workspace` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_department`
--

LOCK TABLES `sd_department` WRITE;
/*!40000 ALTER TABLE `sd_department` DISABLE KEYS */;
INSERT INTO `sd_department` VALUES (7,8,'默认1','0000000','默认',0),(8,9,'','0000000','默认3',0),(9,10,'','0000000','默认2',1),(10,11,'默认','0000000','傻逼',0),(11,12,'默认','0000000','傻逼',1),(12,10,'蛤氏集团第一部门','13560503260','湛江一条街',0),(13,11,'ha','13560503260','00000',0),(14,12,'蛤','13560503260','111111',0),(15,13,'默认','0000000','傻逼',0),(17,14,'工作空间2部门1','13560503260','2.1',0),(19,15,'工作空间3部门','13560503260','3.1',0),(24,19,'默认','0000000','昨木',0),(29,21,'蛤的老家','13560503260','昨木',0),(30,21,'蛤的第二个老家','13560503260','3.2',0),(32,22,'默认部门1','13560503260','1.1',0),(34,24,'工作部门2','13560503260','2.2',0),(35,24,'工作部门3','13160503260','2.3',0);
/*!40000 ALTER TABLE `sd_department` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_job`
--

DROP TABLE IF EXISTS `sd_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_job` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `department_id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `is_delete` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_job_sd_department_1` (`department_id`),
  CONSTRAINT `fk_sd_job_sd_department_1` FOREIGN KEY (`department_id`) REFERENCES `sd_department` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_job`
--

LOCK TABLES `sd_job` WRITE;
/*!40000 ALTER TABLE `sd_job` DISABLE KEYS */;
INSERT INTO `sd_job` VALUES (1,7,'默认职位',0),(2,8,'默认职位',0),(3,9,'默认职位',0),(4,10,'默认职位',0),(5,11,'默认职位',0),(6,14,'李容丞他爸',0),(7,14,'蛤他爸',0),(8,8,'董事长',0),(9,14,'ha',0),(10,15,'默认职位',0),(17,24,'默认职位',0),(19,29,'默认职位',0),(21,30,'蛤他爸',0),(22,30,'蛤他爷',0),(23,32,'默认职位1',0),(25,34,'默认职位',0),(26,35,'董事长',0);
/*!40000 ALTER TABLE `sd_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_order`
--

DROP TABLE IF EXISTS `sd_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `workflow_id` int(11) NOT NULL,
  `create_time` datetime NOT NULL,
  `title` varchar(20) DEFAULT NULL,
  `urge_last_time` datetime DEFAULT NULL,
  `urge_count` int(11) DEFAULT NULL,
  `related_person` json NOT NULL,
  `is_denied` int(11) NOT NULL,
  `is_end` int(11) NOT NULL,
  `state` json DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_order_sd_user_1` (`user_id`),
  KEY `fk_sd_order_sd_workflow_1` (`workflow_id`),
  CONSTRAINT `fk_sd_order_sd_user_1` FOREIGN KEY (`user_id`) REFERENCES `sd_user` (`id`),
  CONSTRAINT `fk_sd_order_sd_workflow_1` FOREIGN KEY (`workflow_id`) REFERENCES `sd_workflow` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_order`
--

LOCK TABLES `sd_order` WRITE;
/*!40000 ALTER TABLE `sd_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_order_circulation_history`
--

DROP TABLE IF EXISTS `sd_order_circulation_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_order_circulation_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `state` varchar(20) NOT NULL,
  `title` varchar(20) DEFAULT NULL,
  `source` varchar(128) DEFAULT NULL,
  `target` varchar(128) DEFAULT NULL,
  `circulation` varchar(128) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `processor` varchar(50) DEFAULT NULL,
  `processor_id` int(11) DEFAULT NULL,
  `cost_duration` datetime DEFAULT NULL,
  `remarks` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_I` (`order_id`),
  KEY `FK` (`processor_id`),
  CONSTRAINT `FK` FOREIGN KEY (`processor_id`) REFERENCES `sd_user` (`id`),
  CONSTRAINT `FK_I` FOREIGN KEY (`order_id`) REFERENCES `sd_order` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_order_circulation_history`
--

LOCK TABLES `sd_order_circulation_history` WRITE;
/*!40000 ALTER TABLE `sd_order_circulation_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_order_circulation_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_order_table`
--

DROP TABLE IF EXISTS `sd_order_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_order_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_history_id` int(11) DEFAULT NULL,
  `form_structure` json DEFAULT NULL,
  `form_data` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ohid` (`order_history_id`),
  CONSTRAINT `fk_ohid` FOREIGN KEY (`order_history_id`) REFERENCES `sd_order` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_order_table`
--

LOCK TABLES `sd_order_table` WRITE;
/*!40000 ALTER TABLE `sd_order_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_order_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_table`
--

DROP TABLE IF EXISTS `sd_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workspace_id` int(11) NOT NULL,
  `data` json NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_table_sd_workspace_1` (`workspace_id`),
  CONSTRAINT `fk_sd_table_sd_workspace_1` FOREIGN KEY (`workspace_id`) REFERENCES `sd_workspace` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_table`
--

LOCK TABLES `sd_table` WRITE;
/*!40000 ALTER TABLE `sd_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_table_draft`
--

DROP TABLE IF EXISTS `sd_table_draft`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_table_draft` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `data` json NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_table_draft_sd_user_1` (`user_id`),
  CONSTRAINT `fk_sd_table_draft_sd_user_1` FOREIGN KEY (`user_id`) REFERENCES `sd_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_table_draft`
--

LOCK TABLES `sd_table_draft` WRITE;
/*!40000 ALTER TABLE `sd_table_draft` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_table_draft` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_user`
--

DROP TABLE IF EXISTS `sd_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `password` varchar(50) CHARACTER SET latin1 NOT NULL,
  `realname` varchar(30) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `wechat` varchar(100) CHARACTER SET latin1 DEFAULT NULL,
  `QQNumber` varchar(100) CHARACTER SET latin1 DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `sex` enum('男','女') DEFAULT NULL,
  `Info` varchar(100) DEFAULT NULL,
  `mail` varchar(50) CHARACTER SET latin1 NOT NULL,
  `company` varchar(30) DEFAULT NULL,
  `department` varchar(30) DEFAULT NULL,
  `vocation` varchar(30) DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET latin1 DEFAULT NULL,
  `photo` varchar(50) CHARACTER SET latin1 DEFAULT NULL,
  `create_date` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mail` (`mail`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_user`
--

LOCK TABLES `sd_user` WRITE;
/*!40000 ALTER TABLE `sd_user` DISABLE KEYS */;
INSERT INTO `sd_user` VALUES (2,'被淹死的虫合','123456','',10,'111','791285634','2021-01-10 08:00:00','女','10102222','791285634@qq.com','湛江海鲜有限公司','蛤的手下','董事长','13100223344','./data/photo/791285634@qq.com','2021-10-27 16:54:08',0),(3,'0000','********','',5,'','','1999-05-30 08:00:00','女','000','2020202@qq.com','','','','','./data/photo/2.png','2021-11-04 16:54:42',0),(4,'被淹死的蛤','00000000','',0,'','',NULL,'男','','792222222@qq.com','','','','','./data/photo/2.png','2021-11-04 16:57:18',0),(5,'被淹死的虫合','********','陈伟良老婆',22,'lrc3260','791285634','2021-11-16 08:00:00','男','440802200008120811','00000000@qq.com','','','','13560503260','./data/photo/00000000@qq.com','2021-11-04 17:04:03',0),(7,'被淹死的虫合','11111111','',0,'','',NULL,'男','','11111111@qq.com','','','','','./data/photo/2.png','2021-11-04 18:09:13',0),(8,'perry','123456','',0,'','',NULL,'男','','798692795@qq.com','','','','','./data/photo/2.png','2021-11-15 15:04:33',0);
/*!40000 ALTER TABLE `sd_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_user_job`
--

DROP TABLE IF EXISTS `sd_user_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_user_job` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `job_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_user_job_sd_user_1` (`user_id`),
  KEY `fk_sd_user_job_sd_job_1` (`job_id`),
  CONSTRAINT `fk_sd_user_job_sd_job_1` FOREIGN KEY (`job_id`) REFERENCES `sd_job` (`id`),
  CONSTRAINT `fk_sd_user_job_sd_user_1` FOREIGN KEY (`user_id`) REFERENCES `sd_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_user_job`
--

LOCK TABLES `sd_user_job` WRITE;
/*!40000 ALTER TABLE `sd_user_job` DISABLE KEYS */;
INSERT INTO `sd_user_job` VALUES (1,2,2),(2,2,3),(3,2,4),(4,2,5),(23,5,25);
/*!40000 ALTER TABLE `sd_user_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_user_workflow_group`
--

DROP TABLE IF EXISTS `sd_user_workflow_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_user_workflow_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `workflow_group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_user_workflow_group_sd_user_1` (`user_id`),
  KEY `fk_sd_user_workflow_group_sd_workflow_group_1` (`workflow_group_id`),
  CONSTRAINT `fk_sd_user_workflow_group_sd_user_1` FOREIGN KEY (`user_id`) REFERENCES `sd_user` (`id`),
  CONSTRAINT `fk_sd_user_workflow_group_sd_workflow_group_1` FOREIGN KEY (`workflow_group_id`) REFERENCES `sd_workflow_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_user_workflow_group`
--

LOCK TABLES `sd_user_workflow_group` WRITE;
/*!40000 ALTER TABLE `sd_user_workflow_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_user_workflow_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow`
--

DROP TABLE IF EXISTS `sd_workflow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `create_user` int(11) NOT NULL,
  `create_time` datetime NOT NULL,
  `is_start` tinyint(1) NOT NULL,
  `ceiling_count` int(11) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `structure` json NOT NULL,
  `tables` json NOT NULL,
  `remarks` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_workflow_sd_user_1` (`create_user`),
  CONSTRAINT `fk_sd_workflow_sd_user_1` FOREIGN KEY (`create_user`) REFERENCES `sd_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow`
--

LOCK TABLES `sd_workflow` WRITE;
/*!40000 ALTER TABLE `sd_workflow` DISABLE KEYS */;
INSERT INTO `sd_workflow` VALUES (3,'test_name1',2,'2021-11-03 16:32:29',0,0,0,'{\"edges\": [{\"sort\": \"2\", \"clazz\": \"flow\", \"label\": \"2\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"start1635927810925\", \"target\": \"receiveTask1635927812066\", \"endPoint\": {\"x\": 313.3125, \"y\": 94, \"index\": 3}, \"startPoint\": {\"x\": 186.3125, \"y\": 95, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}, {\"sort\": \"4\", \"clazz\": \"flow\", \"label\": \"4\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"receiveTask1635927812066\", \"target\": \"userTask1636520102137\", \"endPoint\": {\"x\": 485.40625, \"y\": 98, \"index\": 3}, \"startPoint\": {\"x\": 394.3125, \"y\": 94, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}], \"nodes\": [{\"x\": 170.8125, \"y\": 95, \"id\": \"start1635927810925\", \"size\": [30, 30], \"sort\": \"1\", \"clazz\": \"start\", \"label\": \"1\", \"shape\": \"start-node\"}, {\"x\": 353.8125, \"y\": 94, \"id\": \"receiveTask1635927812066\", \"size\": [80, 44], \"sort\": \"3\", \"clazz\": \"receiveTask\", \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}, {\"x\": 525.90625, \"y\": 98, \"id\": \"userTask1636520102137\", \"size\": [80, 44], \"sort\": \"5\", \"clazz\": \"userTask\", \"label\": \"审批节点\", \"shape\": \"user-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}], \"groups\": []}','[\"1001\"]','备注测试流程1111'),(4,'test_name2',2,'2021-11-04 13:41:03',0,0,0,'{\"edges\": [{\"sort\": \"2\", \"clazz\": \"flow\", \"label\": \"2\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"start1635927810925\", \"target\": \"receiveTask1635927812066\", \"endPoint\": {\"x\": 313.3125, \"y\": 94, \"index\": 3}, \"startPoint\": {\"x\": 186.3125, \"y\": 95, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}], \"nodes\": [{\"x\": 170.8125, \"y\": 95, \"id\": \"start1635927810925\", \"size\": [30, 30], \"sort\": \"1\", \"clazz\": \"start\", \"label\": \"1\", \"shape\": \"start-node\"}, {\"x\": 353.8125, \"y\": 94, \"id\": \"receiveTask1635927812066\", \"size\": [80, 44], \"sort\": \"3\", \"clazz\": \"receiveTask\", \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}], \"groups\": []}','[\"1001\", \"1002\"]','备注测试流程22'),(5,'test',5,'2021-11-10 13:38:16',1,0,0,'{\"edges\": [{\"sort\": \"2\", \"clazz\": \"flow\", \"label\": \"2\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"start1636522640107\", \"target\": \"receiveTask1636522642912\", \"endPoint\": {\"x\": 246.5, \"y\": 95, \"index\": 3}, \"startPoint\": {\"x\": 148.5, \"y\": 95, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}], \"nodes\": [{\"x\": 133, \"y\": 95, \"id\": \"start1636522640107\", \"size\": [30, 30], \"sort\": \"1\", \"clazz\": \"start\", \"label\": \"1\", \"shape\": \"start-node\"}, {\"x\": 287, \"y\": 95, \"id\": \"receiveTask1636522642912\", \"size\": [80, 44], \"sort\": \"3\", \"clazz\": \"receiveTask\", \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}], \"groups\": []}','[]',''),(6,'',8,'2021-12-22 19:05:42',1,0,0,'{\"edges\": [{\"sort\": \"2\", \"clazz\": \"flow\", \"label\": \"2\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"start1640171079833\", \"target\": \"receiveTask1640171080954\", \"endPoint\": {\"x\": 345.5874938964844, \"y\": 127, \"index\": 3}, \"startPoint\": {\"x\": 250.5874938964844, \"y\": 125, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}, {\"sort\": \"4\", \"clazz\": \"flow\", \"label\": \"4\", \"shape\": \"flow-polyline-round\", \"style\": {}, \"source\": \"receiveTask1640171080954\", \"target\": \"userTask1640171092506\", \"endPoint\": {\"x\": 522.5874938964844, \"y\": 123, \"index\": 3}, \"startPoint\": {\"x\": 426.5874938964844, \"y\": 127, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3, \"flowProperties\": \"1\"}], \"nodes\": [{\"x\": 235.0874938964844, \"y\": 125, \"id\": \"start1640171079833\", \"size\": [30, 30], \"sort\": \"2\", \"clazz\": \"start\", \"label\": \"1\", \"shape\": \"start-node\"}, {\"x\": 386.0874938964844, \"y\": 127, \"id\": \"receiveTask1640171080954\", \"size\": [80, 44], \"sort\": \"3\", \"clazz\": \"receiveTask\", \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}, {\"x\": 563.0874938964844, \"y\": 123, \"id\": \"userTask1640171092506\", \"size\": [80, 44], \"sort\": \"5\", \"clazz\": \"userTask\", \"label\": \"审批节点\", \"shape\": \"user-task-node\", \"assignType\": \"workspace\", \"assignValue\": [], \"isCounterSign\": false}], \"groups\": []}','[]','');
/*!40000 ALTER TABLE `sd_workflow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow_application`
--

DROP TABLE IF EXISTS `sd_workflow_application`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow_application` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workflow_id` int(11) DEFAULT NULL,
  `application_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_workflow_application_sd_application_1` (`application_id`),
  KEY `fk_sd_workflow_application_sd_workflow_1` (`workflow_id`),
  CONSTRAINT `fk_sd_workflow_application_sd_application_1` FOREIGN KEY (`application_id`) REFERENCES `sd_application` (`id`),
  CONSTRAINT `fk_sd_workflow_application_sd_workflow_1` FOREIGN KEY (`workflow_id`) REFERENCES `sd_workflow` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow_application`
--

LOCK TABLES `sd_workflow_application` WRITE;
/*!40000 ALTER TABLE `sd_workflow_application` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_workflow_application` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow_draft`
--

DROP TABLE IF EXISTS `sd_workflow_draft`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow_draft` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `is_start` tinyint(1) NOT NULL,
  `ceiling_count` int(11) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `structure` json NOT NULL,
  `tables` json NOT NULL,
  `remarks` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sd_workflow_draft_sd_user_1` (`owner_id`),
  CONSTRAINT `fk_sd_workflow_draft_sd_user_1` FOREIGN KEY (`owner_id`) REFERENCES `sd_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow_draft`
--

LOCK TABLES `sd_workflow_draft` WRITE;
/*!40000 ALTER TABLE `sd_workflow_draft` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_workflow_draft` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow_group`
--

DROP TABLE IF EXISTS `sd_workflow_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow_group`
--

LOCK TABLES `sd_workflow_group` WRITE;
/*!40000 ALTER TABLE `sd_workflow_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_workflow_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow_node`
--

DROP TABLE IF EXISTS `sd_workflow_node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `workflow_id` int(11) NOT NULL,
  `table_id` int(11) NOT NULL,
  `serial_number` int(11) NOT NULL,
  `workflow_group_id` int(11) DEFAULT NULL,
  `
permissions` varchar(255) NOT NULL,
  `is_remind` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_workflow_node_sd_workflow_1` (`workflow_id`),
  KEY `fk_sd_workflow_node_sd_workflow_group_1` (`workflow_group_id`),
  KEY `fk_sd_workflow_node_sd_table_1` (`table_id`),
  CONSTRAINT `fk_sd_workflow_node_sd_table_1` FOREIGN KEY (`table_id`) REFERENCES `sd_table` (`id`),
  CONSTRAINT `fk_sd_workflow_node_sd_workflow_1` FOREIGN KEY (`workflow_id`) REFERENCES `sd_workflow` (`id`),
  CONSTRAINT `fk_sd_workflow_node_sd_workflow_group_1` FOREIGN KEY (`workflow_group_id`) REFERENCES `sd_workflow_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow_node`
--

LOCK TABLES `sd_workflow_node` WRITE;
/*!40000 ALTER TABLE `sd_workflow_node` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_workflow_node` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workflow_node_draft`
--

DROP TABLE IF EXISTS `sd_workflow_node_draft`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workflow_node_draft` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `workflow_id` int(11) NOT NULL,
  `table_id` int(11) NOT NULL,
  `serial_number` int(11) NOT NULL,
  `workflow_group_id` int(11) DEFAULT NULL,
  `permissions` varchar(255) NOT NULL,
  `is_remind` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_sd_workflow_node_draft_sd_workflow_draft_1` (`workflow_id`),
  KEY `fk_sd_workflow_node_draft_sd_table_1` (`table_id`),
  CONSTRAINT `fk_sd_workflow_node_draft_sd_table_1` FOREIGN KEY (`table_id`) REFERENCES `sd_table` (`id`),
  CONSTRAINT `fk_sd_workflow_node_draft_sd_workflow_draft_1` FOREIGN KEY (`workflow_id`) REFERENCES `sd_workflow_draft` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workflow_node_draft`
--

LOCK TABLES `sd_workflow_node_draft` WRITE;
/*!40000 ALTER TABLE `sd_workflow_node_draft` DISABLE KEYS */;
/*!40000 ALTER TABLE `sd_workflow_node_draft` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sd_workspace`
--

DROP TABLE IF EXISTS `sd_workspace`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sd_workspace` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sd_workspace`
--

LOCK TABLES `sd_workspace` WRITE;
/*!40000 ALTER TABLE `sd_workspace` DISABLE KEYS */;
INSERT INTO `sd_workspace` VALUES (1,'111','3260','LZSB',0),(8,'12','3260','LZSB',0),(9,'12','3260','LZSB',0),(10,'34','3260','LZSB',0),(11,'00','13560503260','000',0),(12,'蛤','13560203260','哈哈哈',0),(13,'蛤的狗窝1','13560503260','1',0),(14,'工作空间2','13560503260','2',0),(15,'工作空间3','13560503260','3',0),(16,'工作空间4','13560503260','4',0),(19,'蛤的狗屋2','13560503260','2',0),(21,'蛤的狗屋3','13560503260','3',0),(22,'工作空间1','13560503260','1',0),(24,'工作空间2','13560503260','2',0);
/*!40000 ALTER TABLE `sd_workspace` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-10 17:12:13
