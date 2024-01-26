package main

/*
	EXAMPLE ADMIN CODE

	This quick commandline tool allows you to list and download any available CometBackup Clients
	hosted on the Comet Server you're sending requests to.
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	sdk "github.com/CometBackup/comet-go-sdk"
	"github.com/CometBackup/comet-go-sdk/examples/util"
)

type Client struct {
	client *sdk.CometAPIClient
}

func NewClient(url, username, password string) (*Client, error) {
	client, err := sdk.NewCometAPIClient(url, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) DownloadBrandedClient(platform int) ([]byte, error) {
	return c.client.AdminBrandingGenerateClientByPlatform(platform, nil)
}

func (c *Client) FindPlatform(platform int) (*sdk.AvailableDownload, error) {
	platforms, err := c.client.AdminBrandingAvailablePlatforms()
	if err != nil {
		return nil, err
	}

	p := platforms[platform]
	return &p, nil
}

func (c *Client) ListAndPrintPlatforms() error {
	platforms, err := c.client.AdminBrandingAvailablePlatforms()
	if err != nil {
		return err
	}

	indexes := []int{}
	for i := range platforms {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)

	fmt.Println("| INDEX | RECOMMEND | CATEGORY | DESCRIPTION")
	for _, i := range indexes {
		fmt.Printf("| %5d | %-9t | %-8s | %s\n", i, platforms[i].Recommended, platforms[i].Category, platforms[i].Description)
	}

	return nil
}

func main() {
	url := flag.String("s", "http://localhost:8060", "The URL for the CometServer API")
	username := flag.String("u", "", "The username to authenticate with")
	password := flag.String("p", "", "The password to authenticate with")

	list := flag.Bool("list", false, "List platforms available for download")
	download := flag.Int("download", 0, "The id of the platform to download")
	flag.Parse()

	if (*list && *download != 0) || (!*list && *download == 0) {
		log.Fatal("Error: A command must be chosen. Choose one of either '--list' or '--download #'")
	}
	if *username == "" || *password == "" {
		var err error
		*username, *password, err = util.Credentials()
		if err != nil {
			log.Fatal("Error getting username and password: ", err)
		}
	}
	client, err := NewClient(*url, *username, *password)
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	totp, err := util.Totp()
	if err != nil {
		log.Fatal("Error reading TOTP: ", err)
	}
	client.client.TOTPKey = totp

	if *list {
		err = client.ListAndPrintPlatforms()
		if err != nil {
			log.Fatal("Error listing branded platforms: ", err)
		}
	}

	if *download != 0 {
		data, err := client.DownloadBrandedClient(*download)
		if err != nil {
			log.Fatal("Error downloading branded client (platform ", *download, ")")
		}

		platform, err := client.FindPlatform(*download)
		if err != nil {
			log.Fatal("Failed to find details about platform ", *download)
		}

		var ft string
		switch platform.Category {
		case "Windows":
			if strings.Contains(platform.Description, "zip") {
				ft = "zip"
			} else {
				ft = "exe"
			}
		case "Linux":
			ft = "run"
		case "macOS":
			ft = "pkg"
		case "Synology":
			ft = "pat"
		}

		fileName := "CometBackupClient." + ft
		err = ioutil.WriteFile(fileName, data, 0644)
		if err != nil {
			log.Fatal("Failed to write download client: ", err)
		}

		fmt.Println("Congratulations! You now have a branded Comet client: ", fileName)
	}
}
