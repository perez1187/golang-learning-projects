//https://www.youtube.com/watch?v=9E4UEsWpYvM&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=6
//19:50

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
	fmt.Printf("domain, hasMX, hasSPF, sprRecord,hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input :%v\n", err)
	}
}

func checkDomain(domain string) {

	// we create variables
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// we use net package to lookup MX
	//if something is wrong this function will give us error
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	// we go over text records
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") { // we want to find spf1
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v,", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
