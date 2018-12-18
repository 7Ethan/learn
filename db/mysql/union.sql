
-- UNION 操作符用于合并两个或多个 SELECT 语句的结果集。

-- 请注意，UNION 内部的每个 SELECT 语句必须拥有相同数量的列。列也必须拥有相似的数据类型。同时，每个 SELECT 语句中的列的顺序必须相同。

SELECT column_name(s) FROM table1
UNION
SELECT column_name(s) FROM table2;

注释：默认地，UNION 操作符选取不同的值。如果允许重复的值，请使用 UNION ALL。
 
 --------------------

select name from websites
UNION
select app_name from apps;


mysql> select name from websites
    -> UNION
    -> select app_name from apps;
+---------------+
| name          |
+---------------+
| Google        |
| 淘宝          |
| 个人博客      |
| 微博          |
| Facebook      |
| 百度          |
| stackoverflow |
| QQ APP        |
| 微博 APP      |
| 淘宝 APP      |
+---------------+
10 rows in set (0.00 sec)


------

 UNION ALL 实例

下面的 SQL 语句使用 UNION ALL 从 "Websites" 和 "apps" 表中选取所有的country（也有重复的值）：
