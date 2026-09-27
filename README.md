# tcloud-cns-tool

[<img src="https://img.shields.io/github/license/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/languages/top/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/v/release/esrrhs/tcloud-cns-tool">](https://github.com/esrrhs/tcloud-cns-tool/releases)
[<img src="https://img.shields.io/github/downloads/esrrhs/tcloud-cns-tool/total">](https://github.com/esrrhs/tcloud-cns-tool/releases)
[<img src="https://img.shields.io/docker/pulls/esrrhs/tcloud-cns-tool">](https://hub.docker.com/repository/docker/esrrhs/tcloud-cns-tool)
[<img src="https://img.shields.io/github/actions/workflow/status/esrrhs/tcloud-cns-tool/go.yml?branch=master">](https://github.com/esrrhs/tcloud-cns-tool/actions)
[<img src="https://img.shields.io/github/actions/workflow/status/esrrhs/tcloud-cns-tool/test.yml?branch=master&label=test">](https://github.com/esrrhs/tcloud-cns-tool/actions)

腾讯云 DNS（CNS）解析记录的命令行工具：查询域名下解析，或按名称覆盖写入记录。

基于 [tcloud-cns](https://github.com/sjatsh/tcloud-cns) 修改。

# 安装

从 [Releases](https://github.com/esrrhs/tcloud-cns-tool/releases) 下载对应平台压缩包，或本地编译：

```bash
git clone https://github.com/esrrhs/tcloud-cns-tool.git
cd tcloud-cns-tool
go build
```

Docker：

```bash
docker pull esrrhs/tcloud-cns-tool
```

# 准备密钥

在腾讯云 [API 密钥](https://console.cloud.tencent.com/cam/capi) 页面创建密钥，拿到 `SecretId` 与 `SecretKey`。

# 使用

查询域名下全部解析记录：

```bash
./tcloud-cns-tool -secretId xxxx -secretKey yyyy -domain yourdomain.com -l
```

设置解析记录（同名记录会先删除再添加）：

```bash
./tcloud-cns-tool -secretId xxxx -secretKey yyyy -domain yourdomain.com -s -name test -value 1.2.3.4
```

指定记录类型（默认 `A`）：

```bash
./tcloud-cns-tool -secretId xxxx -secretKey yyyy -domain yourdomain.com -s -name www -value 1.2.3.4 -type A
```

# 参数

| 参数 | 说明 |
| --- | --- |
| `-secretId` | 腾讯云 SecretId（必填） |
| `-secretKey` | 腾讯云 SecretKey（必填） |
| `-domain` | 主域名，例如 `example.com`（必填） |
| `-l` | 列出该域名下的解析记录 |
| `-s` | 设置解析记录 |
| `-name` | 主机记录，例如 `www`、`@` 对应空主机时填 `@` 或按控制台习惯填写 |
| `-value` | 记录值，例如 IP 或 CNAME 目标 |
| `-type` | 记录类型，默认 `A` |
| `-version` | 打印版本并退出 |

`-l` 与 `-s` 二选一使用；使用 `-s` 时必须同时提供 `-name` 与 `-value`。

# 发布

修改 `version.go` 中的版本号并推送到 `master` 后，CI 会自动交叉编译并发布到 GitHub Releases：

```go
const Version = "0.3"
```

# License

[MIT](LICENSE)
