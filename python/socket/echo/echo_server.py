#!/usr/bin/env python3

import socket

HOST = "127.0.0.1"
PORT = 8080

with socket.socket(socket.AF_INET,socket.SOCK_STREAM) as s :
    s.bind((HOST,PORT))
    s.listen()

    conn,err = s.accept()
    with conn:
        while True:
            data = conn.recv(1024)
            print("Receved data: ",data)
            if not data:
                break
            conn.sendall(b'Server reply :' + data)

# 在Windows环境下测试实例 
# =======终端作为client ==========

# 》》》 client 
# $ curl -XPOST http://127.0.0.1:8080 -d'a'
#   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
#                                  Dload  Upload   Total   Spent    Left  Speed
# 100   163    0   162    0     1     11      0 --:--:--  0:00:14 --:--:--     0Server reply :POST / HTTP/1.1
# Host: 127.0.0.1:8080
# User-Agent: curl/7.62.0
# Accept: */*
# Content-Length: 1
# Content-Type: application/x-www-form-urlencoded

# a

# 》》》 server
# $ python echo_server.py
# Traceback (most recent call last):
#   File "echo_server.py", line 15, in <module>
#     data = conn.recv(1024)
# ConnectionResetError: [WinError 10054] 远程主机强迫关闭了一个现有的连接。
# ！！！！！！！！ Python ‘with’ 黑魔法！！