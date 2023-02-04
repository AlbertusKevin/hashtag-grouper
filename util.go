package main

func indexContains(indexs []int, idx int) (bool){
	for _, i := range indexs {
		if i == idx {
				return true
		}
	}

	return false
}