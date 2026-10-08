package utils

import (
	"crypto/tls"
	"net/http"
)

// InsecureClient 是一个忽略 HTTPS 证书验证的 HTTP 客户端。
var InsecureClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}
