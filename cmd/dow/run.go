package dow

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phoebetronic/backup-deribit/pkg/apicliaws"
	"github.com/spf13/cobra"
)

const (
	buc = "phoebetronic"
	pre = "eth-usd"
)

type run struct {
	flags *flags
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	var aws *apicliaws.AWS
	{
		aws = apicliaws.New()
	}

	var lis []string
	{
		lis, err = aws.Lister(buc, pre)
		if err != nil {
			panic(err)
		}
	}

	for i, x := range lis {
		var spl []string
		{
			spl = strings.Split(x, "/")
		}

		var prt string
		{
			prt = spl[len(spl)-1]
		}

		var pat string
		{
			pat = filepath.Join(r.flags.Bas, prt)
		}

		{
			fmt.Printf("downloading %4d/%4d into %s\n", i+1, len(lis), pat)
		}

		var byt []byte
		{
			byt, err = aws.Download(buc, x)
			if err != nil {
				panic(err)
			}
		}

		{
			err := os.WriteFile(pat, byt, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
