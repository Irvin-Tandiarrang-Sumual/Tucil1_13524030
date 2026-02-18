# Queens Solver

## Description
This program uses a brute force algorithm to solve the Queens game (a logic game accessible on LinkedIn). The game is solved when no queens are in the same row/column/region on the board and cannot be diagonally adjacent.

## Requirements
```
Go 1.25.5
gcc
Fyne fyne.io/fyne/v2 v2.7.2
```

## Installation
### Clone the repository
```bash
git clone <repo-url>
cd Tucil_13524030
```

## Compiling the Program
```bash
cd src
go build
```
## Running the Program

### Option 1: Using Executable File 
```bash
cd bin/<folder>
./<filename>
```
Note: the folder can be a Linux or Windows folder, and the file name depends on each folder.

### Option 2: Using go run
```bash
cd src
go run main.go
```

## Usage
1. **Load a File**
   - Click the "Load txt File" button
   - Select a text file containing the puzzle configuration
   - The puzzle board will be displayed in the grid

2. **Solve the Puzzle**
   - Click the "Solve" button
   - The algorithm will solve the Queens puzzle
   - The solution will be displayed on the board
   - A message will show the solving status

3. **Save the Solution**
   - After solving, click the "Save Solution to TXT" button
   - Choose a location and filename to save the solution
   - The solution will be saved as a text file

### Input File Format
Only consists of Alphabet letters from A - Z

**Example (4x4 board):**
```
AABB
AABB
CCDD
CCDD
```

### Output File Format TXT
Only consists of Alphabet letters from A - Z and # that denotes Queen
**Example (4x4 board):**
```
AA#B
#ABB
CCD#
C#DD
```

## Author
```bash
Nama : Irvin Tandiarrang Sumual
NIM  : 13524030
```
