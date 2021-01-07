package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/dmathieu/itree"
	"github.com/goccy/go-graphviz"
)

var (
	FIREHOL_URL = "https://iplists.firehol.org/files/firehol_level1.netset"
	GRAPH_PATH  = "examples/firehol/firehol.gv"
	g           = graphviz.New()
)

func main() {
	content, err := download()
	if err != nil {
		log.Fatal(err)
	}
	data, err := parseData(content)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := itree.NewIPNetTree(data)
	if err != nil {
		log.Fatal(err)
	}

	_ = tree.Contains(net.IPv4(8, 8, 8, 8))

	graph, err := itree.Graphviz(tree.Tree)
	if err != nil {
		log.Fatal(err)
	}

	err = g.RenderFilename(graph, "dot", GRAPH_PATH)
	if err != nil {
		log.Fatal(err)
	}
}

func download() (io.Reader, error) {
	resp, err := http.Get(FIREHOL_URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("got an http error %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func parseData(data io.Reader) ([]*net.IPNet, error) {
	var ranges []*net.IPNet
	scanner := bufio.NewScanner(data)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		r, err := parseLine(line)
		if err != nil {
			return ranges, fmt.Errorf("error parsing line %d: %s", lineNum, err)
		}
		lineNum++

		ranges = append(ranges, r)
	}

	return ranges, scanner.Err()
}

func parseLine(line string) (network *net.IPNet, err error) {
	line = strings.TrimSpace(line)

	if !strings.Contains(line, "/") {
		line += "/32"
	}
	_, network, err = net.ParseCIDR(line)
	return
}
