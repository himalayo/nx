package main

import (
	"github.com/himalayo/nx/cmd/nx/cmd"

	_ "github.com/himalayo/nx/cmd/nx/cmd/figure"
	_ "github.com/himalayo/nx/cmd/nx/cmd/figure/convert"
	_ "github.com/himalayo/nx/cmd/nx/cmd/figure/info"

	_ "github.com/himalayo/nx/cmd/nx/cmd/furni"
	_ "github.com/himalayo/nx/cmd/nx/cmd/furni/info"
	_ "github.com/himalayo/nx/cmd/nx/cmd/furni/search"

	_ "github.com/himalayo/nx/cmd/nx/cmd/get"
	_ "github.com/himalayo/nx/cmd/nx/cmd/get/furni"

	_ "github.com/himalayo/nx/cmd/nx/cmd/profile"

	_ "github.com/himalayo/nx/cmd/nx/cmd/visual"

	_ "github.com/himalayo/nx/cmd/nx/cmd/imager"
	_ "github.com/himalayo/nx/cmd/nx/cmd/imager/avatar"
	_ "github.com/himalayo/nx/cmd/nx/cmd/imager/furni"

	_ "github.com/himalayo/nx/cmd/nx/cmd/texts"

	_ "github.com/himalayo/nx/cmd/nx/cmd/vars"

	_ "github.com/himalayo/nx/cmd/nx/cmd/extract"
)

func main() {
	cmd.Execute()
}
