package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func itemInSlice[T comparable](s T, slice []T) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func filter[T any](function func(arg T) bool, slice []T) []T {
	var newSlice []T
	for _, v := range slice {
		if function(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func validPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// FileOrganizer is the base struct for all the functions to organize files.
type FileOrganizer struct {
	path    string
	actions []string
	regex   *regexp.Regexp
}

func (fo *FileOrganizer) prettify(casing string) {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	for _, file := range filter(func(arg string) bool { return fo.regex.MatchString(arg) }, files) {
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

	if !validPath("./Images") {
		os.Mkdir("Images", os.ModePerm)
		fo.actions = append(fo.actions, "Created Images directory in "+fo.path)
	}
	if !validPath("./Music") {
		os.Mkdir("Music", os.ModePerm)
		fo.actions = append(fo.actions, "Created Music directory in "+fo.path)
	}
	if !validPath("./Videos") {
		os.Mkdir("Videos", os.ModePerm)
		fo.actions = append(fo.actions, "Created Videos directory in "+fo.path)
	}
	if !validPath("./Programs") {
		os.Mkdir("Programs", os.ModePerm)
		fo.actions = append(fo.actions, "Created Programs directory in "+fo.path)
	}
	if !validPath("./Compressed Files") {
		os.Mkdir("Compressed Files", os.ModePerm)
		fo.actions = append(fo.actions, "Created Compressed Files directory in "+fo.path)
	}
	if !validPath(filepath.Join(fo.path, "Others")) {
		os.Mkdir("./Others", os.ModePerm)
		fo.actions = append(fo.actions, "Created Others directory in "+fo.path)
	}
	if !validPath(filepath.Join(fo.path, "PDFs")) {
		os.Mkdir("./PDFs", os.ModePerm)
		fo.actions = append(fo.actions, "Created PDFs directory in "+fo.path)
	}
}

func (fo *FileOrganizer) organize() {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	imageExt := []string{"jpg", "jpeg", "png", "jfif"}
	musicExt := []string{"mp3", "aac", "ogg", "wav"}
	videoExt := []string{"mp4", "webm", "mov", "mkv", "mpv2", "avi", "gif"}
	programExt := []string{"exe", "msi", "msp", "dll", "out"}
	compressedExt := []string{"rar", "zip", "7z", "tar"}
	pdfExt := []string{"pdf"}
	folders := []string{"compressed files", "programs", "videos", "music", "others", "images"}

	fo.createDirs()

	for _, file := range filter(func(arg string) bool { return fo.regex.MatchString(arg) }, files) {
		if itemInSlice(strings.ToLower(file), folders) {
			continue
		}

		split := strings.Split(file, ".")
		if len(split) == 1 {
			continue
		}
		ext := strings.ToLower(split[len(split)-1])

		if itemInSlice(ext, imageExt) {
			os.Rename(file, filepath.Join("Images", file))
		} else if itemInSlice(ext, musicExt) {
			os.Rename(file, filepath.Join("Music", file))
		} else if itemInSlice(ext, videoExt) {
			os.Rename(file, filepath.Join("Videos", file))
		} else if itemInSlice(ext, compressedExt) {
			os.Rename(file, filepath.Join("Compressed Files", file))
		} else if itemInSlice(ext, programExt) {
			os.Rename(file, filepath.Join("Programs", file))
		} else if itemInSlice(ext, pdfExt) {
			os.Rename(file, filepath.Join("PDFs", file))
		} else {
			os.Rename(file, filepath.Join("Others", file))
		}

	}
	fo.actions = append(fo.actions, "Organized the files in "+fo.path)
}

func (fo *FileOrganizer) showActions() string {
	fmt.Println("-----------------------------------------")
	data := fmt.Sprintf("\nAction report for `%s` (%d actions performed): \n", fo.path, len(fo.actions))

	for _, action := range fo.actions {
		data += action + "\n"
	}

	return data
}

func (fo *FileOrganizer) renameDir(newName string) {
	os.Chdir(fo.path)
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}

	for i, file := range filter(func(arg string) bool { return fo.regex.MatchString(arg) }, files) {
		split := strings.Split(file, ".")
		ext := strings.ToLower(split[len(split)-1])

		os.Rename(file, fmt.Sprintf("%v %v.%v", newName, i+1, ext))
	}
	fo.actions = append(fo.actions, "Renamed all files in "+fo.path)
}

func (fo *FileOrganizer) move(targetDir string) {
	files, err := ioutil.ReadDir(fo.path)
	if err != nil {
		panic(err)
	}
	if !validPath(targetDir) {
		fmt.Printf("the target directory `%s` doesnt exist\n", targetDir)
		return
	}

	for _, file := range filter(func(arg fs.FileInfo) bool { return fo.regex.MatchString(arg.Name()) && !arg.IsDir() }, files) {
		newpath := filepath.Join(targetDir, file.Name())
		oldpath := filepath.Join(fo.path, file.Name())
		os.Rename(oldpath, newpath)
	}
}
