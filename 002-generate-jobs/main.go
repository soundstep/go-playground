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

	jobList := []job{
		job{"edna-check-platform-specific-freeview", "features/platform_specific/freeview"},
		job{"edna-check-platform-specific-youview", "features/platform_specific/youview"},
		job{"edna-check-platform-specific-youviewsony", "features/platform_specific/youviewsony"},
		job{"edna-check-tracking", "features features/tracking"},
		job{"edna-check-acceptance-menu-items", "features/acceptance/menu_items"},
		job{"edna-check-acceptance-non-menu-items-advertising", "features/acceptance/non_menu_items/advertising"},
		job{"edna-check-acceptance-non-menu-items-error-handling", "features/acceptance/non_menu_items/error_handling"},
		job{"edna-check-acceptance-non-menu-items-exiting", "features/acceptance/non_menu_items/exiting"},
		job{"edna-check-acceptance-non-menu-items-guidance", "features/acceptance/non_menu_items/guidance"},
		job{"edna-check-acceptance-non-menu-items-more-episodes", "features/acceptance/non_menu_items/more_episodes"},
		job{"edna-check-acceptance-non-menu-notifications", "features/acceptance/non_menu_items/notifications"},
		job{"edna-check-acceptance-non-menu-onward-journey", "features/acceptance/non_menu_items/onward_journey"},
		job{"edna-check-acceptance-non-menu-parental-controls", "features/acceptance/non_menu_items/parental_controls"},
		job{"edna-check-acceptance-non-menu-parental-misc", "features/acceptance/non_menu_items/misc"},
		job{"edna-check-functional-menu-items", "features/functional/menu_items"},
		job{"edna-check-functional-non-menu-items", "features/functional/non_menu_items"},
		job{"edna-check-domain", "features/domain"},
		job{"edna-check-risk", "features/risk"},
	}

	template := getTemplate()
	createJobs(jobList, template)

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

func createJobs(jobList []job, template string) {
	fmt.Println("--- create jobs")
	for i, value := range jobList {
		go createFile(i, value, template, len(jobList))
	}
}

func createFile(index int, jobDesc job, template string, length int) {
	fmt.Println("--- create file", index, jobDesc.name)
	file, err := os.Create("output/" + jobDesc.name + ".yaml.erb")
	if err != nil {
		log.Fatal(err)
	}
	jobContent := strings.Replace(template, "{{SHELL_TOKEN_FEATURE}}", jobDesc.tokenReplacement, -1)
	file.WriteString(jobContent)
	file.Close()
	fileCreated(index, length)
}

func fileCreated(index int, length int) {
	countFile++
	fmt.Println("--- file created", index, strconv.Itoa(countFile)+"/"+strconv.Itoa(length))
	if countFile == length {
		fmt.Println("--- templates creation completed")
		os.Exit(0)
	}
}
