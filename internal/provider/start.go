package provider

import (
	"os"

	"github.com/kairos-io/kairos-sdk/bus"

	"github.com/mudler/go-pluggable"
)

func Start() error {
	factory := pluggable.NewPluginFactory()

	// Input: bus.EventInstallPayload
	// Expected output: map[string]string{}
	factory.Add(bus.EventInstall, Install)

	factory.Add(bus.EventBootstrap, Bootstrap)

	// Input: config
	// Expected output: string
	factory.Add(bus.EventChallenge, Challenge)

	factory.Add(bus.EventRecovery, Recovery)

	factory.Add(bus.EventRecoveryStop, RecoveryStop)

	factory.Add(bus.EventInteractiveInstall, InteractiveInstall)

	factory.Add(bus.EventAvailableReleases, ListVersions)

	return factory.Run(pluggable.EventType(os.Args[1]), os.Stdin, os.Stdout)
}
