package main

import (
	"fmt"
	"os"

	vagrantB "github.com/hashicorp/packer-plugin-vagrant/builder/vagrant"
	vagrantPP "github.com/hashicorp/packer-plugin-vagrant/post-processor/vagrant"
	vagrantCloudPP "github.com/hashicorp/packer-plugin-vagrant/post-processor/vagrant-cloud"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "1.0.0"

	// VersionPrerelease is A pre-release marker for the Version. If this is ""
	// (empty string) then it means that it is a final release. Otherwise, this
	// is a pre-release such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = ""

	// PluginVersion is used by the plugin set to allow Packer to recognize
	// what version this plugin is.
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(vagrantB.Builder))
	pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(vagrantPP.PostProcessor))
	pps.RegisterPostProcessor("cloud", new(vagrantCloudPP.PostProcessor))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
