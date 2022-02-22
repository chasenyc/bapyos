package main

import "fmt"
import "io/ioutil"
import "os"
import "github.com/galihrivanto/svg"
import "strings"
import "bufio"
import "bytes"
// import "encoding/xml"
// import "golang.org/x/net/html/charset"

func main() {
	pyosPath := StringPrompt("What is the filename and path of the PYOS seating chart?")
	fmt.Printf("Loading %s\n", pyosPath)

	pyosSVGString := getFirstSeatingChart(pyosPath)

	baPath := StringPrompt("What is the filename and path of the BA seating chart?")
	fmt.Printf("Loading %s\n", baPath)

	baSVGString := getFirstSeatingChart(baPath)
    // data is the file content, you can use it
	pyosReader := strings.NewReader(pyosSVGString)

	rootElement, _ := svg.Parse(pyosReader, false)
	sections := []*svg.Element{}
	
	for i:=0; i<len(rootElement.Children); i++ {
		pieces := strings.Split(rootElement.Children[i].Attributes["id"], ":")
		if len(pieces) > 3 { // if it's a section
			modifiedElement := rootElement.Children[i]
			modifiedElement.Attributes["style"] = "display: none;"
			sections = append(sections, rootElement.Children[i])
		}

	}
	fmt.Printf("Found %d sections in PYOS chart\n", len(sections))

	baReader := strings.NewReader(baSVGString)
	baRootElement, _ := svg.Parse(baReader, false)
	baRootElement.Children = append(baRootElement.Children, sections...)
	
	outputName := StringPrompt("What would you like to call the new seating chart?")
	file, err := os.Create(outputName)
    if err != nil {
        fmt.Println(err)
    }
	defer file.Close()
	w := &bytes.Buffer{}
	if err := svg.Render(baRootElement, w, false); err != nil {
		fmt.Println(err)
	}
	str := w.String()
	str = strings.Replace(str, "<:", "<", -1)
	str = strings.Replace(str, "</:", "</", -1)
	file.WriteString(str)

	fmt.Printf("Done!\n")
}

func getFirstSeatingChart(pyosPath string) string {
	data, err := ioutil.ReadFile(pyosPath)
    if err != nil {
        fmt.Println("Can't read file:", pyosPath)
        panic(err)
    }
	return string(data)
}

func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}