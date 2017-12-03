package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// TODO: write the go generate line

// This program is a helper to generate all the go files from a truffle build directory containing Truffle Contract Objects

var (
	buildFlag = flag.String("build", "", "Path to the Truffle build repository")
	pkgFlag   = flag.String("pkg", "", "Name of the generated go package")
	outFlag   = flag.String("out", "", "Output folder for the generated go files")
)

func main() {

	flag.Parse()

	if *buildFlag == "" || *outFlag == "" || *pkgFlag == "" {
		fmt.Printf("No build folder (--build), output folder (--out) or package name (--pkg) specified\n")
		os.Exit(-1)
	}

	// Create the output directory if it does not exist
	if _, err := os.Stat(*outFlag); os.IsNotExist(err) {
		err = os.MkdirAll(*outFlag, 0755)
		if err != nil {
			fmt.Printf("Error creating repository: %v\n", err)
			os.Exit(-1)
		}
	}

	// Retrieve Json file names present in the folder
	files, err := ioutil.ReadDir(*buildFlag)
	if err != nil {
		fmt.Printf("Error with build folder's name: %v", err)
		os.Exit(-1)
	}

	for _, v := range files {

		// Check file extension
		if !strings.HasSuffix(v.Name(), ".json") {
			continue
		}

		content, err := ioutil.ReadFile(*buildFlag + "/" + v.Name())
		if err != nil {
			fmt.Printf("Error when reading file %v: %v", v.Name(), err)
			os.Exit(-1)
		}

		// Truffle Contract Object
		var tco map[string]interface{}

		err = json.Unmarshal(content, &tco)
		if err != nil {
			fmt.Printf("Error when unmarshaling file %v: %v", v.Name(), err)
			os.Exit(-1)
		}

		// Name of the type is the name of the file
		types := []string{strings.TrimSuffix(v.Name(), ".json")}

		// Marshall the abi into byte[]
		abi, err := json.Marshal(tco["abi"])
		if err != nil {
			fmt.Printf("Error when marshaling abi %v: %v", v.Name(), err)
			os.Exit(-1)
		}

		bin, ok := tco["bytecode"].(string)
		if !ok {
			fmt.Printf("Error when asserting bin %v: %v", v.Name(), err)
			os.Exit(-1)
		}

		// Call the ethereum Bind function to generate the go file
		code, err := bind.Bind(types, []string{string(abi)}, []string{bin}, *pkgFlag, bind.LangGo)
		if err != nil {
			fmt.Printf("Error when generating go bindings for %v: %v\n", v.Name(), err)
			os.Exit(-1)
		}

		// Create the go file name (2 extra lines to lower the first letter..)
		fileNameUni := []rune(strings.TrimSuffix(v.Name(), ".json") + ".go")
		fileNameUni[0] = unicode.ToLower(fileNameUni[0])
		fileName := string(fileNameUni)

		completePath := *outFlag + "/" + fileName
		if err := ioutil.WriteFile(completePath, []byte(code), 0600); err != nil {
			fmt.Printf("Error when creating the go file %v: %v\n", fileName, err)
			os.Exit(-1)
		}

		fmt.Println(fileName)
	}
	return
}
