package conf

import (
    "github.com/spf13/viper"

    "github.com/shmilwdc/gst/log"
)

type Conf struct {
    Jump       string  `toml:"jump"`
    Key        string  `toml:"key"`
    Passphrase string  `toml:"passphrase"`
    Binds      []*Bind `toml:"binds"`
}

type Bind struct {
    Tag    string `toml:"tag"`
    Local  string `toml:"local"`
    Remote string `toml:"remote"`
}

func LoadConf(fpath string) *Conf {
    var conf *Conf

    viper.SetConfigFile(fpath)

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal("viper.ReadInConfig error: %s", err)
    }

    err = viper.Unmarshal(&conf)
    if err != nil {
        log.Fatal("viper.Unmarshal error: %s", err)
    }

    return conf
}
