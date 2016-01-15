-- SQL LEFT JOIN 关键字
-- LEFT JOIN 关键字从左表（table1）返回所有的行，即使右表（table2）中没有匹配。如果右表中没有匹配，则结果为 NULL


select websites.id, websites.name, access_log.count,access_log.date
  from websites
  left join access_log on websites.id=access_log.site_id 
  order by websites.id ASC;


mysql> select websites.id, websites.name, access_log.count,access_log.date
    ->   from websites
    ->   left join access_log on websites.id=access_log.site_id 
    ->   order by websites.id ASC;
+----+---------------+-------+------------+
| id | name          | count | date       |
+----+---------------+-------+------------+
|  1 | Google        |    45 | 2016-05-10 |
|  1 | Google        |   230 | 2016-05-14 |
|  2 | 淘宝          |    10 | 2016-05-14 |
|  3 | 个人博客      |   220 | 2016-05-15 |
|  3 | 个人博客      |   100 | 2016-05-13 |
|  3 | 个人博客      |   201 | 2016-05-17 |
|  4 | 微博          |    13 | 2016-05-15 |
|  5 | Facebook      |   545 | 2016-05-16 |
|  5 | Facebook      |   205 | 2016-05-14 |
|  6 | 百度          |  NULL | NULL       |
|  7 | stackoverflow |  NULL | NULL       |
+----+---------------+-------+------------+
11 rows in set (0.00 sec)