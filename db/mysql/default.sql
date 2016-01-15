mysql> select * from apps;
+----+------------+-------------------------+------+---------+
| id | app_name   | url                     | w_id | country |
+----+------------+-------------------------+------+---------+
|  1 | QQ APP     | http://im.qq.com/       |    2 | CN      |
|  2 | 微博 APP   | http://weibo.com/       |    3 | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ |    4 | CN      |
+----+------------+-------------------------+------+---------+
3 rows in set (0.00 sec)
    

mysql> alter table apps alter country set default 'zh_cn';
Query OK, 0 rows affected (0.00 sec)
Records: 0  Duplicates: 0  Warnings: 0


mysql> insert into apps (app_name,url,w_id) values ('wechat APP','http://wechat.com',8);
Query OK, 1 row affected (0.01 sec)

mysql> select * from apps;
+----+------------+-------------------------+------+---------+
| id | app_name   | url                     | w_id | country |
+----+------------+-------------------------+------+---------+
|  1 | QQ APP     | http://im.qq.com/       |    2 | CN      |
|  2 | 微博 APP   | http://weibo.com/       |    3 | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ |    4 | CN      |
|  4 | wechat APP | http://wechat.com       |    8 | zh_cn   |
+----+------------+-------------------------+------+---------+
4 rows in set (0.00 sec)


---------------------------------------------------------------------------------------

mysql> alter table apps
    -> alter w_id
    -> set default 1;
Query OK, 0 rows affected (0.01 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> insert into apps (app_name,url) values ('wechat APP','http://wechat.com');
Query OK, 1 row affected (0.01 sec)

mysql> select * from apps;
+----+------------+-------------------------+------+---------+
| id | app_name   | url                     | w_id | country |
+----+------------+-------------------------+------+---------+
|  1 | QQ APP     | http://im.qq.com/       |    2 | CN      |
|  2 | 微博 APP   | http://weibo.com/       |    3 | CN      |
|  3 | 淘宝 APP   | https://www.taobao.com/ |    4 | CN      |
|  4 | wechat APP | http://wechat.com       |    8 | zh_cn   |
|  5 | wechat APP | http://wechat.com       |    1 | zh_cn   |
+----+------------+-------------------------+------+---------+
5 rows in set (0.00 sec)
