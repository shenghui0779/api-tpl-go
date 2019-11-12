#
# SQL Export
# Created by Querious (1068)
# Created: 2017年9月1日 GMT+8 22:20:11
# Encoding: Unicode (UTF-8)
#


SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `book`;


CREATE TABLE `book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '书名',
  `subtitle` varchar(50) NOT NULL DEFAULT '' COMMENT '副标题',
  `author` varchar(50) NOT NULL DEFAULT '' COMMENT '作者',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '版本',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `publisher` varchar(50) NOT NULL DEFAULT '' COMMENT '出版社',
  `publish_date` varchar(50) NOT NULL DEFAULT '' COMMENT '出版日期',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


LOCK TABLES `book` WRITE;
ALTER TABLE `book` DISABLE KEYS;
INSERT INTO `book` (`id`, `title`, `subtitle`, `author`, `version`, `price`, `publisher`, `publish_date`, `created_at`, `updated_at`) VALUES
	(1,'PHP从入门到精通','','明日科技','第三版',69.80,'清华大学出版社','2012年9月',1497864193,1497864193),
	(2,'Go语言编程','','许世伟 吕桂华','第一版',49.00,'人民邮电出版社','2012年8月',1497864193,1497864193),
	(3,'C程序设计','新世纪计算机基础教育丛书','谭浩强','第三版',26.00,'清华大学出版社','2005年7月',1497864193,1497864193),
	(4,'Docker技术入门与实战','','杨保华 戴王剑 曹亚仑','第一版',59.00,'机械工业出版社','2015年1月',1497864193,1497864193),
	(5,'Go Web编程','','谢孟军','第一版',65.00,'电子工业出版社','2013年6月',1497864193,1497864193),
	(6,'鸟哥的Linux私房菜','基础学习篇','鸟哥','第三版',88.00,'人民邮电出版社','2010年8月',1497864193,1497864193);
ALTER TABLE `book` ENABLE KEYS;
UNLOCK TABLES;




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


