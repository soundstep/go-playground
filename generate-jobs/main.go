package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var countFile = 0

func main() {
	tokenReplacements := []job{
		job{"fist-job", "First token"},
		job{"second-job", "Second token"},
	}

	template := getTemplate()
	createJobs(tokenReplacements, template)

	time.Sleep(time.Millisecond * 10000)
}

type job struct {
	name             string
	tokenReplacement string
}

func getTemplate() string {
	fmt.Println("--- read template")
	stream, err := ioutil.ReadFile("job-template.yaml.erb")
	if err != nil {
		log.Fatal(err)
	}
	return string(stream)
}

func createJobs(tokenReplacements []job, template string) {
	fmt.Println("--- create jobs")
	for i, value := range tokenReplacements {
		fmt.Println(i, value)
		go createFile(i, tokenReplacements[i], template, len(tokenReplacements))
	}
}

func createFile(index int, jobDesc job, template string, length int) {
	fmt.Println("--- create file", index, jobDesc.name)
	file, err := os.Create("output/" + jobDesc.name + ".yaml.erb")
	if err != nil {
		log.Fatal(err)
	}
	jobContent := strings.Replace(template, "{{SHELL_COMMAND}}", jobDesc.tokenReplacement, -1)
	file.WriteString(jobContent)
	file.Close()
	fileCreated(index, length)
}

func fileCreated(index int, length int) {
	fmt.Println("--- file created", strconv.Itoa(index+1)+"/"+strconv.Itoa(length))
	if index == length-1 {
		fmt.Println("--- templates creation completed")
		os.Exit(0)
	}
}
