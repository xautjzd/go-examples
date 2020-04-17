package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	hotSearchWord  string
	soarSearchWord string
	srcFile        string
	dstFile        string
)

const (
	csvExt string = ".csv"
)

func init() {
	flag.StringVar(&hotSearchWord, "hotWord", "", "热搜词 csv 文件名称")
	flag.StringVar(&soarSearchWord, "soarWord", "", "飙升词 csv 文件名称")
	flag.StringVar(&srcFile, "src", "", "待处理 csv 文件名称")
	flag.StringVar(&dstFile, "dst", "", "处理完成后生成的 csv 文件名称")
}

func main() {
	flag.Parse()

	if !strings.HasSuffix(hotSearchWord, csvExt) &&
		!strings.HasSuffix(soarSearchWord, csvExt) &&
		!strings.HasSuffix(srcFile, csvExt) &&
		!strings.HasSuffix(dstFile, csvExt) {
		flag.Usage()
		os.Exit(0)
	}

	hotWords, err := readCSV(hotSearchWord)
	if err != nil {
		fmt.Printf("读取热搜词文件失败: %v\n", err)
		os.Exit(0)
	}
	// key 为热搜词，value 为排名
	hot := make(map[string]string, len(hotWords))
	for _, line := range hotWords {
		hot[line[0]] = line[2]
	}

	soar, err := readCSV(soarSearchWord)
	if err != nil {
		fmt.Printf("读取飙升词文件失败: %v\n", err)
		os.Exit(0)
	}
	bs := make(map[string]string, len(soar))
	for _, line := range soar {
		bs[line[0]] = line[2]
	}

	titles, err := readCSV(srcFile)
	if err != nil {
		fmt.Printf("读取待处理文件失败: %v\n", err)
		os.Exit(0)
	}

	target, err := os.Create(dstFile)
	if err != nil {
		fmt.Printf("创建 %s 文件失败: %v\n", dstFile, err)
		os.Exit(0)
	}
	defer target.Close()
	writer := csv.NewWriter(target)

	for i, line := range titles {
		// 包含热搜词
		var hotCol string
		for hk, hv := range hot {
			if strings.Contains(line[0], hk) {
				if hotCol == "" {
					hotCol = fmt.Sprintf("%s %s", hk, hv)
				} else {
					hotCol = hotCol + ", " + fmt.Sprintf("%s %s", hk, hv)
				}
			}
		}

		// 包含飙升词
		var bsCol string
		for bk, bv := range bs {
			if strings.Contains(line[0], bk) {
				if bsCol == "" {
					bsCol = fmt.Sprintf("%s %s", bk, bv)
				} else {
					bsCol = bsCol + ", " + fmt.Sprintf("%s %s", bk, bv)
				}
			}
		}

		newLine := line
		if i == 0 {
			newLine = append(newLine, "热搜", "飙升")
		} else {

			newLine = append(newLine, hotCol, bsCol)
		}
		writer.Write(newLine)
	}
}

func readCSV(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.LazyQuotes = true
	lines, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
