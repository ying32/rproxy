/*
   这是一个自己用的构建工具，一个一个的构建太麻烦了。
   这里主要整合编译，打包相应平台的
*/

package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	switch runtime.GOOS {
	case "windows":
		windowsPkg()

	case "linux":
		linuxPkg()

	case "darwin":
		darwinPkg()
	}
}

func windowsPkg() {

	fmt.Println("编译rproxy-win64-GUI")
	if executeBash("build-win64-GUI.bat") == nil {
		fmt.Println("打包rproxy-win64-GUI")
		createZipFile("rproxy-win64-GUI.zip", true)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-win64")
	if executeBash("build-win64.bat") == nil {
		fmt.Println("打包rproxy-win64")
		createZipFile("rproxy-win64.zip", false)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-win32-GUI")
	if executeBash("build-win32-GUI.bat") == nil {
		fmt.Println("打包rproxy-win32-GUI")
		createZipFile("rproxy-win32-GUI.zip", true)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-win32")
	if executeBash("build-win32.bat") == nil {
		fmt.Println("打包rproxy-win32")
		createZipFile("rproxy-win32.zip", false)
	}
}

func linuxPkg() {

	fmt.Println("编译rproxy-linux64-GUI")
	if executeBash("build-linux64-GUI.sh") == nil {
		fmt.Println("打包rproxy-linux64-GUI")
		createZipFile("rproxy-linux64-GUI.zip", true)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-linux64")
	if executeBash("build-linux64.sh") == nil {
		fmt.Println("打包rproxy-linux64")
		createZipFile("rproxy-linux64.zip", false)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-linux32")
	if executeBash("build-linux32.sh") == nil {
		fmt.Println("打包rproxy-linux32")
		createZipFile("rproxy-linux32.zip", false)
	}
}

func darwinPkg() {

	fmt.Println("编译rproxy-darwin64")
	if executeBash("build-darwin64.sh") == nil {
		fmt.Println("打包rproxy-darwin")
		createZipFile("rproxy-darwin.zip", false)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-darwin32-GUI")
	if executeBash("build-darwin32-GUI.sh") == nil {
		fmt.Println("打包rproxy-darwin32-GUI")
		createZipFile("rproxy-win32-GUI.zip", true)
	}
	fmt.Println("------------------------------")

	fmt.Println("编译rproxy-darwin32")
	if executeBash("build-darwin32.sh") == nil {
		fmt.Println("打包rproxy-darwin32")
		createZipFile("rproxy-darwin32.zip", false)
	}
}

func executeBash(fileName string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", fileName)
	default:
		cmd = exec.Command("sh", "./"+fileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("执行错误：", err)
	}
	return err
}

func createZipFile(zipFileName string, isGUI bool) error {
	f, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()

	zw.RegisterCompressor(zip.Deflate,
		func(out io.Writer) (io.WriteCloser, error) {
			return flate.NewWriter(out, flate.BestCompression)
		})

	compressFile := func(fileName, aliasName string) error {
		ff, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer ff.Close()
		info, _ := ff.Stat()
		header, err := zip.FileInfoHeader(info)
		header.Method = zip.Deflate
		if aliasName != "" {
			header.Name = aliasName
		} else {
			header.Name = info.Name()
		}

		wr, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(wr, ff)
		if err != nil {
			return err
		}
		return nil
	}

	// 复制文档和配置
	compressFile("../README.md", "")
	compressFile("../conf/config.cfg", "conf/config.cfg")
	compressFile("../conf/confighttps.cfg", "conf/confighttps.cfg")

	// 复制可执行文件
	exeExt := ""
	if runtime.GOOS == "windows" {
		exeExt = ".exe"
	}
	fnSuffix := ""
	if isGUI {
		fnSuffix = "_GUI"
	}
	compressFile("../rproxy"+fnSuffix+exeExt, "rproxy"+exeExt)
	// 复制动态链接库
	if isGUI {
		switch runtime.GOOS {
		case "windows":
			liblclPath := "F:\\Golang\\src\\github.com\\ying32\\govcl\\Librarys\\liblcl"
			if runtime.GOARCH == "386" {
				compressFile(liblclPath+"\\win32\\liblcl.dll", "")
			} else if runtime.GOARCH == "amd64" {
				compressFile(liblclPath+"\\win64\\liblcl.dll", "")
			}
		case "linux":
			if runtime.GOARCH == "amd64" {
				compressFile("/usr/lib/liblcl.so", "")
			}
		case "darwin":
			if runtime.GOARCH == "386" {

				// 产生一个app
				//pkgMacOSApp("../rproxy")
			}
		}
	}

	//zw.Flush()

	return nil
}

// --- macOS下的

const (
	infoplist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>zh_CN</string>
	<key>CFBundleExecutable</key>
	<string>%s</string>
	<key>CFBundleName</key>
	<string>%s</string>
	<key>CFBundleIdentifier</key>
	<string>ying32.%s</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleSignature</key>
	<string>proj</string>
	<key>CFBundleShortVersionString</key>
	<string>0.1</string>
	<key>CFBundleVersion</key>
	<string>1</string>
	<key>CSResourcesFileMapped</key>
	<true/>
	<key>CFBundleIconFile</key>
	<string>%s.icns</string>
	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeRole</key>
			<string>Viewer</string>
			<key>CFBundleTypeExtensions</key>
			<array>
				<string>*</string>
			</array>
			<key>CFBundleTypeOSTypes</key>
			<array>
				<string>fold</string>
				<string>disk</string>
				<string>****</string>
			</array>
		</dict>
	</array>
	<key>NSHighResolutionCapable</key>
	<true/>
    <key>NSHumanReadableCopyright</key>
	<string>copyright 2017-2018 ying32.com</string>
</dict>
</plist>`
)

var (
	pkgInfo = []byte{0x41, 0x50, 0x50, 0x4C, 0x3F, 0x3F, 0x3F, 0x3F, 0x0D, 0x0A}
)

func copyFile(src, dest string) error {
	filedest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer filedest.Close()
	filesrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer filesrc.Close()
	_, err = io.Copy(filedest, filesrc)
	return err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getdylib() string {
	env := os.Getenv("GOPATH")
	if env == "" {
		return ""
	}
	for _, s := range strings.Split(env, ":") {
		s += "/bin/liblcl.dylib"
		if fileExists(s) {
			return s
		}
	}
	return ""
}

func pkgMacOSApp(exeFileName string) error {

	execName := "rproxy"
	macContentsDir := execName + ".app/Contents"
	macOSDir := macContentsDir + "/MacOS"
	macResources := macContentsDir + "/Resources"
	execFile := macOSDir + "/" + execName
	if !fileExists(macOSDir) {
		if err := os.MkdirAll(macOSDir, 0755); err != nil {
			return err
		}
	}

	if !fileExists(macResources) {
		os.MkdirAll(macResources, 0755)
	}

	copyFile("../imgs/rproxy.icns", macResources+"/rproxy.icns")

	liblclFileName := macOSDir + "/liblcl.dylib"
	if !fileExists(liblclFileName) {
		libFileName := getdylib()
		if fileExists(libFileName) {
			copyFile(libFileName, liblclFileName)
		}
	}

	plistFileName := macContentsDir + "/Info.plist"
	if !fileExists(plistFileName) {
		ioutil.WriteFile(plistFileName, []byte(fmt.Sprintf(infoplist, execName, execName, execName, execName)), 0666)
	}

	pkgInfoFileName := macContentsDir + "/PkgInfo"
	if !fileExists(pkgInfoFileName) {
		ioutil.WriteFile(pkgInfoFileName, pkgInfo, 0666)
	}

	copyFile(exeFileName, execFile)

	return nil
}
