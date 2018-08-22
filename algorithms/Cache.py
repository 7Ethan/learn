import time
from collections import OrderedDict

class Cache(OrderedDict):
    def __init__(self,expiration=15*60,maxsize=256):
        self.maxsize = maxsize
        self.expiration = expiration
        self.__expire_times = OrderedDict()
        self.__access_time = OrderedDict()
        self.__values = {} #map

    def setItem(self,key,value):
        t = int(time.time())
        self.delItem(key)
        self.__values[key] = value
        self.__access_time[key] = t
        self.__expire_times[key] = t + self.expiration
        self.cleanup()

    def getItem(self,key):
        t = int(time.time())
        # 进行一下操作,因为访问了需要更新过期时间
        del self.__access_time[key]
        self.__access_time[key] = t
        self.cleanup()
        return self.__values[key]

    def delItem(self,key):
        if key in self.__values:
            del self.__values[key]
            del self.__access_time[key]
            del self.__expire_times[key]

    def size(self):
        return len(self.__values)

    def clear(self):
        self.__values.clear()
        self.__access_time.clear()
        self.__expire_times.clear()

    def cleanup(self):
        t = int(time.time())
        for key,expire in self.__expire_times.items():
            if expire < t:
                self.delItem(key)
        
        while self.size() > self.maxsize:
            for key in self.__access_time:
                self.delItem(key)
                break
        
if __name__ == "__main__":
    # cache = Cache(15 * 60,6)
    cache = Cache(1*60,5)
    for i in range(10):
        cache.setItem(i,str(i + 1))

    print(cache.size())
    print(cache.getItem(8))
