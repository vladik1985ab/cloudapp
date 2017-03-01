package main

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"fmt"
	"encoding/json"
)
// for reason the client once work and once not , please check yours server side!!! http://prntscr.com/eeqx5d
func main() {
	c := &cfclient.Config{
		ApiAddress:   "http://api.40.69.197.1.xip.io",
		Username:     "vladislavs",
		Password:     "vladislavs",
		SkipSslValidation: true,
		UserAgent: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		//ApiAddress:   "https://api.10.244.0.34.xip.io",
		//Username:     "admin",
		//Password:     "admin",
	}

	client, e := cfclient.NewClient(c)

	fmt.Println(e)

	buildpack, _ := client.ListBuildpacks()
	fmt.Println(buildpack)

	jsonrsponse, _ := json.Marshal(buildpack);
	fmt.Println(jsonrsponse)


	//client, _ := cfclient.NewClient(c)
	//apps, _ := client.ListApps()
	//fmt.Println(apps)

	//var buildpacks[] string
	//var err error;
	//client.ListBuildpacks()
	//fmt.Println(client)
}