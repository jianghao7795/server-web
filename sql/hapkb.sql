mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: localhost    Database: hapkb
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
-- Table structure for table `annotation_approves`
--

DROP TABLE IF EXISTS `annotation_approves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_approves` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `druginfo_id` int(11) NOT NULL DEFAULT '0',
  `description` text COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_druginfo` (`annotation_id`,`druginfo_id`),
  KEY `druginfo_id` (`druginfo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_approves`
--

LOCK TABLES `annotation_approves` WRITE;
/*!40000 ALTER TABLE `annotation_approves` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_approves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_diseases`
--

DROP TABLE IF EXISTS `annotation_diseases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_diseases` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `disease_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_disease` (`annotation_id`,`disease_id`),
  KEY `disease_id` (`disease_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_diseases`
--

LOCK TABLES `annotation_diseases` WRITE;
/*!40000 ALTER TABLE `annotation_diseases` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_diseases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_drug_groups`
--

DROP TABLE IF EXISTS `annotation_drug_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_drug_groups` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `drug_group_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_drug_group` (`annotation_id`,`drug_group_id`),
  KEY `drug_group_id` (`drug_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_drug_groups`
--

LOCK TABLES `annotation_drug_groups` WRITE;
/*!40000 ALTER TABLE `annotation_drug_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_drug_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_drugs`
--

DROP TABLE IF EXISTS `annotation_drugs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_drugs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `drug_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_drug` (`annotation_id`,`drug_id`),
  KEY `drug_id` (`drug_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_drugs`
--

LOCK TABLES `annotation_drugs` WRITE;
/*!40000 ALTER TABLE `annotation_drugs` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_drugs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_guides`
--

DROP TABLE IF EXISTS `annotation_guides`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_guides` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `druginfo_id` int(11) NOT NULL DEFAULT '0',
  `description` text COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_druginfo` (`annotation_id`,`druginfo_id`),
  KEY `druginfo_id` (`druginfo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_guides`
--

LOCK TABLES `annotation_guides` WRITE;
/*!40000 ALTER TABLE `annotation_guides` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_guides` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_logs`
--

DROP TABLE IF EXISTS `annotation_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `version` bigint(20) unsigned NOT NULL DEFAULT '0',
  `operation` smallint(3) NOT NULL DEFAULT '0',
  `origin` text COLLATE utf8_unicode_ci NOT NULL,
  `revision` text COLLATE utf8_unicode_ci NOT NULL,
  `diff` text COLLATE utf8_unicode_ci NOT NULL,
  `remark` varchar(200) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `username` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `updated_uid` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_username` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `version` (`version`),
  KEY `annotation_id` (`annotation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_logs`
--

LOCK TABLES `annotation_logs` WRITE;
/*!40000 ALTER TABLE `annotation_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_mutations`
--

DROP TABLE IF EXISTS `annotation_mutations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_mutations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `mutation_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `annotation_mutation` (`annotation_id`,`mutation_id`),
  KEY `mutation_id` (`mutation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_mutations`
--

LOCK TABLES `annotation_mutations` WRITE;
/*!40000 ALTER TABLE `annotation_mutations` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_mutations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotation_trials`
--

DROP TABLE IF EXISTS `annotation_trials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotation_trials` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `annotation_id` int(10) unsigned NOT NULL DEFAULT '0',
  `nct_number` varchar(12) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `title` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `phases` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `status` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `url` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `biomarker` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `drug` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `disease` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `location` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `username` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `updated_uid` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_username` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `nct_number` (`nct_number`),
  KEY `annotation_id` (`annotation_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotation_trials`
--

LOCK TABLES `annotation_trials` WRITE;
/*!40000 ALTER TABLE `annotation_trials` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotation_trials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `annotations`
--

DROP TABLE IF EXISTS `annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `annotations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `origin_id` int(11) NOT NULL DEFAULT '0',
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `is_current` tinyint(3) NOT NULL DEFAULT '0',
  `version` bigint(20) unsigned NOT NULL DEFAULT '0',
  `annotation_type` tinyint(4) NOT NULL DEFAULT '0',
  `catalog` tinyint(4) NOT NULL DEFAULT '0',
  `mutation_group_id` int(11) unsigned NOT NULL DEFAULT '0',
  `biomarker` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `variant_group` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `genes_text` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `mutations_text` text COLLATE utf8_unicode_ci NOT NULL,
  `diseases_text` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `drugs_text` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `approve` text COLLATE utf8_unicode_ci NOT NULL,
  `guide` text COLLATE utf8_unicode_ci NOT NULL,
  `clinical` text COLLATE utf8_unicode_ci NOT NULL,
  `prompt` text COLLATE utf8_unicode_ci NOT NULL,
  `comment` text COLLATE utf8_unicode_ci NOT NULL,
  `remark` text COLLATE utf8_unicode_ci NOT NULL,
  `references` text COLLATE utf8_unicode_ci NOT NULL,
  `status` smallint(3) NOT NULL DEFAULT '0',
  `attributes` smallint(3) NOT NULL COMMENT '阳性-2;阴性-1;未知-0',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `username` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `updated_uid` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_username` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `version` (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotations`
--

LOCK TABLES `annotations` WRITE;
/*!40000 ALTER TABLE `annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cosmics`
--

DROP TABLE IF EXISTS `cosmics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cosmics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mutation_id` int(10) unsigned NOT NULL DEFAULT '0',
  `multi` tinyint(3) unsigned NOT NULL,
  `cosmic` varchar(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `cosmic` (`cosmic`) USING BTREE,
  KEY `mutation_id` (`mutation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cosmics`
--

LOCK TABLES `cosmics` WRITE;
/*!40000 ALTER TABLE `cosmics` DISABLE KEYS */;
/*!40000 ALTER TABLE `cosmics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `drug_group_drug`
--

DROP TABLE IF EXISTS `drug_group_drug`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `drug_group_drug` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `drug_group_id` int(11) unsigned NOT NULL DEFAULT '0',
  `drug_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `drug_group_id` (`drug_group_id`),
  KEY `drug_id` (`drug_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `drug_group_drug`
--

LOCK TABLES `drug_group_drug` WRITE;
/*!40000 ALTER TABLE `drug_group_drug` DISABLE KEYS */;
/*!40000 ALTER TABLE `drug_group_drug` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `drug_groups`
--

DROP TABLE IF EXISTS `drug_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `drug_groups` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `annotation_id` int(11) NOT NULL DEFAULT '0',
  `drugs_text` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `drug_ids` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `target` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `significance_type` varchar(30) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `significance` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `level` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `old_level` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `remark` text COLLATE utf8_unicode_ci NOT NULL,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `username` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `updated_uid` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_username` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `annotation_id` (`annotation_id`),
  KEY `level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `drug_groups`
--

LOCK TABLES `drug_groups` WRITE;
/*!40000 ALTER TABLE `drug_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `drug_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `druginfo_diseases`
--

DROP TABLE IF EXISTS `druginfo_diseases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `druginfo_diseases` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `druginfo_id` int(11) NOT NULL DEFAULT '0',
  `disease_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `druginfo_disease` (`disease_id`,`druginfo_id`),
  KEY `druginfo_id` (`druginfo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `druginfo_diseases`
--

LOCK TABLES `druginfo_diseases` WRITE;
/*!40000 ALTER TABLE `druginfo_diseases` DISABLE KEYS */;
/*!40000 ALTER TABLE `druginfo_diseases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `druginfo_genes`
--

DROP TABLE IF EXISTS `druginfo_genes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `druginfo_genes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `druginfo_id` int(11) NOT NULL DEFAULT '0',
  `gene_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `druginfo_gene` (`gene_id`,`druginfo_id`),
  KEY `druginfo_id` (`druginfo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `druginfo_genes`
--

LOCK TABLES `druginfo_genes` WRITE;
/*!40000 ALTER TABLE `druginfo_genes` DISABLE KEYS */;
/*!40000 ALTER TABLE `druginfo_genes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `druginfo_mutations`
--

DROP TABLE IF EXISTS `druginfo_mutations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `druginfo_mutations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `druginfo_id` int(11) NOT NULL DEFAULT '0',
  `mutation_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `druginfo_mutation` (`mutation_id`,`druginfo_id`),
  KEY `druginfo_id` (`druginfo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `druginfo_mutations`
--

LOCK TABLES `druginfo_mutations` WRITE;
/*!40000 ALTER TABLE `druginfo_mutations` DISABLE KEYS */;
/*!40000 ALTER TABLE `druginfo_mutations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gene_transcripts`
--

DROP TABLE IF EXISTS `gene_transcripts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gene_transcripts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gene` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `transcript` varchar(20) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `sign` varchar(1) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `is_exists` tinyint(3) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `gene` (`gene`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gene_transcripts`
--

LOCK TABLES `gene_transcripts` WRITE;
/*!40000 ALTER TABLE `gene_transcripts` DISABLE KEYS */;
/*!40000 ALTER TABLE `gene_transcripts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genes`
--

DROP TABLE IF EXISTS `genes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `genes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gain` tinyint(4) NOT NULL DEFAULT '0',
  `loss` tinyint(4) NOT NULL DEFAULT '0',
  `entrez_id` int(11) NOT NULL DEFAULT '0',
  `civic_gene_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(30) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `alias` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `chrom` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `grch37_start` int(11) NOT NULL DEFAULT '0',
  `grch37_end` int(11) NOT NULL DEFAULT '0',
  `strand` int(11) NOT NULL DEFAULT '0',
  `variant_count` int(11) NOT NULL DEFAULT '0',
  `evidence_count` int(11) NOT NULL DEFAULT '0',
  `transcript` varchar(30) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `oncokb_oncogene` tinyint(3) NOT NULL DEFAULT '0',
  `oncokb_tsg` tinyint(3) NOT NULL DEFAULT '0',
  `cosmic_oncogene` tinyint(3) NOT NULL DEFAULT '0',
  `cosmic_tsg` tinyint(3) NOT NULL DEFAULT '0',
  `tsgs_tsg` tinyint(3) NOT NULL DEFAULT '0',
  `intogen_oncogene` tinyint(3) NOT NULL DEFAULT '0',
  `intogen_tsg` tinyint(3) NOT NULL DEFAULT '0',
  `oncogene` tinyint(3) NOT NULL DEFAULT '0',
  `tsg` tinyint(3) NOT NULL DEFAULT '0',
  `omim_id` int(11) NOT NULL DEFAULT '0',
  `hgnc_id` int(11) NOT NULL DEFAULT '0',
  `nbk_id` int(11) NOT NULL DEFAULT '0',
  `orpha_id` int(11) NOT NULL DEFAULT '0',
  `hpo_id` int(11) NOT NULL DEFAULT '0',
  `description` text COLLATE utf8_unicode_ci NOT NULL,
  `ref_urls` text COLLATE utf8_unicode_ci NOT NULL,
  `interpro` text COLLATE utf8_unicode_ci NOT NULL,
  `pathways` text COLLATE utf8_unicode_ci NOT NULL,
  `mechanism` text COLLATE utf8_unicode_ci NOT NULL,
  `evidence` text COLLATE utf8_unicode_ci NOT NULL,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `username` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `updated_uid` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_username` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `genes_entrez_id_index` (`entrez_id`),
  KEY `genes_civic_gene_id_index` (`civic_gene_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genes`
--

LOCK TABLES `genes` WRITE;
/*!40000 ALTER TABLE `genes` DISABLE KEYS */;
/*!40000 ALTER TABLE `genes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-14  7:48:10
