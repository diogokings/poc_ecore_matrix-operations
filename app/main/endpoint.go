package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"matrix-operations/app/service"
	"net/http"
	"strconv"
	"strings"
)

func echo(w http.ResponseWriter, r *http.Request) {
	var matrix, errorMsg = processFile(w, r)

	if errorMsg == "" {
		fmt.Fprintf(w, service.Echo(matrix))
	} else {
		fmt.Fprintf(w, errorMsg)
	}
}

func flatten(w http.ResponseWriter, r *http.Request) {
	var matrix, errorMsg = processFile(w, r)

	if errorMsg == "" {
		fmt.Fprintf(w, service.Flatten(matrix))
	} else {
		fmt.Fprintf(w, errorMsg)
	}
}

func invert(w http.ResponseWriter, r *http.Request) {
	var matrix, errorMsg = processFile(w, r)

	if errorMsg == "" {
		fmt.Fprintf(w, service.Invert(matrix))
	} else {
		fmt.Fprintf(w, errorMsg)
	}
}

func multiply(w http.ResponseWriter, r *http.Request) {
	var matrix, errorMsg = processFile(w, r)

	if errorMsg == "" {
		fmt.Fprintf(w, service.Multiply(matrix))
	} else {
		fmt.Fprintf(w, errorMsg)
	}
}

func sum(w http.ResponseWriter, r *http.Request) {
	var matrix, errorMsg = processFile(w, r)

	if errorMsg == "" {
		fmt.Fprintf(w, service.Sum(matrix))
	} else {
		fmt.Fprintf(w, errorMsg)
	}
}

func processFile(w http.ResponseWriter, r *http.Request) ([][]int64, string) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer
	file, header, err := r.FormFile("file")

	if err != nil {
		return nil, hasError(err, "Server couldn't process the file")
	}

	defer file.Close()

	log.Printf("File to read: %s\n", header.Filename)

	io.Copy(&buf, file)
	fileContent := buf.String()

	log.Printf("Matrix readed in the file:\n%v\n", fileContent)

	reader := csv.NewReader(strings.NewReader(fileContent))
	records, err := reader.ReadAll()

	if err != nil {
		return nil, hasError(err,
			"Server couldn't read the file! Maybe has empty spaces in the end of file or an extra comma in the end of a line")
	}

	matrix, errorMsg := validateMatrix(records)

	if errorMsg != "" {
		return nil, errorMsg
	}

	log.Println(records)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	buf.Reset()
	// do something else
	// etc write header
	return matrix, ""
}

func validateMatrix(stringMatrix [][]string) ([][]int64, string) {
	var err error
	var xLength = len(stringMatrix)
	var yLength = len(stringMatrix[0])

	if xLength != yLength {
		errorMsg := "It's not square matrix! Verify if there isn't extra commas in the end of the lines"
		log.Println(errorMsg)
		return nil, errorMsg
	}

	intMatrix := make([][]int64, xLength)
	for i := range intMatrix {
		intMatrix[i] = make([]int64, yLength)
	}

	for i := 0; i < xLength; i++ {
		for j := 0; j < yLength; j++ {
			value := strings.TrimSpace(stringMatrix[i][j])
			intMatrix[i][j], err = strconv.ParseInt(value, 10, 64)

			if err != nil {
				errorMsg := fmt.Sprintf(
					"The matrix has a value different of int!\nposition: [%d][%d] - value: %v",
					i, j, value)
				log.Print(errorMsg)
				return intMatrix, errorMsg
			}
		}
	}

	return intMatrix, ""
}

func hasError(err error, errorMsg string) string {
	log.Println(errorMsg)
	log.Printf("ERROR: %v\n", err)
	return errorMsg
}
