package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type FileInfo interface {
	Name() string // base name of the file
	Size() int64  // length in bytes for regular files; system-dependent for others
	// Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}

func main() {
	// return *File
	f, err := os.Open("abcd.exe")
	if err != nil {
		log.Println(err)
	}
	entry, err := f.Readdir(0)
	if err != nil {
		log.Println(err)
	}

	for _, value := range entry {
		fmt.Println(value)
	}

	// 디렉토리의 정보를 읽어
	// dir, err := os.Open("abcd.exe")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// entries, err := dir.Readdir(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(entries)
	//
	// list := []FileInfo{}
	//
	// for _, entry := range entries {
	// 	f := FileInfo{
	// 		Name:    entry.Name(),
	// 		Size:    entry.Size(),
	// 		Mode:    entry.Mode(),
	// 		ModTime: entry.ModTime(),
	// 		IsDir:   entry.IsDir(),
	// 	}
	// 	list = append(list, f)
	// }
	//
	// output, err := json.Marshal(list)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(output))

}
