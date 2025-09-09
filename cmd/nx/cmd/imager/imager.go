package imager

import (
	"github.com/spf13/cobra"

	_root "github.com/himalayo/nx/cmd/nx/cmd"
)

var Cmd = &cobra.Command{
	Use:     "imager",
	Aliases: []string{"img"},
	Short:   "Render resources to images",
}

func init() {
	_root.Cmd.AddCommand(Cmd)
}
