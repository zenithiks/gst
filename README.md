# gst
[![GitHub Release](https://img.shields.io/github/release/shmilwdc/gst.svg?style=popout)](https://github.com/shmilwdc/gst/releases)
[![Actions Status](https://github.com/shmilwdc/gst/workflows/Go/badge.svg)](https://github.com/shmilwdc/gst/actions)
[![GitHub License](https://img.shields.io/github/license/shmilwdc/gst)](LICENSE)

Go SSH Tunnel

## 使用步骤示例

### 以153服务器为例

1. （每个目标服务器仅第一次需要进行该步骤，以后就不用了）登陆249`跳板机`，通过跳板机登陆153`目标服务器`，然后copy本机的~/.ssh/id_rsa.pub文件里的内容到目标服务器的.ssh/authorized_keys文件里
2. ./.github/install（gst.toml中的jump、key、passphrase，改为你自己的用户名、证书名、证书密码）
3. ssh atom_xiejiang@127.0.0.1 -p35322（xiejiang改为自己的用户名），就能连接到153服务器了。或者复制.github/config文件到本机.ssh/config，然后ssh 353就能连接到153服务器了

### 自动安装脚本
```shell script
bash -c "$(curl -sL https://git.io/JezpX)"
```

### 自动卸载脚本
```shell script
bash -c "$(curl -sL https://git.io/Jezp1)"
```

### 自动重启脚本
```shell script
bash -c "$(curl -sL https://git.io/JezpM)"
```