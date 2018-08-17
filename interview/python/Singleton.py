class Singleton(object):
    def __new__(cls,*args,**kw):
        if not hasattr(cls,"_instance"):
            orig = super(Singleton,cls)
            cls._instance = orig.__new__(cls,*args,**kw)
        return cls._instance

class MyClass(Singleton):
    a = 10

# m = MyClass('Ethan',18)
# print(m.name,m.age,id(m))
# # m1 = MyClass('Sergey',23)
# print(m1.name,m1.age,id(m1))