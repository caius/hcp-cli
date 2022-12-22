package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/client/packer_service"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
	"github.com/urfave/cli/v2"
)

/*
Styling after AWS cli tool for now

- List all buckets: list-buckets
- Show bucket details: get-bucket --bucket NAME
- List iteration for a specific bucket: list-iteration --bucket NAME
- List channels for a specific bucket: list-channels --bucket NAME
- Change channel interation: put-channel --bucket NAME --iteration ID

*/

func main() {
	app := &cli.App{
		Name:  "hcp-packer",
		Usage: "Hashicorp Cloud Platform Packer CLI",
		Commands: []*cli.Command{
			{
				Name:  "list-buckets",
				Usage: "List Image Buckets",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("zomg listing buckets")

					httpClient, err := httpclient.New(httpclient.Config{})
					if err != nil {
						log.Fatal(err)
					}

					ctx := context.Background()

					packerClient := packer_service.New(httpClient, nil)
					listBucketsParams := &packer_service.PackerServiceListBucketsParams{
						Context:                ctx,
						LocationOrganizationID: os.Getenv("HCP_ORGANISATION_ID"),
						LocationProjectID:      os.Getenv("HCP_PROJECT_ID"),
					}
					resp, err := packerClient.PackerServiceListBuckets(listBucketsParams, nil)
					if err != nil {
						log.Fatal(err)
					}

					for _, b := range resp.Payload.Buckets {
						fmt.Printf("%s\t%s\n", b.Slug, b.ID)
					}

					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}
