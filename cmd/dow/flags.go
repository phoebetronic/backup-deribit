package dow

import (
	"strings"

	"github.com/spf13/cobra"
)

type flags struct {
	Bas string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Bas, "bas", "b", "", "The base path to write downloaded market data into, e.g. /Users/xh3b4sd/data/.")
}

func (f *flags) Verify() {
	if f.Bas == "" {
		panic("--bas must not be empty")
	}
	if !strings.HasPrefix(f.Bas, "/") {
		panic("--bas must be absolute path with / prefix")
	}
}
