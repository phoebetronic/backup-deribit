package apicliaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xh3b4sd/tracer"
)

func (a *AWS) Download(buc string, key string) ([]byte, error) {
	var log bool

	var siz int64
	{
		inp := &s3.HeadObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
		}

		out, err := a.S3.HeadObject(context.Background(), inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		siz = out.ContentLength
	}

	if log {
		fmt.Printf("fetching %s\n", a.siz(siz))
	}

	var wri *Writer
	{
		wri = &Writer{
			log: log,
			siz: siz,
			wri: manager.NewWriteAtBuffer([]byte{}),
		}
	}

	{
		inp := &s3.GetObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
		}

		_, err := manager.NewDownloader(a.S3).Download(context.Background(), wri, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if log {
		fmt.Printf("\n")
	}

	return wri.wri.Bytes(), nil
}
