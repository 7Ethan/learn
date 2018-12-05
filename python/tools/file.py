import os
'''获得当前路径
'''
cwd=os.getcwd()
print(cwd)


'''
得到当前文件夹下的所有文件和文件夹
'''
print(os.listdir())


'''
delete file
'''
os.remove('a.txt')
print(os.listdir())


'''
删除单个目录和多个目录
'''
# os.removedir()
# os.removedir()
os.removedirs("a")


'''
检查是否是文件／文件夹
'''
print(os.path.isfile('/Users/sergey/GitHub/learn/text.sh'))
print(os.path.isdir('/Users/sergey/GitHub/learn'))



print("检查文件路径是否存在:")
print(os.path.exists(os.getcwd()))

'''
分离文件名
分离扩展名
'''
[dirname,filename]=os.path.split('/Users/sergey/GitHub/learn/README.md')
print(dirname,"\n",filename)

[fname,fename]=os.path.splitext('/Users/sergey/GitHub/learn/README.md')
print(fname,"\n",fename)

'''
获得文件路径
获得文件名
获得当前环境
'''
print("get pathname:",os.path.dirname('/Users/sergey/GitHub/learn/README.md'))
print("get filename:",os.path.basename('/Users/sergey/GitHub/learn/README.md'))
print(os.getenv)

#################################

# > sh-4.4$ python file.py
# C:\Users\sergey\GitHub\learn\python\tools
# ['a', 'a.txt', 'file.py']
# ['a', 'file.py']
# False
# True
# 检查文件路径是否存在:
# True
# /Users/sergey/GitHub/learn
#  README.md
# /Users/sergey/GitHub/learn/README
#  .md
# get pathname: /Users/sergey/GitHub/learn
# get filename: README.md
# <function getenv at 0x00000000027E5950>