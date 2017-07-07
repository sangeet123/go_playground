package mergemeetings

import "sort"

//defining Meeting type
type Meeting struct {
	Start int
	End   int
}

// defining type make us to define collection type
// easier
type Meetings []Meeting

//for sorting implement Len(), Less(i,j int) and Swap(i,j int)
func (slice Meetings) Len() int {
	return len(slice)
}

func (slice Meetings) Less(i, j int) bool {
	return slice[i].Start < slice[j].Start || (slice[i].Start == slice[j].Start && slice[i].End < slice[j].End)
}

func (slice Meetings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//core method for this function
func MergeMeetings(ranges Meetings) Meetings {
	sort.Sort(ranges)
	return MergeSortedMeetings(ranges)
}

func MergeSortedMeetings(ranges Meetings) Meetings {
	result := Meetings{}
	if len(ranges) == 0 {
		return result
	}
	result = append(result, ranges[0])
	cur_pos := 0
	for i := 1; i < len(ranges); i++ {
		if result[cur_pos].End >= ranges[i].Start && result[cur_pos].End < ranges[i].End {
			result[cur_pos].End = ranges[i].End

		} else if result[cur_pos].End < ranges[i].End {
			result = append(result, ranges[i])
			cur_pos++
		}
	}
	return result
}
