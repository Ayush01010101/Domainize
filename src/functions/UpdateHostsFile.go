package functions

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func hostsFilePath() string {
	if runtime.GOOS == "windows" {
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	}
	return "/etc/hosts"
}

func UpdateHostsFile(port int, domain string) {
	entry := "127.0.0.1\t" + domain
	hostsFile := hostsFilePath()

	fmt.Println("now updaet the hosts file")
	content, err := os.ReadFile(hostsFile)
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(content), entry) {
		fmt.Println("Entry already exists")
		return
	}

	f, err := os.OpenFile(
		hostsFile,
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

// RemoveFromHostsFile removes the 127.0.0.1 -> domain mapping this tool added.
// No-op if the entry is absent. Only touches exact loopback entries for this
// domain, so unrelated user lines are preserved. Works on Linux, macOS and
// Windows via hostsFilePath().
func RemoveFromHostsFile(domain string) {
	hostsFile := hostsFilePath()

	content, err := os.ReadFile(hostsFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	kept := make([]string, 0, len(lines))
	removed := false
	for _, line := range lines {
		fields := strings.Fields(line) // Fields treats \r and \t as whitespace
		if len(fields) == 2 && fields[0] == "127.0.0.1" && fields[1] == domain {
			removed = true
			continue
		}
		kept = append(kept, line)
	}

	if !removed {
		fmt.Printf("No host entry found for %s\n", domain)
		return
	}

	if err := os.WriteFile(hostsFile, []byte(strings.Join(kept, "\n")), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Removed host entry for %s\n", domain)
}
