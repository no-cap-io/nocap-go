

# Example
```go
package main

import (
	"github.com/no-cap-io/nocap-go"
	"log"
)

func main() {
	hcap := &NoCap.HCaptcha{
		ApiKey:  "",
		Mode:    NoCap.HCaptchaEnterprise,
		SiteKey: "33f96e6a-38cd-421b-bb68-7806e1764460",
		Proxy:   "http://localhost",
		Host:    "tinytiger.xyz",
	}
	token, err := hcap.SolveHCaptcha()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("success", token)
}
```

	

> Proxy format is `http://user:pass@ip:port`
