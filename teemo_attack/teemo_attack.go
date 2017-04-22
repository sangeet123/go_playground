package teemo

func findPoisonedDuration(timeSeries []int, duration int) int {

	poisoned_for := 0
	affect_till := 0

	for i := 0; i < len(timeSeries); i++ {
		next_affect := timeSeries[i] + duration

		if next_affect <= affect_till {
			continue
		}

		if timeSeries[i] >= affect_till {
			poisoned_for += duration
		} else {
			poisoned_for += next_affect - affect_till
		}
		affect_till = next_affect
	}

	return poisoned_for
}
