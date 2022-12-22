package main

import (
	// "context"
	"fmt"
	// "log"
	"os"

	"github.com/caius/hcp-cli/internal/hcp-packer"
	// "github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/client/packer_service"
	// "github.com/hashicorp/hcp-sdk-go/httpclient"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("hcp-packer", "0.0.1")
	c.Autocomplete = true
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list-buckets": func() (cli.Command, error) {
			return &hcp_packer.ListBucketCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &hcp_packer.VersionCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Println("Error executing cli: ", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}

// app := &cli.App{
// 	Name:  "hcp-packer",
// 	Usage: "Hashicorp Cloud Platform Packer CLI",
// 	Commands: []*cli.Command{
// 		{
// 			Name:  "list-buckets",
// 			Usage: "List Image Buckets",
// 			Action: func(cCtx *cli.Context) error {
// 				fmt.Println("zomg listing buckets")
//
// 				httpClient, err := httpclient.New(httpclient.Config{})
// 				if err != nil {
// 					log.Fatal(err)
// 				}
//
// 				ctx := context.Background()
//
// 				packerClient := packer_service.New(httpClient, nil)
// 				listBucketsParams := &packer_service.PackerServiceListBucketsParams{
// 					Context:                ctx,
// 					LocationOrganizationID: os.Getenv("HCP_ORGANISATION_ID"),
// 					LocationProjectID:      os.Getenv("HCP_PROJECT_ID"),
// 				}
// 				resp, err := packerClient.PackerServiceListBuckets(listBucketsParams, nil)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
//
// 				for _, b := range resp.Payload.Buckets {
// 					fmt.Printf("%s\t%s\n", b.Slug, b.ID)
// 				}
//
// 				return nil
// 			},
// 		},
//
// 		{
// 			Name: "list-iterations",
// 			Flags: []cli.Flag{
// 				&cli.StringFlag{
// 					Name:  "bucket",
// 					Usage: "Bucket ID",
// 				},
// 			},
// 			Action: func(cCtx *cli.Context) error {
// 				fmt.Println("zomg listing iterations for ", cCtx.String("bucket"))
// 				return nil
// 			},
// 		},
// 	},
// }
// app.Run(os.Args)
