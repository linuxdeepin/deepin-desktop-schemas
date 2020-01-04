package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"pkg.deepin.io/lib/keyfile"
)

var optArch string

func init() {
	flag.StringVar(&optArch, "arch", "", "")
}

var products = []string{"pro", "server"}

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	files, err := getArchFiles(optArch)
	if err != nil {
		log.Fatal(err)
	}
	err = combineFiles(files, "90_deepin_base.gschema.override")
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		productFiles, err := getProductFiles(optArch, product)
		if err != nil {
			log.Fatal(err)
		}

		err = combineFiles(productFiles, product+"-override")
		if err != nil {
			log.Fatal(err)
		}
	}
}

const extOverride = ".override"

func getDirFiles(dir string) (files []string, err error) {
	dir = filepath.Join("overrides", dir)
	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}

	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			continue
		}
		if filepath.Ext(fileInfo.Name()) == extOverride {
			files = append(files, filepath.Join(dir, fileInfo.Name()))
		}
	}
	return
}

func getArchFiles(arch string) (files []string, err error) {
	// 返回　common/*.override, $arch/*.override
	commonFiles, err := getDirFiles("common")
	if err != nil {
		return
	}
	files = append(files, commonFiles...)

	archFiles, err := getDirFiles(arch)
	if err != nil {
		return
	}
	files = append(files, archFiles...)
	return
}

func getProductFiles(arch, product string) (files []string, err error) {
	// 返回　common/$product/*.override，$arch/$product/*.override
	commonFiles, err := getDirFiles(filepath.Join("common", product))
	if err != nil {
		return
	}
	files = append(files, commonFiles...)

	archFiles, err := getDirFiles(filepath.Join(arch, product))
	if err != nil {
		return
	}
	files = append(files, archFiles...)
	return
}

func getValue(value string, kf *keyfile.KeyFile, section0 string) string {
	const refPrefix = "@ref:"
	if strings.HasPrefix(value, refPrefix) {
		value = strings.TrimPrefix(value, refPrefix)
		fields := strings.Split(value, ";")

		refMap := make(map[string]string)
		// 支持的键有：FILE,SECTION,KEY
		for _, field := range fields {
			parts := strings.SplitN(field, "=", 2)
			if len(parts) != 2 {
				continue
			}
			refMap[parts[0]] = parts[1]
		}

		// TODO: refMap["FILE"]

		section := refMap["SECTION"]
		if section == "" {
			section = section0
		}

		key := refMap["KEY"]
		if key == "" {
			log.Println("WARN: key is empty")
			return ""
		}

		newVal, _ := kf.GetValue(section, key)
		return newVal
	}
	return value
}

func combineFiles(inputFiles []string, outputFile string) (err error) {
	log.Printf("inputFiles: %+v -> outputFile: %s\n", inputFiles, outputFile)
	combinedKf := keyfile.NewKeyFile()
	outputFile = filepath.Join("result", outputFile)

	for _, inputFile := range inputFiles {
		kf := keyfile.NewKeyFile()
		err = kf.LoadFromFile(inputFile)
		if err != nil {
			return
		}
		sections := kf.GetSections()
		for _, section := range sections {
			sectionMap, _ := kf.GetSection(section)
			for key, val := range sectionMap {
				val = getValue(val, kf, section)
				combinedKf.SetValue(section, key, val)
			}
		}
	}

	outputDir := filepath.Dir(outputFile)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return
	}

	err = combinedKf.SaveToFile(outputFile)
	if err != nil {
		return
	}

	return
}

