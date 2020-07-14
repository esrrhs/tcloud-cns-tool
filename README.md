# tcloud-cns-tool
腾讯云-域名解析-查询设置-工具，
代码修改自[tcloud-cns](https://github.com/sjatsh/tcloud-cns)

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
