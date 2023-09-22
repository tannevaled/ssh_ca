package utils

import "os"

/*
 * cf https://www.tutorialspoint.com/how-to-check-if-a-file-exists-in-golang
 *
 * File exists and is not a directory
 *
 */
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
