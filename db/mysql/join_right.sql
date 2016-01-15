


select websites.id, websites.name, access_log.count,access_log.date
  from websites
  right join access_log on websites.id=access_log.site_id 
  order by websites.id ASC;




mysql> select websites.id, websites.name, access_log.count,access_log.date
    ->   from websites
    ->   right join access_log on websites.id=access_log.site_id 
    ->   order by websites.id ASC;
+------+--------------+-------+------------+
| id   | name         | count | date       |
+------+--------------+-------+------------+
|    1 | Google       |    45 | 2016-05-10 |
|    1 | Google       |   230 | 2016-05-14 |
|    2 | 淘宝         |    10 | 2016-05-14 |
|    3 | 个人博客     |   220 | 2016-05-15 |
|    3 | 个人博客     |   100 | 2016-05-13 |
|    3 | 个人博客     |   201 | 2016-05-17 |
|    4 | 微博         |    13 | 2016-05-15 |
|    5 | Facebook     |   545 | 2016-05-16 |
|    5 | Facebook     |   205 | 2016-05-14 |
+------+--------------+-------+------------+
9 rows in set (0.00 sec)