package tunnel

import (
    "io/ioutil"

    "golang.org/x/crypto/ssh"

    "github.com/shmilwdc/gst/log"
)

func PrivateKeyFile(file, passphrase string) ssh.AuthMethod {
    buffer, err := ioutil.ReadFile(file)
    if err != nil {
        return nil
    }

    var key ssh.Signer

    if passphrase != "" {
        key, err = ssh.ParsePrivateKeyWithPassphrase(buffer, []byte(passphrase))
        if err != nil {
            log.Errorf("ssh.ParsePrivateKeyWithPassphrase error: %s", err)
            return nil
        }
    } else {
        key, err = ssh.ParsePrivateKey(buffer)
        if err != nil {
            log.Errorf("ssh.ParsePrivateKey error: %s", err)
            return nil
        }
    }

    return ssh.PublicKeys(key)
}
