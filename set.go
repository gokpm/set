package set

func Convert(numArr []int) []int {
	setMap := map[int]bool{}
	setArr := []int{}
	for _, value := range numArr {
		setMap[value] = true
	}
	for key, _ := range setMap {
		setArr = append(setArr, key)
	}
	return setArr	
}

func Union(numMat ...[]int) []int {
	setMap := map[int]bool{}
	setArr := []int{}
	for _, row := range numMat {
		for _, value := range row {
			setMap[value] = true
		}
	}
	for key, _ := range setMap {
		setArr = append(setArr, key)
	}
	return setArr
}

func Intersection(numMat ...[]int) []int {
	setMap := map[int]bool{}
	setArr := []int{}
	for _, value := range numMat[0] {
		setMap[value] = true
	}
	numMat = numMat[1:]
	for _, row := range numMat {
		tmpMap := map[int]bool{}
		for _, value := range row{
			tmpMap[value] = true
		}
		for key, _ := range setMap {
			if !tmpMap[key] {
				delete(setMap, key)
				continue
			}
		}
	}
	for key, _ := range setMap {
		setArr = append(setArr, key)
	}
	return setArr
}
