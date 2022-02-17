package NoCap

import (
	"testing"
)

const (
	ApiKey = "TEST-TEST-TEST"
	Proxy = "http://user:pass@proxy.no-cap.io:2612"
)

func TestHCaptchaRegular(t *testing.T) {
	hcap := &HCaptcha{
		ApiKey:  ApiKey,
		Mode:    HCaptchaRegular,
		SiteKey: "33f96e6a-38cd-421b-bb68-7806e1764460",
		Proxy:   Proxy,
		Host:    "tinytiger.xyz",
	}
	token, err := hcap.SolveHCaptcha()

	if err != nil {
		t.Errorf("error retrieving token \nreason: %s", err)
	}

	t.Logf(token)
}

func TestHCaptchaEnterprise(t *testing.T) {
	hcap := &HCaptcha{
		ApiKey:  ApiKey,
		Mode:    HCaptchaEnterprise,
		SiteKey: "4c672d35-0701-42b2-88c3-78380b0db560",
		Proxy:   Proxy,
		Host:    "discord.com",
	}
	token, err := hcap.SolveHCaptcha()

	if err != nil {
		t.Errorf("error retrieving token \nreason: %s", err)
	}

	t.Logf(token)
}
