from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
engine = create_engine('mysql://root:shangan@localhost:3306/hotHeap?charset=utf8', encoding="utf-8", echo=True)
BaseDB = declarative_base()

ERROR_CODE = {
    "0": "ok",
    #Users error code
    "1001": "入参非法",
    "1002": "用户已注册，请直接登录",
}