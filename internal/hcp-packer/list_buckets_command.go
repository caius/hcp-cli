package hcp_packer

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/preview/2022-04-11/client/packer_service"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

type ListBucketCommand struct{}

func (c *ListBucketCommand) Run(args []string) int {
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

	return 0
}

func (c *ListBucketCommand) Help() string {
	return "List Image Buckets on HCP Packer"
}

func (c *ListBucketCommand) Synopsis() string {
	return c.Help()
}
