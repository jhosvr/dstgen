package main

import(
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	// "gopkg.in/go-ini/ini.v1"
)


func parseYaml(f string) map[interface{}]interface{}{
	contents := make(map[interface{}]interface{})
	yamlFile, err := ioutil.ReadFile(f)
    if err != nil {
        fmt.Println(err)
	}	
    err = yaml.Unmarshal(yamlFile, &contents)
    if err != nil {
        fmt.Println(err)
	}
	return contents
}


func main(){
	// https://stackoverflow.com/questions/51947879/read-and-merge-two-yaml-files
	config := parseYaml("./defaults.yaml")
	overrides := parseYaml("./overrides.yaml")
}
