package utils

import (
	"managementInvestment/models"
)

func SelectionSort(investments []models.Investment, count int, getPriceFunc func(string) int, ascending bool) {
	for i := 0; i < count-1; i++ {
		selectedIndex := i
		for j := i + 1; j < count; j++ {
			currentValue := investments[j].Value(getPriceFunc(investments[j].Symbol))
			selectedValue := investments[selectedIndex].Value(getPriceFunc(investments[selectedIndex].Symbol))
			if ascending {
				if currentValue < selectedValue {
					selectedIndex = j
				}
			} else {
				if currentValue > selectedValue {
					selectedIndex = j
				}
			}
		}
		if selectedIndex != i {
			investments[i], investments[selectedIndex] = investments[selectedIndex], investments[i]
		}
	}
}

func InsertionSort(investments []models.Investment, count int, getPriceFunc func(string) int, ascending bool) {
	for i := 1; i < count; i++ {
		key := investments[i]
		j := i - 1
		for j >= 0 {
			currentValue := investments[j].Value(getPriceFunc(investments[j].Symbol))
			keyValue := key.Value(getPriceFunc(key.Symbol))
			if ascending {
				if currentValue > keyValue {
					investments[j+1] = investments[j]
				} else {
					break
				}
			} else {
				if currentValue < keyValue {
					investments[j+1] = investments[j]
				} else {
					break
				}
			}
			j--
		}
		investments[j+1] = key
	}
}

func SortBySymbol(investments []models.Investment, ascending bool) {
	n := len(investments)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			shouldSwap := false
			if ascending {
				if investments[j].Symbol > investments[j+1].Symbol {
					shouldSwap = true
				}
			} else {
				if investments[j].Symbol < investments[j+1].Symbol {
					shouldSwap = true
				}
			}
			if shouldSwap {
				investments[j], investments[j+1] = investments[j+1], investments[j]
			}
		}
	}
}
