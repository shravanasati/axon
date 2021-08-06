package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	path    string
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
	compressedExt := []string{"rar", "zip", "7z", "tar.gz"}
	folders := []string{"compressed files", "programs", "videos", "music", "others", "images"}

	fo.createDirs()

	for _, file := range files {
		if stringInSlice(strings.ToLower(file), folders) {
			continue
		}

		split := strings.Split(file, ".")
		if len(split) == 1 {
			continue
		}
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

	for i, file := range files {
		split := strings.Split(file, ".")
		ext := strings.ToLower(split[len(split)-1])

		os.Rename(file, fmt.Sprintf("%v %v.%v", newName, i+1, ext))
	}
	fo.actions = append(fo.actions, "Renamed all files in "+fo.path)
}
