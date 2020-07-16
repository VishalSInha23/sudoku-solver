package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"sudoku/utils"
)

const (
	address = ":3001"
)

func main() {

	httpHandler := httpHandler{}

	fmt.Println("Starting server with address " + address)

	err := http.ListenAndServe(address, httpHandler)
	if err != nil {
		log.Fatal(err)
	}
}

type httpHandler struct {
}

// ServeHTTP : Request handler method for server http requests
func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	path := req.URL.String()
	fmt.Println("Received a new request. Path : " + path)

	switch path {
	case "/":
		fmt.Println("Serving index.html.")
		http.ServeFile(res, req, "./app")

	case "/solved.html":
		fmt.Println("Serving solved.html.")
		http.ServeFile(res, req, "./app/solved.html")

	case "/styles.css":
		fmt.Println("Serving styles.css.")
		http.ServeFile(res, req, "./app/styles.css")

	case "/script.js":
		fmt.Println("Serving script.js.")
		http.ServeFile(res, req, "./app/script.js")

	case "/sudoku":
		fmt.Println("Passing control to sudoku handler.")
		sudokuHandler(res, req)

	default:
		fmt.Printf("Invalid path called %s. Sending 404 response.\n", path)
		sendErrorResponse(res, http.StatusNotFound, "Invalid path.")
	}
}

type body struct {
	Sudoku utils.Sudoku `json:"sudoku"`
}

// sudokuHandler : Handles sudoku request
func sudokuHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		fmt.Printf("Request method not supported : %s.\n", req.Method)
		sendErrorResponse(res, http.StatusBadRequest, "Request method not supported.")
		return
	}

	fmt.Println("Reading request body.")

	reqBodyBytes, err := readRequestBody(req)
	if err != nil {
		sendErrorResponse(res, http.StatusInternalServerError, "Error in reading body.")
		return
	}

	var reqBody body
	err = json.Unmarshal(reqBodyBytes, &reqBody)
	if err != nil {
		sendErrorResponse(res, http.StatusInternalServerError, "Error in converting JSON to request body struct.")
		return
	}

	fmt.Println("Input sudoku :")
	utils.PrintSudoku(reqBody.Sudoku)

	respBody := body{}
	respBody.Sudoku, err = utils.SolveSudoku(reqBody.Sudoku, 0, 0)
	if err != nil {
		sendErrorResponse(res, http.StatusInternalServerError, "Error in solving sudoku.")
		return
	}

	fmt.Println("Output sudoku :")
	utils.PrintSudoku(respBody.Sudoku)

	respBodyBytes, err := json.Marshal(respBody)
	if err != nil {
		sendErrorResponse(res, http.StatusInternalServerError, "Error in converting request body struct to JSON.")
		return
	}

	fmt.Printf("Sending response : %v\n", respBody)

	res.Header().Set("content-length", strconv.Itoa(len(respBodyBytes)))
	res.WriteHeader(http.StatusOK)
	res.Write(respBodyBytes)
}

// sendErrorResponse : Sends error response
func sendErrorResponse(res http.ResponseWriter, errCode int, message string) {

	fmt.Printf("Sending error message : %v : %s", errCode, message)

	res.Header().Set("content-length", strconv.Itoa(len(message)))
	res.Header().Set("connection", "close")
	res.WriteHeader(errCode)
	res.Write([]byte(message))
}

// readRequestBody : Reads request body in the right way
func readRequestBody(req *http.Request) ([]byte, error) {

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
