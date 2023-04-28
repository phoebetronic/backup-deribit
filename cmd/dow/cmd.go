package dow

import (
	"github.com/spf13/cobra"
)

const (
	use = "dow"
	sho = "Download historical market data from S3."
	lon = "Download historical market data from S3."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *flags
	{
		f = &flags{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{flags: f}).run,
		}
	}

	{
		f.Create(c)
	}

	return c, nil
}
