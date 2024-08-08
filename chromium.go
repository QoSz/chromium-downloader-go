package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/schollz/progressbar/v3"
)

const url = "https://chromium.woolyss.com/"

func main() {
	// Fetch the download link
	link, filename, err := getDownloadLink()
	if err != nil {
		fmt.Println("Error getting download link:", err)
		return
	}

	// Display file information and ask for user confirmation
	fmt.Println("File Name:", filename)
	if !confirmDownload() {
		fmt.Println("Download canceled.")
		return
	}

	// Download the file
	if err := downloadFile(link, filename); err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	fmt.Printf("\nDownloaded: %s\n", filename)

	// Ask to delete the file
	if confirmDelete(filename) {
		if err := os.Remove(filename); err != nil {
			fmt.Println("Error deleting file:", err)
		} else {
			fmt.Printf("File '%s' has been deleted.\n", filename)
		}
	}
}

func getDownloadLink() (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}

	re := regexp.MustCompile(`_ungoogled_mini_installer\.exe`)
	var link string
	doc.Find("a[title]").Each(func(_ int, s *goquery.Selection) {
		if title, exists := s.Attr("title"); exists && re.MatchString(title) {
			link, _ = s.Attr("href")
		}
	})

	if link == "" {
		return "", "", fmt.Errorf("download link not found")
	}

	filename := filepath.Base(link)
	return link, filename, nil
}

func confirmDownload() bool {
	var choice string
	fmt.Print("Do you want to download this file? (y/n): ")
	fmt.Scanln(&choice)
	return strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes"
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading",
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	return err
}

func confirmDelete(filename string) bool {
	var choice string
	fmt.Printf("Do you want to delete the file '%s'? (y/n): ", filename)
	fmt.Scanln(&choice)
	return strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes"
}
