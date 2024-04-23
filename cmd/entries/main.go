package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var (
	_, b, _, _  = runtime.Caller(0)
	ProjectRoot = filepath.Dir(b)
	FILENAME    = filepath.Join(ProjectRoot, "entries.html")
	OUTFILE     = "entries.csv"
)

func main() {
	var reEntries = regexp.MustCompile(`^(\S+) entries$`)
	var rePlayer = regexp.MustCompile(`^\s+\d+\s+(\d+)\s+([A-Z\- ]+)\s+(\d+).*$`)
	var (
		sectionID  string
		uscfID     string
		playerName string
		rating     string
	)
	
	// Get the contents of the plain text of entries.html
	cmd := exec.Command("pandoc", FILENAME, "-t", "plain")
	body, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Create(OUTFILE)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	bb := bytes.NewBuffer(body)
	scanner := bufio.NewScanner(bb)	
	for scanner.Scan() {
		line := scanner.Text()
		m := reEntries.FindStringSubmatch(line)
		if m != nil {
			sectionID = m[1]
		}
		m = rePlayer.FindStringSubmatch(line)
		if m != nil {
			uscfID = m[1]
			playerName = strings.TrimSpace(m[2])
			rating = m[3]
			fmt.Fprintf(fp, "%s,%s,%s,%s\n", sectionID, uscfID, playerName, rating)
		}
	}
}
