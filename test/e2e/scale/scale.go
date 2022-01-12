package scale

import (
	"fmt"
	"time"

	"github.com/fatedier/frp/test/e2e/framework"
	"github.com/fatedier/frp/test/e2e/framework/consts"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("[Feature: Scale]", func() {
	f := framework.NewDefaultFramework()

	Describe("Memory Test", func() {
		FIt("Multiple Client With One Server", func() {
			httpsPort := f.AllocPort()
			serverConf := consts.DefaultServerConfig
			serverConf += fmt.Sprintf(`
			vhost_https_port = %d
			log_level = debug
			log_file = /tmp/frp/frps.log
			`, httpsPort)

			clientConf := consts.DefaultClientConfig

			clientConfs := []string{}
			for i := 0; i < 1000; i++ {
				tmpConf := clientConf + fmt.Sprintf(`
				heartbeat_interval = -1
				log_file = /tmp/frp/frpc%d.log
				[https-%d]
				type = https
				local_port = {{ .%s }}
				custom_domains = test-%d.com
				`, i, i, framework.TCPEchoServerPort, i)

				clientConfs = append(clientConfs, tmpConf)
			}

			f.RunProcesses([]string{serverConf}, clientConfs)

			time.Sleep(3000 * time.Second)
		})
	})
})
