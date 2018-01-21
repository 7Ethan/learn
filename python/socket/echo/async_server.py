import asyncio
import socket

HOST = "127.0.0.1"
PORT = 8080

# async def echo_server():
#     reader,writer = await asyncio.open_connection(HOST,PORT)

#     readData = await reader.read(100)
#     print(f'Receved data : {readData.decode() !r}')

#     writer.write(readData.encode())

#     print("Close the connection!")
#     writer.close()
#     await writer.wait_closed()

async def echo_server():
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


asyncio.run(echo_server())

