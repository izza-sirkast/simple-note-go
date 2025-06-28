package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//text := "This is midnight coding session"

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Input: ")
	// input, err_read_input := reader.ReadString('\n')
	// if err_read_input != nil {
	// 	fmt.Println(err_read_input)
	// 	return
	// }
	// input = strings.TrimSpace(input)
	// fmt.Println(input)

	// err := os.WriteFile("test.txt", []byte(input), 0644)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// data, err := os.ReadFile("test.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(string(data))

	reader := bufio.NewReader(os.Stdin)

	// Read text from the input file
	noteData, errNoteData := os.ReadFile("note.txt")
	if errNoteData != nil {
		fmt.Println(errNoteData)
		return
	}
	noteDataString := string(noteData)

	// Append index for every line
	noteLines := strings.Split(noteDataString, "\n")
	for i, line := range noteLines {
		noteLines[i] = fmt.Sprintf("%d %s\n", i+1, line)
	}
	for _, line := range noteLines {
		fmt.Print(line)
	}

	// Change / insert a line with user input
	fmt.Print("\n\n\n")
	fmt.Print("Which line to be inserted: ")
	insertLineNum, errInsertLineNum := reader.ReadString('\n')
	if errInsertLineNum != nil {
		fmt.Println(errInsertLineNum)
		return
	}
	insertLineNum = strings.TrimSpace(insertLineNum)
	insertLineNumInt, errInsertLineNumInt := strconv.Atoi(insertLineNum)
	if errInsertLineNumInt != nil {
		fmt.Println(errInsertLineNumInt)
		return
	}

	fmt.Print("Input: ")
	insertLine, errInsertLine := reader.ReadString('\n')
	if errInsertLine != nil {
		fmt.Println(errInsertLine)
		return
	}

	noteLinesUpdated := strings.Split(noteDataString, "\n")
	noteLinesUpdated[insertLineNumInt-1] = insertLine
	var updatedNoteStrings string
	for _, line := range noteLinesUpdated {
		updatedNoteStrings += fmt.Sprintf("%s\n", line)
	}

	// Write it to the file
	errWriteNoteData := os.WriteFile("note.txt", []byte(updatedNoteStrings), 0644)

	if errWriteNoteData != nil {
		fmt.Println(errWriteNoteData)
		return
	}
}
