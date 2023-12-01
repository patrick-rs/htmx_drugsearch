package data

import (
	"encoding/json"
	"os"
	"strings"
)

var drugs []string

func Setup() {
	f, err := os.ReadFile("drug_names.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(f, &drugs)

	if err != nil {
		panic(err)
	}
}

func SearchDrugs(s string) []string {
	matchingDrugs := []string{}

	for _, drug := range drugs {
		if strings.HasPrefix(strings.ToLower(drug), strings.ToLower(s)) {
			matchingDrugs = append(matchingDrugs, drug)
		}

		if len(matchingDrugs) == 10 {
			break
		}
	}

	return matchingDrugs
}
