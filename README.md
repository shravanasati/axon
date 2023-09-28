# axon

## Overview
*axon* is a cross platform, command line utility for organizing and prettifying cluttered folders, written in Go.



## Features
- Cross platform
- Regex Selection
- Prettify files in a folder
- Move files to a folder
- Organize files in a folder
- Rename files in a folder
- Fast and reliable


## Installation

**For Linux users:**

Execute the following command in bash:

```bash
curl https://raw.githubusercontent.com/Shravan-1908/axon/master/scripts/linux_install.sh > axon_install.sh

chmod +x ./axon_install.sh

bash ./axon_install.sh
```


**For MacOS users:**

Execute the following command in bash:

```bash
curl https://raw.githubusercontent.com/Shravan-1908/axon/master/scripts/macos_install.sh > axon_install.sh

chmod +x ./axon_install.sh

bash ./axon_install.sh
```

**For Windows users:**

Open Powershell **as Admin** and execute the following command:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; (Invoke-WebRequest -Uri https://raw.githubusercontent.com/Shravan-1908/axon/master/scripts/windows_install.ps1 -UseBasicParsing).Content | powershell -
```

To verify the installation of *axon*, restart the shell and execute `axon -v`. You should see output like this:

```
axon 1.0.0

Version: 1.0.0
```

If the output isn't something like this, you need to repeat the above steps carefully.


<br>


## Usage

<!-- todo document regex and move features -->

```
axon 1.1.0

axon is a command line utility to organise and pretty your file system quickly and reliably.

Usage:
   axon [dirs] {flags}
   axon <command> {flags}

Commands: 
   help                          displays usage informationn
   up                            Update axon.
   version                       displays version number

Arguments: 
   dirs                          The directory to be organized. (default: ./) {variadic}

Flags: 
   -h, --help                    displays usage information of the application or a command (default: false)
   -m, --move                    Move selected files to a directiry. (default: :none:)
   -o, --organise                Organise the directory. (default: false)
   -p, --prettify                Prettify all files with a desired casing. (default: none)
   -x, --regex                   Filter files using regular expressions. (default: none)
   -r, --rename                  Rename the files numerically with a certain alias. (default: none)
   -v, --version                 displays version number (default: false)

```

axon can accept multiple directories as arguments and will concurrently work on them with the provided options.