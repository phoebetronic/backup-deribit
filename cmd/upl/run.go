package upl

import (
	"fmt"
	"time"

	"github.com/phoebetronic/backup-deribit/pkg/apiclideribit"
	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	client := apiclideribit.New()

	start := time.Date(2023, 4, 26, 12, 5, 0, 0, time.UTC)
	end := time.Date(2023, 4, 26, 12, 10, 0, 0, time.UTC)
	candles, err := client.GetCandles(start, end, "ETH-PERPETUAL")
	if err != nil {
		fmt.Printf("Error getting candles: %v", err)
		return
	}

	for _, c := range candles {
		fmt.Printf(
			"%s - O: %.2f, H: %.2f, L: %.2f, C: %.2f\n",
			c.Time.UTC(),
			c.Open,
			c.High,
			c.Low,
			c.Close,
		)
	}
}
