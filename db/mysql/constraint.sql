-- 在 SQL 中，我们有如下约束：
    NOT NULL - 指示某列不能存储 NULL 值。
    UNIQUE - 保证某列的每行必须有唯一的值。
    PRIMARY KEY - NOT NULL 和 UNIQUE 的结合。确保某列（或两个列多个列的结合）有唯一标识，有助于更容易更快速地找到表中的一个特定的记录。
    FOREIGN KEY - 保证一个表中的数据匹配另一个表中的值的参照完整性。
    CHECK - 保证列中的值符合指定的条件。
    DEFAULT - 规定没有给列赋值时的默认值。


-- 约束的目的就是确保表中的数据的完整性。
-- 常用的约束类型如下:
-- 主键约束:(Primary Key constraint)      要求主键列唯一，并且不允许为空
-- 唯一约束:(Unique Constraint)              要求该列唯一，允许为空，但只能出现一个空值
-- 检查约束:(Check Constraint)                某列取值范围限制、格式限制等。如有关年龄的限制
-- 默认约束:(Default Constraint)               某列的默认值，如我们的男性学员比较多，性别默认为男
-- 外键约束:(Foreign Key Constraint)       用于在两表之间建立关系，需要指定引用主表的哪一列

UNIQUE 约束唯一标识数据库表中的每条记录。

UNIQUE 和 PRIMARY KEY 约束均为列或列集合提供了唯一性的保证。

PRIMARY KEY 约束拥有自动定义的 UNIQUE 约束。

请注意，每个表可以有多个 UNIQUE 约束，但是每个表只能有一个 PRIMARY KEY 约束。


------------


alter table apps ADD column_test int default 0;

--------------

mysql> alter table apps ADD CONSTRAINT UNIQUE(id,url);

mysql> alter table apps ADD CONSTRAINT UNIQUE(id,url);
CREATE TABLE `apps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_name` char(20) NOT NULL DEFAULT '' COMMENT '站点名称',
  `url` varchar(255) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0',
  `country` char(10) NOT NULL DEFAULT '' COMMENT '国家',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`,`url`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 


ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)
ALTER TABLE xxxx ADD UNIQUE (xxxx)