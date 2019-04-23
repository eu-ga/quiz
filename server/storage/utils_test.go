package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := []UserStatistics{}
	el1 := UserStatistics{
		SuccessRate:60,
	}
	el2 := UserStatistics{
		SuccessRate:70,
	}
	el3 := UserStatistics{
		SuccessRate:10,
	}
	data, _ = InsertSort(data,el1)
	data, _ = InsertSort(data,el2)
	data, _ = InsertSort(data,el3)
	assert.True(t, data[0].SuccessRate==el3.SuccessRate,fmt.Sprintf("First element should be %d, but has %d",el3.SuccessRate,data[0].SuccessRate))
}