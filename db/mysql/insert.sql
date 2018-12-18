-- INSERT INTO 语句可以有两种编写形式。

-- 第一种形式无需指定要插入数据的列名，只需提供被插入的值即可：
-- INSERT INTO table_name
-- VALUES (value1,value2,value3,...);

-- 第二种形式需要指定列名及被插入的值：
-- INSERT INTO table_name (column1,column2,column3,...)
-- VALUES (value1,value2,value3,...);

insert into websites (name, url, alexa, country)
values ('百度','https://www.baidu.com/','4','CN');  

insert into websites (name,url,country)
values ('stackoverflow', 'http://stackoverflow.com/', 'IND');


-- insert into select 和select into from 的区别
--      insert into scorebak select * from socre where neza='neza'   --插入一行,要求表scorebak 必须存在
--      select *  into scorebak from score  where neza='neza'  --也是插入一行,要求表scorebak 不存在



-- SQL SELECT INTO 语法
我们可以复制所有的列插入到新表中：
SELECT *
INTO newtable [IN externaldb]
FROM table1; 

-- MySQL 数据库不支持 SELECT ... INTO 语句，但支持 INSERT INTO ... SELECT 。
-- 我们可以从一个表中复制所有的列插入到另一个已存在的表中：
INSERT INTO table2(已存在)
SELECT * FROM table1; 

 INSERT INTO Websites (name, country)
SELECT app_name, country FROM apps; 