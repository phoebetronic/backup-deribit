package upl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/phoebetronic/backup-deribit/pkg/apicliaws"
	"github.com/phoebetronic/backup-deribit/pkg/apiclideribit"
	"github.com/phoebetronic/backup-deribit/pkg/candle"
	"github.com/spf13/cobra"
)

const (
	buc = "phoebetronic"
	pre = "eth-usd"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	der := apiclideribit.New()
	aws := apicliaws.New()

	var lis []string
	{
		lis, err = aws.Lister(buc, pre)
		if err != nil {
			panic(err)
		}
	}

	var sta time.Time
	if len(lis) == 0 {
		sta = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	} else {
		sta = latest(lis).UTC().Add(24 * time.Hour)
	}

	var end time.Time
	{
		end = sta.Add(24 * time.Hour)
	}

	var now time.Time
	{
		now = time.Now().UTC()
	}

	if end.After(now) {
		fmt.Printf(
			"cannot backup historical candles between %s and %s as current day is not complete at %s\n",
			sta.Format("2006-01-02"),
			end.Format("2006-01-02"),
			now,
		)

		return
	} else {
		fmt.Printf(
			"backing up historical candles between %s and %s\n",
			sta.Format("2006-01-02"),
			end.Format("2006-01-02"),
		)
	}

	var can []candle.Candle
	{
		can, err = der.Candles(sta, end, "ETH-PERPETUAL")
		if err != nil {
			panic(err)
		}
	}

	var pat string
	{
		pat = filepath.Join(pre, sta.Format("2006-01-02")+".json")
	}

	var byt []byte
	{
		byt, err = json.Marshal(can)
		if err != nil {
			panic(err)
		}
	}

	var rea bytes.Reader
	{
		rea = *bytes.NewReader(byt)
	}

	{
		err := aws.Upload(buc, pat, rea)
		if err != nil {
			panic(err)
		}
	}
}

func latest(input []string) time.Time {
	var err error
	var lat time.Time

	for _, s := range input {
		var spl []string
		{
			spl = strings.Split(s, "/")
		}

		var prt string
		{
			prt = spl[len(spl)-1]
		}

		var dat time.Time
		{
			dat, err = time.Parse("2006-01-02", strings.ReplaceAll(prt, filepath.Ext(prt), ""))
			if err != nil {
				panic(err)
			}
		}

		if dat.After(lat) {
			lat = dat
		}
	}

	return lat
}
