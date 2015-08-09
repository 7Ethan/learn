cc=g++

SRCS= students.cpp main.cpp

OBJS = $(SRCS:.cpp=.0)

EXEc = main.exe 

# 执行链接操作
start:$(OBJS)
	$(cc) $(OBJS) -o $(EXEc)

.cpp.o:
	$(cc) -o $@ -c $<

#声明clean是一个虚目标
.PHONY:clean
	echo "All irrelvant files has been clean."
# 清理所有中间件.
clean:
	rm -f $(OBJS) $(EXEc)