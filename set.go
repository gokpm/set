package set

type Set []int

func Convert(numArr []int) []int {
	setArr := []int{}
	for _, i := range numArr {
		flag := false
		for _, j := range setArr {
			if i == j {
				flag = true
			}
		}
		if flag {
			continue
		} else {
			setArr = append(setArr, i)
		}
	}
	return setArr
}

func Union(numMatrix ...[]int) []int {
	setArr := []int{}
	for _, i := range numMatrix {
		for _, j := range i {
			flag := false
			for _, z := range setArr {
				if z == j {
					flag = true
				}
			if flag {
				continue
			} else {
				setArr = append(setArr, j)
			}
			}
		}
	}
	return setArr
}

