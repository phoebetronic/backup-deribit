package apicliaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xh3b4sd/tracer"
)

func (a *AWS) Lister(buc string, pre string) ([]string, error) {
	var lis []string
	{
		pag := s3.NewListObjectsV2Paginator(a.S3, &s3.ListObjectsV2Input{
			Bucket: aws.String(buc),
			Prefix: aws.String(pre),
		})

		for pag.HasMorePages() {
			out, err := pag.NextPage(context.Background())
			if err != nil {
				return nil, tracer.Mask(err)
			}

			for _, c := range out.Contents {
				lis = append(lis, *c.Key)
			}
		}
	}

	return lis, nil
}
