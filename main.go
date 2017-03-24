package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type Store struct {
	storedTimes int64
	storedRegex string 
}

const FEATURE_1 = "avgtimes"
const FEATURE_2 = "requests"

func main() {
	// check # of inputted arguments and which argument
	if len(os.Args) < 2 || os.Args[1] != FEATURE_1 && os.Args[1] != FEATURE_2 {
		panic("please run with arguments `avgtimes <filename>` or `requests <filename>`")
	}


	data := make([]map[string]string, 0)
	// open the file
	if f, err := os.Open(os.Args[2]); err != nil {
		panic(err.Error())
	} else {
	// decode file as JSON into the data variable
		if err := json.NewDecoder(f).Decode(&data); err != nil {
			panic(err.Error())
		}
	}

	counts := make(map[string][]Store, 0)

	for _, k := range data {

		var tmpStore Store

		if os.Args[1] == FEATURE_1 {
			if strings.Contains(k["timing"], "ms") {
			timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "ms"), 10, 64)
			tmpStore.storedTimes = timing
			} else if strings.Contains(k["timing"], "s"){
				timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "s"), 10, 64)
				tmpStore.storedTimes = timing * 1e3
			}
		} else if os.Args[1] == FEATURE_2 {
			tmpStore.storedRegex = k["regex"]
		}

		counts[k["category"]] = append(counts[k["category"]], tmpStore)

	}

	if os.Args[1] == FEATURE_1 {
		// times
		for cat, stores := range counts {
			accumTimes := 0.0

			for _, t := range stores {
				accumTimes += float64(t.storedTimes)
			}

			fmt.Printf("category: %s\t%fms\n", cat, accumTimes/float64(len(stores)))
		}

	} else if os.Args[1] == FEATURE_2 {
		// regex

		// iterate over the main category map
		for mainCategory, stores := range counts {

			accumRegEx := ""

			// accumulate all regex into a string
			for _, t := range stores {
				accumRegEx += t.storedRegex + " "
			}

			mapOfRegEx := make (map[string]int, 0) 

			// iterate over the main category again to compare the regex with
			for subCategory, _ := range counts {

				// find which category the regex belongs and how many there are
				regExSubStr := regexp.MustCompile(subCategory)
				numOfMatchesForSubStr := regExSubStr.FindAllStringIndex(accumRegEx, -1)

				if mainCategory != subCategory &&  len(numOfMatchesForSubStr) > 0{
					// make sure not comparing the same category
					mapOfRegEx[subCategory] = len(numOfMatchesForSubStr)
				}

			}

			// have a map of how many regex there are in each main category
			// and ofc the name of the main category, which is just the mainCategory var in the current iteration
			// fmt.Println(mainCategory, mapOfRegEx)

			// now iterate over each regex, find the main category
			// get the main category's regex 
			for eachRegExToCompare, numOfRegEx := range mapOfRegEx {


				for categoryToCompare, subStores := range counts {

					// check if the regex we are looking is the same the category
					if eachRegExToCompare == categoryToCompare {

						accumSubRegEx := ""

						// accumulate all regex into a string
						for _, t := range subStores {
							accumSubRegEx += t.storedRegex + " "
						}

						regExSubStr2 := regexp.MustCompile (mainCategory)
						numOfMatchesForSubStr2 := regExSubStr2.FindAllStringIndex (accumSubRegEx, -1)

						if len (numOfMatchesForSubStr2) > 0 {
							fmt.Println(mainCategory, "has", numOfRegEx, " ", eachRegExToCompare)
							fmt.Println(eachRegExToCompare, "has", len(numOfMatchesForSubStr2), " ", mainCategory)
							fmt.Println()
						}					
					}
				}
			}
		}
	}
}







