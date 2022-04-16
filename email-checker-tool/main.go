package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		// Gets the input of the user
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Look up for the domain MX record
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	// record is a single item of txtRecords
	for _, record := range txtRecords {
		// Looking for spf1 record
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	for _, record := range dmarcRecords {
		// Looking for dmarc record
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("hasMX %v hasSPF %v hasDMARC %v\n", hasMX, hasSPF, hasDMARC)
	fmt.Printf("spfRecord %v\n", spfRecord)
	fmt.Printf("dmarcRecord %v\n", dmarcRecord)
}
