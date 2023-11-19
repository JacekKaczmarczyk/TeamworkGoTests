// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

type Domain struct {
	DomainName string
	Count      int
}

// interface to make testing possible without the need for file
type FileReader interface {
	Open(fileName string) (io.ReadCloser, error)
}

const csvFileName = "customers.csv"

// This functions counts unique email domains and returs sorted domain list
func CountAndSortEmailDomains(fileRader FileReader) ([]Domain, error) {
	csvInput, err := fileRader.Open(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("Error has occured while trying to open file: %v", err)
	}
	defer csvInput.Close()

	reader := csv.NewReader(csvInput)

	// Skipping header line
	reader.Read()

	domains := make(map[string]int)

	for {
		// Read csv file line by line until EOF or error
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("Error has occured while reading file: %v", err)
		}

		email := record[2]
		if isValidEmail(email) {
			domain := extractDomain(email)
			domains[domain]++
		}
	}

	domainList := convertDomainsMapToStruct(domains)

	// Sort domainList alphabetically
	sort.Slice(domainList, func(i, j int) bool {
		return domainList[i].DomainName < domainList[j].DomainName
	})

	return domainList, nil
}

// Converts domains map into []Domains
func convertDomainsMapToStruct(domains map[string]int) []Domain {
	var domainList []Domain

	for domain, count := range domains {
		domainList = append(domainList, Domain{
			DomainName: domain,
			Count:      count,
		})
	}
	return domainList
}

// Checks if provided strings is an email
func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[\w.-]+@([\w-]+\.)+[\w-]{2,4}$`)
	return regex.MatchString(email)
}

// Extracts domain form email
func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	return parts[1]
}
