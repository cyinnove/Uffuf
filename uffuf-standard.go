package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const defaultWordlist = "sallam/Wordlists/elisa_fuzz.txt"
const defaultOutputFile = "output.txt"

// Function to print usage
func printUsage() {
	fmt.Printf("Usage: %s -w <wordlist> -l <list_of_urls> -o <output_file>\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Println("  -w <wordlist>: Path to the wordlist file (default: sallam/Wordlists/elisa_fuzz.txt)")
	fmt.Println("  -l <list_of_urls>: File containing list of URLs")
	fmt.Println("  -o <output_file>: File to write unique results (default: output.txt)")
	os.Exit(1)
}

// Function to execute ffuf and capture unique results
func makeFFUFResultsUnique(urlList, wordlist, outputFile string) {
	file, err := os.Open(urlList)
	if err != nil {
		fmt.Println("Error opening URL list file:", err)
		os.Exit(1)
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outFile.Close()

	seenSizes := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := runFFUF(line, wordlist, "GET", seenSizes, outFile); err != nil {
			fmt.Println("Error running ffuf for GET:", err)
			os.Exit(1)
		}
		if err := runFFUF(line, wordlist, "POST", seenSizes, outFile); err != nil {
			fmt.Println("Error running ffuf for POST:", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading URL list file:", err)
		os.Exit(1)
	}
}

// Function to run ffuf with specified method and capture unique results
func runFFUF(url, wordlist, method string, seenSizes map[string]bool, outFile *os.File) error {
	cmd := exec.Command("ffuf", "-u", fmt.Sprintf("%s/FUZZ", url), "-w", wordlist, "-X", method)
	fmt.Printf("[DEBUG] Running command: ffuf -u %s/FUZZ -w %s -X %s\n", url, wordlist, method) // Debug output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	re := regexp.MustCompile(`Words: \d+`)
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if match := re.FindString(line); match != "" {
			size := strings.Split(match, " ")[1]
			if _, seen := seenSizes[size]; !seen {
				seenSizes[size] = true
				outFile.WriteString(line + "\n")
				fmt.Printf("[DEBUG] Found unique result: %s\n", line) // Debug output
			}
		}
	}

	return scanner.Err()
}

func main() {
	var wordlist, urlList, outputFile string

	flag.StringVar(&wordlist, "w", defaultWordlist, "Path to the wordlist file")
	flag.StringVar(&urlList, "l", "", "File containing list of URLs")
	flag.StringVar(&outputFile, "o", defaultOutputFile, "File to write unique results")
	flag.Parse()

	if urlList == "" {
		fmt.Println("Error: -l <list_of_urls> is required.")
		printUsage()
	}

	makeFFUFResultsUnique(urlList, wordlist, outputFile)
}
