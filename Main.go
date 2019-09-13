package dutch_flag_sort

func sort(arr []int) {
	i, j, n := 0, 0, len(arr)-1

	for j <= n {
		if arr[j] == 1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		} else if arr[j] == 3 {
			arr[j], arr[n] = arr[n], arr[j]
			n--
		} else {
			j++
		}
	}
}
