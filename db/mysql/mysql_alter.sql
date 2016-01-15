alter table apps ADD column_test int default 0;

--------------

mysql> alter table apps ADD CONSTRAINT UNIQUE(id,url);

mysql> alter table apps ADD CONSTRAINT UNIQUE(id,url);
CREATE TABLE `apps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_name` char(20) NOT NULL DEFAULT '' COMMENT '站点名称',
  `url` varchar(255) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0',
  `country` char(10) NOT NULL DEFAULT '' COMMENT '国家',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`,`url`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 

--------------------


如下命令使用了 ALTER 命令及 DROP 子句来删除以上创建表的 i 字段：

mysql> ALTER TABLE testalter_tbl  DROP i;

mysql> alter table apps drop column_test;
Query OK, 0 rows affected (0.08 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> select * from apps;
+----+------------+-------------------------+---------+
| id | app_name   | url                     | country |
+----+------------+-------------------------+---------+
|  1 | QQ APP     | http://im.qq.com/       | CN      |
|  2 | 微博 APP   | http://weibo.com/       | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ | CN      |
+----+------------+-------------------------+---------+

-------------------------



mysql> ALTER TABLE apps ADD count INT NOT NULL DEFAULT 0 AFTER url;
Query OK, 0 rows affected (0.08 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> select * from apps;
+----+------------+-------------------------+-------+---------+
| id | app_name   | url                     | count | country |
+----+------------+-------------------------+-------+---------+
|  1 | QQ APP     | http://im.qq.com/       |     0 | CN      |
|  2 | 微博 APP   | http://weibo.com/       |     0 | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ |     0 | CN      |
+----+------------+-------------------------+-------+---------+
3 rows in set (0.00 sec)
