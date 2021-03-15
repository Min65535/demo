package apollo

import (
	"fmt"
	"github.com/philchia/agollo"
)

func MyApolloDemo() {
	if err := agollo.StartWithConf(&agollo.Conf{
		AppID:          "dfm",
		Cluster:        "dfm-test",
		NameSpaceNames: []string{"manage", "default"},
		IP:             "172.30.9.76:8090",
		//IP:             "172.30.9.75:20180",
	}); err != nil {
		fmt.Println("err:", err.Error())
		return
		//panic(err)
	}
	fmt.Println("namespace:", agollo.GetAllKeys("manage"))
	fmt.Println("msg:", agollo.GetStringValue("msg", ""))
}
