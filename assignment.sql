-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: localhost    Database: assignment
-- ------------------------------------------------------
-- Server version	8.0.23

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
-- Table structure for table `inventories`
--

DROP TABLE IF EXISTS `inventories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inventories` (
  `sku` varchar(36) NOT NULL,
  `name` varchar(200) NOT NULL,
  `price` decimal(10,2) unsigned NOT NULL,
  `qty` int unsigned NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`sku`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
INSERT INTO `inventories` VALUES ('120P90','Google Home',49.99,10,'2022-04-14 14:59:17.154371','2022-04-14 14:59:17.154371'),('234234','Raspberry Pi B',30.00,2,'2022-04-14 14:59:17.154371','2022-04-14 14:59:17.154371'),('43N23P','Macbook Pro',5399.99,5,'2022-04-14 14:59:17.154371','2022-04-14 14:59:17.154371'),('A304SD','Alexa Speaker',109.50,10,'2022-04-14 14:59:17.154371','2022-04-14 14:59:17.154371');
/*!40000 ALTER TABLE `inventories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotions`
--

DROP TABLE IF EXISTS `promotions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `promotions` (
  `id` varchar(36) NOT NULL,
  `item_sku` varchar(6) NOT NULL,
  `min_qty` int unsigned NOT NULL,
  `typ` varchar(45) NOT NULL,
  `uom` varchar(6) DEFAULT NULL,
  `value` decimal(4,2) unsigned NOT NULL,
  `is_active` tinyint NOT NULL DEFAULT '1',
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  KEY `idx` (`item_sku`,`is_active`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotions`
--

LOCK TABLES `promotions` WRITE;
/*!40000 ALTER TABLE `promotions` DISABLE KEYS */;
INSERT INTO `promotions` VALUES ('7465bac3-6825-487a-86af-1b10c997f045','A304SD',3,'DISCOUNT',NULL,10.00,1,'2022-04-14 15:06:45.182390','2022-04-14 15:06:45.182390'),('8efbff81-cf27-45bf-bb54-3a71a9b75372','43N23P',1,'FREE_ITEM','234234',1.00,1,'2022-04-14 15:06:45.182390','2022-04-14 15:06:45.182390'),('b5b0a253-7ca1-4b5a-8e71-ed0edc892dc0','120P90',3,'BOGO_SCHEME',NULL,1.00,1,'2022-04-14 15:06:45.182390','2022-04-14 15:06:45.182390');
/*!40000 ALTER TABLE `promotions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-16 16:51:29
