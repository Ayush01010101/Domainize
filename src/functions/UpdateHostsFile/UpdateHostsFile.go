package hostfile

import (
	"fmt"
	"os"
	"strings"
)

func UpdateHostsFile(port int, domain string) {
	entry := "127.0.0.1\t" + domain

	fmt.Println("now updaet the hosts file")
	content, err := os.ReadFile("/etc/hosts")
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(content), entry) {
		fmt.Println("Entry already exists")
		return
	}

	f, err := os.OpenFile(
		"/etc/hosts",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("\n" + entry + "\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Entry added successfully of %s", domain)

}
