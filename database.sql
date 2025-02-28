-- MySQL dump 10.13  Distrib 8.0.41, for Win64 (x86_64)
--
-- Host: localhost    Database: kcal_calculator
-- ------------------------------------------------------
-- Server version	8.0.41

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `foods`
--

DROP TABLE IF EXISTS `foods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `foods` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `calories` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=319 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foods`
--

LOCK TABLES `foods` WRITE;
/*!40000 ALTER TABLE `foods` DISABLE KEYS */;
INSERT INTO `foods` VALUES (1,'Apple',52),(2,'Banana',89),(3,'Orange',47),(4,'Strawberry',32),(5,'Grapes',69),(6,'Watermelon',30),(7,'Blueberry',57),(8,'Pineapple',50),(9,'Mango',60),(10,'Peach',39),(11,'Cherry',50),(12,'Pear',57),(13,'Plum',46),(14,'Kiwi',61),(15,'Papaya',43),(16,'Lemon',29),(17,'Lime',30),(18,'Avocado',160),(19,'Tomato',18),(20,'Cucumber',16),(21,'Carrot',41),(22,'Broccoli',34),(23,'Spinach',23),(24,'Lettuce',15),(25,'Potato',77),(26,'Sweet Potato',86),(27,'Onion',40),(28,'Garlic',149),(29,'Bell Pepper',31),(30,'Zucchini',17),(31,'Eggplant',25),(32,'Mushroom',22),(33,'Pumpkin',26),(34,'Corn',86),(35,'Peas',81),(36,'Green Beans',31),(37,'Asparagus',20),(38,'Cauliflower',25),(39,'Cabbage',25),(40,'Brussels Sprouts',43),(41,'Kale',49),(42,'Celery',16),(43,'Radish',16),(44,'Beetroot',43),(45,'Turnip',28),(46,'Parsnip',75),(47,'Rutabaga',37),(48,'Artichoke',47),(49,'Leek',61),(50,'Okra',33),(51,'Bok Choy',13),(52,'Swiss Chard',19),(53,'Collard Greens',32),(54,'Mustard Greens',27),(55,'Arugula',25),(56,'Endive',17),(57,'Radicchio',23),(58,'Watercress',11),(59,'Chicory',23),(60,'Dandelion Greens',45),(61,'Fennel',31),(62,'Sorrel',22),(63,'Taro',142),(64,'Yam',118),(65,'Cassava',160),(66,'Plantain',122),(67,'Breadfruit',103),(68,'Jackfruit',95),(69,'Durian',147),(70,'Lychee',66),(71,'Rambutan',68),(72,'Longan',60),(73,'Mangosteen',73),(74,'Passion Fruit',97),(75,'Guava',68),(76,'Pomegranate',83),(77,'Fig',74),(78,'Date',282),(79,'Raisin',299),(80,'Prune',240),(81,'Apricot',48),(82,'Nectarine',44),(83,'Persimmon',81),(84,'Cantaloupe',34),(85,'Honeydew',36),(86,'Cranberry',46),(87,'Blackberry',43),(88,'Raspberry',52),(89,'Gooseberry',44),(90,'Currant',56),(91,'Mulberry',43),(92,'Elderberry',73),(93,'Huckleberry',37),(94,'Boysenberry',50),(95,'Cloudberry',51),(96,'Salmonberry',47),(97,'Lingonberry',50),(98,'Sea Buckthorn',82),(99,'Acai Berry',70),(100,'Maqui Berry',65),(101,'Chicken breast',164),(102,'Chicken thigh',209),(103,'Chicken drumstick',119),(104,'Chicken wing',42),(105,'Turkey breast',135),(106,'Turkey thigh',161),(107,'Turkey drumstick',119),(108,'Turkey wing',42),(109,'Duck breast',337),(110,'Duck thigh',318),(111,'Duck drumstick',119),(112,'Duck wing',42),(113,'Goose breast',161),(114,'Goose thigh',161),(115,'Goose drumstick',119),(116,'Goose wing',42),(117,'Pork chop',206),(118,'Pork loin',143),(119,'Pork tenderloin',143),(120,'Pork ribs',172),(121,'Pork belly',518),(122,'Pork shoulder',212),(123,'Pork leg',242),(124,'Beef tenderloin',143),(125,'Beef sirloin',143),(126,'Beef ribeye',143),(127,'Beef striploin',143),(128,'Beef flank',143),(129,'Beef brisket',143),(130,'Beef chuck',143),(131,'Beef shank',143),(132,'Beef round',143),(133,'Beef short ribs',143),(134,'Beef tongue',143),(135,'Beef liver',143),(136,'Beef heart',143),(137,'Beef kidney',143),(138,'Beef tripe',143),(139,'Beef oxtail',143),(140,'Beef bone marrow',143),(141,'Lamb chop',206),(142,'Lamb loin',143),(143,'Lamb tenderloin',143),(144,'Lamb ribs',172),(145,'Lamb belly',518),(146,'Lamb shoulder',212),(147,'Lamb leg',242),(148,'Venison loin',143),(149,'Venison ribs',172),(150,'Venison belly',518),(151,'Venison shoulder',212),(152,'Venison leg',242),(153,'Rabbit loin',143),(154,'Rabbit ribs',172),(155,'Vegi bag',200),(156,'Vegi bag',200),(157,'Vegi bag',200),(158,'Vegi bag',200),(159,'Vegi bag',200),(160,'Vegi bag',200),(161,'Vegi bag',200),(162,'Vegi bag',100),(163,'Vegi bag',1900),(164,'Apple',52),(165,'Banana',89),(166,'Orange',47),(167,'Strawberry',32),(168,'Grapes',69),(169,'Watermelon',30),(170,'Blueberry',57),(171,'Pineapple',50),(172,'Mango',60),(173,'Peach',39),(174,'Cherry',50),(175,'Pear',57),(176,'Plum',46),(177,'Kiwi',61),(178,'Papaya',43),(179,'Lemon',29),(180,'Lime',30),(181,'Avocado',160),(182,'Tomato',18),(183,'Cucumber',16),(184,'Carrot',41),(185,'Broccoli',34),(186,'Spinach',23),(187,'Lettuce',15),(188,'Potato',77),(189,'Sweet Potato',86),(190,'Onion',40),(191,'Garlic',149),(192,'Bell Pepper',31),(193,'Zucchini',17),(194,'Eggplant',25),(195,'Mushroom',22),(196,'Pumpkin',26),(197,'Corn',86),(198,'Peas',81),(199,'Green Beans',31),(200,'Asparagus',20),(201,'Cauliflower',25),(202,'Cabbage',25),(203,'Brussels Sprouts',43),(204,'Kale',49),(205,'Celery',16),(206,'Radish',16),(207,'Beetroot',43),(208,'Turnip',28),(209,'Parsnip',75),(210,'Rutabaga',37),(211,'Artichoke',47),(212,'Leek',61),(213,'Okra',33),(214,'Bok Choy',13),(215,'Swiss Chard',19),(216,'Collard Greens',32),(217,'Mustard Greens',27),(218,'Arugula',25),(219,'Endive',17),(220,'Radicchio',23),(221,'Watercress',11),(222,'Chicory',23),(223,'Dandelion Greens',45),(224,'Fennel',31),(225,'Sorrel',22),(226,'Taro',142),(227,'Yam',118),(228,'Cassava',160),(229,'Plantain',122),(230,'Breadfruit',103),(231,'Jackfruit',95),(232,'Durian',147),(233,'Lychee',66),(234,'Rambutan',68),(235,'Longan',60),(236,'Mangosteen',73),(237,'Passion Fruit',97),(238,'Guava',68),(239,'Pomegranate',83),(240,'Fig',74),(241,'Date',282),(242,'Raisin',299),(243,'Prune',240),(244,'Apricot',48),(245,'Nectarine',44),(246,'Persimmon',81),(247,'Cantaloupe',34),(248,'Honeydew',36),(249,'Cranberry',46),(250,'Blackberry',43),(251,'Raspberry',52),(252,'Gooseberry',44),(253,'Currant',56),(254,'Mulberry',43),(255,'Elderberry',73),(256,'Huckleberry',37),(257,'Boysenberry',50),(258,'Cloudberry',51),(259,'Salmonberry',47),(260,'Lingonberry',50),(261,'Sea Buckthorn',82),(262,'Acai Berry',70),(263,'Maqui Berry',65),(264,'Chicken breast',164),(265,'Chicken thigh',209),(266,'Chicken drumstick',119),(267,'Chicken wing',42),(268,'Turkey breast',135),(269,'Turkey thigh',161),(270,'Turkey drumstick',119),(271,'Turkey wing',42),(272,'Duck breast',337),(273,'Duck thigh',318),(274,'Duck drumstick',119),(275,'Duck wing',42),(276,'Goose breast',161),(277,'Goose thigh',161),(278,'Goose drumstick',119),(279,'Goose wing',42),(280,'Pork chop',206),(281,'Pork loin',143),(282,'Pork tenderloin',143),(283,'Pork ribs',172),(284,'Pork belly',518),(285,'Pork shoulder',212),(286,'Pork leg',242),(287,'Beef tenderloin',143),(288,'Beef sirloin',143),(289,'Beef ribeye',143),(290,'Beef striploin',143),(291,'Beef flank',143),(292,'Beef brisket',143),(293,'Beef chuck',143),(294,'Beef shank',143),(295,'Beef round',143),(296,'Beef short ribs',143),(297,'Beef tongue',143),(298,'Beef liver',143),(299,'Beef heart',143),(300,'Beef kidney',143),(301,'Beef tripe',143),(302,'Beef oxtail',143),(303,'Beef bone marrow',143),(304,'Lamb chop',206),(305,'Lamb loin',143),(306,'Lamb tenderloin',143),(307,'Lamb ribs',172),(308,'Lamb belly',518),(309,'Lamb shoulder',212),(310,'Lamb leg',242),(311,'Venison loin',143),(312,'Venison ribs',172),(313,'Venison belly',518),(314,'Venison shoulder',212),(315,'Venison leg',242),(316,'Rabbit loin',143),(317,'Rabbit ribs',172),(318,'Apple',52);
/*!40000 ALTER TABLE `foods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `targetkcal`
--

DROP TABLE IF EXISTS `targetkcal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `targetkcal` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `targetKcal` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `targetkcal`
--

LOCK TABLES `targetkcal` WRITE;
/*!40000 ALTER TABLE `targetkcal` DISABLE KEYS */;
INSERT INTO `targetkcal` VALUES (1,948);
/*!40000 ALTER TABLE `targetkcal` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-27 14:13:49
