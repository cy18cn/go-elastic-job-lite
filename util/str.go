package util

import (
	"fmt"
	"strconv"
)

func ToStringSlice(items interface{}) ([]string, error) {
	var s []string
	switch items.(type) {
	case []string:
		return items.([]string), nil
	case []bool:
		for _, item := range items.([]bool) {
			s = append(s, strconv.FormatBool(item))
		}
	case []float64:
		for _, item := range items.([]float64) {
			s = append(s, strconv.FormatFloat(item, 'f', -1, 64))
		}
	case []float32:
		for _, item := range items.([]float32) {
			s = append(s, strconv.FormatFloat(float64(item), 'f', -1, 64))
		}
	case []int:
		for _, item := range items.([]int) {
			s = append(s, strconv.Itoa(item))
		}
	case []int64:
		for _, item := range items.([]int64) {
			s = append(s, strconv.FormatInt(item, 10))
		}
	case []int32:
		for _, item := range items.([]int32) {
			s = append(s, strconv.Itoa(int(item)))
		}
	case []int16:
		for _, item := range items.([]int16) {
			s = append(s, strconv.FormatInt(int64(item), 10))
		}
	case []int8:
		for _, item := range items.([]int8) {
			s = append(s, strconv.FormatInt(int64(item), 10))
		}
	case []uint:
		for _, item := range items.([]uint) {
			s = append(s, strconv.FormatUint(uint64(item), 10))
		}
	case []uint64:
		for _, item := range items.([]uint64) {
			s = append(s, strconv.FormatUint(item, 10))
		}
	case []uint32:
		for _, item := range items.([]uint32) {
			s = append(s, strconv.FormatUint(uint64(item), 10))
		}
	case []uint16:
		for _, item := range items.([]uint16) {
			s = append(s, strconv.FormatUint(uint64(item), 10))
		}
	case []uint8:
		for _, item := range items.([]uint8) {
			s = append(s, strconv.FormatUint(uint64(item), 10))
		}
	case [][]byte:
		for _, item := range items.([][]byte) {
			s = append(s, string(item))
		}
	case nil:
		s = nil
	default:
		return nil, fmt.Errorf("unable to cast %#v to []string", items)
	}

	return s, nil
}
