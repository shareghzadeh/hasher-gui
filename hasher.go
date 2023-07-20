package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"html"
	"net/url"
	"regexp"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input2 := widget.NewEntry()
	result := widget.NewEntry()

	// Labels

	inputLabel := widget.NewLabel("Raw Text: ")

	options := []string{"MD5", "HTML", "URL", "BASE32", "BASE64", "SHA1", "SHA256", "SHA512"}

	myDropDown := widget.NewSelect(options, func(selected string) {})
	encode := widget.NewButton("Encode/Copy", func() {
		if myDropDown.SelectedIndex() == 0 {
			result.Text = md5ToString(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 1 {
			result.Text = html.EscapeString(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 2 {
			result.Text = url.QueryEscape(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 3 {
			result.Text = base32.StdEncoding.EncodeToString([]byte(input.Text))
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 3 {
			result.Text = base64.StdEncoding.EncodeToString([]byte(input.Text))
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 5 {
			result.Text = Sha1ToString(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 6 {
			result.Text = Sha256ToString(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else if myDropDown.SelectedIndex() == 7 {
			result.Text = Sha512ToString(input.Text)
			result.Refresh()
			if result.Text != "Please Select Something" {
				myWindow.Clipboard().SetContent(result.Text)
			}
		} else {
			result.Text = "Please Select Something"
			result.Refresh()
		}

	})

	decode := widget.NewButton("Decode/Copy", func() {
		if isMd5(input.Text) && md5ToString(input2.Text) == input.Text {
			if result.Text != "NOT FOUND" {
				myWindow.Clipboard().SetContent(result.Text)
			}
			result.Text = input2.Text
			result.Refresh()

		} else if isMd5(input2.Text) && md5ToString(input.Text) == input2.Text {
			if result.Text != "NOT FOUND" {
				myWindow.Clipboard().SetContent(result.Text)
			}
			result.Text = input.Text
			result.Refresh()
		} else if isSha1(input2.Text) && Sha1ToString(input.Text) == input2.Text {
			if result.Text != "NOT FOUND" {
				myWindow.Clipboard().SetContent(result.Text)
			}
			result.Text = input.Text
			result.Refresh()
		} else if isSha256(input2.Text) && Sha256ToString(input.Text) == input2.Text {
			if result.Text != "NOT FOUND" {
				myWindow.Clipboard().SetContent(result.Text)
			}
			result.Text = input.Text
			result.Refresh()
		} else if isSha512(input2.Text) && Sha512ToString(input.Text) == input2.Text {
			if result.Text != "NOT FOUND" {
				myWindow.Clipboard().SetContent(result.Text)
			}
			result.Text = input.Text
			result.Refresh()
		} else if input.Text != "" || input2.Text != "" {
			if isHtmlEncode(input.Text) {
				result.Text = html.UnescapeString(input.Text)
				result.Refresh()
			} else if isHtmlEncode(input2.Text) {
				result.Text = html.UnescapeString(input2.Text)
				result.Refresh()
			} else if isUrlEncode(input.Text) {
				result.Text, _ = url.QueryUnescape(input.Text)
				result.Refresh()
			} else if isUrlEncode(input2.Text) {
				result.Text, _ = url.QueryUnescape(input2.Text)
				result.Refresh()
			} else {
				result.Text = "NOT FOUND"
				result.Refresh()
			}
		} else {
			result.Text = "NOT FOUND"
			result.Refresh()
		}
	})

	input.SetPlaceHolder("Enter Your Text")
	input2.SetPlaceHolder("Enter Your Hash")

	content := container.NewVBox(inputLabel, input, input2, myDropDown, encode, decode, result)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func isMd5(s string) bool {
	// Compile a regular expression to match a 32-character hexadecimal string
	pattern := "^[a-f0-9]{32}$"
	re := regexp.MustCompile(pattern)

	// Check if the input matches the pattern
	if !re.MatchString(s) {
		return false
	}
	return true
}

func isHtmlEncode(s string) bool {
	// Test HTML-encoded pattern
	htmlPattern := "&[a-zA-Z]+;|&#\\d+;"
	htmlRegex := regexp.MustCompile(htmlPattern)
	// Check if the input matches the pattern
	if !htmlRegex.MatchString(s) {
		return false
	}
	return true

}
func isUrlEncode(s string) bool {
	// Test URL-encoded pattern
	urlPattern := "%[0-9a-fA-F]{2}"
	urlRegex := regexp.MustCompile(urlPattern)
	if !urlRegex.MatchString(s) {
		return false
	}
	return true
}
func isSha1(s string) bool {
	// Test URL-encoded pattern
	sha1Pattern := "^[a-fA-F0-9]{40}$"
	sha1Regex := regexp.MustCompile(sha1Pattern)
	if !sha1Regex.MatchString(s) {
		return false
	}
	return true
}
func isSha256(s string) bool {
	// Test URL-encoded pattern
	sha256Pattern := "^[a-fA-F0-9]{64}$"
	sha256Regex := regexp.MustCompile(sha256Pattern)
	if !sha256Regex.MatchString(s) {
		return false
	}
	return true
}
func isSha512(s string) bool {
	// Test URL-encoded pattern
	sha512Pattern := "^[0-9a-fA-F]{128}$"
	sha512Regex := regexp.MustCompile(sha512Pattern)
	if !sha512Regex.MatchString(s) {
		return false
	}
	return true
}
func md5ToString(a string) string {
	// Text to byte
	hash := md5.Sum([]byte(a))
	// md5 byte to string
	hashToString := fmt.Sprintf("%x", hash)
	// return the md5 string
	return hashToString
}
func Sha1ToString(a string) string {
	hash := sha1.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func Sha256ToString(a string) string {
	hash := sha256.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func Sha512ToString(a string) string {
	hash := sha512.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}
