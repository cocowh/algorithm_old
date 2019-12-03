package _134_canCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	totalLast,currLast,point,length := 0,0,0,len(gas)
	for i := 0; i < length ; i++ {
		currLast += gas[i] - cost[i]
		totalLast += gas[i] - cost[i]
		if currLast < 0 {
			point = i + 1
			currLast = 0
		}
	}
	if totalLast < 0 {
		return -1
	}
	return point
}

