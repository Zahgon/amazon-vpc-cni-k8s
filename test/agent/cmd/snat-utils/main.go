package main

import (
	"flag"
	"log"
)

func main() {
	var testIPTableRules bool
	var testExternalDomainConnectivity bool
	var randomizedSNATValue string
	var numOfCidrs int
	var url string

	flag.BoolVar(&testIPTableRules, "testIPTableRules", false, "bool flag when set to true tests validate if IPTable has required rules")
	flag.StringVar(&randomizedSNATValue, "randomizedSNATValue", "prng", "value for AWS_VPC_K8S_CNI_RANDOMIZESNAT")
	flag.IntVar(&numOfCidrs, "numOfCidrs", 1, "Number of CIDR blocks in customer VPC")
	flag.BoolVar(&testExternalDomainConnectivity, "testExternalDomainConnectivity", false, "bool flag when set to true tests if the pod has internet access")
	flag.StringVar(&url, "url", "https://aws.amazon.com/", "url to check for connectivity")

	flag.Parse()

	if testIPTableRules {
		err := validateIPTableRules(randomizedSNATValue, numOfCidrs)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Randomized SNAT test passed for AWS_VPC_K8S_CNI_RANDOMIZESNAT: %s\n", randomizedSNATValue)
	}

	if testExternalDomainConnectivity {
		err := validateExternalDomainConnectivity(url)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("External Domain Connectivity test passed")
	}
}

func validateExternalDomainConnectivity(url string) error { _ = "STUB: not implemented"; return nil }

func validateIPTableRules(randomizedSNATValue string, numOfCidrs int) error {
	_ = "STUB: not implemented"
	// Check IPTable rules corresponding to AWS_VPC_K8S_CNI_RANDOMIZESNAT
	return nil
}

// If AWS-SNAT-CHAIN-1 exists, we run the old logic

// One rule per cidr + SNAT rule + chain creation rule

// Fetch rules from lastChain

// Check for rule with following pattern
