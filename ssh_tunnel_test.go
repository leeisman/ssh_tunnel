package ssh_tunnel

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTunnel(t *testing.T) {
	type args struct {
		localPort  int
		sshHost    string
		sshPort    int
		remoteHost string
		remotePort int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ssh tunnel",
			args: args{
				localPort:  3307,
				sshHost:    "123.123.123.123",
				sshPort:    30020,
				remoteHost: "172.31.4.210",
				remotePort: 3306,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Tunnel(tt.args.localPort, tt.args.sshHost, tt.args.sshPort, tt.args.remoteHost, tt.args.remotePort, "root")
		})
		time.Sleep(time.Second * 100)
	}
}

func TestGetTunnelConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ssh config test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTunnelConf(getWorkingDirPath())
		})
	}
}

func TestTunnelByConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestTunnelByConf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TunnelByConf(getWorkingDirPath())
		})
	}
	time.Sleep(time.Second * 100)
}

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workingDirPath:", dir)
	return dir
}
