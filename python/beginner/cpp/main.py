import os

cpptest = "./cppexec.exe"   #当前目录下的可执行文件
# if os.path.exists(cpptest):
#     f = os.popen(cpptest)
#     data = f.readlines()
#     f.close()
#     print(data)

print("python execute cpp program:")
os.system(cpptest)  #python系统调用，直接执行cpp编译出来的可执行文件。