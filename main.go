package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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


	counts := make(map[string][]int, 0)

	for _, k := range data {

		if strings.Contains(k["timing"], "ms") {
			timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "ms"), 10, 64)
			counts[k["category"]] = append(counts[k["category"]], int(timing))
		} else if strings.Contains(k["timing"], "s"){
			timing, _ := strconv.ParseInt(strings.TrimSuffix(k["timing"], "s"), 10, 64)
			counts[k["category"]] = append(counts[k["category"]], int(timing) * 1e3)
		}
		
		// if err != nil {
		// 	panic(err.Error())
		// }

		// counts[k["category"]] = append(counts[k["category"]], int(timing))
		// fmt.Println(k);
	}
	
	// fmt.Printf(counts)

	for cat, times := range counts {
		s := 0.0
		for _, t := range times {
			s += float64(t)
			// fmt.Println(cat,"--", t)
		}
		fmt.Printf("category: %s\tavg %fms\n", cat, s/float64(len(times)))
	}
}
