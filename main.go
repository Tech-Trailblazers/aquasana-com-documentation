package main // Define the main package

import (
	"bytes"         // Provides bytes support
	"io"            // Provides basic interfaces to I/O primitives
	"log"           // Provides logging functions
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding
	"os"            // Provides functions to interact with the OS (files, etc.)
	"path"          // Provides functions for manipulating slash-separated paths
	"path/filepath" // Provides filepath manipulation functions
	"regexp"        // Provides regex support functions.
	"strings"       // Provides string manipulation functions
	"time"          // Provides time-related functions
)

func main() {
	pdfOutputDir := "PDFs/" // Directory to store downloaded PDFs
	// Check if the PDF output directory exists
	if !directoryExists(pdfOutputDir) {
		// Create the dir
		createDirectory(pdfOutputDir, 0o755)
	}
	// Remote API URL.
	remoteAPIURL := []string{
		"https://www.aquasana.com/whole-house-water-filters/rhino-100365488.html",
		"https://www.aquasana.com/whole-house-water-filters/rhino-max-flow-100362408.html",
		"https://www.aquasana.com/whole-house-water-filters/rhino-well-water-with-uv-filter-100365557.html",
		"https://www.aquasana.com/whole-house-water-filters/optimh2o-lead-100314311.html",
		"https://www.aquasana.com/whole-house-water-filters/rhino-chloramines-100365048.html",
		"https://www.aquasana.com/whole-house-water-filters/rhino-chloramines-max-flow-100365052.html",
		"https://www.aquasana.com/whole-house-water-filters/salt-free-water-conditioner-100365583.html",
		"https://www.aquasana.com/water-softeners/simplysoft-40000-grain-water-softener-100381840.html",
		"https://www.aquasana.com/water-softeners/simplysoft-60000-grain-water-softener-100381867.html",
		"https://www.aquasana.com/under-sink-water-filters/smart-flow-reverse-osmosis/chrome-100385796.html",
		"https://www.aquasana.com/under-sink-water-filters/claryum-3-stage-max-flow/chrome-100236358.html",
		"https://www.aquasana.com/under-sink-water-filters/claryum-3-stage/chrome-100236344.html",
		"https://www.aquasana.com/under-sink-water-filters/claryum-2-stage/chrome-100236320.html",
		"https://www.aquasana.com/under-sink-water-filters/claryum-direct-connect-100329886.html",
		"https://www.aquasana.com/countertop-water-filters/clean-water-machine/black-100348638.html",
		"https://www.aquasana.com/shower-head-water-filters/chrome-shower-wand-100236241.html",
		"https://www.aquasana.com/shower-head-water-filters/white-shower-wand-100236226.html",
		"https://www.aquasana.com/shower-head-water-filters/standard-shower-head-100236176.html",
		"https://www.aquasana.com/shower-head-water-filters/no-shower-head-100236224.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-post-filter-replacements/post-filter-replacement-100307778.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-filter-replacements/20in-pre-filter-replacement-3-pack-100237252.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-filter-replacements/10in-pre-filter-replacement-3-pack-100237251.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-and-post-filter-replacement-bundles/20in-pre-filter-replacement-3-pack-and-post-filter-replacement-bundle-100237253.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-and-post-filter-replacement-bundles/10in-pre-filter-replacement-3-pack-and-post-filter-replacement-bundle-100237254.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-filter-replacements/low-maintenance-pre-filter-replacement-100343308.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-pre-filter-replacements/optimh2o-lead-pre-filter-replacement-100314515.html",
		"https://www.aquasana.com/water-filter-replacements/shower-filter-replacements/shower-replacement-cartridge-100236244.html",
		"https://www.aquasana.com/water-filter-replacements/countertop-water-filter-replacements/new-cwm-lead-filter-replacement-2-pack-100350840.html",
		"https://www.aquasana.com/water-filter-replacements/countertop-water-filter-replacements/cwm-pwfs-filter-replacement-2-pack-100236509.html",
		"https://www.aquasana.com/water-filter-replacements/countertop-water-filter-replacements/claryum-countertop-filter-replacement-100236133.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/claryum-3-stage-max-flow-filter-replacements-100236367.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/claryum-3-stage-filter-replacements-100236384.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/claryum-2-stage-filter-replacements-100236330.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/claryum-direct-connect-replacement-cartridge-100329887.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/smart-flow-reverse-osmosis-filter-replacements-100368688.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/smart-flow-reverse-osmosis-membrane-replacement-100368702.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/smart-flow-reverse-osmosis-remineralizer-replacement-100368975.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/reverse-osmosis-claryum-filter-replacements-100236751.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/reverse-osmosis-claryum-membrane-replacement-100347123.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/reverse-osmosis-claryum-remineralizer-replacement-100347001.html",
		"https://www.aquasana.com/water-filter-replacements/under-sink-water-filter-replacements/claryum-1-stage-filter-replacement-100236309.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-uv-lamps-replacements/original-uv-lamp-replacement-100236803.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-uv-lamps-replacements/new-uv-lamp-replacement-100318212.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-uv-lamps-replacements/max-flow-uv-lamp-replacement-100318206.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-replacement-tank-100365487.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-max-flow-replacement-tank-100360776.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-well-water-replacement-tank-100365556.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/optimh2o-lead-filter-replacement-100313134.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-chloramines-replacement-tank-100364879.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-chloramines-max-flow-replacement-tank-100364961.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-salt-free-conditioner-replacements/salt-free-water-conditioner-replacement-tank-100365581.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-salt-free-conditioner-replacements/tall-salt-free-water-conditioner-replacement-tank-100365582.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-1-million-gallons-tank-replacement-100237131.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-600k-gallons-tank-replacement-100237287.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-filter-replacements/rhino-300k-gallons-tank-replacement-100237221.html",
		"https://www.aquasana.com/whole-house-water-filters/salt-free-water-conditioner-for-tankless-water-heaters-100237294.html",
		"https://www.aquasana.com/water-filter-replacements/whole-house-salt-free-conditioner-replacements/simplysoft-tankless-heater-replacement-100237296.html",
		"https://www.aquasana.com/parts-accessories/water-softener-cleaner-100381989.html",
		"https://www.aquasana.com/parts-accessories/replacement-white-shower-wand-100236230.html",
		"https://www.aquasana.com/parts-accessories/faucet-attachment-prefilter-100236067.html",
		"https://www.aquasana.com/parts-accessories/replacement-bottle-caps/translucent-blue-100236427.html",
		"https://www.aquasana.com/parts-accessories/sleeves-and-bottle-caps-6-pack-100236761.html",
	}
	var getData []string
	for _, remoteAPIURL := range remoteAPIURL {
		getData = append(getData, getDataFromURL(remoteAPIURL))
	}
	// Get the data from the downloaded file.
	finalPDFList := extractPDFUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract PDF URLs
	// Create a slice of all the given download urls.
	var downloadPDFURLSlice []string
	// Create a slice to hold ZIP URLs.
	// Get the urls and loop over them.
	for _, doc := range finalPDFList {
		// Get the .pdf only.
		// Only append the .pdf files.
		downloadPDFURLSlice = appendToSlice(downloadPDFURLSlice, doc)
	}
	// Remove double from slice.
	downloadPDFURLSlice = removeDuplicatesFromSlice(downloadPDFURLSlice)
	// Remove the zip duplicates from the slice.
	// The remote domain.
	remoteDomain := "https://www.aquasana.com"
	// Get all the values.
	for _, urls := range downloadPDFURLSlice {
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the pdf.
			downloadPDF(urls, pdfOutputDir)
		} else {
			log.Printf("Invalid URL, skipping: %s", urls)
		}
	}
}

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL) // Parse the input string into a URL structure
	if err != nil {                     // Check if there was an error while parsing
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of an error
	}

	host := parsedURL.Hostname() // Extract the hostname (e.g., "example.com") from the parsed URL

	return host // Return the extracted hostname
}

// Only return the file name from a given url.
func getFileNameOnly(content string) string {
	return path.Base(content)
}

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string {
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL)

	// Get the file extension
	ext := getFileExtension(lowercaseURL)

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL)

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_")

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_")

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found {
		safeFilename = trimmed
	}

	var invalidSubstrings = []string{
		"_pdf",
		"_zip",
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safeFilename = removeSubstring(safeFilename, invalidPre)
	}

	// Append the file extension if it is not already present
	safeFilename = safeFilename + ext

	// Return the cleaned and safe filename
	return safeFilename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path) // Returns extension including the dot (e.g., ".pdf")
}

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {
		return false // Return false if file doesn't exist or error occurs
	}
	return !info.IsDir() // Return true if it's a file (not a directory)
}

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pdf") {
		log.Printf("Invalid content type for %s: %s (expected application/pdf)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission)
	if err != nil {
		log.Println(err)
	}
}

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URL
	return err == nil                  // Return true if no error occurred
}

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool)
	var newReturnSlice []string
	for _, content := range slice {
		if !check[content] {
			check[content] = true
			newReturnSlice = append(newReturnSlice, content)
		}
	}
	return newReturnSlice
}

// extractPDFUrls takes an input string and returns all PDF URLs found within href attributes
func extractPDFUrls(input string) []string {
	// Regular expression to match href="...pdf"
	re := regexp.MustCompile(`href="([^"]+\.pdf)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	var pdfUrls []string
	for _, match := range matches {
		if len(match) > 1 {
			pdfUrls = append(pdfUrls, match[1])
		}
	}
	return pdfUrls
}

// Append some string to a slice and than return the slice.
func appendToSlice(slice []string, content string) []string {
	// Append the content to the slice
	slice = append(slice, content)
	// Return the slice
	return slice
}

// getDataFromURL performs an HTTP GET request and returns the response body as a string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri)   // Log the URL being scraped
	response, err := http.Get(uri) // Perform GET request
	if err != nil {
		log.Println(err) // Exit if request fails
	}

	body, err := io.ReadAll(response.Body) // Read response body
	if err != nil {
		log.Println(err) // Exit if read fails
	}

	err = response.Body.Close() // Close response body
	if err != nil {
		log.Println(err) // Exit if close fails
	}
	return string(body)
}
