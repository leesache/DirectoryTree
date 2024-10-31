
# Directory Tree Structure

This program generates a tree-like representation of a directory structure, displaying files and folders hierarchically.

## Features

- **Directory traversal**: Displays folders and files within a specified directory.
- **File size display**: Shows the size of each file, indicating empty files as `(empty)`.
- **Optional file display**: Only displays files if the `-f` flag is passed, otherwise only directories are shown.

## Usage

To run this program, use the following command:

```bash
go run main.go <directory-path> [-f]
```

### Parameters

- `<directory-path>`: Required. The path of the directory you want to display.
- `[-f]`: Optional. If included, files within the directory are displayed along with their sizes. If omitted, only directories are shown.

### Example

```bash
go run main.go . -f
```

This will display the contents of the current directory (and its subdirectories) in a tree format, showing both folders and files along with their sizes.

## Example Output

Here's what a sample output might look like when run on a directory:

```
├───dir1
│   ├───subdir1
│   └───file1.txt (12b)
└───dir2
    ├───file2.txt (empty)
    └───subdir2
        └───file3.txt (20b)
```

In this example:
- Directories are displayed with indentation and branching lines.
- Files display their sizes in bytes, or `(empty)` if the file has no content.

## Notes

- The program panics if the command-line arguments are not in the correct format.
- Ensure you have permission to read all files in the specified directory.
