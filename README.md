# Example

在執行的路徑放置ssh.yaml
``` go
package main

import (
	"fmt"
	"github.com/leeisman/ssh_tunnel"
	"os"
	"time"
)

func main() {
	ssh_tunnel.TunnelByConf(getWorkingDirPath())
	time.Sleep(time.Second * 1000)
}

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workingDirPath:", dir)
	return dir
}

```
ssh yaml
``` yaml
ssh:
  - [ {localPort},"ssh server host",{ssh server port},"remote server host",{remote server port} ]

```