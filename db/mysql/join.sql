-- SQL JOIN 类型：

--     INNER JOIN：如果表中有至少一个匹配，则返回行
--     LEFT JOIN：即使右表中没有匹配，也从左表返回所有的行
--     RIGHT JOIN：即使左表中没有匹配，也从右表返回所有的行
--     FULL JOIN：只要其中一个表中存在匹配，则返回行

-- SQL JOIN

-- SQL JOIN 子句用于把来自两个或多个表的行结合起来，基于这些表之间的共同字段。

-- 最常见的 JOIN 类型：SQL INNER JOIN（简单的 JOIN）。 SQL INNER JOIN 从多个表中返回满足 JOIN 条件的所有行。

-- mysql> select * from websites;
-- +----+---------------+---------------------------+-------+---------+
-- | id | name          | url                       | alexa | country |
-- +----+---------------+---------------------------+-------+---------+
-- |  1 | Google        | https://www.google.cm/    |     1 | USA     |
-- |  2 | 淘宝          | https://www.taobao.com/   |    13 | CN      |
-- |  3 | 个人博客      | http://github.7Ethan.io   |  4689 | CN      |
-- |  4 | 微博          | http://weibo.com/         |    20 | CN      |
-- |  5 | Facebook      | https://www.facebook.com/ |     3 | USA     |
-- |  6 | 百度          | https://www.baidu.com/    |     4 | CN      |
-- |  7 | stackoverflow | http://stackoverflow.com/ |     0 | USA     |
-- +----+---------------+---------------------------+-------+---------+
-- 7 rows in set (0.00 sec)


-- mysql> mysql> select * from access_log;
-- +-----+---------+-------+------------+
-- | aid | site_id | count | date       |
-- +-----+---------+-------+------------+
-- |   1 |       1 |    45 | 2016-05-10 |
-- |   2 |       3 |   100 | 2016-05-13 |
-- |   3 |       1 |   230 | 2016-05-14 |
-- |   4 |       2 |    10 | 2016-05-14 |
-- |   5 |       5 |   205 | 2016-05-14 |
-- |   6 |       4 |    13 | 2016-05-15 |
-- |   7 |       3 |   220 | 2016-05-15 |
-- |   8 |       5 |   545 | 2016-05-16 |
-- |   9 |       3 |   201 | 2016-05-17 |
-- +-----+---------+-------+------------+
-- 9 rows in set (0.00 sec)


-- "Websites" 表中的 "id" 列指向 "access_log" 表中的字段 "site_id"。上面这两个表是通过 "site_id" 列联系起来的。


select websites.id, websites.name, access_log.count,access_log.date
  from websites
  INNER join access_log on websites.id=access_log.site_id;

-- mysql> select websites.id, websites.name, access_log.count,access_log.date
--     ->   from websites
--     ->   INNER join access_log on websites.id=access_log.site_id;
-- +----+--------------+-------+------------+
-- | id | name         | count | date       |
-- +----+--------------+-------+------------+
-- |  1 | Google       |    45 | 2016-05-10 |
-- |  3 | 个人博客     |   100 | 2016-05-13 |
-- |  1 | Google       |   230 | 2016-05-14 |
-- |  2 | 淘宝         |    10 | 2016-05-14 |
-- |  5 | Facebook     |   205 | 2016-05-14 |
-- |  4 | 微博         |    13 | 2016-05-15 |
-- |  3 | 个人博客     |   220 | 2016-05-15 |
-- |  5 | Facebook     |   545 | 2016-05-16 |
-- |  3 | 个人博客     |   201 | 2016-05-17 |
-- +----+--------------+-------+------------+
-- 9 rows in set (0.00 sec)
