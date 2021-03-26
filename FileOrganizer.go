package main

import (
	"fmt"
	"strings"
	"path/filepath"
	"os"
)

func stringInSlice(s string, slice []string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func validPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}


// FileOrganizer is the base struct for all the functions to organize files.
type FileOrganizer struct {
	path string
	actions []string
}

func (fo *FileOrganizer) prettify(casing string) {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if casing == "lower" {
			os.Rename(file, strings.ToLower(file))
		} else if casing == "upper" {
			os.Rename(file, strings.ToUpper(file))
		} else {
			os.Rename(file, strings.Title(file))
		}
	}
	fo.actions = append(fo.actions, "Prettified the directory to "+casing)
}

func (fo *FileOrganizer) createDirs() {
	os.Chdir(fo.path)

	if validPath("Images") == false {
		os.Mkdir("Images", os.ModePerm)
		fo.actions = append(fo.actions, "Created Images directory in "+fo.path)
	}
	if validPath("Music") == false {
		os.Mkdir("Music", os.ModePerm)
		fo.actions = append(fo.actions, "Created Music directory in "+fo.path)
	}
	if validPath("Videos") == false {
		os.Mkdir("Videos", os.ModePerm)
		fo.actions = append(fo.actions, "Created Videos directory in "+fo.path)
	}
	if validPath("Programs") == false {
		os.Mkdir("Programs", os.ModePerm)
		fo.actions = append(fo.actions, "Created Programs directory in "+fo.path)
	}
	if validPath("Compressed Files") == false {
		os.Mkdir("Compressed Files", os.ModePerm)
		fo.actions = append(fo.actions, "Created Compressed Files directory in "+fo.path)
	}
	if validPath(filepath.Join(fo.path, "Others")) == false {
		os.Mkdir("Others", os.ModePerm)
		fo.actions = append(fo.actions, "Created Others directory in "+fo.path)
	}
}

func (fo *FileOrganizer) organize() {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	imageExt := []string{"jpg", "jpeg", "png"}
	musicExt := []string{"mp3", "aac", "ogg", "wav"}
	videoExt := []string{"mp4", "webm", "mov", "mkv", "mpv2", "avi"}
	programExt := []string{"exe", "msi", "msp", "dll"}
	compressedExt := []string{"rar", "zip", "7z", "tar.gz"}
	folders := []string{"compressed files", "programs", "videos", "music", "others", "images"}

	fo.createDirs()

	for _, file := range files {
		if stringInSlice(strings.ToLower(file), folders) {
			continue
		}

		split := strings.Split(file, ".")
		ext := strings.ToLower(split[len(split)-1])

		if stringInSlice(ext, imageExt) {
			os.Rename(file, filepath.Join("Images", file))
		} else if stringInSlice(ext, musicExt) {
			os.Rename(file, filepath.Join("Music", file))
		} else if stringInSlice(ext, videoExt) {
			os.Rename(file, filepath.Join("Videos", file))
		} else if stringInSlice(ext, compressedExt) {
			os.Rename(file, filepath.Join("Compressed Files", file))
		} else if stringInSlice(ext, programExt) {
			os.Rename(file, filepath.Join("Programs", file))
		} else if ext == "pdf" {
			os.Rename(file, filepath.Join("PDFs", file))
		} else {
			os.Rename(file, filepath.Join("Others", file))
		}

	}
	fo.actions = append(fo.actions, "Organized the files in "+fo.path)
}

func (fo *FileOrganizer) showActions() {
	fmt.Println("-----------------------------------------")
	data := "\nAction report " + "(" + fmt.Sprint(len(fo.actions)) + " actions performed):\n"
	for _, action := range fo.actions {
		data += action + "\n"
	}
	fmt.Println(data)
}

func (fo *FileOrganizer) renameDir(newName string) {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		split := strings.Split(file, ".")
		ext := strings.ToLower(split[len(split)-1])

		os.Rename(file, fmt.Sprintf("%v %v.%v", newName, i+1, ext))
	}
	fo.actions = append(fo.actions, "Renamed all files in "+fo.path)
}

func main()  {
	fmt.Println("Welcome to the GO FILE ORGANIZER!")
	
	var dir string
	fmt.Print("Enter the directory: ")
	fmt.Scanln(&dir)
	fo := FileOrganizer{path:dir, actions:[]string{}}


	if !(validPath(fo.path)) {
		panic("Invalid path to the directory!")
	}

	var task int
	fmt.Println("What do you want to do?\n 1. Prettify files \n 2. Organize files \n 3. Rename all files \n 4. Master plan (prettify + organize)")
	fmt.Scanln(&task)

	if task == 1 {
		for {
			var casing string
			fmt.Print("Enter the conversion case (lower/upper/title): ")
			fmt.Scanln(&casing)

			if stringInSlice(strings.ToLower(casing), []string{"lower", "upper", "title"}) {
				fo.prettify(casing)
				fo.showActions()
				break
			} else {
				fmt.Println("Invalid casing! It must be one from", []string{"lower", "upper", "title"})
			}
		}
	
	} else if task == 2 {
		fo.organize()
		fo.showActions()
	
	} else if task == 3 {
		var name string
		fmt.Print("Enter the new name for all the files in the directory: ")
		fmt.Scanln(&name)
		fo.renameDir(name)
		fo.showActions()
	
	} else if task == 4 {
		fo.prettify("title")
		fo.organize()
		fo.showActions()
	
	} else {
		fmt.Println("Invalid input!")
	}
	fmt.Scanln()
}