package apiclideribit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/phoebetronic/backup-deribit/pkg/candle"
)

type Client struct {
	api string
	res int
}

func New() *Client {
	return &Client{
		api: "https://deribit.com",
		res: 1, // 1 minute
	}
}

func (c *Client) Candles(sta time.Time, end time.Time, ins string) ([]candle.Candle, error) {
	var err error

	// We remove 1 nanosecond from the end time so that the end of the time
	// range is exclusive. That way we fetch candles for minutes 0, 1, 2, 3 and
	// 4 when we provide sta/end 0/5, because the candle for the fifth minute
	// starts at minute 4 and ends at minute 5.
	{
		end = end.Add(-1)
	}

	var que url.Values
	{
		que = url.Values{}

		que.Set("instrument_name", ins)
		que.Set("resolution", fmt.Sprintf("%d", c.res))
		que.Set("start_timestamp", fmt.Sprintf("%d", sta.UnixNano()/int64(time.Millisecond)))
		que.Set("end_timestamp", fmt.Sprintf("%d", end.UnixNano()/int64(time.Millisecond)))
	}

	var uri string
	{
		uri = fmt.Sprintf(
			"%s/api/v2/public/get_tradingview_chart_data?%s",
			c.api,
			que.Encode(),
		)
	}

	var res *http.Response
	{
		res, err = http.Get(uri)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}

	var dat struct {
		Result struct {
			Ticks []int64   `json:"ticks"`
			Open  []float64 `json:"open"`
			High  []float64 `json:"high"`
			Low   []float64 `json:"low"`
			Close []float64 `json:"close"`
		} `json:"result"`
	}

	{
		err = json.NewDecoder(res.Body).Decode(&dat)
		if err != nil {
			return nil, err
		}
	}

	can := make([]candle.Candle, len(dat.Result.Ticks))
	for i, tick := range dat.Result.Ticks {
		can[i] = candle.Candle{
			Time:  time.UnixMilli(tick).UTC(),
			Open:  dat.Result.Open[i],
			High:  dat.Result.High[i],
			Low:   dat.Result.Low[i],
			Close: dat.Result.Close[i],
		}
	}

	return can, nil
}
