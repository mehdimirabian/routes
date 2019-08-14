package main

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Routs struct {
	Domains []*Domain
}

type Domain struct {
	DomainName	string
	SubDomains  []*Domain
}

var fileName string

func main(){
	fileName = os.Getenv("FILE_NAME")
	//check if environment is correctly set
	err := checkEnv()
	if err != nil {
		log.Fatal("error while checking environment: ", err.Error())
	}
	loadConfig()
}

func checkEnv() error{
	if fileName == "" {
		return errors.New("FILE_NAME environment variable not set")
	}
	return nil
}

func loadConfig() (*Routs, error) {
	//Read the routing config file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("error while reading the file: ", err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	log.Println(txtlines)
	//Do the following
	return nil, nil
}
