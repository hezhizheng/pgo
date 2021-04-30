package pgo

import (
	"sync"
)

var mutex = sync.Mutex{}

func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		if _, ok := hystack.([]string); ok == true {
			for _, item := range hystack.([]string) {
				if key == item {
					return true
				}
			}
		}
		if _, ok := hystack.([]interface{}); ok == true {
			for _, item := range hystack.([]interface{}) {
				if key == item {
					return true
				}
			}
		}
	case int:
		if _, ok := hystack.([]int); ok == true {
			for _, item := range hystack.([]int) {
				if key == item {
					return true
				}
			}
		}
		if _, ok := hystack.([]interface{}); ok == true {
			for _, item := range hystack.([]interface{}) {
				if key == item {
					return true
				}
			}
		}
	case int64:
		if _, ok := hystack.([]int64); ok == true {
			for _, item := range hystack.([]int64) {
				if key == item {
					return true
				}
			}
		}
		if _, ok := hystack.([]interface{}); ok == true {
			for _, item := range hystack.([]interface{}) {
				if key == item {
					return true
				}
			}
		}
	default:
		return false
	}
	return false
}

func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			mutex.Lock()
			columns = append(columns, v)
			mutex.Unlock()
		}
	}
	return columns
}

func ArrayPush(s *[]interface{}, elements ...interface{}) int {
	mutex.Lock()
	*s = append(*s, elements...)
	mutex.Unlock()
	return len(*s)
}

func ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}


func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	mutex.Lock()
	*s = append(elements, *s...)
	mutex.Unlock()
	return len(*s)
}

func ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}