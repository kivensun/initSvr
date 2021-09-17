package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type NetInfo struct {
	Type               string
	Proxy_Method       string
	Browser_Only       string
	BootProto          string
	Defroute           string
	IPV4_Failure_Fatal string
	IPV6Init           string
	IPV6_AutoConf      string
	IPV6_Defroute      string
	IPV6_Failure_Fatal string
	IPV6_Addr_Gen_Mode string
	Name               string
	UUID               string
	Device             string
	OnBoot             string
	IPAddr             string
	Prefix             string
	Gateway            string
	DNS                []string
}

func GetOSVersion() (string, error) {
	vs, err := os.ReadFile("./redhat-release")
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(vs), "6.") {
		return "6", nil
	} else if strings.Contains(string(vs), "7.") {
		return "7", nil
	} else {
		return "", errors.New("未知Centos版本")
	}
}

func readC6NetInfo() {

}

func readC7NetInfo() (netFiles []string, err error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "ifcfg") {
			netFiles = append(netFiles, file.Name())
		}
	}
	return netFiles, nil
}

func readNetInfo(vs string) ([]string, error) {
	switch vs {
	case "6":
		fmt.Println("6")
	case "7":
		files, err := readC7NetInfo()
		if err != nil {
			return nil, err
		}
		return files, nil
	default:
		return nil, errors.New("不正确的系统版本")
	}
	return nil, errors.New("不正确的系统版本")
}

func main() {
	//pwd, _ := os.Getwd()

	// fileInofoList, err := ioutil.ReadDir("/etc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(len(fileInofoList))
	// for i := range fileInofoList {
	// 	fmt.Println(fileInofoList[i].Name())
	// }

	//(&cli.App{}).Run(os.Args)

	osVer, err := GetOSVersion()
	if err != nil {
		log.Fatal(err)
	}
	netFiles, err := readNetInfo(osVer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("请输入你要修改的端口顺序号:")
	for i, file := range netFiles {
		fmt.Printf("%d.%s\n", i+1, file)
	}
	sel := 0
	_, err = fmt.Scanf("%d", &sel)
	if err != nil {
		log.Fatal(err)
	}
	if sel >= len(netFiles) || sel < 1 {
		log.Panic("输入的序号不正确")
	}
	fmt.Printf("你选择的端口是:%s,其配置信息是:\n", netFiles[sel-1])
	//netinfo := NetInfo{}

	f, err := os.Open("./" + netFiles[sel-1])

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)

	for {
		line, err := bfRd.ReadBytes('\n')
		fmt.Print(string(line))
		if err != nil {
			if err == io.EOF {
				//
			}
			log.Fatal(err)
		}

	}

}
