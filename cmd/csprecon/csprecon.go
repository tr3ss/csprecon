/*
csprecon - Discover new target domains using Content Security Policy

This repository is under MIT License https://github.com/tr3ss/csprecon/blob/main/LICENSE
*/

package main

import (
	"github.com/tr3ss/csprecon/pkg/csprecon"
	"github.com/tr3ss/csprecon/pkg/input"
)

func main() {
	options := input.ParseOptions()
	runner := csprecon.New(options)
	runner.Run()
}
