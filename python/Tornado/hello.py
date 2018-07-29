import tornado.ioloop 
import tornado.web 
import time

class MainHandler(tornado.web.RequestHandler): 
    def get(self): 
        self.write("Hello, Tornado .") 
        
application = tornado.web.Application([
    (r"/", MainHandler),
]) 

if __name__ == "__main__":
    application.listen(8080)
    tornado.ioloop.IOLoop.current().start()
