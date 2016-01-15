-- IN 操作符允许您在 WHERE 子句中规定多个值。

-- SELECT column_name(s)
-- FROM table_name
-- WHERE column_name IN (value1,value2,...);

SELECT * FROM Websites
WHERE name IN ('Google','facebook');

-- mysql> SELECT * FROM websites WHERE name IN ('Google','facebook');
-- +----+----------+---------------------------+-------+---------+
-- | id | name     | url                       | alexa | country |
-- +----+----------+---------------------------+-------+---------+
-- |  1 | Google   | https://www.google.cm/    |     1 | USA     |
-- |  5 | Facebook | https://www.facebook.com/ |     3 | USA     |
-- +----+----------+---------------------------+-------+---------+
-- 2 rows in set (0.01 sec)


