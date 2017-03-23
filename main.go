package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Store struct {
	storedTimes int64
	storedRegex string 
}

func main() {
	// check # of inputted arguments and which argument
	if len(os.Args) < 2 || os.Args[1] != "avgtimes" {
		panic("please run with arguments `avgtimes <filename>`")
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

		if strings.Contains(k["timing"], "ms") {
			timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "ms"), 10, 64)
			// counts[k["category"]] = append(counts[k["category"]], int(timing))
			tmpStore.storedTimes = timing
		} else if strings.Contains(k["timing"], "s"){
			timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "s"), 10, 64)
			// counts[k["category"]] = append(counts[k["category"]], int(timing) * 1e3)
			tmpStore.storedTimes = timing * 1e3
		}

		tmpStore.storedRegex = k["regex"]
		counts[k["category"]] = append(counts[k["category"]], tmpStore)
		
		// if err != nil {
		// 	panic(err.Error())
		// }

		// counts[k["category"]] = append(counts[k["category"]], int(timing))
		// fmt.Println(k);
	}
	
	// fmt.Printf(counts)

	for cat, stores := range counts {
		accumTimes := 0.0
		accumRegEx := ""

		for _, t := range stores {
			// s += float64(t)
			accumTimes += float64(t.storedTimes)
			accumRegEx += t.storedRegex + " "
			// fmt.Println(cat,"--", t)
		}
		fmt.Printf("category: %s\tavg %fms\t%s \n", cat, accumTimes/float64(len(stores)), accumRegEx)
	}
}







