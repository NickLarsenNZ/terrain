package utils

func StringInSlice(search string, slice *[]string) bool {
	for _, item := range *slice {
		if search == item {
			return true
		}
	}
	return false
}

func AppendError(slice *[]error, item error) {
	if item != nil {
		*slice = append(*slice, item)
	}
}
