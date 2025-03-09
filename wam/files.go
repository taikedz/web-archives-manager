package wam

func requireFileData(path string) string {
	if fh, err := openFile(path); err == nil {
		return fh.read()
	} else {
		fail(FILE_ERROR, "Could not open %s", prefix_definitions_path)
	}
}

func removeFile(path string) {
	// Remove the file, or fail
}

func writeFile(path string, data string) {
	// write, or fail
}