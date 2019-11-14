package tunnel

import (
    "io"
    "net"

    "golang.org/x/crypto/ssh"

    "github.com/shmilwdc/gst/log"
)

type Tunnel struct {
    Local  *Endpoint
    Jump   *Endpoint
    Remote *Endpoint
    Config *ssh.ClientConfig
}

func (t *Tunnel) Dial() {
    addr, err := net.ResolveTCPAddr("tcp", t.Local.String())
    if err != nil {
        log.Errorf("local net.ResolveTCPAddr error: %s\n", err)
        return
    }

    listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        log.Errorf("local net.ListenTCP error: %s\n", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.AcceptTCP()
        if err != nil {
            log.Errorf("local listener.Accept error: %s\n", err)
            return
        }

        err = conn.SetKeepAlive(true)
        if err != nil {
            log.Errorf("local conn.SetKeepAlive error: %s\n", err)
        }

        go t.forward(conn)
    }
}

func (t *Tunnel) forward(localConn net.Conn) {
    jumpConn, err := ssh.Dial("tcp", t.Jump.String(), t.Config)
    if err != nil {
        log.Errorf("jump ssh.Dial error: %s", err)
        return
    }

    remoteConn, err := jumpConn.Dial("tcp", t.Remote.String())
    if err != nil {
        log.Errorf("remote ssh.Dial error: %s", err)
        return
    }

    go copyConn(localConn, remoteConn)
    go copyConn(remoteConn, localConn)
}

func copyConn(dst, src net.Conn) {
    defer func() {
        _ = src.Close()
        _ = dst.Close()
    }()

    _, err := io.Copy(dst, src)
    if err != nil {
        log.Errorf("io.Copy error: %s\n", err)
    }
}

func NewTunnel(local, jump, remote string, auth ssh.AuthMethod) *Tunnel {
    localEndpoint := NewEndpoint(local)
    jumpEndpoint := NewEndpoint(jump)
    remoteEndpoint := NewEndpoint(remote)

    tunnel := &Tunnel{
        Local:  localEndpoint,
        Jump:   jumpEndpoint,
        Remote: remoteEndpoint,
        Config: &ssh.ClientConfig{
            User:            jumpEndpoint.User,
            Auth:            []ssh.AuthMethod{auth},
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        },
    }

    return tunnel
}
