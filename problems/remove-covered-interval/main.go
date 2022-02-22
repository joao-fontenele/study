package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func (in Intervals) Swap(i, j int) {
	aux := in[i]
	in[i] = in[j]
	in[j] = aux
}

func (in Intervals) Less(i, j int) bool {
	x := in[j][0] - in[i][0]
	if x != 0 {
		return x > 0
	}

	y := in[j][1] - in[i][1]
	return y > 0
}

func (in Intervals) Len() int {
	return len(in)
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Sort(Intervals(intervals))

	res := 1
	prev := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[prev][0] && intervals[i][1] > intervals[prev][1] {
			res++
		}
		if intervals[i][1] > intervals[prev][1] {
			prev = i
		}
	}

	return res
}

func main() {
	var r int
	r = removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}})
	fmt.Printf("2 = r=%v\n", r)

	r = removeCoveredIntervals([][]int{{1, 4}, {2, 3}})
	fmt.Printf("1 = r=%v\n", r)

	r = removeCoveredIntervals([][]int{{1, 2}, {1, 4}, {3, 4}})
	fmt.Printf("1 = r=%v\n", r)

	r = removeCoveredIntervals([][]int{{3, 10}, {4, 10}, {5, 11}})
	fmt.Printf("2 = r=%v\n", r)

	r = removeCoveredIntervals([][]int{
		{97744, 99177},
		{9782, 42547},
		{21210, 35161},
		{31377, 85790},
		{53330, 82476},
		{59552, 64449},
		{4177, 4511},
		{22686, 79581},
		{7900, 55898},
		{70317, 75508},
		{48660, 60445},
		{4175, 59106},
		{64406, 97296},
		{2547, 35392},
		{24716, 42920},
		{69598, 78736},
		{74744, 91826},
		{66305, 69290},
		{34631, 95035},
		{23099, 86779},
		{23707, 83804},
		{84597, 91731},
		{76336, 97281},
		{4507, 19729},
		{68007, 87741},
		{47660, 72540},
		{20096, 72534},
		{79422, 89929},
		{66650, 94270},
		{17827, 46319},
		{5584, 44234},
		{52418, 53669},
		{29550, 79734},
		{4525, 37837},
		{26458, 54655},
		{10550, 97776},
		{19571, 49453},
		{62428, 94877},
		{41642, 76480},
		{22741, 44283},
		{40915, 83070},
		{41016, 59403},
		{49628, 97532},
		{2874, 25053},
		{902, 4508},
		{17388, 46321},
		{64582, 70841},
		{27836, 36686},
		{29291, 82231},
		{34819, 96651},
	})
	fmt.Printf("5 = r=%v\n", r)
}
