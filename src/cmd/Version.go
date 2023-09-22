package Cmd

import (
	"fmt"
	"log"

	"runtime/debug"

	"github.com/spf13/cobra"
)

func Version(
	version string,
) *cobra.Command {
	__cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Run: func(cmd *cobra.Command, args []string) {
			var OperatingSystem string
			var Arch string
			//var LastCommit time.Time
			var LC string
			var GitHash string

			fmt.Printf("version: '%s'\n", version)

			if buildInfo, ok := debug.ReadBuildInfo(); ok {
				//log.Printf("Version: %#v\n", buildInfo)
				//log.Printf("Version: %#v\n", buildInfo.Deps)
				//log.Printf("Version: %#v\n", buildInfo.Settings)
				for _, kv := range buildInfo.Settings {
					switch kv.Key {
					case "GOOS":
						OperatingSystem = kv.Value
					case "GOARCH":
						Arch = kv.Value
					case "vcs.time":
						//LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
						LC = kv.Value
					case "vcs.revision":
						GitHash = kv.Value
					}
				}
				log.Printf("Main.Version: %#v\n", buildInfo.Main.Version)
				log.Printf("OperatingSystem: %#v\n", OperatingSystem)
				log.Printf("Arch: %#v\n", Arch)
				log.Printf("Last commit: %#v\n", LC)
				log.Printf("Git Hash: %#v\n", GitHash)
				log.Printf("%s\n", "Module dependencies")
				for _, mod := range buildInfo.Deps {
					log.Printf("\t%s %s\n", mod.Path, mod.Version)
				}
			}
		},
	}

	return __cmd
}
