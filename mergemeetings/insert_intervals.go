package mergemeetings

// ranges has non overlapping intervals
// toInsert may overlap with existing interval
// use MergeSortedMeetings to merge the sorted one
// complexicity is O(2n)
func InsertInterval(ranges Meetings, toInsert Meeting) Meetings {
	result := []Meeting{}
	startFrom := 0
	//Get the position for the start to append
	for i := 0; i < len(ranges); i++ {
		if toInsert.Start < ranges[i].Start {
			break
		}
		startFrom++
	}

	result = append(result, ranges[0:startFrom]...)
	result = append(result, toInsert)
	result = append(result, ranges[startFrom:]...)
	return MergeSortedMeetings(result)

}
