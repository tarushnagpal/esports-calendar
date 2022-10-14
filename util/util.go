package util

import (
	"strings"
	"log"
	"strconv"
)

func ArrayContains(arr []string, element string) bool {
	var result bool = false
    for _, x := range arr {
        if x == element {
            result = true
            break
        }
    }
	return result
}

func GetQueryParams(url string) []string {
	if !strings.Contains(url, "?") {
		return []string{}
	}
	return strings.SplitAfter(strings.SplitAfter(url, "?")[1], "&")
}

// Accepts time in format hh:mm
func AddOneHourToTime(time string) string {
	hh_mm := strings.Split(time, ":")
	log.Println("HHH", hh_mm)
	hh, _ := strconv.Atoi(hh_mm[0])
	hh = hh + 1
	mm := hh_mm[1]
	log.Println("HHH", hh)
	log.Println("HHH", mm)
	return "" + strconv.Itoa(hh) + ":" + mm
}