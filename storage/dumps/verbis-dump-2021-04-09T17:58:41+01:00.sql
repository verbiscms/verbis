-- MySQL dump 10.13  Distrib 8.0.23, for osx10.16 (x86_64)
--
-- Host: 127.0.0.1    Database: verbis
-- ------------------------------------------------------
-- Server version	8.0.23

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

DROP TABLE IF EXISTS `categories`;
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
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `form_fields`
--

DROP TABLE IF EXISTS `form_fields`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `form_fields` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `form_id` int NOT NULL,
  `key` varchar(150) NOT NULL,
  `label` varchar(150) NOT NULL,
  `type` varchar(150) NOT NULL,
  `validation` varchar(500) NOT NULL,
  `required` bit(1) NOT NULL DEFAULT b'0',
  `options` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `form_id_index` (`form_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `form_fields`
--

LOCK TABLES `form_fields` WRITE;
/*!40000 ALTER TABLE `form_fields` DISABLE KEYS */;
/*!40000 ALTER TABLE `form_fields` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `form_submissions`
--

DROP TABLE IF EXISTS `form_submissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `form_submissions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `form_id` int NOT NULL,
  `fields` json DEFAULT NULL,
  `ip_address` varchar(20) NOT NULL,
  `user_agent` varchar(1000) NOT NULL,
  `sent_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `form_id_index` (`form_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `form_submissions`
--

LOCK TABLES `form_submissions` WRITE;
/*!40000 ALTER TABLE `form_submissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `form_submissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forms`
--

DROP TABLE IF EXISTS `forms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `forms` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `name` varchar(150) NOT NULL,
  `email_send` bit(1) NOT NULL DEFAULT b'0',
  `email_message` mediumtext,
  `email_subject` varchar(78) NOT NULL,
  `recipients` varchar(1000) NOT NULL,
  `store_db` bit(1) NOT NULL DEFAULT b'0',
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forms`
--

LOCK TABLES `forms` WRITE;
/*!40000 ALTER TABLE `forms` DISABLE KEYS */;
/*!40000 ALTER TABLE `forms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `media`
--

DROP TABLE IF EXISTS `media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `media` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `url` varchar(255) NOT NULL,
  `title` varchar(150) NOT NULL,
  `alt` varchar(150) NOT NULL,
  `description` text NOT NULL,
  `file_path` varchar(255) NOT NULL,
  `file_size` int NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `sizes` json DEFAULT NULL,
  `mime` varchar(150) NOT NULL,
  `user_id` int NOT NULL,
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `media`
--

LOCK TABLES `media` WRITE;
/*!40000 ALTER TABLE `media` DISABLE KEYS */;
/*!40000 ALTER TABLE `media` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `options`
--

DROP TABLE IF EXISTS `options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `option_name` varchar(255) NOT NULL,
  `option_value` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `options`
--

LOCK TABLES `options` WRITE;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
INSERT INTO `options` VALUES (1,'site_logo','\"/verbis/images/verbis-logo.svg\"'),(2,'footer_disclosure','\"\"'),(3,'seo_robots_serve','true'),(4,'media_compression','80'),(5,'minify_js','false'),(6,'minify_css','false'),(7,'contact_address','\"\"'),(8,'seo_enforce_slash','false'),(9,'cache_frontend_extensions','[\"jpg\", \"jpeg\", \"gif\", \"png\", \"ico\", \"cur\", \"webp\", \"jxr\", \"svg\", \"css\", \"js\", \"htc\", \"ttf\", \"tt\", \"otf\", \"eot\", \"woff\", \"woff2\", \"webm\"]'),(10,'gzip_use_paths','false'),(11,'seo_sitemap_serve','true'),(12,'seo_sitemap_excluded','[]'),(13,'minify_xml','false'),(14,'meta_facebook_title','\"\"'),(15,'cache_server_templates','false'),(16,'minify_svg','false'),(17,'site_description','\"A Verbis website. Publish online, build a business, work from home\"'),(18,'general_locale','\"en_GB\"'),(19,'contact_telephone','\"\"'),(20,'footer_text','\"\"'),(21,'gzip','true'),(22,'gzip_compression','\"default-compression\"'),(23,'social_instagram','\"\"'),(24,'codeinjection_head','\"\"'),(25,'site_title','\"Verbis\"'),(26,'meta_facebook_description','\"\"'),(27,'media_upload_max_size','100000'),(28,'media_upload_max_height','0'),(29,'cache_frontend_seconds','31536000'),(30,'gzip_excluded_paths','[]'),(31,'social_pinterest','\"\"'),(32,'codeinjection_foot','\"\"'),(33,'media_images_sizes','{\"hd\": {\"crop\": false, \"name\": \"HD Size\", \"width\": 1920, \"height\": 0}, \"large\": {\"crop\": false, \"name\": \"Large Size\", \"width\": 1280, \"height\": 0}, \"medium\": {\"crop\": false, \"name\": \"Medium Size\", \"width\": 992, \"height\": 0}, \"thumbnail\": {\"crop\": true, \"name\": \"Thumbnail Size\", \"width\": 550, \"height\": 300}}'),(34,'gzip_excluded_extensions','[]'),(35,'social_facebook','\"\"'),(36,'social_youtube','\"\"'),(37,'meta_twitter_title','\"\"'),(38,'media_serve_webp','true'),(39,'minify_html','false'),(40,'meta_twitter_description','\"\"'),(41,'meta_twitter_image_id','0'),(42,'seo_private','false'),(43,'homepage','0'),(44,'media_organise_year_month','true'),(45,'cache_frontend','true'),(46,'cache_frontend_request','\"max-age\"'),(47,'active_theme','\"Verbis\"'),(48,'contact_email','\"\"'),(49,'seo_robots','\"User-agent: *\\nAllow: /\"'),(50,'site_url','\"http://127.0.0.1:8080\"'),(51,'social_twitter','\"\"'),(52,'seo_sitemap_redirects','true'),(53,'meta_title','\"\"'),(54,'media_convert_webp','true'),(55,'media_upload_max_width','0'),(56,'minify_json','false'),(57,'social_linkedin','\"\"'),(58,'meta_description','\"\"'),(59,'meta_facebook_image_id','0'),(60,'cache_server_field_layouts','false');
/*!40000 ALTER TABLE `options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `password_resets`
--

DROP TABLE IF EXISTS `password_resets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_resets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_resets`
--

LOCK TABLES `password_resets` WRITE;
/*!40000 ALTER TABLE `password_resets` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_resets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_categories`
--

DROP TABLE IF EXISTS `post_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_categories` (
  `category_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`category_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_categories`
--

LOCK TABLES `post_categories` WRITE;
/*!40000 ALTER TABLE `post_categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_fields`
--

DROP TABLE IF EXISTS `post_fields`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_fields` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `uuid` varchar(36) NOT NULL,
  `type` varchar(250) NOT NULL,
  `name` varchar(250) NOT NULL,
  `field_key` varchar(500) NOT NULL,
  `value` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `post_fields_unique` (`uuid`,`field_key`,`post_id`,`name`),
  KEY `post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_fields`
--

LOCK TABLES `post_fields` WRITE;
/*!40000 ALTER TABLE `post_fields` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_fields` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_options`
--

DROP TABLE IF EXISTS `post_options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post_options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `seo` json DEFAULT NULL,
  `meta` json DEFAULT NULL,
  `edit_lock` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_options`
--

LOCK TABLES `post_options` WRITE;
/*!40000 ALTER TABLE `post_options` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `title` varchar(500) NOT NULL,
  `slug` varchar(150) NOT NULL,
  `status` varchar(150) NOT NULL DEFAULT 'draft',
  `resource` varchar(150) NOT NULL,
  `page_template` varchar(150) NOT NULL DEFAULT 'default',
  `layout` varchar(150) NOT NULL DEFAULT 'default',
  `codeinjection_head` longtext NOT NULL,
  `codeinjection_foot` longtext NOT NULL,
  `user_id` int NOT NULL,
  `published_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `redirects`
--

DROP TABLE IF EXISTS `redirects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `redirects` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_path` varchar(255) NOT NULL,
  `to_path` varchar(255) NOT NULL,
  `code` int NOT NULL,
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `redirect_from_path` (`from_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `redirects`
--

LOCK TABLES `redirects` WRITE;
/*!40000 ALTER TABLE `redirects` DISABLE KEYS */;
/*!40000 ALTER TABLE `redirects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Banned','The user has been banned from the system.'),(2,'Contributor','The user can create and edit their own draft posts, but they are unable to edit drafts of users or published posts.'),(3,'Author','The user can write, edit and publish their own posts.'),(4,'Editor','The user can do everything defined in the Author role but they can also edit and publish posts of others, as well as their own.'),(5,'Administrator','The user can do everything defined in the Editor role but they can also edit site settings and data. Additionally they can manage users'),(6,'Owner','The user is a special user with all of the permissions as an Administrator however they cannot be deleted');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,6);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `first_name` varchar(150) NOT NULL,
  `last_name` varchar(150) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(60) NOT NULL,
  `website` text NOT NULL,
  `facebook` text NOT NULL,
  `twitter` text NOT NULL,
  `linked_in` text NOT NULL,
  `instagram` text NOT NULL,
  `biography` longtext NOT NULL,
  `profile_picture_id` int DEFAULT NULL,
  `token` varchar(38) NOT NULL,
  `token_last_used` timestamp NULL DEFAULT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'384afec1-94bd-46b1-9125-b853fa8f5a1f','AInsley','Clark','ainsley@reddico.co.uk','$2a$10$w6r7/4Wu8tuwgx0vcrs8vuPpQJ2JkGdBMp1dnEsSt8XiWERlDo2eG','','AInsley','','','','',NULL,'34b9c4c7fe05c89e20cda06a772cbbc1eb141',NULL,NULL,'2021-04-01 14:07:07','2021-04-01 14:07:07');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-04-09 17:58:41
