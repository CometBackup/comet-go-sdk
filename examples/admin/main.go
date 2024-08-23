package main

/*
	EXAMPLE ADMIN CODE

	This quick commandline tool allows you to list and download any available CometBackup Clients
	hosted on the Comet Server you're sending requests to.
	It also demonstrates how to upload a resource file using AdminMetaResourceNew.
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
	if _, errIn := client.AdminAccountSessionStart(nil); errIn != nil {
		return retryWithTOTP(client, errIn, url)
	}

	return &Client{
		client: client,
	}, nil
}

func retryWithTOTP(c *sdk.CometAPIClient, errIn error, serverURL string) (client *Client, err error) {
	if strings.Contains(errIn.Error(), "449") {
		fmt.Printf("re-connecting to server: %s\n", serverURL)
		totp, err := util.Totp()
		if err != nil {
			log.Fatal("Error reading TOTP: ", err)
		}
		client.client.TOTPKey = totp
		// ReuseSessionKey is especially useful when using TOTP based authentication, otherwise you need to enter a
		// new TOTPKey for every api call.
		client.client.ReuseSessionKey = true
		fmt.Println("Connection successful.")
		return &Client{c}, nil
	}
	return nil, fmt.Errorf("error starting admin account session: %w", errIn)
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

func (c *Client) UploadResource(filePath string) error {
	response, err := c.client.AdminMetaResourceNew(filePath)
	if err != nil {
		return err
	}
	fmt.Printf("Resource uploaded successfully. response: %s\n", response)
	return nil
}

func main() {
	url := flag.String("s", "http://localhost:8060", "The URL for the CometServer API")
	username := flag.String("u", "", "The username to authenticate with")
	password := flag.String("p", "", "The password to authenticate with")

	list := flag.Bool("list", false, "List platforms available for download")
	download := flag.Int("download", 0, "The id of the platform to download")
	upload := flag.String("upload", "README.md", "Path to the file to upload as a resource")
	flag.Parse()

	if (*list && *download != 0) || (!*list && *download == 0 && *upload == "") {
		log.Fatal("Error: A command must be chosen. Choose one of '--list', '--download #', or '--upload <file>'")
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

	if *upload != "" {
		err := client.UploadResource(*upload)
		if err != nil {
			log.Fatal("Error uploading resource: ", err)
		}
	}
}
