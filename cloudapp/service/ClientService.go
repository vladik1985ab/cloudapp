package service

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func GetClient() *cfclient.Client {

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

	client, _ := cfclient.NewClient(c);
	return client;

}
