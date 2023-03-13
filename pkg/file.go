package pkg

import "io/fs"

func IsFileExist(files []fs.FileInfo, login string) bool {
	for _, file := range files {
		if file.Name() == login {
			return true
		}
	}
	return false
}
