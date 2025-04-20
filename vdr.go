package vdr_core

import (
	"os"
	pluginUtil "plugin"

	"github.com/eclipse-xfsc/ssi-vdr-core/types"
)

const (
	VdrTypeKey             = "VDR_TYPE"
	PluginsLocationKey     = "PLUGINS_LOCATION"
	PluginSymbolKey        = "Plugin"
	DefaultPluginsLocation = "etc/plugins/"
)

var provider types.VerifiableDataRegistry

func initialize() {
	vdrType := os.Getenv(VdrTypeKey)
	pluginsLocation := os.Getenv(PluginsLocationKey)
	if pluginsLocation == "" {
		pluginsLocation = DefaultPluginsLocation
	}
	plugin, err := pluginUtil.Open(pluginsLocation + vdrType)
	if err != nil {
		panic(err)
	}

	v, err := plugin.Lookup(PluginSymbolKey)
	if err != nil {
		panic(err)
	}
	provider, err = v.(types.VerifiableDataRegistryModule).GetVerifiableDataRegistry()
	if err != nil {
		panic(err)
	}
}

func VerifiableDataRegistryInitializer() types.VerifiableDataRegistry {
	initialize()
	return provider
}
