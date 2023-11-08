package release

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Release struct {
	Commit string `json:"commit"`
	Image  string `json:"image"`
}

func Handler(response http.ResponseWriter, request *http.Request) {

	var event Release

	err := json.NewDecoder(request.Body).Decode(&event)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	go CodeRunner(event.Image)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(event)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Macro struct {
	Rule []string
}

func CodeRunner(image string) {
	dat, err := os.ReadFile("config/rules.json")
	check(err)

	var x map[string][]string

	json.Unmarshal(dat, &x)

	for a, b := range x {

		var re = regexp.MustCompile(`(?m)` + strings.Join(strings.Split(a, "/"), `\/`))

		for i, match := range re.FindAllString(image, -1) {
			fmt.Println(match, "found at index", i)
		}

		if re.MatchString(image) {

			fmt.Println("regex: ", a, " - ", image)

			for _, v := range b {
				fmt.Println("\t CMD: ", v)
				execCommand(v + " " + image)
			}
		}
	}

	fmt.Print()
}

func execCommand(cmds string) {

	cmdss := strings.Split(cmds, " ")

	comando, args := cmdss[0], cmdss[1:]

	cmd := exec.Command(comando, args...)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))
}
