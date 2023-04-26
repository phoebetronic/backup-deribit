package upl

import (
	"github.com/spf13/cobra"
)

const (
	use = "upl"
	sho = "Upload historical market data."
	lon = "Upload historical market data."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
		}
	}

	return c, nil
}
