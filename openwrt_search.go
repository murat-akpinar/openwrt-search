package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

type TOH struct {
	Columns []string        `json:"columns"`
	Entries [][]interface{} `json:"entries"`
}

func main() {
	resp, err := http.Get("https://openwrt.org/toh.json")
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Okuma hatası:", err)
		os.Exit(1)
	}

	var toh TOH
	if err := json.Unmarshal(body, &toh); err != nil {
		fmt.Println("JSON ayrıştırma hatası:", err)
		os.Exit(1)
	}

	var brandIndex, modelIndex, versionIndex int = -1, -1, -1
	for i, col := range toh.Columns {
		if col == "brand" {
			brandIndex = i
		}
		if col == "model" {
			modelIndex = i
		}
		if col == "supportedcurrentrel" {
			versionIndex = i
		}
	}
	if brandIndex == -1 || modelIndex == -1 || versionIndex == -1 {
		fmt.Println("Gerekli sütun(lar) bulunamadı")
		os.Exit(1)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "Marka\tModel\tSürüm\n")
	fmt.Fprintf(w, "------\t-----\t-----\n")
	for _, entry := range toh.Entries {
		version, _ := entry[versionIndex].(string)
		if strings.HasPrefix(version, "24.") {
			brand, _ := entry[brandIndex].(string)
			model, _ := entry[modelIndex].(string)
			fmt.Fprintf(w, "%s\t%s\t%s\n", brand, model, version)
		}
	}
	w.Flush()
}
