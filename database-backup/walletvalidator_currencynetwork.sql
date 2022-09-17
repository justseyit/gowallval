CREATE DATABASE  IF NOT EXISTS `walletvalidator` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `walletvalidator`;
-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: localhost    Database: walletvalidator
-- ------------------------------------------------------
-- Server version	8.0.30

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
-- Table structure for table `currencynetwork`
--

DROP TABLE IF EXISTS `currencynetwork`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `currencynetwork` (
  `currencyID` int NOT NULL,
  `networkID` int NOT NULL,
  PRIMARY KEY (`currencyID`,`networkID`),
  KEY `networkID` (`networkID`),
  CONSTRAINT `currencynetwork_ibfk_1` FOREIGN KEY (`currencyID`) REFERENCES `currency` (`currencyID`),
  CONSTRAINT `currencynetwork_ibfk_2` FOREIGN KEY (`networkID`) REFERENCES `network` (`networkID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `currencynetwork`
--

LOCK TABLES `currencynetwork` WRITE;
/*!40000 ALTER TABLE `currencynetwork` DISABLE KEYS */;
INSERT INTO `currencynetwork` VALUES (3,1),(2,2),(4,3),(5,4),(1,5),(7,5),(8,5),(9,5),(14,5),(17,5),(20,5),(21,5),(22,5),(24,5),(26,5),(1,6),(1,7),(7,7),(8,7),(9,7),(17,7),(22,7),(1,8),(9,8),(17,8),(21,8),(22,8),(23,8),(24,8),(25,8),(26,8),(1,9),(6,10),(7,11),(8,12),(10,13),(11,14),(12,15),(13,16),(14,17),(15,18),(16,19),(17,20),(19,22),(18,23),(20,24),(24,25);
/*!40000 ALTER TABLE `currencynetwork` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-17 18:27:16
