//rft -- register type file
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/xLib"

	"golang.org/x/sys/windows/registry"
)

//testInputParams - test input params
//4 parameters required
//5-th parameter optional
func testInputParams() error {
	//rft (i) (infArc) "описание архива" "notepad++ %1" "c:\windows\ShellNew\(arcInfo).(i)"
	//0    1     2             3              4                   5
	if len(os.Args) < 2 {
		return errors.New("required 4 parameters")
	}
	if !xLib.StrIsPrintRune(os.Args[1]) {
		return errors.New("1 arg: '" + os.Args[1] + "' contain non printable symbols!")
	}
	if xLib.StrContainBackSlash(os.Args[1]) {
		return errors.New("1 arg: '" + os.Args[1] + "' contain symbols '\\'")
	}
	if !xLib.StrIsPrintRune(os.Args[2]) {
		return errors.New("2 arg: '" + os.Args[2] + "' contain non printable symbols!")
	}
	if xLib.StrContainBackSlash(os.Args[2]) {
		return errors.New("2 arg: '" + os.Args[2] + "' contain symbols '\\'")
	}
	if !xLib.StrIsPrintRune(os.Args[3]) {
		return errors.New("3 arg: '" + os.Args[3] + "' contain non printable symbols!")
	}
	if xLib.StrContainBackSlash(os.Args[3]) {
		return errors.New("3 arg: '" + os.Args[3] + "' contain symbols '\\'")
	}

	if len(os.Args) > 5 {
		testFile, err := os.Open(os.Args[5]) //Open file to READ
		defer testFile.Close()
		if err != nil {
			return errors.New("Attantion file: " + os.Args[5] + " can't open to read!")
		}
	}

	return nil
}

func main() {

	err := testInputParams()
	if err != nil {
		fmt.Println(err)
		fmt.Println("using: >rft ext regKey newFielName openCommand pathToTmplt")
		os.Exit(1)
	}
	fmt.Println(len(os.Args))
	fmt.Println("1 arg: '" + os.Args[1] + "' ok.")
	fmt.Println("2 arg: '" + os.Args[2] + "' ok.")
	fmt.Println("3 arg: '" + os.Args[3] + "' ok.")
	fmt.Println("4 arg: '" + os.Args[4] + "' ok.")
	pathToTemplate := ""
	if len(os.Args) > 5 {
		fmt.Println("5 arg: '" + os.Args[5] + "' ok.")
		pathToTemplate = os.Args[5]
	}
	err = createFileAssocRegKey(os.Args[1],
		os.Args[2],
		os.Args[3],
		os.Args[4],
		pathToTemplate)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("ext added successful")
	}
}

//parameters: ext, regKeyName, fTypeName & openCommand can't be empty
//ext:			file extention to want register in system
//regKeyName:	name of registry key in HKEY_CLASSES_ROOT
//fTypeName:	friendly type name, using on create new file
//openCommand:	command to open file
//newFile:		the path to the template of the new file of this type, like - C:\Windows\ShellNew\arcInfo.ext
func createFileAssocRegKey(ext, regKeyName, fTypeName, openCommand, newFile string) error {
	if ext == "" {
		return errors.New("parameter 'ext' is empty")
	}
	if regKeyName == "" {
		return errors.New("parameter 'regKeyName' is empty")
	}
	if openCommand == "" {
		return errors.New("parameter 'openCommand' is empty")
	}

	key, _, err := registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\.`+ext,
		registry.ALL_ACCESS)
	defer key.Close()
	if err != nil {
		return err
	}
	err = key.SetStringValue("", regKeyName)
	if err != nil {
		return err
	}
	key.Close()

	key, _, err = registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\.`+ext+`\ShellNew`,
		registry.ALL_ACCESS)
	if newFile != "" {
		err = key.SetStringValue("FileName", newFile)
		if err != nil {
			return err
		}
	}
	key.Close()

	key, _, err = registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\`+regKeyName,
		registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	//key.SetStringValue("", "text document")
	err = key.SetStringValue("FriendlyTypeName", fTypeName)
	if err != nil {
		return err
	}
	key.Close()

	key, _, err = registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\`+regKeyName+`\shell`,
		registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	key, _, err = registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\`+regKeyName+`\shell\open`,
		registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	key, _, err = registry.CreateKey(registry.CURRENT_USER,
		`SOFTWARE\Classes\`+regKeyName+`\shell\open\command`,
		registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	err = key.SetStringValue("", openCommand)
	if err != nil {
		return err
	}
	return nil
}
