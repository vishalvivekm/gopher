/*
Meshery Component Updater
Uses a spreadsheet of centralized information about MeshModel components and their metadata like color, icon, and so on. Script is used to update components metada (svgs, icons etc) for Meshery, Websites (Layer5.io, Meshery.io), and Remote Provider.

Secret - this script expects that an environment variable `CRED` is available
and it contains a token for Google Sheets API interactions.

Example:
	export CRED='{
		"type": "service_account",
		"project_id": "",
		"private_key_id": "",
		"private_key": "-----BEGIN PRIVATE KEY-----\nn-----END PRIVATE KEY-----\n",
		"client_email": "",
		"client_id": "",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "",
		"universe_domain": "googleapis.com"
		}'

Usage: (order of flags matters)

    ./main [path-to-spreadsheet] [--system] [<system-name>] [relative path to docs in layer5 website] [relative path to docs in meshery website] [--only-published]

Examples:

	1. ./main https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system docs layer5/src/collections/integrations meshery.io/integrations docs/ --published-only
	2. ./main https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system remote-provider <remote-provider>/meshmodels/models <remote-provider>/ui/public/img/meshmodels
	3. ./main https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system meshery ../../server/meshmodel

The flags are:

  --system
        defined type of system to update. Can be one of "meshery", "docs", or "remote-provider".

	--only-published
        Only handle components that have a value of "true" under the "Published?" column in spreadsheet.
*/

package main

import (
	"encoding/csv"
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	//	"bytes"
	//"encoding/xml"
	"io"
	"net/http"
)

type TemplateAttributes struct {
	Title                   string
	ModelName               string
	Subtitle                string
	DocURL                  string
	Category                string
	Subcategory             string
	FeatureList             string
	HowItWorks              string
	HowItWorksDetails       string
	AboutProject            string
	StandardBlurb           string
	WorkingSlides           string
	Published               string
	IntegrationIcon         string
	DarkModeIntegrationIcon string
	FullPage                string
	ColorSVG                string
	WhiteSVG                string
}

func GetEntries(r *csv.Reader, columnNames []string) ([]map[string]string, error) {
	final := make([]map[string]string, 0)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	mapColumns := make(map[string]bool)
	for _, c := range columnNames {
		mapColumns[c] = true
	}
	var internalIndex map[int]string
	var start bool
	var startReadingFrom int
	for i, record := range records {
		start, internalIndex = startRecording(record, mapColumns)
		if start {
			startReadingFrom = i + 1
			break
		}
	}
	if internalIndex == nil {
		return nil, fmt.Errorf("could not populate entries")
	}
	for startReadingFrom < len(records) {
		record := records[startReadingFrom]
		temp := make(map[string]string)
		for i, cell := range record {
			key := internalIndex[i]
			value := cell
			temp[key] = value
		}
		final = append(final, temp)
		startReadingFrom++
	}
	return final, nil
}

func startRecording(record []string, columnNames map[string]bool) (bool, map[int]string) {
	for _, r := range record {
		if columnNames[r] { //even if one column is detected, recording the further rows will start
			m := make(map[int]string)
			for i, r0 := range record {
				m[i] = r0
			}
			return true, m
		}
	}
	return false, nil
}

func DownloadCSV(url string) (string, error) {
	file, err := os.CreateTemp("./", "*.csv")
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	byt, _ := io.ReadAll(resp.Body)
	file.WriteString(string(byt))
	byt, _ = io.ReadAll(file)
	path, _ := filepath.Abs(file.Name())
	return path, nil
}

func (t TemplateAttributes) CreateMarkDown() string {

	markdown := t.FullPage
	markdown = strings.ReplaceAll(markdown, "\r", "\n")
	return markdown
}

func WriteToFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
	// Close the file to save the changes.
	err = file.Close()
	if err != nil {
		panic(err)
	}
	return nil
}

var (
	ColumnNamesToExtract        = []string{"modelDisplayName", "model", "category", "subCategory", "shape", "primaryColor", "secondaryColor", "logoURL", "svgColor", "svgWhite", "Publish?", "CRDs", "component", "svgComplete", "genealogy", "styleOverrides"}
	ColumnNamesToExtractForDocs = []string{"modelDisplayName", "Page Subtitle", "Docs URL", "category", "subCategory", "Feature 1", "Feature 2", "Feature 3", "howItWorks", "howItWorksDetails", "Publish?", "About Project", "Standard Blurb", "svgColor", "svgWhite", "Full Page", "model"}
	PrimaryColumnName           = "model"
	OutputPath                  = ""
	ExcludeDirs                 = []string{"relationships", "policies"}
)

//var System string

//const (
//	SVG_WIDTH  = 20
//	SVG_HEIGHT = 20
//)

func main() {
	url := "https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub?gid=0&single=true&output=csv"
	integer := 1
	filep, err := DownloadCSV(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
		return
	}
	csvReader := csv.NewReader(file)
	output, err := GetEntries(csvReader, ColumnNamesToExtractForDocs)
	if err != nil {
		log.Fatal(err)
		return
	}
	file.Close()
	os.Remove(file.Name())
	docsUpdater(output, integer)

}

// returns the index of column. Returns -1 if doesn't exist
//func contains(key string, col []string) int {
//	for i, n := range col {
//		if n == key {
//			return i
//		}
//	}
//	return -1
//}

// For Docs:: entries with empty Component field are preferred as they are considered general
// In other words, the absence of a component name indicates that a given row is a Model-level entry.
// And that for docs/websites updates, the values found in this row should be used to represent the
// integration overall (whether there is 1 or many 10s of components contained in the package / in the integration).
func cleanupDuplicatesAndPreferEmptyComponentField(out []map[string]string, groupBykey string) (out2 []map[string]string) {
	keyToEntry := make(map[string]map[string]string)
	for _, o := range out {
		gkey := o[groupBykey]
		//If the row with given gkey is encountered for the first time, or the given row already exists but with a non-empty component field then use the new entry.
		//This logic will prioritize empty component fields to not be overriden
		if keyToEntry[gkey] == nil || keyToEntry[gkey]["component"] != "" {
			keyToEntry[gkey] = o
		}

	}
	for _, entry := range keyToEntry {
		out2 = append(out2, entry)
	}
	return out2
}

func docsUpdater(output []map[string]string, integer int) {
	//if len(os.Args) < 7 {
	//	log.Fatal("docsUpdater: invalid number of arguments; missing website and docs path")
	//	return
	//}
	//	pathToIntegrationsLayer5 := os.Args[4]
	//pathToIntegrationsMeshery := os.Args[5]
	//pathToIntegrationsMesheryDocs := os.Args[6]
	//updateOnlyPublished := true
	//if len(os.Args) > 6 {
	//	if os.Args[6] == "--published-only" {
	//		updateOnlyPublished = true
	//	}
	//}
	output = cleanupDuplicatesAndPreferEmptyComponentField(output, "model")
	//mesheryDocsJSON := "const data = ["
	for _, out := range output {
		var t TemplateAttributes

		publishValue, err := strconv.ParseBool(out["Publish?"])
		if err != nil {
			publishValue = false
		}
		if !publishValue {
			continue
		}
		for key, val := range out {

			switch key {
			case "modelDisplayName":
				t.Title = val
			case "model":
				t.ModelName = val
			case "Page Subtitle":
				t.Subtitle = val
			case "Docs URL":
				t.DocURL = val
			case "category":
				t.Category = val
			case "subCategory":
				t.Subcategory = val
			case "howItWorks":
				t.HowItWorks = val
			case "hotItWorksDetails":
				t.HowItWorksDetails = val
			case "Publish?":
				t.Published = val
			case "About Project":
				t.AboutProject = val
			case "Standard Blurb":
				t.StandardBlurb = val
			case "Full Page":
				t.FullPage = val
			case "svgColor":
				t.ColorSVG = val
			case "svgWhite":
				t.WhiteSVG = val
			}
			integer++
			//if integer <= 13 {
			//	WriteToFile("info.txt", t.FullPage)
			//	integer += 1
			//}

		}
		//	t.FeatureList = "[" + strings.Join([]string{out["Feature 1"], out["Feature 2"], out["Feature 3"]}, ",") + "]"
		//	t.WorkingSlides = `[
		//			../_images/meshmap-visualizer.png,
		//			../_images/meshmap-designer.png]`
		//
		//	md := t.CreateMarkDown()
		//	jsonItem := t.CreateJSONItem()
		//	mesheryDocsJSON += jsonItem + ","
		//	modelName := strings.TrimSpace(out["model"])
		//	pathToIntegrationsLayer5, _ := filepath.Abs(filepath.Join("../../../", pathToIntegrationsLayer5, modelName))
		//	pathToIntegrationsMeshery, _ := filepath.Abs(filepath.Join("../../../", pathToIntegrationsMeshery))
		//	pathToIntegrationsMesheryDocs, _ := filepath.Abs(filepath.Join("../../", pathToIntegrationsMesheryDocs, "assets/img/meshmodel/", modelName))
		//	err = os.MkdirAll(pathToIntegrationsLayer5, 0777)
		//	if err != nil {
		//		panic(err)
		//	}
		//	_ = pkg.WriteToFile(filepath.Join(pathToIntegrationsLayer5, "index.mdx"), md)
		//	svgcolor := out["svgColor"]
		//	svgwhite := out["svgWhite"]
		//
		//	// Write SVGs to Layer5 docs
		//	err = os.MkdirAll(filepath.Join(pathToIntegrationsLayer5, "icon", "color"), 0777)
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsLayer5, "icon", "color", modelName+"-color.svg"), svgcolor) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//	err = os.MkdirAll(filepath.Join(pathToIntegrationsLayer5, "icon", "white"), 0777)
		//	if err != nil {
		//		panic(err)
		//	}
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsLayer5, "icon", "white", modelName+"-white.svg"), svgwhite) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	// Write SVGs to Meshery website
		//	err = os.MkdirAll(filepath.Join(pathToIntegrationsMeshery, "../", "images"), 0777)
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsMeshery, "../", "images", modelName+"-color.svg"), svgcolor) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsMeshery, "../", "images", modelName+"-white.svg"), svgwhite) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	// Write SVGs to Meshery docs
		//	err = os.MkdirAll(filepath.Join(pathToIntegrationsMesheryDocs), 0777)
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsMesheryDocs, modelName+"-color.svg"), svgcolor) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	err = pkg.WriteSVG(filepath.Join(pathToIntegrationsMesheryDocs, modelName+"-white.svg"), svgwhite) //CHANGE PATH
		//	if err != nil {
		//		panic(err)
		//	}
		//}

		//mesheryDocsJSON = strings.TrimSuffix(mesheryDocsJSON, ",")
		//mesheryDocsJSON += "]; export default data"
		//if err := pkg.WriteToFile(filepath.Join("../../../", pathToIntegrationsMeshery, "data.js"), mesheryDocsJSON); err != nil {
		//	log.Fatal(err)
		//}
		//
		//if err := pkg.WriteToFile(filepath.Join("../../", pathToIntegrationsMesheryDocs, "_data/integrations/", "data.js"), mesheryDocsJSON); err != nil {
		//	log.Fatal(err)
		//}

	}
	fmt.Println(integer)
}
