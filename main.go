package main

import (
	"fmt"

	"github.com/imhshekhar47/go-api-core/config"
)

func main() {
	appConfig := config.GetApplicationConfig()
	fmt.Println(appConfig.Yaml())
}
