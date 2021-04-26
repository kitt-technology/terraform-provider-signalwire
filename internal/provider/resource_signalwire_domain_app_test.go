package provider

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSignalwireDomainApp_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSignalwireDomainAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSignalwireDomainAppConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSignalwireDomainAppExists("signalwire_domain_app.test_app"),
				),
			},
		},
	})
}

func testAccCheckSignalwireDomainAppDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "signalwire_domain_app" {
			continue
		}

		resp, err := client.Req("GET", os.Getenv("SIGNALWIRE_SPACE"), "domain_applications", nil)
		if err != nil {
			return err
		}

		endpoints := resp["data"].([]interface{})

		if len(endpoints) > 0 {
			return errors.New("Domain App still exists")
		}
		return nil
	}
	return nil
}

func testAccCheckSignalwireDomainAppExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No key ID is set")
		}

		client := testAccProvider.Meta().(*Client)

		resp, err := client.Req("GET", os.Getenv("SIGNALWIRE_SPACE"), "domain_applications/"+rs.Primary.ID, nil)

		if err != nil {
			return err
		}

		if resp["id"] != rs.Primary.ID {
			return errors.New("SIP Endpoint does not match")
		}

		return nil
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func testAccSignalwireDomainAppConfig() string {
	name := RandStringRunes(10)
	return fmt.Sprintf(`
		resource "signalwire_domain_app" "test_app" {
			space = "%[1]s"
			name = "%s"
			identifier = "%s"
			ip_auth_enabled = true
			ip_auth = ["8.8.8.8", "4.4.4.4"]
			encryption = "required"
			call_handler = "relay_context"
			call_relay_context = "incoming"
			ciphers = [
				"AEAD_AES_256_GCM_8",
			]
			codecs = [
				"PCMU",
				"PCMA",
			]
		}
	`, testSignalwireSpace, name, name)
}
