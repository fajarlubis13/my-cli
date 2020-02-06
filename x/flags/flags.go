package flags

import (
	"flag"
	"log"
	"os"
)

type Flags struct {
}

var (
	Create = flag.NewFlagSet("create", flag.ExitOnError)
)

func (m *Flags) New() {
	Create.Parse(os.Args)

	// CreateTextPtr := Create.String("stringx", "valuex", "usageee")

	if Create.Parsed() {
		Create.PrintDefaults()
	}

	log.Println(*Create)
}

func (m *Flags) Init() {

}
