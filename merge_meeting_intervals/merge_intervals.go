package mergeMeetings

import "sort"

//defining Meeting type
type Meeting struct{
	start int
	end int
}

// defining type make us to define collection type
// easier
type Meetings []Meeting


//for sorting implement Len(), Less(i,j int) and Swap(i,j int)
func (slice Meetings) Len() int{
	return len(slice)
}

func (slice Meetings) Less(i, j int) bool{
	return slice[i].start < slice[j].start ||  (slice[i].start == slice[j].start && slice[i].end < slice[j].end)
}

func(slice Meetings) Swap(i, j int){
	slice[i], slice[j] = slice[j], slice[i]
}

//core method for this function
func MergeMeetings(ranges Meetings)Meetings{
	result := Meetings{}
	sort.Sort(ranges)
	if len(ranges) == 0{
		return result
	}

	cur_pos := 0
	result = append(result, ranges[0])
	for i := 1 ; i < len(ranges); i++{
		if result[cur_pos].end >= ranges[i].start && result[cur_pos].end < ranges[i].end{
			result[cur_pos].end = ranges[i].end

		}else if result[cur_pos].end < ranges[i].end{
			result = append(result, ranges[i])
			cur_pos++
		}
	}

	return result
}
