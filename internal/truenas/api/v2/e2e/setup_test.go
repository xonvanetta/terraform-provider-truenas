package e2e

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
)

func getAccClient() v2.Client {
	acc := os.Getenv(resource.TestEnvVar)
	if acc != "1" {
		return nil
	}
	//
	host := os.Getenv("TRUENAS_HOST")
	if host == "" {
		host = "http://localhost"
	}

	apiKey := os.Getenv("TRUENAS_API_KEY")
	if apiKey == "" {
		panic("can't run e2e tests with api accClient, missing TRUENAS_API_KEY in env")
	}

	return v2.NewClient(host, apiKey)
}
