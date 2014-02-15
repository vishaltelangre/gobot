package main

// containsString is the function to lookup the slice of strings for an passd
// item; if the item is found in the slice, then it returns true, neither it
// returns false value.
func containsString(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}