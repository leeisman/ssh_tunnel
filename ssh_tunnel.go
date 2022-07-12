package ssh_tunnel

import (
	"fmt"
	"github.com/rgzr/sshtun"
	"github.com/spf13/viper"
	"log"
	"time"
)

type SSHConfigs struct {
	SSH [][]interface{}
}

func TunnelByConf(configPath string) {
	sshConfigs := GetTunnelConf(configPath)
	if sshConfigs == nil {
		return
	}
	for _, sshConfig := range sshConfigs.SSH {
		localPort := sshConfig[0].(int)
		sshHost := fmt.Sprint(sshConfig[1])
		sshHostPort := sshConfig[2].(int)
		remoteHost := fmt.Sprint(sshConfig[3])
		remoteHostPort := sshConfig[4].(int)
		Tunnel(localPort, sshHost, sshHostPort, remoteHost, remoteHostPort)
	}
}

func GetTunnelConf(configPath string) *SSHConfigs {
	var sshConfig *SSHConfigs
	viper.SetConfigName("ssh")      // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath) // optionally look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil
	}
	err = viper.Unmarshal(&sshConfig)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil
	}
	fmt.Println(sshConfig)
	return sshConfig
}


func Tunnel(localPort int, sshHost string, sshPort int, remoteHost string, remotePort int, user string) {
	// We want to connect to port 8080 on our machine to access port 80 on my.super.host.com
	sshTun := sshtun.New(localPort, sshHost, remotePort)
	sshTun.SetPort(sshPort)
	sshTun.SetRemoteHost(remoteHost)
	// We enable debug messages to see what happens
	sshTun.SetDebug(true)
	sshTun.SetUser(user)
	// We set a callback to know when the tunnel is ready
	sshTun.SetConnState(func(tun *sshtun.SSHTun, state sshtun.ConnState) {
		switch state {
		case sshtun.StateStarting:
			log.Printf("STATE is Starting")
		case sshtun.StateStarted:
			log.Printf("STATE is Started")
		case sshtun.StateStopped:
			log.Printf("STATE is Stopped")
		}
	})

	// We start the tunnel (and restart it every time it is stopped)
	go func() {
		if err := sshTun.Start(); err != nil {
			log.Printf("SSH tunnel stopped: %s", err.Error())
			time.Sleep(time.Second) // don't flood if there's a start error :)
		}
	}()

	// We stop the tunnel every 20 seconds (just to see what happens)
	//for {
	//	time.Sleep(time.Second * time.Duration(20))
	//	log.Println("Lets stop the SSH tunnel...")
	//	sshTun.Stop()
	//}
}
