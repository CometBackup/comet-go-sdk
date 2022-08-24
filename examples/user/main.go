package main

/*
	EXAMPLE USER CODE

	This quick commandline tool allows you get and set your own user profile
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	cometsdk "github.com/CometBackup/comet-go-sdk"
)

type Client struct {
	client *cometsdk.CometAPIClient
}

func NewClient(url, username, password string) (*Client, error) {
	client, err := cometsdk.NewCometAPIClient(url, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetUserProfile() (*cometsdk.GetProfileAndHashResponseMessage, error) {
	return c.client.UserWebGetUserProfileAndHash()
}

func (c *Client) SetUserProfile(config cometsdk.UserProfileConfig, hash string) error {
	_, err := c.client.UserWebSetProfileHash(config, hash)
	return err
}

func main() {
	url := flag.String("s", "http://localhost:8060", "The URL for the CometServer API")
	username := flag.String("u", "test", "The username to authenticate with")
	password := flag.String("p", "testtest", "The password to authenticate with")

	get := flag.Bool("get", false, "Get a user's profile")
	set := flag.Bool("set", false, "Set a user's profile")
	hash := flag.String("hash", "", "The hash of the profile meant to be edited (typically obtained with --get)")
	jstr := flag.String("json", "", "The raw JSON string with the user's configuration")
	file := flag.String("file", "", "The file path to read/write user's configuration in JSON to")
	flag.Parse()

	if (!*set && !*get) || (*get && *set) {
		log.Fatal("Error: A command must be chosen. Choose either '--get' or '--set'")
	}
	if *set && (*jstr == "" && *file == "") {
		log.Fatal("Error: '--set' requires '--json' or '--file' to be set as well")
	}

	client, err := NewClient(*url, *username, *password)
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}

	if *get {
		message, err := client.GetUserProfile()
		if err != nil {
			log.Fatal("Error calling UserWebGetUserProfileAndHash: ", err)
		}
		result, err := json.MarshalIndent(message, "", "  ")
		if err != nil {
			log.Fatal("Failed to marshal user profile from API response. Raw message:\n", message)
		}
		if *file != "" {
			ioutil.WriteFile(*file, result, 0644)
		} else {
			fmt.Println(string(result))
		}
	}

	if *set {
		bytes := []byte{}
		if *file != "" {
			bytes, err = ioutil.ReadFile(*file)
			if err != nil {
				log.Fatal("Error attempting to read file (--file ", *file, "): ", err)
			}
		} else {
			bytes = []byte(*jstr)
		}

		config := cometsdk.UserProfileConfig{}
		err := json.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatal("Error marshaling passed JSON to UserProfileConfig: ", err)
		}

		err = client.SetUserProfile(config, *hash)
		if err != nil {
			log.Fatal("Error calling UserWebSetProfileHash: ", err)
		}
	}
}
