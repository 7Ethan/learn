#coding:utf-8
for i in range(2,101):
    flag = False
    for j in range(2,i - 1):
        if i %j == 0:
            flag = True
            break
    
    if flag == False:
        print(i)