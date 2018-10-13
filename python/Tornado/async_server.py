import time
import tornado.ioloop
import tornado.web
import tornado.httpclient
import tornado.gen

class echo_handler(tornado.web.RequestHandler):
    async def get(self):
        params = self.get_argument("echo")
        self.write("Now time is : [{}],".format(time.strftime("%Y-%m-%d %H:%M:%S",time.localtime())))
        self.write("echo server return:{}".format(params))
    def post(self):
        params = self.get_argument("echo_post")
        self.write("Echo server return from POST request:{}".format(params))
 
class time_consuming_handler(tornado.web.RequestHandler):
	async def get(self):
		t1 =  time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 
		result = await tornado.ioloop.IOLoop.current().run_in_executor(None, self.do_task,5)
		self.write("Task start at:[{}], return {}".format(t1,result))
	def do_task(self,sec):
		time.sleep(sec)
		return "done"
 
class fetch1_handler(tornado.web.RequestHandler):
	@tornado.gen.coroutine
	def get(self):
		t1 =  time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 
		http_client = tornado.httpclient.AsyncHTTPClient()
		response = yield http_client.fetch("http://www.taobao.com")
		self.write("Task start at:[{}]\n".format(t1))		
		self.write(response.body)
 
class fetch2_handler(tornado.web.RequestHandler):
	async def get(self):
		t1 =  time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 
		http_client = tornado.httpclient.AsyncHTTPClient()
		response = await http_client.fetch("http://www.baidu.com")
		self.write("Task start at:[{}]\n".format(t1))		
		self.write(response.body)
		
if __name__ == "__main__":
	webapp=tornado.web.Application([
		(r"/fetch1", fetch1_handler),
		(r"/fetch2", fetch2_handler),		
		(r"/heavy", time_consuming_handler),
        (r"/echo",echo_handler),
	])
	webapp.listen(8080)
	tornado.ioloop.IOLoop.current().start()
