package tests

import (
	"os"
	"os/exec"
	"testing"

	core "github.com/eclipse-xfsc/ssi-vdr-core"
)

const PluginType = "test"
const PluginsLocation = "./data/plugins/"

func TestVerifiableDataRegistryInitializer(t *testing.T) {
	err := buildTestPlugin()
	if err != nil {
		t.Error(err)
		return
	}
	defer clearTestPlugin(t)
	os.Setenv(core.VdrTypeKey, PluginType)
	os.Setenv(core.PluginsLocationKey, PluginsLocation)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("LDFLAGS", "-Wl,-ld_classic")
	vdr := core.VerifiableDataRegistryInitializer()
	if vdr == nil {
		t.Error("did not load plugin")
	}
}

func buildTestPlugin() error {
	cmd := exec.Command("go", "build", "-C", "./data", "-o", "plugins/"+PluginType, "-x", "-buildmode=plugin")
	return cmd.Run()
}

func clearTestPlugin(t *testing.T) error {
	if r := recover(); r != nil {
		t.Errorf("did not load plugin. Error: %s", r)
	}
	return os.RemoveAll(PluginsLocation)
}
