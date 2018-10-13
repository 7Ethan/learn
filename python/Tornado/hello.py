import tornado.ioloop 
import tornado.web 
import time

class MainHandler(tornado.web.RequestHandler): 
    def get(self): 
        self.write("Hello, Tornado .\n") 
        self.write(str(time.localtime()))

class EchoHandler(tornado.web.RequestHandler):
    def get(self):
        args = str(self.get_argument("str"))
        self.write(str(args))

def make_app():
    return tornado.web.Application([
        (r"/", MainHandler),
        (r"/echo",EchoHandler),
        ]) 

if __name__ == "__main__":
    application = make_app()
    application.listen(8080)
    tornado.ioloop.IOLoop.current().start()
