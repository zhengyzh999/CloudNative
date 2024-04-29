# jwt无状态token的签发与验证
# 公钥、私钥、证书之间的关系
1. 公钥: 对外公开。
2. 私钥: 不公开。
3. 作用: 加密与解密、签名与验签
4. 使用方式: 客户端会拿到公钥，通过公钥加密消息，将加密后的消息传送到服务端。服务端使用私钥进行解密。

# jwt的无状态特性的两个风险
1. 一旦签发，是无法更改的。token直到失效前，都是可以使用的，有被拦截使用的风险
2. 公钥通过物理或网络方式传输到客户端。
3. 通过网络传输，服务端第一次将公钥传递到客户端时。有可能被第三方获取，第三方自己生成一个公钥，进行每次的客户端与服务端数据交互，产生风险

# root ca 权威机构的公钥私钥
1. root ca(签名，权威机构签发的)预装到操作系统中的，公钥+签名=证书。
2. ca.key: 签名将要签发出去的公钥
3. 签名是，私钥签名，公钥验签;解密是，公钥加密，私钥解密

# root ca如何保证服务器第一次向客户端发送的公钥不被篡改
1. 权威机构在系统或浏览器中预装自己的公钥私钥
2. 服务向权威机构申请证书
3. 权威机构生成服务器使用的公钥私钥
4. 权威机构使用自身的私钥签名服务器公钥，产生服务器证书
5. 向申请方返回服务器证书和私钥(证书: 服务器公钥+服务器公钥的签名)
6. 服务器部署申请到的证书
7. 客户端第一次请求服务器，服务器向客户端发送证书
8. 服务端使用权威机构的公钥(root ca)对证书进行验，即可确认服务器发送的公钥是否被篡改
# jwt的两种签名与验签实践
1. 加盐
2. 公钥私钥

# jwt签发后的token部分
1. eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
2. 头部;携带的信息;签名
3. 头部和携带的信息为明文传输，不适合传递敏感信息