package provider_test

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kairos-io/kairos-sdk/bus"

	. "github.com/kairos-io/provider-kairos/internal/provider"
	providerConfig "github.com/kairos-io/provider-kairos/internal/provider/config"
	"github.com/mudler/go-pluggable"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Bootstrap provider", func() {
	Context("logging", func() {
		e := &pluggable.Event{}

		BeforeEach(func() {
			e = &pluggable.Event{}
		})

		It("logs to file", func() {
			f, err := ioutil.TempFile(os.TempDir(), "tests")
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(f.Name())

			cfg := &providerConfig.Config{
				P2P: &providerConfig.P2P{
					NetworkToken: "foo",
				},
			}
			dat, err := yaml.Marshal(cfg)
			Expect(err).ToNot(HaveOccurred())
			payload := &bus.BootstrapPayload{Logfile: f.Name(), Config: string(dat)}

			dat, err = json.Marshal(payload)
			Expect(err).ToNot(HaveOccurred())

			e.Data = string(dat)
			resp := Bootstrap(e)
			dat, _ = json.Marshal(resp)
			Expect(resp.Errored()).To(BeTrue(), string(dat))

			data, err := ioutil.ReadFile(f.Name())
			Expect(err).ToNot(HaveOccurred())

			Expect(string(data)).Should(ContainSubstring("Configuring VPN"))
		})
	})
})
