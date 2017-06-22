# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.18)
# Database: yiigo
# Generation Time: 2017-06-20 08:40:36 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table yii_book
# ------------------------------------------------------------

DROP TABLE IF EXISTS `yii_book`;

CREATE TABLE `yii_book` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '书名',
  `subtitle` varchar(50) NOT NULL DEFAULT '' COMMENT '副标题',
  `author` varchar(50) NOT NULL DEFAULT '' COMMENT '作者',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '版本',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `publisher` varchar(50) NOT NULL DEFAULT '' COMMENT '出版社',
  `publish_date` varchar(50) NOT NULL DEFAULT '' COMMENT '出版日期',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `yii_book` WRITE;
/*!40000 ALTER TABLE `yii_book` DISABLE KEYS */;

INSERT INTO `yii_book` (`id`, `title`, `subtitle`, `author`, `version`, `price`, `publisher`, `publish_date`, `created_at`, `updated_at`)
VALUES
	(1,'PHP从入门到精通','','明日科技','第三版',69.80,'清华大学出版社','2012年9月','2017-06-19 17:23:13','2017-06-19 17:36:21'),
	(2,'Go语言编程','','许世伟 吕桂华','第一版',49.00,'人民邮电出版社','2012年8月','2017-06-19 17:23:51','2017-06-19 17:36:45'),
	(3,'C程序设计','新世纪计算机基础教育丛书','谭浩强','第三版',26.00,'清华大学出版社','2005年7月','2017-06-19 17:24:28','2017-06-19 17:35:22'),
	(4,'Docker技术入门与实战','','杨保华 戴王剑 曹亚仑','第一版',59.00,'机械工业出版社','2015年1月','2017-06-19 17:27:23','2017-06-19 17:34:10'),
	(5,'Go Web编程','','谢孟军','第一版',65.00,'电子工业出版社','2013年6月','2017-06-19 17:29:54','2017-06-19 17:33:57'),
	(6,'鸟哥的Linux私房菜','基础学习篇','鸟哥','第三版',88.00,'人民邮电出版社','2010年8月','2017-06-19 17:32:29','2017-06-19 17:38:08');

/*!40000 ALTER TABLE `yii_book` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
