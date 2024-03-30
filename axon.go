package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/thatisuday/commando"
)

const (
	NAME    = "axon"
	VERSION = "1.2.0"
)

func main() {
	fmt.Println(NAME, VERSION)
	deletePreviousInstallation()

	// setting up executable details
	commando.
		SetExecutableName(NAME).
		SetVersion(VERSION).
		SetDescription("axon is a command line utility to organise and pretty your file system quickly and reliably.")

	// root command
	commando.
		Register(nil).
		SetShortDescription("Run axon with default options.").
		SetDescription("Run axon with default options, organising all files and folders.").
		AddArgument("dirs...", "The directory to be organized.", "./").
		AddFlag("prettify,p", "Prettify all files with a desired casing.", commando.String, "none").
		AddFlag("organise,o", "Organise the directory.", commando.Bool, true).
		AddFlag("rename,r", "Rename the files numerically with a certain alias.", commando.String, "none").
		AddFlag("regex,x", "Filter files using regular expressions.", commando.String, ":_none_:").
		AddFlag("insensitive,i", "Make the provided regex case-insensitive.", commando.Bool, false).
		AddFlag("move,m", "Move selected files to a directory.", commando.String, ":_none_:").
		AddFlag("copy,c", "Copy selected files to a directory.", commando.String, ":_none_:").
		AddFlag("verbose,V", "Enable verbose output.", commando.Bool, false).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			// getting all arg and flag values
			dirsGiven := strings.Split(args["dirs"].Value, ",")
			dirs := []string{}
			for _, d := range dirsGiven {
				norm, err := filepath.Abs(d)
				if err != nil {
					fmt.Printf("unable to get absolute path of `%s`, error: %v, skipping it \n", d, err)
					continue
				}
				dirs = append(dirs, norm)
			}
			verboseOutput, e := flags["verbose"].GetBool()
			if e != nil {
				verboseOutput = false
			}
			prettify, e := flags["prettify"].GetString()
			if e != nil {
				prettify = "none"
			}
			prettify = strings.ToLower(prettify)
			if !itemInSlice(prettify, []string{"lower", "upper", "title"}) {
				fmt.Println("invalid prettify casing: must be one of {lower, upper, title}")
				return
			}
			organise, e := flags["organise"].GetBool()
			if e != nil {
				organise = false
			}

			rename, e := flags["rename"].GetString()
			if e != nil {
				rename = "none"
			}

			moveToDir, e := flags["move"].GetString()
			if e != nil {
				moveToDir = ":_none_:"
			} else {
				moveToDir, e = filepath.Abs(moveToDir)
				if e != nil {
					fmt.Printf("unable to get absolute path of `%s`, error: %v, skipping moving operation \n", moveToDir, e)
					moveToDir = ":_none_:"
				}
			}

			copyToDir, e := flags["copy"].GetString()
			if e != nil {
				copyToDir = ":_none_:"
			} else {
				copyToDir, e = filepath.Abs(copyToDir)
				if e != nil {
					fmt.Printf("unable to get absolute path of `%s`, error: %v, skipping copying operation \n", copyToDir, e)
					copyToDir = ":_none_:"
				}
			}

			regexPattern, e := flags["regex"].GetString()
			if e != nil {
				regexPattern = ""
			}
			caseInsensitive, e := flags["insensitive"].GetBool()
			if e != nil {
				caseInsensitive = false
			}
			if caseInsensitive {
				regexPattern = "(?i)" + regexPattern
			}
			regex, err := regexp.Compile(regexPattern)
			if err != nil {
				fmt.Printf("unable to parse the given regex. please check it again.\n%v", err)
				return
			}
			// making a buffered channel
			ch := make(chan string, len(dirs))
			// todo better actions report
			// organising the files
			for _, dir := range dirs {
				go func(dir string) {
					if validPath(dir) {
						fo := NewFileOrganizer(dir, regex)
						if verboseOutput {
							fmt.Println("Selected files:")
							files, e := fo.getFilesInFolder()
							if e != nil {
								ch <- fmt.Sprintf("unable to get files in the folder, error: %v\n", e)
								return
							}
							for i, v := range files {
								fmt.Printf("%v. %v \n", i+1, v.Name())
							}
						}

						if prettify != "none" {
							fo.prettify(prettify)
						}

						if moveToDir != ":_none_:" {
							fo.move(moveToDir)
						}

						if copyToDir != ":_none_:" {
							fo.copy(copyToDir)
						}

						// todo better renaming
						if rename != "none" {
							fo.renameDir(rename)
						}

						if organise {
							fo.organize()
						}

						ch <- fo.showActions()

					} else {
						ch <- fmt.Sprintf("Skipping `%s` since it's not a valid directory.", dir)
					}

				}(dir)
			}

			// waiting for all the goroutines to finish
			for i := 0; i < len(dirs); i++ {
				fmt.Println(<-ch)
			}

		})

	// up command
	commando.
		Register("up").
		SetShortDescription("Update axon.").
		SetDescription("Update axon to the latest version.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			update()
		})

	commando.Parse(nil)
}
