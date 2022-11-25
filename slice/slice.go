package slice

// DifferenceSet 返回两个slice切片的差集
func DifferenceSet(a []int, b []int) []int {
	var c []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}

	return c
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
// 结构体切片去重
func removeRepByMap(slc []*SeriesRes) []*SeriesRes {
	resultMap := make(map[string]*SeriesRes, len(slc))
	for _, v := range slc {
		resultMap[v.SeriesId] = v
	}
	var result []*SeriesRes
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
