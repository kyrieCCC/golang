## 记录一个echo服务器的细节部分
### net.Listen(network, address string) Listener, error
该方法表示**在本地网络地址上收听通告**，注意这里是本地网络地址上，该方法返回两个参数，第一个参数表示server，我们可以针对server进行一系列操作，第二个参数为错误信息

该方法有两个参数：
1. network，第一个参数表示网络，网络必须是“tcp”、“tcp4”、“tcp6”、“unix”或“unixpacket”的其中一个。

2. address，第二个参数表示ip地址，对于 TCP 网络，如果 address 参数中的主机为空或文字未指定的 IP 地址，Listen 将侦听本地系统的所有可用单播和任播 IP 地址。如果地址参数中的端口为空或“0”，如“127.0.0.1:”或“[::1]:0”，则会自动选择一个端口号。 Listener 的 Addr 方法可用于发现所选端口。

### server.Accept() (Conn, err)
Accept实现了Listener接口中的Accept方法； 它等待下一次调用并返回一个通用的 Conn。

该方法可以不传入参数，每次进行执行成功后都会返回一个**连接（Conn）**，我们可以在后面定义方法来处理这个连接，错误后会返回错误信息