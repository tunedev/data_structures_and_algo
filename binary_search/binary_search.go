package binarysearch

import("fmt")

func BinarySearch(haystack []int, needle int) bool {
	start := 0
	end := len(haystack)

	for start < end {
		mid := (end + (end - start))/2
		fmt.Println("mid ===>", mid)
		curr := haystack[mid]

		if(curr == needle) {
			return true
		}else {
			if curr > needle {
				start = mid + 1
			} else {
				end = mid
			}
		}
	}
	return false
}
