package storage

import (
	"sort"
)

func InsertSort(data []UserStatistics, el UserStatistics) ([]UserStatistics,int){
	index := sort.Search(len(data), func(i int) bool { return data[i].SuccessRate > el.SuccessRate })
	data = append(data, UserStatistics{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data,index
}
