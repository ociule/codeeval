package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

type Object struct {
	User int
	File int
}

type Permission struct {
	Grant bool
	Read  bool
	Write bool
}

type FilePermissions struct {
	Permissions map[Object]Permission
}

func InitFilePermissions() FilePermissions {
	p := FilePermissions{}
	p.Permissions = make(map[Object]Permission, 1000)
	p.Setup(1, 1, true, true, true)
	p.Setup(1, 2, true, false, true)
	p.Setup(2, 1, false, true, true)
	p.Setup(2, 2, false, false, true)
	p.Setup(2, 3, false, true, false)
	p.Setup(3, 1, true, true, false)
	p.Setup(3, 2, true, false, false)
	p.Setup(3, 3, true, true, false)
	p.Setup(4, 1, true, false, true)
	p.Setup(4, 2, true, true, true)
	p.Setup(4, 3, true, false, false)
	p.Setup(5, 1, false, true, true)
	p.Setup(5, 3, false, false, true)
	p.Setup(6, 1, false, true, false)
	p.Setup(6, 2, false, false, true)
	p.Setup(6, 3, false, true, true)
	return p
}

func (p *FilePermissions) Setup(user, file int, grant, read, write bool) {
	p.Permissions[Object{user, file}] = Permission{Grant: grant, Read: read, Write: write}
}

func checkStringPerm(perm string) bool {
	return perm == "grant" || perm == "read" || perm == "write"
}

func (p *FilePermissions) CheckPerm(user, file int, perm string) bool {
	// perm MUST be one of "grant", "read", "write"
	if !checkStringPerm(perm) {
		panic("Wrong perm")
	}
	permsUserOnFile := p.Permissions[Object{User: user, File: file}]
	switch perm {
	case "grant":
		return permsUserOnFile.Grant
	case "read":
		return permsUserOnFile.Read
	case "write":
		return permsUserOnFile.Write
	}
	return false
}

func (p *FilePermissions) GrantWithString(user, file int, perm string) {
	// perm MUST be one of "grant", "read", "write"
	if !checkStringPerm(perm) {
		panic("Wrong perm")
	}
	permsUserOnFile := p.Permissions[Object{User: user, File: file}]
	switch perm {
	case "grant":
		permsUserOnFile.Grant = true
	case "read":
		permsUserOnFile.Read = true
	case "write":
		permsUserOnFile.Write = true
	}
	p.Permissions[Object{User: user, File: file}] = permsUserOnFile
}

func (p *FilePermissions) Run(in string) string {
	split := strings.Fields(in)
	for _, inst := range split {
		user, file, perm, granted_perm, grantee := 0, 0, "", "", 0
		_, _ = fmt.Sscanf(inst, "user_%d=>file_%d=>%5s=>%s", &user, &file, &perm, &granted_perm)
		if !p.CheckPerm(user, file, perm) {
			return "False"
		}
		if perm == "grant" {
			splitGP := strings.Split(granted_perm, "=>user_")
			granted_perm = splitGP[0]
			_, _ = fmt.Sscanf(splitGP[1], "%d", &grantee)
			//fmt.Println(user, file, perm, granted_perm, grantee)
			p.GrantWithString(grantee, file, granted_perm)
		}
		//fmt.Println(user, file, perm, granted_perm, grantee)
	}
	return "True"
}

func (p *FilePermissions) PPrint() {
	for n := 0; n < 200; n += 1 {
	}
	//fmt.Println(strings.Join(temp, ""))
}

func solve(in string) string {
	p := InitFilePermissions()
	//p.Run()
	return p.Run(in)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(solve(line))
	}
}
