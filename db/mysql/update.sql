-- UPDATE 语法
--     UPDATE table_name
--     SET column1=value1,column2=value2,...
--     WHERE some_column=some_value;

-- 执行没有 WHERE 子句的 UPDATE 要慎重，再慎重。

-- 在 MySQL 中可以通过设置 sql_safe_updates 这个自带的参数来解决，当该参数开启的情况下，你必须在update 语句后携带 where 条件，否则就会报错。

-- set sql_safe_updates=1; 表示开启该参数。

update websites
   set country='USA'
 where id=7;