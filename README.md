# Chromium Downloader

A Go script that downloads the latest ungoogled Chromium mini installer from the [chromium.woolyss.com](https://chromium.woolyss.com/) website.

## Features

1. Automatically finds the download link for the latest ungoogled Chromium mini installer.
2. Downloads the file with a progress bar.
3. Allows the user to delete the downloaded file after it has been saved.

## Usage

1. Ensure you have Go installed on your system.
2. Clone the repository or download the `chromium.go` file.
3. Open a terminal/command prompt and navigate to the directory containing the `chromium.go` file.
4. Run the script using the following command:
go run chromium.go

5. The script will display the filename of the file it found and ask if you want to download it. Enter "y" or "yes" to confirm.
6. The file will be downloaded with a progress bar.
7. After the download is complete, the script will ask if you want to delete the downloaded file.

## Dependencies

This script uses the following Go packages:

- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - for parsing HTML content
- [github.com/schollz/progressbar/v3](https://github.com/schollz/progressbar) - for displaying a download progress bar

You can install these dependencies by running the following command:
go get github.com/PuerkitoBio/goquery github.com/schollz/progressbar/v3

## Limitations

- The script does not execute the downloaded file. It only downloads and saves the file to the current directory.
- The script does not perform any safety checks on the downloaded file before saving it.

## License

This project is licensed under the [MIT License](LICENSE).