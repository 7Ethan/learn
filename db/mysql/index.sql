索引查询：

      select * 

      from table

     indexed by index_name

     where  xxx;


------------------------=======================================================


sql语句优化 

性能不理想的系统中除了一部分是因为应用程序的负载确实超过了服务器的实际处理能力外,更多的是因为系统存在大量的SQL语句需要优化。
为了获得稳定的执行性能，SQL语句越简单越好。对复杂的SQL语句，要设法对之进行简化。
常见的简化规则如下：
 
1）不要有超过5个以上的表连接（JOIN）
2）考虑使用临时表或表变量存放中间结果。
3）少用子查询
4）视图嵌套不要过深,一般视图嵌套不要超过2个为宜。


CREATE INDEX 语句用于在表中创建索引。

在不读取整个表的情况下，索引使数据库应用程序可以更快地查找数据。


用户无法看到索引，它们只能被用来加速搜索/查询。

注释：更新一个包含索引的表需要比更新一个没有索引的表花费更多的时间，这是由于索引本身也需要更新。因此，理想的做法是仅仅在常常被搜索的列（以及表）上面创建索引。


----------------
在表上创建一个简单的索引。允许使用重复的值：
CREATE INDEX index_name
ON table_name (column_name)


--------------------
表上创建一个唯一的索引。不允许使用重复的值：唯一的索引意味着两个行不能拥有相同的索引值。
CREATE UNIQUE INDEX index_name
ON table_name (column_name)

===============================================================================================

使用CREATE 语句创建索引

CREATE INDEX index_name ON table_name(column_name,column_name) include(score)

普通索引

CREATE UNIQUE INDEX index_name ON table_name (column_name) ;

非空索引

CREATE PRIMARY KEY INDEX index_name ON table_name (column_name) ;

主键索引
 
使用ALTER TABLE语句创建索引

alter table table_name add index index_name (column_list) ;
alter table table_name add unique (column_list) ;
alter table table_name add primary key (column_list) ;


删除索引

drop index index_name on table_name ;
alter table table_name drop index index_name ;
alter table table_name drop primary key ;