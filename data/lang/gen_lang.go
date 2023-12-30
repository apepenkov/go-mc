//go:build generate
// +build generate

// This program can automatically download language.json file and convert into .go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"
)

// language=gohtml
var langTmpl = `// Code generated by downloader.go; DO NOT EDIT.

package {{.Name}}
{{if ne .Name "en_us"}}
import "github.com/apepenkov/go-mc/chat"

func init() { chat.SetLanguage(Map) }
{{end}}
var Map = {{.LangMap | printf "%#v"}}
`

//go:generate go run $GOFILE
//go:generate go fmt ./...
func main() {
	if len(os.Args) == 2 {
		fmt.Println("generating en-us lang")
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		readLang("en_us", f)
		return
	} else {
		fmt.Println("generating langs except en-us")
		fmt.Println("WARN: You should also set the secondary argument to en-us's json file")
	}

	versionURL, err := assetIndexURL()
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(versionURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var list struct {
		Objects map[string]struct {
			Hash string `json:"hash"`
			Size int64  `json:"size"`
		} `json:"objects"`
	}

	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		panic(err)
	}

	tasks := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range tasks {
				v := list.Objects[i]
				if strings.HasPrefix(i, "minecraft/lang/") {
					name := i[len("minecraft/lang/") : len(i)-len(".json")]
					lang(name, v.Hash)
				}
			}
		}()
	}
	for i := range list.Objects {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}

func lang(name, hash string) {
	// download language
	LangURL := "https://resources.download.minecraft.net/" + hash[:2] + "/" + hash
	fmt.Println(name, ":", LangURL)
	resp, err := http.Get(LangURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	readLang(name, resp.Body)
}

// read one language translation
func readLang(name string, r io.Reader) {
	var LangMap map[string]string
	err := json.NewDecoder(r).Decode(&LangMap)
	if err != nil {
		panic(err)
	}
	trans(LangMap)

	pName := strings.ReplaceAll(name, "_", "-")

	// mkdir
	err = os.Mkdir(pName, 0o777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	f, err := os.OpenFile(filepath.Join(pName, name+".go"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	genData := struct {
		PkgName string
		Name    string
		LangMap map[string]string
	}{
		PkgName: pName,
		Name:    name,
		LangMap: LangMap,
	}

	tmpl := template.Must(template.New("").Parse(langTmpl))
	if err := tmpl.Execute(f, genData); err != nil {
		panic(err)
	}
}

var javaN = regexp.MustCompile(`%[0-9]\$s`)

// Java use %2$s to refer to the second arg, but Golang use %2s, so we need this
func trans(m map[string]string) {
	// replace "%[0-9]\$s" with "%[0-9]s"
	for i := range m {
		c := m[i]
		if javaN.MatchString(c) {
			m[i] = javaN.ReplaceAllStringFunc(c, func(s string) string {
				var index int
				_, err := fmt.Sscanf(s, "%%%d$s", &index)
				if err != nil {
					panic(err)
				}
				return fmt.Sprintf("%%[%d]s", index)
			})
		}
	}
}

func assetIndexURL() (string, error) {
	// Pseudo code for get versionURL:
	// $manifest = {https://piston-meta.mojang.com/mc/game/version_manifest_v2.json}
	// $latest = $manifest.latest.release
	// $versionURL = {$manifest.versions[where .id == $latest ].url}
	// $assetIndexURL = $version.assetIndex.url
	var manifest struct {
		Latest struct {
			Release string `json:"release"`
		} `json:"latest"`
		Versions []struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"versions"`
	}

	manifestRes, err := http.Get("https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")
	if err != nil {
		return "", fmt.Errorf("could not reach version manifest: %w", err)
	}
	defer manifestRes.Body.Close()

	if err := json.NewDecoder(manifestRes.Body).Decode(&manifest); err != nil {
		return "", fmt.Errorf("could not decode manifest JSON: %w", err)
	}

	var versionURL string
	for _, v := range manifest.Versions {
		if manifest.Latest.Release == v.ID {
			versionURL = v.URL
			break
		}
	}
	if versionURL == "" {
		return "", errors.New("could not determine versionURL")
	}

	var version struct {
		AssetIndex struct {
			URL string `json:"url"`
		} `json:"assetIndex"`
	}

	versionRes, err := http.Get(versionURL)
	if err != nil {
		return "", fmt.Errorf("could not reach versionURL: %w", err)
	}
	defer versionRes.Body.Close()

	if err := json.NewDecoder(versionRes.Body).Decode(&version); err != nil {
		return "", fmt.Errorf("could not decode version JSON: %w", err)
	}

	return version.AssetIndex.URL, nil
}
