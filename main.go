/*
   Copyright 2022 Ludger Pottmeier

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main


import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"os"
)

func getEnvMap() map[string]string {
	res := make(map[string]string)
	r := os.Environ()

	for i := 0; i < len(r); i++ {
		res1 := strings.Split(r[i], "=")
		res[res1[0]] = res1[1]
	}
	return res
}

func main() {

	var file string
	flag.StringVar(&file, "t", "", "Go template file")

	flag.Parse()

	if len(file) == 0 {
		fmt.Println("No file provided ! Usage: fill")
		flag.PrintDefaults()
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Cannot open " + file)
		flag.PrintDefaults()
		os.Exit(1)

	}
	// Create a new template and parse the filecontent into it.
	t := template.Must(template.New("fill").Parse(string(content)))

	// Execute the template using the environment variable
	env := getEnvMap()
	if err = t.Execute(os.Stdout, env); err != nil {
		fmt.Println("Unable to process template")
		os.Exit(1)
	}

}
