# Generic Makefile for C/C++ Program
#
# License: GPL (General Public License)
# Author:  Sergey <sergeychang@gmail.com>

# 别名编译器
cc=clang++

# SRCS 变量代表所有需要编译的源文件
SRCS=main.cpp shape.cpp rect.cpp \
				circle.cpp

# OBJS 是所有编译中间件,一般为.o文件
OBJS=$(SRCS:.cpp=.o)

# EXEc代表可执行文件,可随意命名.
EXEc=main.exe


# ------- simple way ---------
# 执行链接操作
start:$(OBJS)
	$(cc) $(OBJS) -o $(EXEc)

# 执行编译所有源文件操作.
#如果在依赖关系中，有多个需要编译的.cpp文件，那么这个语句就需要执行多次。
# -c $<指的是需要编译的.cpp文件,-o $@指这个cpp文件编译后的中间文件名。
# 比如在依赖关系中，有a.cpp和b.cpp，即$(OBJS)的值为a.cpp b.cpp，那么这条语句需要执行2次，
# 第一次的$@为a.o,$<为a.cpp，第二次的$@为b.o,$<为b.cpp。-c表示只编译不链接，-o表示生成的目标文件

# -DMYLINUX:-D为参数，MYLINUX为在cpp源文件中定义的宏名称，
# 如#ifdef MYLINUX。注意-D和宏名称中间没有空格
.cpp.o:
	$(cc) -o $@ -c $< -DMYLINUX

#声明clean是一个虚目标
.PHONY:clean
# 清理所有中间件.
clean:
	rm -f $(OBJS) $(EXEc)