# # Generic Makefile for Haskell Program
# #
# # License: GPL (General Public License)
# # Author:  Sergey <sergeychang@gmail.com>


# cc=ghc

# # SRCS 变量代表所有需要编译的源文件
# SRCS=*.hs

# # OBJS 是所有编译中间件,一般为.o文件
# OBJS=$(SRCS:.hs=.o)

# # EXEc代表可执行文件,可随意命名.
# EXEc=main.exe


# # ------- simple way ---------
# # 执行链接操作
# start:$(OBJS)
# 	$(cc) $(OBJS) -o $(EXEc)

# .hs.o:
# 	$(cc) -o $@ -c $< -DMYLINUX

# #声明clean是一个虚目标
# .PHONY:clean
# # 清理所有中间件.
# clean:
# 	rm -f $(OBJS) $(EXEc) *.hi