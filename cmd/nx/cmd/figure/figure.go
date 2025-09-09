package figure

import (
	"github.com/spf13/cobra"

	_root "github.com/himalayo/nx/cmd/nx/cmd"
)

var Cmd = &cobra.Command{
	Use: "figure",
}

func init() {
	_root.Cmd.AddCommand(Cmd)
}
