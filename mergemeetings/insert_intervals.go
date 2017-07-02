package mergemeetings

//import "fmt"

func InsertInterval(ranges Meetings, toInsert Meeting) Meetings {

	if len(ranges) <= 0 {
		return Meetings{toInsert}
	}

	mergePoint := 0
	isOverLapping := false
	for i := 0; i < len(ranges) && toInsert.End > ranges[i].Start && !isOverLapping; i++ {
		isOverLapping = overlaps(ranges[i],toInsert)
		if !isOverLapping{
			mergePoint = i + 1
		}
	}

	if !isOverLapping{
		result := Meetings{}
		result = append(result, ranges[0:mergePoint]...)
		result = append(result,Meetings{toInsert}...)
		result = append(result, ranges[mergePoint:]...)
		return result
	}

	if toInsert.Start < ranges[mergePoint].Start{
		ranges[mergePoint].Start = toInsert.Start
	}

	if toInsert.End > ranges[mergePoint].End{
		ranges[mergePoint].End = toInsert.End
	}

	mergeTill := mergePoint + 1
	for ;mergeTill < len(ranges) && ranges[mergePoint].End >= ranges[mergeTill].End; mergeTill++ {
	}

	return append(ranges[0:mergePoint+1],ranges[mergeTill:len(ranges)]...)

}

func overlaps(m1 Meeting, m2 Meeting) bool{
	isOverLapping := m2.End >= m1.End && m2.Start <= m1.Start
	isOverLapping = isOverLapping || m2.Start >= m1.Start && m2.Start < m1.End
	isOverLapping = isOverLapping || m2.End > m1.Start && m2.End <= m1.End
	return isOverLapping
}
