package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "golang.org/x/crypto/ssh"

    "github.com/shmilwdc/gst/conf"
    "github.com/shmilwdc/gst/log"
    "github.com/shmilwdc/gst/tunnel"
)

var fpath string

func main() {
    flag.StringVar(&fpath, "f", "conf/gst.toml", "config file location.")
    flag.Parse()

    // 加载配置文件
    cfg := conf.LoadConf(fpath)

    // 加载私钥文件
    auth := tunnel.PrivateKeyFile(cfg.Key, cfg.Passphrase)
    if auth == nil {
        log.Fatal("tunnel.PrivateKeyFile error: ssh.AuthMethod is nil")
    }

    // 开始建立隧道
    dial(cfg, auth)

    // 监控配置文件
    watch(fpath)

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
}

func dial(config *conf.Conf, auth ssh.AuthMethod) {
    for _, v := range config.Binds {
        log.Debugf("tag: %-50s\tforward: %-50s\n", v.Tag, fmt.Sprintf("[%s -> %s]", v.Local, v.Remote))
        go tunnel.NewTunnel(v.Local, config.Jump, v.Remote, auth).Dial()
    }
}

func watch(fpath string) {
    log.Debugf("start watch conf file: %s", fpath)
    viper.WatchConfig()
    viper.OnConfigChange(func(in fsnotify.Event) {
        log.Debugf("receive watch conf file event: %s", in)
        if in.Op&fsnotify.Create == fsnotify.Create || in.Op&fsnotify.Write == fsnotify.Write {
            os.Exit(1)
        }
    })
}
