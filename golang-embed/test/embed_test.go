package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T)  {
	fmt.Println(version)
}


//go:embed profile.jpeg
var logo []byte
func TestByteSlice(t *testing.T)  {
	err := ioutil.WriteFile("profile_new.jpeg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}


//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var file embed.FS

func TestEmbedMultipleFile(t *testing.T)  {
	a, _ := file.ReadFile("file/a.txt")
	fmt.Println(a)

	b, _ := file.ReadFile("file/c.txt")
	fmt.Println(b)

	c, _ := file.ReadFile("file/b.txt")
	fmt.Println(c)
}



//go:embed files/*.txt
var path embed.FS

func TestEmbedPatchMatcher(t *testing.T)  {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir(){
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content : " + string(content))
		}
	}
}