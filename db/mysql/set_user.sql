USE mysql;

-- INSERT INTO user 
--           (host, user, password, 
--            select_priv, insert_priv, update_priv) 
--            VALUES ('localhost', 'shangan', 
--            PASSWORD('123456'), 'Y', 'Y', 'Y');
desc user;

set names utf8;

-- new an suer
CREATE USER 'shangan'@'localhost' IDENTIFIED BY "xxxxxx";

-- Grant all access to the account.
GRANT ALL ON *.* TO 'shangan'@'localhost';

show databases;

show tables;

CREATE DATABASE shangan;

drop database <数据库名>;

-- CHAR 	0-255字节 	定长字符串
-- VARCHAR 	0-65535 字节 	变长字符串 

-- DATE 	    3 	    1000-01-01/9999-12-31 	    YYYY-MM-DD 	    日期值
-- TIME 	    3 	    '-838:59:59'/'838:59:59' 	HH:MM:SS 	    时间值或持续时间
-- YEAR 	    1 	    1901/2155 	                YYYY 	        年份值
-- DATETIME 	8 	    1000-01-01 00:00:00/9999-12-31 23:59:59 	YYYY-MM-DD HH:MM:SS 	混合日期和时间值
-- TIMESTAMP 	4 	    1970-01-01 00:00:00/2038        结束时间是第 2147483647 秒，北京时间 2038-1-19 11:14:07，格林尼治时间 2038年1月19日 凌晨 03:14:07   YYYYMMDD HHMMSS 	混合日期和时间值，时间戳 

--new a database