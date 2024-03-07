package main

import (
	"fmt"
	"os"
	"log"
	"net"
	"net/http"
	"time"
	"io"
	"context"
)

func main() {
	const name string = "goraga"
	log.SetPrefix(name + ":  ")

	if len(os.Args) != 2 {
		log.Fatal("no url specified")
	}

	hostname := os.Args[1]
	ips, err := net.LookupIP(hostname)

	if err != nil {
		log.Fatalf("lookup is: %s: %v", hostname, err)
	}
	if len(ips) == 0 {
		log.Fatalf("no ips found for %s", hostname)
	}

	c := http.Client{Timeout: time.Second*5}
	
	url := "http://" + hostname
	if err := downloadAndSave(context.TODO(), &c, url); err != nil {
		log.Fatal(err)
	}
}

func downloadAndSave(ctx context.Context, c *http.Client, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: GET %q: %v", url, err)
	}
	resp, err := c.Do(req)

	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status: %s", resp.Status)
	}
	
	dstPath := "./index.html"
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("creating file %v", err)
	}
	defer dstFile.Close()

	io.Copy(dstFile, resp.Body)

	return nil
}