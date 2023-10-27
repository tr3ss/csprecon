/*
csprecon - Discover new target domains using Content Security Policy

This repository is under MIT License https://github.com/tr3ss/csprecon/blob/main/LICENSE
*/

package output

import "github.com/projectdiscovery/gologger"

// nolint: gochecknoglobals
var printed = false

const (
	Version = "v0.0.8"
	banner  = `    ______________  ________  _________  ____ 
   / ___/ ___/ __ \/ ___/ _ \/ ___/ __ \/ __ \
  / /__(__  ) /_/ / /  /  __/ /__/ /_/ / / / /
  \___/____/ .___/_/   \___/\___/\____/_/ /_/ 
          /_/                                   `
)

func ShowBanner() {
	if !printed {
		gologger.Print().Msgf("%s%s\n\n", banner, Version)
		gologger.Print().Msgf("\t\t@edoardottt, https://www.edoardoottavianelli.it/\n")
		gologger.Print().Msgf("\t\t             https://github.com/edoardottt/\n\n")

		printed = true
	}
}
