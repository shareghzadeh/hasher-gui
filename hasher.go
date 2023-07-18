package main

// import (
// 	"crypto/md5"
// 	"encoding/base32"
// 	"encoding/base64"
// 	"fmt"
// 	"net/url"

// 	// HTML Encode(Escape/UnEscape)
// 	"html"
// 	"os"
// )

// const (
// 	Red    = "\033[31m"
// 	Green  = "\033[32m"
// 	Yellow = "\033[33m"
// 	Blue   = "\033[34m"
// 	Purple = "\033[35m"
// 	Cyan   = "\033[36m"
// 	White  = "\033[37m"
// 	Reset  = "\033[0m"
// )

// func main() {

// 	// this if/else if will check if the passed arguments are less than 2 argument
// 	// and check if the passed arguments are more that 7 or equal to it
// 	if len(os.Args) <= 1 {
// 		fmt.Printf("%sUSAGE:\n\tHTML:\n\t\tEscape --> ./hasher html -e \"<script>alert('Hacked')</script>\"\n\t\tUnEscape --> ./hasher html -d \"&lt;script&gt;alert('Hacked')&lt;/script&gt;\"\n\tMD5:\n\t\tHash --> ./hasher md5 -h \"Hi\"\n\t\tDeHash --> ./hasher md5 -d \"Hi\" \"c1a5298f939e87e8f962a5edfc206918\"\n\tBASE64:\n\t\tEncode --> ./hasher base64 -e \"Hi\"\n\t\tDecode --> ./hasher base64 -d \"SGk=\"\n\tBASE32:\n\t\tEncode --> ./hasher base32 -e \"Hi\"\n\t\tDecode --> ./hasher base32 -d \"JBUQ====\"\n\tURL:\n\t\tEncode --> ./hasher url -e \"Hello, World\"\n\t\tDecode --> ./hasher url -d \"Hello%%2C+World\"\n", Green)
// 		os.Exit(1)

// 	} else if len(os.Args) >= 7 {
// 		fmt.Printf("%sTo much arguments\n", Red)
// 		os.Exit(1)
// 	}

// 	// Arguments
// 	arg1 := os.Args[1]
// 	arg2 := os.Args[2]
// 	arg3 := os.Args[3]

// 	// HTML Escape/UnEscape
// 	if arg1 == "html" {
// 		switch arg2 {
// 		// Encode(Escape) html special characters like: <>"'&
// 		case "e", "-e", "--encode", "encode", "escape", "--escape":
// 			fmt.Printf("%s%s\n", Green, html.EscapeString(arg3))
// 		case "d", "-d", "--decode", "decode", "unescape", "--unescape":
// 			fmt.Printf("%s\n", html.UnescapeString(arg3))
// 		default:
// 			fmt.Printf("%sNOTHING", Red)
// 		}

// 	} else if arg1 == "md5" {
// 		switch arg2 {
// 		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
// 			hash := md5.Sum([]byte(arg3))
// 			hashToString := fmt.Sprintf("%x\n", hash)
// 			fmt.Println(hashToString)
// 		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
// 			arg4 := os.Args[4]
// 			// This if statement checks if passed tow md5 are the same or not(first turn the text to md5)
// 			if md5ToString(arg3) == arg4 || arg3 == md5ToString(arg4) {
// 				fmt.Printf("%s%s --> %s\n", Green, arg3, arg4)
// 			} else {
// 				fmt.Printf("%sHash Not Found!\n", Red)
// 			}
// 		default:
// 			fmt.Printf("%sUSAGE:\n\t\tHash --> ./hasher md5 -e <YOUR_TEXT>\n\t\tDehash--> ./hasher md5 -d <YOUR_TEXT> <YOUR_MD5>\n", Red)
// 		}

// 	} else if arg1 == "base64" {
// 		switch arg2 {
// 		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
// 			encode := base64.StdEncoding.EncodeToString([]byte(arg3))
// 			fmt.Println(encode)
// 		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
// 			decode, err := base64.StdEncoding.DecodeString(arg3)
// 			if err != nil {
// 				fmt.Printf("%sNot a Valid base64", Red)
// 				os.Exit(1)
// 			}
// 			fmt.Println(string(decode))
// 		default:
// 			fmt.Println("USAGE:\n\t\tEncode --> ./hasher base64 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base64 -d <YOUR_BASE64>")
// 		}

// 	} else if arg1 == "base32" {
// 		switch arg2 {
// 		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
// 			encode := base32.StdEncoding.EncodeToString([]byte(arg3))
// 			fmt.Println(encode)
// 		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
// 			decode, err := base32.StdEncoding.DecodeString(arg3)
// 			if err != nil {
// 				fmt.Println("Not a Valid base32")
// 				return
// 			}
// 			fmt.Println(string(decode))
// 		default:
// 			fmt.Println("USAGE:\n\t\tEncode --> ./hasher base32 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base32 -d <YOUR_BASE32>")
// 		}

// 	} else if arg1 == "url" {
// 		switch arg2 {
// 		case "e", "-e", "--encode", "encode":
// 			encode := url.QueryEscape(arg3)
// 			fmt.Println(encode)
// 		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
// 			decode, err := url.QueryUnescape(arg3)
// 			if err != nil {
// 				fmt.Printf("%sNot a Valid URL Encode", Red)
// 				return
// 			}
// 			fmt.Println(string(decode))
// 		default:
// 			fmt.Println("USAGE:\n\t\tEncode --> ./hasher url -e <YOUR_TEXT>\n\t\tDecode --> ./hasher url -d <YOUR_ENCODED_TEXT>")
// 		}

// 	} else {
// 		fmt.Println("argument not satisfied")
// 		os.Exit(1)
// 	}
// }

// // This function will turn text to md5 and return string of it.
// // I wrote this function for compare Text vs Md5 to check if
// // its the same(decode)
// func md5ToString(a string) string {
// 	// Text to byte
// 	hash := md5.Sum([]byte(a))
// 	// md5 byte to string
// 	hashToString := fmt.Sprintf("%x", hash)
// 	// return the md5 string
// 	return hashToString
// }

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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

	options := []string{"MD5", "HTML", "URL"}

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
			}
		} else {
			result.Text = "NOT FOUND"
			result.Refresh()
		}
	})

	input.SetPlaceHolder("Enter Your Text OR Your Encode Here")
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
func md5ToString(a string) string {
	// Text to byte
	hash := md5.Sum([]byte(a))
	// md5 byte to string
	hashToString := fmt.Sprintf("%x", hash)
	// return the md5 string
	return hashToString
}
func sha1ToString(a string) string {
	hash := sha1.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func sha256ToString(a string) string {
	hash := sha256.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func sha512ToString(a string) string {
	hash := sha512.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}
