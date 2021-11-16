package set

type Set []int

func Convert(numArr []int) Set {
	setMap := map[int]bool{}
	setArr := Set{}
	for _, value := range numArr {
		setMap[value] = true
	}
	for key, _ := range setMap {
		setArr = append(setArr, key)
	}
	return setArr	
}

func (setArr *Set) Convert() {
	setMap := map[int]bool{}
	for _, value := range *setArr {
		setMap[value] = true
	}
	*setArr = Set{}
	for key, _ := range setMap {
		*setArr = append(*setArr, key)
	}	
}

func Union(numMat ...Set) Set {
	setMap := map[int]bool{}
	setArr := Set{}
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

func Intersection(numMat ...Set) Set {
	setMap := map[int]bool{}
	setArr := Set{}
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

func (setArr *Set) Append(numbers ...int) {
	*setArr = append(*setArr, numbers...)
	setArr.Convert()
}

func (setArr *Set) Delete(numbers ...int) {
	setMap := map[int]bool{}
	for _, value := range *setArr {
		setMap[value] = true
	}
	*setArr = Set{}
	for _, value := range numbers {
		delete(setMap, value)
	}
	for key, _ := range setMap {
		*setArr = append(*setArr, key)
	}
}
