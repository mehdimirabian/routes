package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

type Routs struct {
	Domains []*Domain
}

type Domain struct {
	DomainName string
	SubDomains []*Domain
}

var (
	nodeTable = map[string]*Domain{}
	root      []*Domain
)

var fileName string

func add(name, parentId string) {
	fmt.Printf("add:name=%v parentId=%v\n", name, parentId)

	node := &Domain{DomainName: name, SubDomains: []*Domain{}}

	if parentId == "" && nodeTable[name] == nil {
		root = append(root, node)
	} else {

		parent, ok := nodeTable[parentId]
		if !ok {
			fmt.Printf("add: parentId=%v: not found\n", parentId)
			return
		}

		parent.SubDomains = append(parent.SubDomains, node)
	}

	nodeTable[name] = node
}

func showNode(node *Domain, prefix string) {
	if prefix == "" {
		fmt.Printf("%v\n\n", node.DomainName)
	} else {
		fmt.Printf("%v %v\n\n", prefix, node.DomainName)
	}
	for _, n := range node.SubDomains {
		showNode(n, prefix+"--")
	}
}

func show() {
	if root == nil {
		fmt.Printf("show: root node not found\n")
		return
	}
	fmt.Printf("RESULT:\n")
	for _, node := range root {
		showNode(node, "")
	}
}

func main() {
	fileName = os.Getenv("FILE_NAME")
	//check if environment is correctly set
	err := checkEnv()
	if err != nil {
		log.Fatal("error while checking environment: ", err.Error())
	}
	fmt.Printf("main: reading input from stdin\n")
	scan()
	fmt.Printf("main: reading input from stdin -- done\n")
	show()
	fmt.Printf("main: end\n")
	for _, dom := range root {
		fmt.Println(dom)
		for _, next := range dom.SubDomains {
			fmt.Println(next)
		}
	}
}

func checkEnv() error {
	if fileName == "" {
		return errors.New("FILE_NAME environment variable not set")
	}
	return nil
}

//
//func loadConfig() (*Routs, error) {
//	//Read the routing config file
//	file, err := os.Open(fileName)
//	if err != nil {
//		log.Fatalln("error while reading the file: ", err.Error())
//	}
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanLines)
//	var txtlines []string
//
//	for scanner.Scan() {
//		txtlines = append(txtlines, scanner.Text())
//		line := strings.Split(scanner.Text(), ".")
//		for _, dom := range(line) {
//			if myMap[dom] != nil {
//				myMap = myMap[dom].Subset
//			}
//			}
//		}
//	}
//	log.Println(txtlines)
//	//Do the following
//	return nil, nil
//}

func scan() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("error while reading the file: ", err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
		line := strings.Split(scanner.Text(), ".")
		for i, dom := range line {
			if i == 0 {
				add(dom, "")
			} else {
				add(dom, line[i-1])
			}
		}
	}

}
