


SELECT DISTINCT country FROM Websites;
-- 仅从 "Websites" 表的 "country" 列中选取唯一不同的值，也就是去掉 "country" 列重复值

SELECT * from websites where country='cn';  --大小写不区分，cn 与 CN 都可以。
SELECT * from websites where country='cn' and alexa >= 20;  


SELECT * FROM websites WHERE country='CN' AND alexa > 50;


SELECT * FROM websites WHERE country='USA' OR country='CN';


SELECT * FROM websites WHERE alexa > 15 AND (country='CN' OR country='USA');

SELECT * FROM websites WHERE alexa > 15 AND (country='CN' OR country='USA') ORDER BY alexa desc; --desc降序 ASC升序


--  逻辑运算的优先级：
--         ()    not        and         or


--  特殊条件

-- 1.空值判断： is null

--         Select * from websites where name is null;

-- 2.between and (在 之间的值)

-- 3.In


-- ORDER BY 多列的时候，先按照第一个column name排序，在按照第二个column name排序；如上述教程最后一个例子：
--     1）、先将country值这一列排序，同为CN的排前面，同属USA的排后面；
--     2）、然后在同属CN的这些多行数据中，再根据alexa值的大小排列。
--     3）、ORDER BY 排列时，不写明ASC DESC的时候，默认是ASC。 