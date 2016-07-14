package main // ParseJSON parses the given bytes as a VersionInfo JSON.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (vi *VersionInfo) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal([]byte(jsonBytes), &vi)
}

// VersionInfo data container
type VersionInfo struct {
	FixedFileInfo  `json:"FixedFileInfo"`
	StringFileInfo `json:"StringFileInfo"`
	VarFileInfo    `json:"VarFileInfo"`
	Timestamp      bool
	Buffer         bytes.Buffer
	Structure      VSVersionInfo
	IconPath       string
	ManifestPath   string
}

// Translation with langid and charsetid.
type Translation struct {
	LangID    `json:"LangID"`
	CharsetID `json:"CharsetID"`
}

// FileVersion with 3 parts.
type FileVersion struct {
	Major int
	Minor int
	Patch int
	Build int
}

// FixedFileInfo contains file characteristics - leave most of them at the defaults.
type FixedFileInfo struct {
	FileVersion    `json:"FileVersion"`
	ProductVersion FileVersion
	FileFlagsMask  string
	FileFlags      string
	FileOS         string
	FileType       string
	FileSubType    string
}

// VarFileInfo is the translation container.
type VarFileInfo struct {
	Translation `json:"Translation"`
}

// StringFileInfo is what you want to change.
type StringFileInfo struct {
	Comments         string
	CompanyName      string
	FileDescription  string
	FileVersion      string
	InternalName     string
	LegalCopyright   string
	LegalTrademarks  string
	OriginalFilename string
	PrivateBuild     string
	ProductName      string
	ProductVersion   string
	SpecialBuild     string
}

func main() {
	file, err := os.Open("abcd.exe")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	byteFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(byteFile)
}
