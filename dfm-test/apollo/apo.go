package apollo

import (
	"fmt"
	"github.com/philchia/agollo"
)

func MyApolloDemo() {
	if err := agollo.StartWithConf(&agollo.Conf{
		AppID:          "demo",
		Cluster:        "default",
		NameSpaceNames: []string{"xl-dev"},
		//IP:             "172.30.9.76:10071",
		IP:             "172.30.9.75:20180",
	}); err != nil {
		panic(err)
	}

	fmt.Println("msg:",agollo.GetStringValue("msg", ""))
}