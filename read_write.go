package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Application loop
	programState := 1

	for programState != 0 {
		// Show current note
		// Read text from the input file
		noteData, errNoteData := os.ReadFile("note.txt")
		if errNoteData != nil {
			fmt.Println(errNoteData)
			return
		}
		noteDataString := string(noteData)

		// Format the note by appending index for every line
		noteLines := strings.Split(noteDataString, "\n")
		for i, line := range noteLines {
			noteLines[i] = fmt.Sprintf("%d %s\n", i+1, line)
		}
		// Print the formatted note
		for _, line := range noteLines {
			fmt.Print(line)
		}

		// Get user input for program option
		fmt.Print("\n\n" +
			"What do you want to do?\n" +
			"1) Update a line\n" +
			"2) Delete a line\n" +
			"0) Exit\n" +
			"Enter your option [1/2/0]: ")

		userOptionPick, errUserOptionPick := reader.ReadString('\n')
		if errUserOptionPick != nil {
			fmt.Println(errUserOptionPick)
			return
		}
		userOptionPick = strings.TrimSpace(userOptionPick)

		userOptionPickInt, err := strconv.Atoi(userOptionPick)
		if err != nil {
			fmt.Println(err)
			return
		}

		programState = userOptionPickInt

		if programState == 1 {
			// Change / insert a line with user input
			fmt.Print("Which line to be updated: ")
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
	}
}
