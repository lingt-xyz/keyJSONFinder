package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type (
	// types that are mapping vex json json
	GraphJson struct {
		Directed       int    `json:"directed"`
		ProjectName    string `json:"projectname"`
		ProjectAddress int    `json:"projectaddress"`
		NumberOfBlocks int    `json:"numberofblocks"`
	}

	EdgeJson struct {
		Source int `json:"source"`
		Target int `json:"target"`
	}

	NodeJson struct {
		Id        int    `json:"id"`
		InDegree  int    `json:"indegree"`
		OutDegree int    `json:"outdegree"`
		StartAdd  string `json:"startadd"`
		Code      string `json:"code"`
		JumpKind  string `json:"jumpkind"`
	}

	NormalizedCfgJson struct {
		Graph GraphJson
		Edges []EdgeJson
		Nodes []NodeJson
	}
)

func findJSON(folder string, keywords []string, top int) {
	outputDir, err := prepareDir()
	if err != nil {
		log.Fatalf("Cannot create output folder %q", outputDir)
	}
	topK := 0
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_normalized_cfg.json") {
			if found, err := containsKeywords(path, keywords); found && err == nil {
				_ = copyFile(path, filepath.Join(outputDir, info.Name()))
				topK++
				if topK >= top {
					return io.EOF
				}
			}
		}
		return nil
	})
	if err != nil {
		if err == io.EOF {
		} else {
			fmt.Printf("error walking the path %q: %v\n", folder, err)
		}
	}
}

func prepareDir() (string, error) {
	path, _ := os.Getwd()
	outputDir := filepath.Join(path, "output")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return outputDir, nil
}

func containsKeywords(filePath string, keywords []string) (bool, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	data := &NormalizedCfgJson{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return false, err
	}
	for i := range data.Nodes {
		for j := range keywords {
			if strings.ToUpper(data.Nodes[i].JumpKind) == strings.ToUpper(keywords[j]) {
				log.Printf("Found file %q, keyword \"%v\"", filePath, keywords[j])
				return true, nil
			}
		}
	}

	return false, nil
}

func copyFile(srcFile, destFile string) error {
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
