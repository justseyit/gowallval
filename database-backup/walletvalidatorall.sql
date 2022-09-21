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
-- Table structure for table `currency`
--

DROP TABLE IF EXISTS `currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `currency` (
  `currencyID` int NOT NULL AUTO_INCREMENT,
  `currencySymbol` text NOT NULL,
  `currencyName` text NOT NULL,
  PRIMARY KEY (`currencyID`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `currency`
--

LOCK TABLES `currency` WRITE;
/*!40000 ALTER TABLE `currency` DISABLE KEYS */;
INSERT INTO `currency` VALUES (1,'BTC','Bitcoin'),(2,'ALGO','Algorand'),(3,'AE','Æternity'),(4,'ARDR','Ardor'),(5,'ATOM','Cosmos'),(6,'BTS','BitShare'),(7,'ADA','Cardano'),(8,'EOS','EOS'),(9,'ETH','Ethereum'),(10,'HBAR','Hedera Hashgraph'),(11,'ICX','ICON'),(12,'IOST','IOST'),(13,'LSK','Lisk'),(14,'XLM','Stellar Lumens'),(15,'XNO','Nano'),(16,'XEM','NEM'),(17,'XRP','Ripple'),(18,'SYS','Syscoin'),(19,'STEEM','Steem'),(20,'ZIL','Zilliqa'),(21,'1INCH','1INCH'),(22,'AAVE','AAVE'),(23,'ANT','Aragon'),(24,'AXS','Axie Infinity'),(25,'BADGER','Badger DAO'),(26,'BAL','Balancer'),(27,'SC','Siacoin'),(28,'BAND','BAND'),(29,'BAT','Basic Attention Token'),(30,'BCH','Bitcoin Cash'),(31,'BNT','Bancor'),(32,'CHZ','Chiliz'),(33,'COMP','Compound'),(34,'CRV','Curve'),(35,'DAI','Dai'),(36,'ENJ','Enjin Coin'),(37,'GNO','Gnosis'),(38,'GRT','The Graph'),(39,'KNC','KyberNetwork'),(40,'LINK','ChainLink'),(41,'LRC','Loopring'),(42,'LTC','Litecoin'),(43,'MATIC','Polygon'),(44,'MIR','Mirror Protocol'),(45,'MKR','Maker'),(46,'OCEAN','Ocean Protocol'),(47,'OGN','Origin Token'),(48,'OMG','OMG Network'),(49,'OXT','Orchid'),(50,'PAXG','PAX Gold'),(51,'REP','Augur v2'),(52,'SAND','The Sandbox'),(53,'SNX','Synthetix Network Token'),(54,'STORJ','Storj'),(55,'SUSHI','Sushi'),(56,'UNI','Uniswap'),(57,'USDT','TetherUS'),(58,'YFI','yearn.finance'),(59,'ZEC','Zcash'),(60,'ZRX','0x'),(61,'XMR','Monero'),(62,'XTZ','Tezos'),(63,'CQT','Covalent'),(64,'KEEP','Keep Network');
/*!40000 ALTER TABLE `currency` ENABLE KEYS */;
UNLOCK TABLES;

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
INSERT INTO `currencynetwork` VALUES (3,1),(2,2),(4,3),(5,4),(1,5),(7,5),(8,5),(9,5),(14,5),(17,5),(20,5),(21,5),(22,5),(24,5),(26,5),(28,5),(29,5),(30,5),(31,5),(33,5),(35,5),(39,5),(40,5),(41,5),(42,5),(43,5),(44,5),(45,5),(46,5),(47,5),(52,5),(53,5),(55,5),(56,5),(57,5),(58,5),(59,5),(62,5),(1,6),(1,7),(7,7),(8,7),(9,7),(17,7),(22,7),(28,7),(29,7),(30,7),(32,7),(33,7),(35,7),(40,7),(42,7),(45,7),(53,7),(55,7),(56,7),(57,7),(58,7),(59,7),(62,7),(1,8),(9,8),(17,8),(21,8),(22,8),(23,8),(24,8),(25,8),(26,8),(28,8),(29,8),(30,8),(31,8),(32,8),(33,8),(34,8),(35,8),(36,8),(37,8),(38,8),(39,8),(40,8),(41,8),(43,8),(44,8),(45,8),(46,8),(47,8),(48,8),(49,8),(50,8),(51,8),(52,8),(53,8),(54,8),(55,8),(56,8),(57,8),(58,8),(60,8),(63,8),(64,8),(1,9),(6,10),(7,11),(8,12),(10,13),(11,14),(12,15),(13,16),(14,17),(15,18),(16,19),(17,20),(19,22),(18,23),(20,24),(24,25),(30,26),(43,27),(52,27),(57,27),(42,28),(59,29),(57,30),(57,31),(57,32),(62,32),(61,35),(28,36);
/*!40000 ALTER TABLE `currencynetwork` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `network`
--

DROP TABLE IF EXISTS `network`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `network` (
  `networkID` int NOT NULL AUTO_INCREMENT,
  `networkSymbol` text NOT NULL,
  `networkName` text NOT NULL,
  PRIMARY KEY (`networkID`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `network`
--

LOCK TABLES `network` WRITE;
/*!40000 ALTER TABLE `network` DISABLE KEYS */;
INSERT INTO `network` VALUES (1,'AE','Æternity'),(2,'ALGO','Algorand'),(3,'ARDR','Ardor'),(4,'ATOM','Cosmos'),(5,'BSC','Binance Smart Chain (BEP20)'),(6,'BTC','Bitcoin'),(7,'BNB','Binance Beacon Chain (BEP2)'),(8,'ETH','Ethereum (ERC20)'),(9,'SEGWITBTC','Bitcoin (SegWit)'),(10,'BTS','BitShares'),(11,'ADA','Cardano'),(12,'EOS','EOS'),(13,'HBAR','Hedera Hashgraph'),(14,'ICX','ICON'),(15,'IOST','IOST'),(16,'LSK','Lisk'),(17,'XLM','Stellar Lumens'),(18,'NANO','Nano'),(19,'XEM','NEM'),(20,'XRP','Ripple'),(21,'SC','Siacoin'),(22,'STEEM','Steem'),(23,'SYS','Syscoin'),(24,'ZIL','Zilliqa'),(25,'RON','Ronin'),(26,'BCH','Bitcoin Cash'),(27,'MATIC','Polygon'),(28,'LTC','Litecoin'),(29,'ZEC','Zcash'),(30,'AVAXC','AVAX C-Chain'),(31,'SOL','Solana'),(32,'XTZ','Tezos'),(33,'TRX','Tron (TRC20)'),(34,'DASH','Dash'),(35,'XMR','Monero'),(36,'BAND','BAND');
/*!40000 ALTER TABLE `network` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'walletvalidator'
--

--
-- Dumping routines for database 'walletvalidator'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-22  2:50:35
