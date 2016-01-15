
SQL FOREIGN KEY 约束:一个表中的 FOREIGN KEY 指向另一个表中的 UNIQUE KEY(唯一约束的键)。


"Persons" 表：
-- P_Id 	LastName 	FirstName 	Address 	    City
-- 1 	    Hansen 	    Ola 	    Timoteivn 10 	Sandnes
-- 2 	    Svendson 	Tove 	    Borgvn 23 	    Sandnes
-- 3 	    Pettersen 	Kari 	    Storgt 20 	    Stavanger

"Orders" 表：
-- O_Id 	OrderNo 	P_Id
-- 1 	    77895 	    3
-- 2 	    44678 	    3
-- 3 	    22456 	    2
-- 4 	    24562 	    1

请注意，"Orders" 表中的 "P_Id" 列指向 "Persons" 表中的 "P_Id" 列。

"Persons" 表中的 "P_Id" 列是 "Persons" 表中的 PRIMARY KEY。

"Orders" 表中的 "P_Id" 列是 "Orders" 表中的 FOREIGN KEY。

FOREIGN KEY 约束用于预防破坏表之间连接的行为。

FOREIGN KEY 约束也能防止非法数据插入外键列，因为它必须是它指向的那个表中的值之一。

--------------------


在创建表的时候指定外键约束

CREATE TABLE 表名
    (
        column1 datatype null/not null,
        column2 datatype null/not null,
        ...
        CONSTRAINT 外键约束名 FOREIGN KEY  (column1,column2,... column_n) 
        REFERENCES 外键依赖的表 (column1,column2,...column_n)
        ON DELETE CASCADE--级联删除
    );

在创建表后增加外键约束

ALTER TABLE 表名
    ADD CONSTRAINT 外键约束名
    FOREIGN KEY (column1, column2,...column_n) 
    REFERENCES 外键所依赖的表 (column1,column2,...column_n)
    ON DELETE CASCADE;--级联删除

使用工具plsql来新增外键约束

注意，在创建外键约束时，必须先创建外键约束所依赖的表，并且该列为该表的主键.


-----------------

mysql> alter table apps
    -> ADD constraint app_w_id
    -> foreign key (w_id)
    -> references websites(id);
Query OK, 3 rows affected (0.09 sec)
Records: 3  Duplicates: 0  Warnings: 0

-- mysql> show create table apps;

-- CREATE TABLE `apps` (
--   `id` int(11) NOT NULL AUTO_INCREMENT,
--   `app_name` char(20) NOT NULL DEFAULT '' COMMENT '站点名称',
--   `url` varchar(255) NOT NULL DEFAULT '',
--   `w_id` int(11) NOT NULL,
--   `country` char(10) NOT NULL DEFAULT '' COMMENT '国家',
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `id` (`id`,`url`),
--   KEY `app_w_id` (`w_id`),
--   CONSTRAINT `app_w_id` FOREIGN KEY (`w_id`) REFERENCES `websites` (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 


-----------------


mysql> alter table apps ADD FOREIGN KEY (w_id) REFERENCES websites(id);
Query OK, 3 rows affected (0.10 sec)
Records: 3  Duplicates: 0  Warnings: 0

mysql> show create table apps;
apps  | CREATE TABLE `apps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_name` char(20) NOT NULL DEFAULT '' COMMENT '站点名称',
  `url` varchar(255) NOT NULL DEFAULT '',
  `w_id` int(11) NOT NULL,
  `country` char(10) NOT NULL DEFAULT '' COMMENT '国家',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`,`url`),
  KEY `w_id` (`w_id`),
  CONSTRAINT `apps_ibfk_1` FOREIGN KEY (`w_id`) REFERENCES `websites` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 

mysql> select * from apps;
+----+------------+-------------------------+------+---------+
| id | app_name   | url                     | w_id | country |
+----+------------+-------------------------+------+---------+
|  1 | QQ APP     | http://im.qq.com/       |    2 | CN      |
|  2 | 微博 APP   | http://weibo.com/       |    3 | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ |    4 | CN      |
+----+------------+-------------------------+------+---------+
3 rows in set (0.00 sec)

mysql> select * from websites;
+----+---------------+---------------------------+-------+---------+
| id | name          | url                       | alexa | country |
+----+---------------+---------------------------+-------+---------+
|  1 | Google        | https://www.google.cm/    |     1 | USA     |
|  2 | 淘宝          | https://www.taobao.com/   |    13 | CN      |
|  3 | 个人博客      | http://github.7Ethan.io   |  4689 | CN      |
|  4 | 微博          | http://weibo.com/         |    20 | CN      |
|  5 | Facebook      | https://www.facebook.com/ |     3 | USA     |
|  6 | 百度          | https://www.baidu.com/    |     4 | CN      |
|  7 | stackoverflow | http://stackoverflow.com/ |     0 | USA     |
|  8 | QQ APP        |                           |     0 | CN      |
|  9 | 微博 APP      |                           |     0 | CN      |
| 10 | 淘宝 APP      |                           |     0 | CN      |
+----+---------------+---------------------------+-------+---------+
10 rows in set (0.00 sec)