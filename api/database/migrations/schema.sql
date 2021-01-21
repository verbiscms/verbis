-- MySQL dump 10.13  Distrib 8.0.21, for osx10.15 (x86_64)
--
-- ------------------------------------------------------
-- Server version	8.0.21

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
-- Table structure for table `categories`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36) NOT NULL,
    `slug` varchar(150) NOT NULL,
    `name` varchar(150) NOT NULL,
    `description` text,
    `resource` varchar(150) NOT NULL,
    `parent_id` int DEFAULT NULL,
    `archive_id` int DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `forms`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `forms` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36) NOT NULL,
    `name` varchar(150) NOT NULL,
    `email_send` bit DEFAULT 0 NOT NULL,
    `email_message` mediumtext NULL,
    `email_subject` varchar(78) NOT NULL,
    `store_db` bit DEFAULT 0 NOT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `form_fields`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `form_fields` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36) NOT NULL,
    `form_id` int NOT NULL,
    `key` varchar(150) NOT NULL,
    `label` varchar(150) NOT NULL,
    `type` varchar(150) NOT NULL,
    `validation` varchar(500) NULL,
    `required` bit DEFAULT 0 NOT NULL,
    `options` json,
    PRIMARY KEY (`id`),
    KEY `form_id_index` (`form_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `form_submissions`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `form_submissions` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36) NOT NULL,
    `form_id` int NOT NULL,
    `fields` json,
    `ip_address` varchar(20) NOT NULL,
    `user_agent` varchar(36) NOT NULL,
    `sent_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `form_id_index` (`form_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `media`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
    `updated_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `options`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `options` (
    `id` int NOT NULL AUTO_INCREMENT,
    `option_name` varchar(255) NOT NULL,
    `option_value` json DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `password_resets`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_resets` (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL,
    `token` varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `post_categories`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_categories` (
    `category_id` int NOT NULL,
    `post_id` int NOT NULL,
    PRIMARY KEY (`category_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `post_fields`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_fields` (
    `id` int NOT NULL AUTO_INCREMENT,
    `post_id` int NOT NULL,
    `uuid` varchar(36) NOT NULL,
    `type` varchar(250) NOT NULL,
    `name` varchar(250) NOT NULL,
    `field_key` varchar(500) NOT NULL,
    `value` longtext NULL,
    PRIMARY KEY (`id`),
    KEY `post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;



--
-- Table structure for table `posts`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36) NOT NULL,
    `title` varchar(500) NOT NULL,
    `slug` varchar(150) NOT NULL,
    `status` varchar(150) NOT NULL DEFAULT 'draft',
    `resource` varchar(150) NULL,
    `page_template` varchar(150) NOT NULL DEFAULT 'default',
    `layout` varchar(150)  NOT NULL DEFAULT 'default',
    `codeinjection_head` longtext,
    `codeinjection_foot` longtext,
    `user_id` int NOT NULL,
    `published_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `slug` (`slug`),
    KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `post_options`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_options` (
    `id` int NOT NULL AUTO_INCREMENT,
    `post_id` int NOT NULL,
    `seo` json DEFAULT NULL,
    `meta` json DEFAULT NULL,
    `edit_lock` varchar(150) NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `roles`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_roles`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
    `user_id` int NOT NULL,
    `role_id` int NOT NULL,
    PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
    `biography` longtext,
    `profile_picture_id` INT NULL,
    `token` varchar(38) NOT NULL,
    `token_last_used` timestamp NULL DEFAULT NULL,
    `email_verified_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-09-22 14:13:46
