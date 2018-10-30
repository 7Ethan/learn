-- CREATE TABLE student(
--     id serial PRIMARY KEY,
--     name VARCHAR(20) NOT NULL,
--     email VARCHAR(50) NOT NULL,
-- )

CREATE TABLE equipment(
    id serial PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    color VARCHAR(25) NOT NULL,
    location VARCHAR(25) CHECK (location in ('north','west','east','south')),
    install_date DATE
);

-- 添加列
ALTER TABLE equipment ADD COLUMN functioning bool;
-- 列默认值
ALTER TABLE equipment ALTER COLUMN functioning SET DEFAULT 'true';
ALTER TABLE equipment ALTER COLUMN functioning SET NOT NULL;


-- 查看表
-- \d 表名

-- 重命名列名
ALTER TABLE equipment RENAME COLUMN functioning TO inorder;

-- 删除列
ALTER TABLE equipment DELETE COLUMN inorder;

-- 插入语句/
INSERT INTO equipment(id,type,color,location,inorder) VALUES(001,'book','blue','north',false);

