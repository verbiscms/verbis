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
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
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
  `updated_at` timestamp NULL DEFAULT NULL,
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
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table options
--

LOCK TABLES options WRITE;
/*!40000 ALTER TABLE options DISABLE KEYS */;

INSERT INTO options VALUES ('21','gzip_compression','true'),('22','site_title','"Verbis"'),('23','site_description','"A Verbis website. Publish online, build a business, work from home"'),('24','site_url','"http://localhost:8080"'),('25','media_convert_webp','true'),('26','media_upload_max_size','100000'),('27','media_images_sizes','{"hd": {"url": "", "crop": false, "name": "HD Size", "uuid": "00000000-0000-0000-0000-000000000000", "width": 1920, "height": 0, "file_size": 0}, "large": {"url": "", "crop": false, "name": "Large Size", "uuid": "00000000-0000-0000-0000-000000000000", "width": 1280, "height": 0, "file_size": 0}, "medium": {"url": "", "crop": false, "name": "Medium Size", "uuid": "00000000-0000-0000-0000-000000000000", "width": 992, "height": 0, "file_size": 0}, "thumbnail": {"url": "", "crop": true, "name": "Thumbnail Size", "uuid": "00000000-0000-0000-0000-000000000000", "width": 550, "height": 300, "file_size": 0}}'),('28','cache_global','true'),('29','media_upload_max_width','0'),('30','cache_layout','true'),('31','cache_site','true'),('32','media_serve_webp','true'),('33','media_upload_max_height','0'),('34','media_organise_year_month','true'),('35','cache_fields','true'),('36','cache_resources','true'),('37','site_logo','"/verbis/images/verbis-logo.svg"'),('38','media_compression','80'),('39','cache_templates','true'),('40','gzip_files','["text/css", "text/javascript", "text/xml", "text/plain", "text/x-component", "application/javascript", "application/json", "application/xml", "application/rss+xml", "font/truetype", "font/opentype", "application/vnd.ms-fontobject", "image/svg+xml"]');

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
  `title` varchar(500) NOT NULL,
  `slug` varchar(150) NOT NULL,
  `status` varchar(150) NOT NULL DEFAULT 'draft',
  `resource` varchar(150) DEFAULT NULL,
  `page_template` varchar(150) NOT NULL DEFAULT 'default',
  `layout` varchar(150) NOT NULL DEFAULT 'default',
  `fields` json DEFAULT NULL,
  `codeinjection_head` longtext,
  `codeinjection_foot` longtext,
  `user_id` int NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table posts
--

LOCK TABLES posts WRITE;
/*!40000 ALTER TABLE posts DISABLE KEYS */;

INSERT INTO posts VALUES ('1','56f4d11f-d173-4ebd-a627-280064793402','Title','url','','','','default','{}','','','2','2020-10-05T11:22:13Z','2020-10-05T11:22:13Z'),('2','3e9fb973-446f-48a6-a7d3-4c20db16e82d','Titlehhh','urlhh','','','','default','{}','','','2','2020-10-05T11:26:01Z','2020-10-05T11:26:01Z'),('3','6530b15a-dcf1-4803-8b92-9274cff436a0','Titlehhhgjhkhjkhjk','gurlhh','','','','default','{"text": "hello mothAppend", "number": "2045243545345", "flexible": {"layoutkey1": {"text2": "jhgghg"}, "layoutkey2": {"text1": "hjghjg", "text2": "ghvghhggh"}}, "repeater": [{"text": "hello mother fucker", "text2": "hjvgdsjfgasdfjhsadgfjgjg"}, {"text": "fucker fuck", "text2": "kjhkjhjkhkjh"}, {}], "richtext": "<p>kjhhkhkjkjkkjhkjh</p>", "textarea": "ljok;jhlj"}','','','2','2020-10-05T19:53:50Z','2020-10-05T11:26:29Z'),('4','07db5214-7b65-41f0-9c07-b57095f5997a','the page title','flexibletest','','','','default','{"text": "Append", "email": "ffggfdgfdgfdsggfdAppend", "range": "32010", "number": "20", "flexible": [], "repeater": [{"text": "hhhh", "text2": "hhh"}], "richtext": "", "textarea": "", "post_object": ["verbis_post_1", "verbis_post_4"]}','','','2','2020-10-06T17:47:24Z','2020-10-06T09:48:10Z'),('5','9e6c62d6-aeb9-4d7d-8b1b-93dca8dd2a53','gghjhjggh','thenewurl','','','','default','{"user": ["verbis_user_2"], "email": "Append", "flexible": [{"type": "layoutkey1", "fields": {"text": "jhjhhgjhjg", "text2": ""}}, {"type": "layoutkey2", "fields": {"text": "ghjjhghj", "text1": "", "text2": "ghhjghjjhg"}}], "repeater": [{}]}','','','2','2020-10-07T14:48:30Z','2020-10-06T10:18:29Z');

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

INSERT INTO roles VALUES ('1','Banned','The user has been banned from the system.'),('2','Contributor','The user can create and edit their own draft posts, but they are unable to edit drafts of users or published posts.'),('3','Author','The user can write, edit and publish their own posts.'),('4','Editor','The user can do everything defined in the Author role but they can also edit and publish posts of others, as well as their own.'),('5','Administrator','The user can do everything defined in the Editor role but they can also edit site settings and data. Additionally they can manage users'),('6','Owner','The user is a special user with all of the permissions as an Administrator however they cannot be deleted');

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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table seo_meta_options
--

LOCK TABLES seo_meta_options WRITE;
/*!40000 ALTER TABLE seo_meta_options DISABLE KEYS */;

INSERT INTO seo_meta_options VALUES ('1','1','',''),('2','2','',''),('3','3','',''),('4','4','',''),('5','5','','');

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
  `updated_at` timestamp NULL DEFAULT NULL,
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

INSERT INTO user_roles VALUES ('2','6');

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
  `login_time` timestamp NULL DEFAULT NULL,
  `last_seen_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table user_sessions
--

LOCK TABLES user_sessions WRITE;
/*!40000 ALTER TABLE user_sessions DISABLE KEYS */;

INSERT INTO user_sessions VALUES ('1','2','e5d857cf26f2ca3dac98f810cabbce43','2020-10-05T11:22:02Z','2020-10-05T11:22:02Z'),('2','2','434631d24a12a08e26a4629d59fc92f3','2020-10-08T14:03:09Z','2020-10-08T14:03:09Z');

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
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Dumping data for table users
--

LOCK TABLES users WRITE;
/*!40000 ALTER TABLE users DISABLE KEYS */;

INSERT INTO users VALUES ('2','fb41286d-3033-4aa0-94c7-60f2a01352b3','Ainsley','Clark','ainsley@reddico.co.uk','$2a$10$s2UO4AoRikdeZDUsvlkU/.zhAmtgXzP5uSNOAMHZUqNJnAWR0sMBy','','','','','','345ca2b77f1bf7e58432d6f2016176f67a141','','2020-10-05T11:21:34Z','2020-10-05T11:21:34Z');

/*!40000 ALTER TABLE users ENABLE KEYS */;
UNLOCK TABLES;

-- Dump completed on 2020-10-08 14:46:10.558786 +0100 BST m=+2.051871086
