package main

import (
	"os"

	_ "github.com/daeuniverse/outbound/dialer/juicity"
	_ "github.com/daeuniverse/outbound/protocol/juicity"
	_ "github.com/daeuniverse/outbound/transport/tls"
)

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
