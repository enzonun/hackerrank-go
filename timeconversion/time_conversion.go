package timeconversion

import (
	"fmt"
	"strconv"
)

func TimeConversion(time string) string {
	fmt.Println(time)
	var h, m, s int
	var suffix string
	//     var out string

	h, _ = strconv.Atoi(time[0:2])
	m, _ = strconv.Atoi(time[3:5])
	s, _ = strconv.Atoi(time[6:8])
	suffix = time[8:10]

	fmt.Println(h, m, s, suffix)

	if suffix == "AM" && h == 12 {
		h = 0
	}
	if suffix == "PM" && h != 12 {
		h += 12
	}

	var out = fmt.Sprintf("%02d:%02d:%02d", h, m, s)

	return out

}
