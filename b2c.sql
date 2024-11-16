-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: B2CShop
-- ------------------------------------------------------
-- Server version	8.0.36

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
-- Table structure for table `address`
--

DROP TABLE IF EXISTS `address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `address` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` int DEFAULT '0' COMMENT '用户编号',
  `phone` varchar(30) DEFAULT '' COMMENT '用户手机',
  `name` varchar(30) DEFAULT '' COMMENT '用户名字',
  `zipcode` varchar(20) DEFAULT '' COMMENT '邮政编码',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `default_address` int DEFAULT '0' COMMENT '默认地址',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COMMENT='地址信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address`
--

LOCK TABLES `address` WRITE;
/*!40000 ALTER TABLE `address` DISABLE KEYS */;
INSERT INTO `address` VALUES (3,4,'15674683467','曾佳晨','410000','湖南省长沙市中南大学',0,0),(4,4,'01023465873','华为','123456','上海青浦',1,0),(5,6,'18507446563','tss','','湖南省益阳市沅江市沅江三中教师公寓',1,0);
/*!40000 ALTER TABLE `address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `administrator`
--

DROP TABLE IF EXISTS `administrator`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `administrator` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT '' COMMENT '用户名',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `role_id` int DEFAULT '0' COMMENT '角色编号',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `is_super` tinyint DEFAULT '0' COMMENT '是否超级管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `administrator`
--

LOCK TABLES `administrator` WRITE;
/*!40000 ALTER TABLE `administrator` DISABLE KEYS */;
INSERT INTO `administrator` VALUES (1,'admin','e10adc3949ba59abbe56e057f20f883e','15674683467','15674683467@163.com',1,1,0,1),(7,'admin_jsb','e10adc3949ba59abbe56e057f20f883e','15674683467','15674683467@163.com',1,2,1717214046,0),(8,'admin_yyb','e10adc3949ba59abbe56e057f20f883e','15674683467','1208041200@qq.com',1,3,1717294967,0);
/*!40000 ALTER TABLE `administrator` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth`
--

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth` (
  `id` int NOT NULL AUTO_INCREMENT,
  `module_name` varchar(80) NOT NULL DEFAULT '',
  `action_name` varchar(80) DEFAULT '' COMMENT '操作名称',
  `type` tinyint(1) DEFAULT '0' COMMENT '节点类型',
  `url` varchar(250) DEFAULT '' COMMENT '跳转地址',
  `module_id` int DEFAULT '0' COMMENT '模块编号',
  `sort` int DEFAULT '0' COMMENT '排序',
  `description` varchar(250) DEFAULT '' COMMENT '描述',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `checked` tinyint(1) DEFAULT '0' COMMENT '是否检验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb3 COMMENT='权限控制';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth`
--

LOCK TABLES `auth` WRITE;
/*!40000 ALTER TABLE `auth` DISABLE KEYS */;
INSERT INTO `auth` VALUES (1,'系统管理员','最高权限',3,'/admin',0,0,'',0,0,0),(2,'组织部门','',3,'',0,0,'',0,0,0),(3,'权限管理','',3,'',0,0,'',0,0,0),(4,'Banner管理','',3,'',0,0,'',0,0,0),(5,'商品管理','',3,'',0,0,'',0,0,0),(6,'订单管理','',3,'',0,0,'',0,0,0),(7,'设置管理','',3,'',0,0,'',0,0,0),(8,'管理员列表','管理员列表',2,'/administrator',1,0,'',1,0,1),(9,'','新增管理员',2,'/administrator/add',1,0,'',1,0,0),(10,'','部门列表',2,'/role',2,0,'',1,0,0),(11,'','新增部门',2,'/role/add',2,0,'',0,0,0),(12,'','新增权限',2,'/auth/add',3,0,'',0,0,0),(13,'','权限列表',2,'/auth',3,0,'',0,0,0),(14,'','Banner列表',2,'/banner',4,0,'',0,0,0),(15,'','新增Banner',2,'/banner/add',4,0,'',0,0,0),(16,'','商品列表',2,'/product',5,0,'',0,0,0),(17,'','商品分类',2,'/productCate',5,0,'',0,0,0),(18,'','商品类型',2,'/productType',5,0,'',0,0,0),(19,'','订单列表',2,'/order',6,0,'',0,0,0),(20,'','导航管理',2,'/menu',7,0,'',0,0,0),(21,'','商城设置',2,'/setting',7,0,'',0,0,0),(23,'','删除管理员',3,'/administrator/delete',1,0,'',1,0,0),(24,'','修改管理员',3,'/administrator/edit',1,1,'',1,0,0),(25,'','删除部门',3,'/role/delete',2,10,'',1,0,0),(26,'','修改部门',3,'/role/edit',2,10,'',1,0,0),(27,'','修改权限列表',3,'/auth/edit',3,10,'',1,0,0),(28,'','删除权限',3,'/auth/delete',3,1,'',1,0,0),(29,'','修改Banner',3,'/banner/edit',4,10,'',1,0,0),(30,'','删除Banner',1,'/banner/delete',4,10,'',1,0,0),(31,'','新增商品分类',3,'/productCate/add',5,10,'',1,0,0),(32,'','修改商品分类',3,'/productCate/edit',5,10,'',1,0,0),(33,'','删除商品分类',3,'/productCate/delete',5,10,'',1,0,0),(34,'','新增商品类型',3,'/productType/add',5,10,'',1,0,0),(35,'','修改商品类型',3,'/productType/edit',5,100,'',1,0,0),(36,'','删除产品类型',3,'/productType/delete',5,100,'',1,0,0),(37,'','新增商品',3,'/product/add',5,10,'',1,0,0),(38,'','修改商品',3,'/product/edit',5,10,'',1,0,0),(39,'','删除商品',3,'/product/delete',5,10,'',1,0,0),(40,'','新增商品属性',3,'/productTypeAttribute/add',5,10,'',1,0,0),(41,'','修改商品属性',3,'/productTypeAttribute/edit',5,10,'',1,0,0),(42,'','删除商品属性',3,'/productTypeAttribute/delete',5,10,'',1,0,0),(43,'','修改订单',3,'/order/edit',6,10,'',1,0,0),(44,'','删除订单',1,'/order/delete',6,10,'',1,0,0),(45,'','查看商品属性',3,'/productTypeAttribute',5,10,'',1,0,0),(48,'','查看订单详情',3,'/order/detail',6,10,'',1,0,0);
/*!40000 ALTER TABLE `auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `banner`
--

DROP TABLE IF EXISTS `banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `banner` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(50) DEFAULT '' COMMENT '标题',
  `banner_type` tinyint DEFAULT '0' COMMENT '类型',
  `banner_img` varchar(100) DEFAULT '' COMMENT '图片地址',
  `link` varchar(200) DEFAULT '' COMMENT '连接',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` int DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3 COMMENT='焦点图表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `banner`
--

LOCK TABLES `banner` WRITE;
/*!40000 ALTER TABLE `banner` DISABLE KEYS */;
INSERT INTO `banner` VALUES (2,'banner2',1,'static/upload/20240602/1717316554319312220.jpg','/category_8.html',1000,1,1603504839),(9,'外星人',1,'static/upload/20240602/1717316479584604688.jpg','/category_10.html',100000,1,1717316415),(10,'华为',1,'static/upload/20240602/1717316673741427694.jpg','/category_3.html',1000,1,1717316673);
/*!40000 ALTER TABLE `banner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cart`
--

DROP TABLE IF EXISTS `cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cart` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(250) DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) DEFAULT '0.00',
  `goods_version` varchar(50) DEFAULT '' COMMENT '版本',
  `num` int DEFAULT '0' COMMENT '数量',
  `product_gift` varchar(100) DEFAULT '' COMMENT '商品礼物',
  `product_fitting` varchar(100) DEFAULT '' COMMENT '商品搭配',
  `product_color` varchar(50) DEFAULT '' COMMENT '商品颜色',
  `product_img` varchar(150) DEFAULT '' COMMENT '商品图片',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='购物车';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart`
--

LOCK TABLES `cart` WRITE;
/*!40000 ALTER TABLE `cart` DISABLE KEYS */;
/*!40000 ALTER TABLE `cart` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `link` varchar(250) DEFAULT '' COMMENT '连接',
  `position` int DEFAULT '0' COMMENT '位置',
  `is_opennew` int DEFAULT '0' COMMENT '是否新打开',
  `relation` varchar(100) DEFAULT '' COMMENT '关系',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` int DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (9,'手机','',2,1,'5,6,10',10,1,1717315793),(10,'电脑','',2,1,'7,8,9',10,1,1717315813);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(100) DEFAULT '' COMMENT '订单编号',
  `uid` int DEFAULT '0' COMMENT '用户编号',
  `all_price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `phone` varchar(30) DEFAULT '' COMMENT '电话',
  `name` varchar(100) DEFAULT '' COMMENT '名字',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `zipcode` varchar(30) DEFAULT '' COMMENT '邮编',
  `pay_status` tinyint DEFAULT '0' COMMENT '支付状态',
  `pay_type` tinyint DEFAULT '0' COMMENT '支付类型',
  `order_status` tinyint DEFAULT '0' COMMENT '订单状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` VALUES (8,'2024060216295674',4,19999.00,'15674683467','曾佳晨','湖南省长沙市中南大学','410000',1,0,1,1717316951),(9,'2024060217126428',4,4199.00,'15674683469','张三','成都市xxxx区xxxx街道xxxx号','123456',0,0,0,1717319571),(11,'2024071123307821',4,19999.00,'15674683467','曾佳晨','湖南省长沙市中南大学','410000',0,0,0,1720711852),(12,'2024111416297506',4,3699.00,'01023465873','华为','上海青浦','123456',1,0,1,1731572956),(15,'2024111620593653',6,5398.00,'18507446563','tss','湖南省益阳市沅江市沅江三中教师公寓','',1,1,1,1731761988);
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_item`
--

DROP TABLE IF EXISTS `order_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_item` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '订单编号',
  `order_id` int DEFAULT '0' COMMENT '订单编号',
  `uid` int DEFAULT '0' COMMENT '用户编号',
  `product_title` varchar(100) DEFAULT '' COMMENT '商品标题',
  `product_id` int DEFAULT '0' COMMENT '商品编号',
  `product_img` varchar(200) DEFAULT '' COMMENT '商品图片',
  `product_price` decimal(10,2) DEFAULT '0.00' COMMENT '商品价格',
  `product_num` int DEFAULT '0' COMMENT '商品数量',
  `goods_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `goods_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_item`
--

LOCK TABLES `order_item` WRITE;
/*!40000 ALTER TABLE `order_item` DISABLE KEYS */;
INSERT INTO `order_item` VALUES (1,1,4,'go web from ',1,'static/upload/20201023/1603440321653795000.png',66.66,1,'2','red',1603442151);
/*!40000 ALTER TABLE `order_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `product_sn` varchar(50) DEFAULT '',
  `cate_id` int DEFAULT '0' COMMENT '分类id',
  `click_count` int DEFAULT '0' COMMENT '点击数',
  `product_number` int DEFAULT '0' COMMENT '商品编号',
  `price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `market_price` decimal(10,2) DEFAULT '0.00' COMMENT '市场价格',
  `relation_product` varchar(100) DEFAULT '' COMMENT '关联商品',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  `product_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `product_img` varchar(100) DEFAULT '' COMMENT '商品图片',
  `product_gift` varchar(100) DEFAULT '',
  `product_fitting` varchar(100) DEFAULT '',
  `product_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `product_keywords` varchar(100) DEFAULT '' COMMENT '关键词',
  `product_desc` varchar(50) DEFAULT '' COMMENT '描述',
  `product_content` varchar(100) DEFAULT '' COMMENT '内容',
  `is_delete` tinyint DEFAULT '0' COMMENT '是否删除',
  `is_hot` tinyint DEFAULT '0' COMMENT '是否热门',
  `is_best` tinyint DEFAULT '0' COMMENT '是否畅销',
  `is_new` tinyint DEFAULT '0' COMMENT '是否新品',
  `product_type_id` tinyint DEFAULT '0' COMMENT '商品类型编号',
  `sort` int DEFAULT '0' COMMENT '商品分类',
  `status` tinyint DEFAULT '0' COMMENT '商品状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='商品';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (5,'华为nova 12','华为nova 12 Pro 前置6000万人像追焦双摄 512GB 12号色物理可变光圈 鸿蒙智慧通信智能手机nova系列','',3,100,10,4199.00,5999.00,'','','nova 12','static/upload/20240602/1717313897313843503.jpg','','','1,2,3','','','华为nova 12 Pro 前置6000万人像追焦双摄 512GB 12号色物理可变光圈 鸿蒙智慧通信智能手机nova系列',0,1,1,1,5,10,1,1717313897),(6,'华为畅享 70 Pro','华为畅享 70 Pro 1亿像素超清影像40W超级快充5000mAh大电池长续航 256GB 翡冷翠 鸿蒙智能手机','',3,100,10,1999.00,3999.00,'','','','static/upload/20240602/1717314304194322269.jpg','','','1,3','','','华为畅享 70 Pro 1亿像素超清影像40W超级快充5000mAh大电池长续航 256GB 翡冷翠 鸿蒙智能手机',0,1,1,1,5,1,1,1717314304),(7,'外星人（Alienware）','外星人（Alienware）【2024】m16 R2 16英寸游戏本酷睿Ultra 7 16G 512G RTX4060 240Hz AI高性能笔记本电脑4760QB','',10,100,10,19999.00,39999.00,'','','','static/upload/20240602/1717314460421696719.jpg','','','1','','','外星人（Alienware）【2024】m16 R2 16英寸游戏本酷睿Ultra 7 16G 512G RTX4060 240Hz AI高性能笔记本电脑4760QB',0,1,1,1,6,2,1,1717314460),(8,'华为MateBook D ','华为MateBook D 16 SE 2024笔记本电脑 13代酷睿标压处理器/16英寸护眼大屏 i5 16G 512G 皓月银','',4,100,12,3699.00,3999.00,'','','','static/upload/20240602/1717314761823583349.jpg','','','2','','','华为MateBook D 16 SE 2024笔记本电脑 13代酷睿标压处理器/16英寸护眼大屏 i5 16G 512G 皓月银',0,1,1,1,6,1,1,1717314761),(9,' 戴尔（DELL）',' 戴尔（DELL）灵龙笔记本电脑AI轻薄本灵越14-5445锐龙版高性能商务办公学生 R7-8840HS 16G 512G 2.2K','',9,100,8,3799.00,3999.00,'','','','static/upload/20240602/1717314964159897559.jpg','','','1,2','','',' 戴尔（DELL）灵龙笔记本电脑AI轻薄本灵越14-5445锐龙版高性能商务办公学生 R7-8840HS 16G 512G 2.2K',0,1,1,1,6,2,1,1717314964),(10,'小米Redmi K70E','小米Redmi K70E 天玑8300-Ultra小米澎湃OS 12GB+256GB墨羽 AI功能 红米5G手机','',8,100,3,1599.00,1999.00,'','','','static/upload/20240602/1717315151663194719.jpg','','','1,2,3','','','小米Redmi K70E 天玑8300-Ultra小米澎湃OS 12GB+256GB墨羽 AI功能 红米5G手机',0,1,1,0,5,1,1,1717315151);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_attr`
--

DROP TABLE IF EXISTS `product_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_attr` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT '0' COMMENT '商品编号',
  `attribute_cate_id` int DEFAULT '0' COMMENT '属性分类编号',
  `attribute_id` int DEFAULT '0' COMMENT '属性编号',
  `attribute_title` varchar(100) DEFAULT '' COMMENT '属性标题',
  `attribute_type` int DEFAULT '0' COMMENT '属性类型',
  `attribute_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `sort` int DEFAULT '0' COMMENT '排序',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb3 COMMENT='商品属性';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_attr`
--

LOCK TABLES `product_attr` WRITE;
/*!40000 ALTER TABLE `product_attr` DISABLE KEYS */;
INSERT INTO `product_attr` VALUES (14,2,1,1,'平板电脑',1,'',10,1717311135,1),(55,6,5,12,'续航',3,'大于12小时',10,1717317417,1),(56,6,5,13,'充电时间',1,'12h',10,1717317417,1),(57,7,6,14,'重量',1,'2kg',10,1717317445,1),(58,7,6,15,'CPU型号',1,'i7',10,1717317445,1),(59,7,6,16,'内存',1,'16G',10,1717317445,1),(60,7,6,17,'磁盘',1,'1T',10,1717317445,1),(65,8,6,14,'重量',1,'1.3kg',10,1717317871,1),(66,8,6,15,'CPU型号',1,'i5',10,1717317871,1),(67,8,6,16,'内存',1,'8',10,1717317871,1),(68,8,6,17,'磁盘',1,'SSD 512',10,1717317871,1),(69,9,6,14,'重量',1,'2.2kg',10,1717317966,1),(70,9,6,15,'CPU型号',1,'i9',10,1717317966,1),(71,9,6,16,'内存',1,'32G',10,1717317966,1),(72,9,6,17,'磁盘',1,'1T',10,1717317966,1),(73,10,5,12,'续航',3,'大于6小时\r\n',10,1717317990,1),(74,10,5,13,'充电时间',1,'12小时',10,1717317990,1),(75,5,5,12,'续航',3,'大于12小时',10,1717319537,1),(76,5,5,13,'充电时间',1,'6',10,1717319537,1);
/*!40000 ALTER TABLE `product_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_cate`
--

DROP TABLE IF EXISTS `product_cate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_cate` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(200) DEFAULT '' COMMENT '分类名称',
  `cate_img` varchar(200) DEFAULT '' COMMENT '分类图片',
  `link` varchar(250) DEFAULT '' COMMENT '链接',
  `template` text COMMENT '模版',
  `pid` int DEFAULT '0' COMMENT '父编号',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `keywords` varchar(250) DEFAULT '' COMMENT '关键字',
  `description` text COMMENT '描述',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3 COMMENT='商品分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_cate`
--

LOCK TABLES `product_cate` WRITE;
/*!40000 ALTER TABLE `product_cate` DISABLE KEYS */;
INSERT INTO `product_cate` VALUES (1,'手机','static/upload/20240602/1717312291262204147.jpg','','',0,'手机','手机','手机',0,1,0),(2,'电脑','','','',0,'PC电脑','','',0,1,0),(3,'华为手机','static/upload/20240530/1717043158963485900.jpg','','',1,'nove 7','手机,华为','华为手机',1,1,0),(4,'华为电脑','static/upload/20240530/1717043240054305397.jpg','','',2,'','','',1,1,1717043240),(8,'小米手机','static/upload/20240602/1717312400448541494.jpg','','',1,'','','',10,1,1717312400),(9,'戴尔电脑','static/upload/20240602/1717312501168415369.jpg','','',2,'戴尔电脑','戴尔电脑','戴尔电脑',10,1,1717312501),(10,'外星人','static/upload/20240602/1717312698591940269.jpg','','',2,'外星人','外星人','外星人',10,1,1717312698);
/*!40000 ALTER TABLE `product_cate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_collect`
--

DROP TABLE IF EXISTS `product_collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_collect` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `product_id` int NOT NULL,
  `add_time` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户收藏商品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_collect`
--

LOCK TABLES `product_collect` WRITE;
/*!40000 ALTER TABLE `product_collect` DISABLE KEYS */;
INSERT INTO `product_collect` VALUES (12,4,7,'20241112');
/*!40000 ALTER TABLE `product_collect` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_color`
--

DROP TABLE IF EXISTS `product_color`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_color` (
  `id` int NOT NULL AUTO_INCREMENT,
  `color_name` varchar(100) DEFAULT '' COMMENT '颜色名字',
  `color_value` varchar(100) DEFAULT '' COMMENT '颜色值',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `checked` tinyint DEFAULT '0' COMMENT '是否检验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_color`
--

LOCK TABLES `product_color` WRITE;
/*!40000 ALTER TABLE `product_color` DISABLE KEYS */;
INSERT INTO `product_color` VALUES (1,'黑色','#ffffff',1,0),(2,'白色','#ffffff',1,0),(3,'蓝色','#ffff',1,0);
/*!40000 ALTER TABLE `product_color` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_image`
--

DROP TABLE IF EXISTS `product_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_image` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT '0' COMMENT '商品编号',
  `img_url` varchar(250) DEFAULT '' COMMENT '图片地址',
  `color_id` int DEFAULT '0' COMMENT '颜色编号',
  `sort` int DEFAULT '0' COMMENT '排序',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_image`
--

LOCK TABLES `product_image` WRITE;
/*!40000 ALTER TABLE `product_image` DISABLE KEYS */;
INSERT INTO `product_image` VALUES (1,1,'/static/upload/20201024/1603519200684359000.jpg',1,10,1603519201,1),(2,1,'/static/upload/20201024/1603519285204437000.jpg',1,10,1603519291,1),(6,2,'/static/upload/20210118/1610940522542324000.jpg',2,10,1610940523,1),(7,2,'/static/upload/20210118/1610940522573123000.jpg',2,10,1610940523,1),(8,3,'/static/upload/20210118/1610940548355473000.jpg',1,10,1610940548,1),(10,5,'/static/upload/20240602/1717313949736562187.jpg',3,10,1717313950,1),(11,6,'/static/upload/20240602/1717314351980865986.jpg',3,10,1717314353,1),(12,7,'/static/upload/20240602/1717314642220693928.jpg',1,10,1717314644,1),(13,8,'/static/upload/20240602/1717314843811701284.jpg',2,10,1717314845,1),(14,9,'/static/upload/20240602/1717315069190529790.jpg',0,10,1717315069,1),(15,10,'/static/upload/20240602/1717315567764709143.jpg',1,10,1717315572,1);
/*!40000 ALTER TABLE `product_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_type`
--

DROP TABLE IF EXISTS `product_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_type`
--

LOCK TABLES `product_type` WRITE;
/*!40000 ALTER TABLE `product_type` DISABLE KEYS */;
INSERT INTO `product_type` VALUES (5,'电子产品（手机）','',1,1717313683),(6,'电子产品（电脑）','',1,1717313696);
/*!40000 ALTER TABLE `product_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_type_attribute`
--

DROP TABLE IF EXISTS `product_type_attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_type_attribute` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cate_id` int DEFAULT '0' COMMENT '分类编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `attr_type` tinyint DEFAULT '0' COMMENT '属性类型',
  `attr_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `sort` int DEFAULT '0' COMMENT '排序',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_type_attribute`
--

LOCK TABLES `product_type_attribute` WRITE;
/*!40000 ALTER TABLE `product_type_attribute` DISABLE KEYS */;
INSERT INTO `product_type_attribute` VALUES (1,1,'平板电脑',1,'',1,10,1603440086),(2,1,'台式电脑',1,'',1,1,1717046424),(3,2,'苹果手机',1,'',1,10,1717046444),(4,2,'华为手机',1,'',1,5,1717046457),(6,4,'中国邮电出版社',1,'',1,10,1717310572),(12,5,'续航',3,'大于6小时\r\n小于6小时\r\n大于12小时',1,10,1717317270),(13,5,'充电时间',1,'',1,10,1717317292),(14,6,'重量',1,'',1,10,1717317328),(15,6,'CPU型号',1,'',1,10,1717317338),(16,6,'内存',1,'',1,10,1717317346),(17,6,'磁盘',1,'',1,10,1717317353);
/*!40000 ALTER TABLE `product_type_attribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题名称',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'超级管理员','超级管理员',1,0),(2,'技术部','技术部',1,1603518015),(3,'运营部','运营部',1,1603518054);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_auth`
--

DROP TABLE IF EXISTS `role_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_auth` (
  `auth_id` int NOT NULL COMMENT '权限编号',
  `role_id` int DEFAULT '0' COMMENT '角色编号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_auth`
--

LOCK TABLES `role_auth` WRITE;
/*!40000 ALTER TABLE `role_auth` DISABLE KEYS */;
INSERT INTO `role_auth` VALUES (1,1),(1,2),(8,2),(2,2),(10,2),(4,2),(14,2),(15,2),(29,2),(30,2),(5,2),(16,2),(17,2),(18,2),(31,2),(32,2),(33,2),(34,2),(35,2),(36,2),(37,2),(38,2),(39,2),(40,2),(41,2),(42,2),(45,2),(6,2),(19,2),(48,2),(7,2),(20,2),(1,3),(8,3),(2,3),(10,3),(4,3),(14,3),(15,3),(29,3),(30,3),(5,3),(16,3),(17,3),(18,3),(31,3),(32,3),(33,3),(34,3),(35,3),(36,3),(37,3),(38,3),(39,3),(40,3),(41,3),(42,3),(45,3),(6,3),(19,3),(43,3),(44,3),(48,3),(7,3),(20,3),(21,3);
/*!40000 ALTER TABLE `role_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting`
--

DROP TABLE IF EXISTS `setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `site_title` varchar(100) DEFAULT '' COMMENT '商城名称',
  `site_logo` varchar(250) DEFAULT '' COMMENT '商城图标',
  `site_keywords` varchar(100) DEFAULT '' COMMENT '商城关键字',
  `site_description` varchar(500) DEFAULT '' COMMENT '商城描述',
  `no_picture` varchar(100) DEFAULT '' COMMENT '没有图片显示',
  `site_icp` varchar(50) DEFAULT '' COMMENT '商城ICP',
  `site_tel` varchar(50) DEFAULT '' COMMENT '商城手机号',
  `search_keywords` varchar(250) DEFAULT '' COMMENT '搜索关键字',
  `tongji_code` varchar(500) DEFAULT '' COMMENT '统计编码',
  `appid` varchar(50) DEFAULT '' COMMENT 'oss appid',
  `app_secret` varchar(80) DEFAULT '' COMMENT 'oss app_secret',
  `end_point` varchar(200) DEFAULT '' COMMENT 'oss 终端点',
  `bucket_name` varchar(200) DEFAULT '' COMMENT 'oss 桶名称',
  `oss_status` tinyint DEFAULT '0' COMMENT 'oss 状态',
  `ip` char(16) DEFAULT NULL COMMENT '修改的ip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting`
--

LOCK TABLES `setting` WRITE;
/*!40000 ALTER TABLE `setting` DISABLE KEYS */;
INSERT INTO `setting` VALUES (28,'FCB Happy Shop','static/upload/20240530/1717057279403675250.jpg','','','static/upload/20240530/1717057279404605020.jpg','','','','','','','','',0,'120.227.57.37'),(29,'FCB Happy Shop','static/upload/20240530/1717057279403675250.jpg','FCB Happy Shop','FCB 网上商城','static/upload/20240530/1717057279404605020.jpg','FCB Happy Shop','15674683467','FCB Happy Shop','1208041200','','','','',0,'120.227.57.37');
/*!40000 ALTER TABLE `setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(30) DEFAULT '' COMMENT '手机号',
  `password` varchar(80) DEFAULT '' COMMENT '密码',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `last_ip` varchar(50) DEFAULT '' COMMENT '最近ip',
  `email` varchar(80) DEFAULT '' COMMENT '邮编',
  `status` tinyint DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (4,'15674683467','e10adc3949ba59abbe56e057f20f883e',0,'120.227.56.82','',0),(5,'15674683468','e10adc3949ba59abbe56e057f20f883e',0,'120.227.56.82','',0),(6,'18507446563','4a6881b0a5615a4e4eafb888962cbbd6',0,'113.242.104.30','',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info`
--

DROP TABLE IF EXISTS `user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `gender` int DEFAULT NULL COMMENT '性别',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `phone` varchar(30) DEFAULT NULL COMMENT '电话',
  `username` varchar(255) DEFAULT NULL COMMENT '昵称',
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info`
--

LOCK TABLES `user_info` WRITE;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_sms`
--

DROP TABLE IF EXISTS `user_sms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_sms` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) DEFAULT '' COMMENT 'ip地址',
  `phone` varchar(50) DEFAULT '' COMMENT '手机号',
  `send_count` int DEFAULT '0' COMMENT '发送统计',
  `add_day` varchar(200) DEFAULT '' COMMENT '添加日期',
  `add_time` int DEFAULT '0' COMMENT '添加时间',
  `sign` varchar(80) DEFAULT '' COMMENT '签名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_sms`
--

LOCK TABLES `user_sms` WRITE;
/*!40000 ALTER TABLE `user_sms` DISABLE KEYS */;
INSERT INTO `user_sms` VALUES (9,'120.227.56.82','15674683467',1,'20240526',1716690675,'fac9b2547e9a71c3470a6d0fa851842b'),(10,'120.227.56.82','15674683468',1,'20240526',1716709638,'e9ad5af3647b68532ab9b87be077a6ee'),(11,'113.242.104.30','18507446563',1,'20241116',1731761783,'43b61a9d814e1684ff08e60934622367');
/*!40000 ALTER TABLE `user_sms` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-16 21:14:33
