# File Organizer

## Overview
This is a simple command-line tool written in **Go** that organizes files in a given directory by their extensions. It automatically creates categorized subdirectories and moves files into their respective folders based on their types.

## Features
- Scans a directory for files
- Organizes files by their extensions
- Creates a unique directory for organized files
- Handles duplicate directory names safely

## Run the project
To run this project, you need to have **Go installed**. If you haven't installed Go yet, follow the instructions [here](https://go.dev/doc/install).

Clone the repository and navigate into the project folder:
```sh
$ git clone https://github.com/yourusername/file-organizer.git
$ cd file-organizer
```

## Build the project:
```sh
$ go build -o file_organizer
```

## Installation

To install the latest version of the File Organizer, run:

```sh
$ go install github.com/Daniel-Sogbey/file_organizer@latest
```

## Usage
Run the program using the following command:
```sh
$ file_organizer -dir /path/to/directory
```
For example:
```sh
$ file_organizer -dir ~/Downloads
```

### Expected Behavior
- The script generates a unique folder inside the target directory (e.g., `XK23J9_organized`)
- Subdirectories for different file types are created (e.g., `pdf`, `images`, `txt`)
- Files are moved into the appropriate subdirectories

## Example
### Before:
```
Downloads/
├── document.pdf
├── image.jpg
├── script.py
```
### After Running:
```
Downloads/
├── XK23J9_organized/
│   ├── pdf/
│   │   ├── document.pdf
│   ├── images/
│   │   ├── id_card.jpg
│   │   ├── avatar.png
│   ├── py/
│   │   ├── script.py
```

## Error Handling
- If the directory does not exist, an error will be displayed.
- If permissions are insufficient, an error will be logged.
- If a file with the same name exists in the new directory, it will **not** be overwritten.

## Future Enhancements
- Add logging for better debugging
- Implement file **moving** instead of **creating empty files**
- Add support for **custom file organization rules**
- Add **multi-threading** for large directories

## License
This project is open-source under the MIT License. Feel free to modify and improve it!