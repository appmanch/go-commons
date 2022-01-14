package fsutils

import (
	"net/http"
	"os"
	"path/filepath"
)

//KnownFileTypes map holds the common extensions that are known to map against the mime type
var KnownFileTypes = map[string]string{
	".css":  "text/css",
	".gif":  "image/gif",
	".htm":  "text/html",
	".html": "text/html",
	".jpeg": "image/jpeg",
	".jpg":  "image/jpeg",
	".js":   "text/javascript",
	".mjs":  "text/javascript",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".svg":  "image/svg+xml",
	".wasm": "application/wasm",
	".webp": "image/webp",
	".xml":  "text/xml",
	".json": "application/json",
	".yaml": "text/yaml",
	".yml":  "text/yaml",
}

//FileExists function will check if the file exists in the specified path and if it is a file indeed
func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !fileInfo.IsDir()
}

//DirExists function will check if the Directory exists in the specified path
func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

//LookupContentType will lookup ContentType based on the file extension.
//This function will only check based on the name of the file and use the file extension.
func LookupContentType(path string) string {
	if val, ok := KnownFileTypes[filepath.Ext(path)]; ok {
		return val
	}
	return "application/octet-stream"
}

//DetectContentType will detect the content type of a file.
func DetectContentType(path string) (string, error) {
	file, err := os.Open(path)
	if err == nil {
		buffer := make([]byte, 512)
		n, err := file.Read(buffer)
		if err == nil {
			return http.DetectContentType(buffer[:n]), nil
		}
	}
	return "", err

}
