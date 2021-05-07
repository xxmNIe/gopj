package util

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
)
//总的来说该函数是用于处理从证书凭证文件（PEM），最终获取tls.Config作为HTTP2的使用参数
func GetTLSConfig(certPemPath,certKeyPath string) *tls.Config{
	var certkeyPair *tls.Certificate

	cert,_ := ioutil.ReadFile(certPemPath)
	key,_ := ioutil.ReadFile(certKeyPath)
	//从一对PEM编码的数据中解析公钥/私钥对。成功则返回公钥/私钥对
	pair,err := tls.X509KeyPair(cert,key)
	if err != nil {
		log.Println("TLS KeyPair err: %v\n", err)
	}
	certkeyPair = &pair

	return &tls.Config{
		Certificates: []tls.Certificate{*certkeyPair},
		//NextProtoTLS是谈判期间的NPN/ALPN协议，用于HTTP/2的TLS设置
		NextProtos: []string{http2.NextProtoTLS},
	}
}
