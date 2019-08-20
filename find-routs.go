package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

//TODO: start with the way you are doing and adding roots for the ones that have no parent
//Then for everything else run dfs and if dfs returns false add it to the subdomains
//this way you will have a tree per root

type Routs struct {
	Domains []*Domain
}

type Domain struct {
	DomainName string
	SubDomains []*Domain
}

var (
	rootTable   = map[string]*Domain{}
	existingSub = map[string][]string{}
	roots       []*Domain
)

var fileName string

func addRoot(name string) {
	fmt.Printf("adding root:name=%v\n", name)
	node := &Domain{DomainName: name, SubDomains: []*Domain{}}
	if rootTable[name] == nil {
		roots = append(roots, node)
		rootTable[name] = node
	}
}

func addLeafDfs(name, root *Domain) {

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
	if roots == nil {
		fmt.Printf("show: root node not found\n")
		return
	}
	fmt.Printf("RESULT:\n")
	for _, node := range roots {
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
	loadConfig()
	fmt.Printf("main: reading input from stdin -- done\n")
	show()
	fmt.Printf("main: end\n")
	for _, dom := range roots {
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

func loadConfig() {
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
		root := rootTable[line[0]]
		if root == nil {
			addRoot(line[0])
		}
		for i, dom := range line {
			if i == 0 {
				add(dom, "")
			} else {
				add(dom, line[i-1])
			}
		}
	}

}
