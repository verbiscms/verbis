-- Go SQL Dump 0.2.2
--
-- ------------------------------------------------------
-- Server version	8.0.21

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
-- Table structure for table categories
--

DROP TABLE IF EXISTS categories;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `slug` varchar(150) NOT NULL,
  `name` varchar(150) NOT NULL,
  `description` text,
  `hidden` tinyint(1) NOT NULL DEFAULT '0',
  `parent_id` int DEFAULT NULL,
  `page_template` varchar(150) DEFAULT NULL,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table categories
--

LOCK TABLES categories WRITE;
/*!40000 ALTER TABLE categories DISABLE KEYS */;

/*!40000 ALTER TABLE categories ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table media
--

DROP TABLE IF EXISTS media;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `media` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  `title` varchar(150) DEFAULT NULL,
  `alt` varchar(150) DEFAULT NULL,
  `description` text,
  `file_path` varchar(255) NOT NULL,
  `file_size` int NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `sizes` json DEFAULT NULL,
  `type` varchar(150) NOT NULL,
  `user_id` int DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table media
--

LOCK TABLES media WRITE;
/*!40000 ALTER TABLE media DISABLE KEYS */;

/*!40000 ALTER TABLE media ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table migrations
--

DROP TABLE IF EXISTS migrations;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `migrations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) DEFAULT NULL,
  `batch` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table migrations
--

LOCK TABLES migrations WRITE;
/*!40000 ALTER TABLE migrations DISABLE KEYS */;

/*!40000 ALTER TABLE migrations ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table options
--

DROP TABLE IF EXISTS options;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `option_name` varchar(255) NOT NULL,
  `option_value` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table options
--

LOCK TABLES options WRITE;
/*!40000 ALTER TABLE options DISABLE KEYS */;

/*!40000 ALTER TABLE options ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table password_resets
--

DROP TABLE IF EXISTS password_resets;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `password_resets` (
  `email` varchar(255) NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table password_resets
--

LOCK TABLES password_resets WRITE;
/*!40000 ALTER TABLE password_resets DISABLE KEYS */;

/*!40000 ALTER TABLE password_resets ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table post_categories
--

DROP TABLE IF EXISTS post_categories;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post_categories` (
  `category_id` int NOT NULL,
  `post_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table post_categories
--

LOCK TABLES post_categories WRITE;
/*!40000 ALTER TABLE post_categories DISABLE KEYS */;

/*!40000 ALTER TABLE post_categories ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table posts
--

DROP TABLE IF EXISTS posts;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `slug` varchar(150) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `status` varchar(150) NOT NULL DEFAULT 'draft',
  `resource` varchar(150) DEFAULT NULL,
  `page_template` varchar(150) NOT NULL,
  `layout` varchar(150) NOT NULL,
  `fields` json DEFAULT NULL,
  `codeinjection_head` longtext,
  `codeinjection_foot` longtext,
  `user_id` int NOT NULL,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table posts
--

LOCK TABLES posts WRITE;
/*!40000 ALTER TABLE posts DISABLE KEYS */;

/*!40000 ALTER TABLE posts ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table roles
--

DROP TABLE IF EXISTS roles;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table roles
--

LOCK TABLES roles WRITE;
/*!40000 ALTER TABLE roles DISABLE KEYS */;

/*!40000 ALTER TABLE roles ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table seo_meta_options
--

DROP TABLE IF EXISTS seo_meta_options;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seo_meta_options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `page_id` int NOT NULL,
  `seo` json DEFAULT NULL,
  `meta` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table seo_meta_options
--

LOCK TABLES seo_meta_options WRITE;
/*!40000 ALTER TABLE seo_meta_options DISABLE KEYS */;

/*!40000 ALTER TABLE seo_meta_options ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table subscribers
--

DROP TABLE IF EXISTS subscribers;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subscribers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table subscribers
--

LOCK TABLES subscribers WRITE;
/*!40000 ALTER TABLE subscribers DISABLE KEYS */;

/*!40000 ALTER TABLE subscribers ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table user_roles
--

DROP TABLE IF EXISTS user_roles;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table user_roles
--

LOCK TABLES user_roles WRITE;
/*!40000 ALTER TABLE user_roles DISABLE KEYS */;

/*!40000 ALTER TABLE user_roles ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table user_sessions
--

DROP TABLE IF EXISTS user_sessions;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_sessions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `session_key` varchar(38) NOT NULL,
  `login_time` timestamp NOT NULL,
  `last_seen_time` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table user_sessions
--

LOCK TABLES user_sessions WRITE;
/*!40000 ALTER TABLE user_sessions DISABLE KEYS */;

/*!40000 ALTER TABLE user_sessions ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table users
--

DROP TABLE IF EXISTS users;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `first_name` varchar(150) NOT NULL,
  `last_name` varchar(150) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(60) NOT NULL,
  `website` text,
  `facebook` text,
  `twitter` text,
  `linked_in` text,
  `instagram` text,
  `token` varchar(38) NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table users
--

LOCK TABLES users WRITE;
/*!40000 ALTER TABLE users DISABLE KEYS */;

/*!40000 ALTER TABLE users ENABLE KEYS */;
UNLOCK TABLES;

-- Dump completed on 2020-09-23 09:16:07.742532 +0100 BST m=+1.010977452
