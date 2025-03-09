package wam


import(
	"fmt"
)

var active_prefix_path string = "~/.config/wam/current"
var prefix_definitions_path string = "~/.config/wam/prefixes.json"

func Prefix_GetCurrent() (string, error) {
	if fileExists(active_prefix_path) {
		return requireFileData(active_prefix_path), nil
	} else {
		fail(PREFIX_ERROR, "No prefix currently active")
	}
}

func Prefix_SetValue(name string, path string) {
	prefixes := loadPrefixes()
	prefixes[name] = path
	savePrefixes(prefixes)
}

func Prefix_SetNone() {
	removeFile(active_prefix_path)
}

func Prefix_Choose(name string) {
	prefixes := loadPrefixes()
	if name in prefixes {
		writeFile(active_prefix_path, name)
	} else {
		fail(PREFIX_ERROR, "Invalid prefix name")
	}
}

func Prefix_List() {
	for name, _ := range loadPrefixes() {
		fmt.Printf("%s\n", name)
	}
}

// ====

func loadPrefixes() map[string]string {
	return parsePrefixes(requireFileData(prefix_definitions_path))
}

func savePrefixes(prefixes map[string]string) {
	writeFile(prefix_definitions_path, json_MapConvert(prefixes))
}

func parsePrefixes(data string) (map[string]string, error) {
	// FIXME placeholder data!
	result := make(map[string]string)
	result["data"] = data
	return result, nil
}

func json_MapConvert(prefixes map[string]string) string {
}