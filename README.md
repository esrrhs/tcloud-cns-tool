# tcloud-cns-tool

[<img src="https://img.shields.io/github/license/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/languages/top/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool)
[![Go Report Card](https://goreportcard.com/badge/github.com/esrrhs/tcloud-cns-tool)](https://goreportcard.com/report/github.com/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/v/release/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool/releases)
[<img src="https://img.shields.io/github/downloads/esrrhs/tcloud-cns-tool/total">](https://github.com/esrrhs/tcloud-cns-tool/releases)
[<img src="https://img.shields.io/docker/pulls/esrrhs/tcloud-cns-tool">](https://hub.docker.com/repository/docker/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/workflow/status/esrrhs/tcloud-cns-tool/Go">](https://github.com/esrrhs/tcloud-cns-tool/actions)

腾讯云-域名解析-查询设置-工具，
代码修改自[tcloud-cns](https://github.com/sjatsh/tcloud-cns)

# 编译
```
# git clone https://github.com/esrrhs/tcloud-cns-tool.git
# cd tcloud-cns-tool
# go build
```

# 使用
* 去腾讯云申请API密钥，地址：[capi](https://console.cloud.tencent.com/cam/capi)，拿到```SecretId```、```SecretKey```
* 查询名下域名解析
```
# ./tcloud-cns-tool -secretId xxxx -secretKey yyyy -domain yourdomain.com -l
```
* 设置域名解析，内部实现会先删除再添加
```
# ./tcloud-cns-tool -secretId xxxx -secretKey yyyy -domain yourdomain.com -s -name test -value 1.2.3.4
```
