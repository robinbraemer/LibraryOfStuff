package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
}

const (
	file = "all.yaml"
)

func Main() error {
	all, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(all), "\n")
	changes := enableSSL(lines)

	if len(changes) == 0 {
		fmt.Printf("All SSL enabled in %s!\n", file)
		return nil
	}

	if err = ioutil.WriteFile(file, []byte(strings.Join(lines, "\n")), 0); err != nil {
		return err
	}
	fmt.Printf("Rewritten %d lines %s to enable SSL!\n", changes, file)
	return nil
}

func enableSSL(lines []string) (changedLines []int) {
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if strings.Contains(l, "name: ENABLE_SSL") && strings.Contains(lines[i+1], `value: "false"`) {
			i++
			lines[i] = strings.ReplaceAll(lines[i], `value: "false"`, `value: "true"`)
			changedLines = append(changedLines, i)
		}
	}
	return
}
