# RPC学习项目

<a href="https://www.youtube.com/watch?v=2Sm_O75I7H0">原视频</a>

## 序列化对象为二进制和Json

### Go

1. 将protobuf信息写入二进制文件
2. 从二进制文件中读取protobuf信息
3. 写入json文件并比较大小

## gRPC 的四种通信方式

1. 类似REST的单请求+单回复
2. 客户端多请求+服务端单回复
3. 客户端单请求+服务端多回复
4. 客户端多请求+服务端多回复

## gRPC 反射

<a href="https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md">grpc反射</a>
gRPC 服务器反射提供有关服务器上可公开访问的 gRPC 服务的信息，并帮助客户端在运行时构造 RPC 请求和响应，而无需预编译的服务信息。
它由 gRPC CLI 使用，可用于自省服务器原型和发送/接收测试 RPC。

### evans

一个grpc客户端,在服务端开启反射并运行时，通过`evans -r repl -p 端口`进入shell，
通过`show package`查看反射的包信息，通过`package 包名`选择不同的包,通过`show service`查看反射的服务信息,通过`service 服务`选择服务，通过`call CreateLaptop`
调用服务，中途使用`ctrl+D`取消重复字段的输入...
<a href="https://github.com/ktr0731/evans">evans</a>

## gRPC 拦截器

类似于中间件，可以在服务端和客户端之间添加的额外功能，服务器端拦截器是gRPC服务器在调用实际RPC方法前将调用的函数，可以用于日志记录，跟踪，限流，身份验证，限流等。
客户端拦截器是gRPC客户端在调用实际RPC方法前将调用的函数.

服务器端拦截器将采用JWT来验证，客户端拦截器将添加JWT到请求。

## SSL/TLS

### 详解

TLS 是传输层安全协议，它用于实现客户端和服务器之间的加密通信。SSL是TLS的前身。

![image-20220504204623342](https://gitee.com/ORaja/picture/raw/master/img/image-20220504204623342.png)

TLS在网络（HTTPS = HTTP+TLS），邮件（SMTPS = SMTP+TLS），文件传输（FTPS = FTP+TLS）等中使用。

作用：

1. 身份验证 证明访问的网站不是伪造的

   将服务器公钥放入到**数字证书**中，解决了冒充的风险

2. 信息加密 交互信息无法被窃取

   通过**混合加密**的方式可以保证信息的**机密性**，解决了窃听的风险

   HTTPS 采用的是**对称加密**和**非对称加密**结合的「混合加密」方式：

   - 在通信建立前采用**非对称加密**的方式交换「会话秘钥」，后续就不再使用非对称加密。
   - 在通信过程中全部使用**对称加密**的「会话秘钥」的方式加密明文数据。

3. 校验机制 无法篡改通信内容

   **摘要算法**用来实现**完整性**，能够为数据生成独一无二的「指纹」，用于校验数据的完整性，解决了篡改的风险。

一般采用ECDHE密钥协商算法生成会话密钥

1. TLS第一次握手

   客户端首先会发一个「**Client Hello**」消息，消息里面有客户端使用的 TLS 版本号、支持的密码套件列表，以及生成的**随机数（\*Client Random\*）**。

2. TLS第二次握手

   服务端收到客户端的「打招呼」返回「**Server Hello**」消息，消息面有服务器确认的 TLS 版本号，也给出了一个**随机数（\*Server Random\*）**，然后从客户端的密码套件列表选择了一个合适的密码套件。接着，服务端为了证明自己的身份，发送「**Certificate**」消息，会把证书也发给客户端。

   因为服务端选择了 ECDHE 密钥协商算法，所以会在发送完证书后，发送「**Server Key Exchange**」消息。

   - 选择了**椭圆曲线**，选好了椭圆曲线相当于椭圆曲线基点 G 也定好了，这些都会公开给客户端；
   - 生成随机数作为服务端椭圆曲线的私钥，保留到本地；
   - 根据基点 G 和私钥计算出**服务端的椭圆曲线公钥**，这个会公开给客户端。

   为了保证这个椭圆曲线的公钥不被第三方篡改，服务端会用 RSA 签名算法给服务端的椭圆曲线公钥做个签名。

3. TLS第三次握手

   客户端收到了服务端的证书后，校验证书是否合法。

   客户端会生成一个随机数作为客户端椭圆曲线的私钥，然后再根据服务端前面给的信息，生成**客户端的椭圆曲线公钥**，然后用「**Client Key Exchange**」消息发给服务端

   **最终的会话密钥，就是用「客户端随机数 + 服务端随机数 + x（ECDHE 算法算出的共享密钥） 」三个材料生成的**。

   算好会话密钥后，客户端会发一个「**Change Cipher Spec**」消息，告诉服务端后续改用对称算法加密通信。

   接着，客户端会发「**Encrypted Handshake Message**」消息，把之前发送的数据做一个摘要，再用对称密钥加密一下，让服务端做个验证，验证下本次生成的对称密钥是否可以正常使用。

4. TLS第四次握手

   最后，服务端也会有一个同样的操作，发「**Change Cipher Spec**」和「**Encrypted Handshake Message**」消息，如果双方都验证加密和解密没问题，那么握手正式完成。于是，就可以正常收发加密的 HTTP 请求和响应了。

2. INSECURE 无安全验证
3. SERVER-SIDE TLS 服务端证书
4. MUTUAL SSL 客户端与服务端证书

