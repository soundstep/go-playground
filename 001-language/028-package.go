/*
http://golangtutorials.blogspot.co.uk/2011/10/go-packages-and-goinstall-creating-and.html

Code and package structure
We shall now write a illustrative code example for which we shall use identifiable names for the folder, package name, and file name. Please note that this is not the convention used in Go. Typically in Go, one would have the package name the same as the directory it is in. So, for now for the purpose of learning, go ahead and do the following.

* Under $GOPATH/src, create a directory pkgdirname. So I now have $GOPATH/src/pkgdirname
* Within $GOPATH/src/pkgdirname, create the file fileinpkgdir.go. So I now have $GOPATH/src/pkgdirname/fileinpkgdir.go

* Add the code below into the file fileinpkgdir.go
Full program - $GOPATH/src/pkgdirname/fileinpkgdir.go
package pkgnameinfile

func Add2Ints(i, j int) int {
    return i + j
}

At this point, if you want, you can check whether your syntax is right by compiling the file thus: 6g fileinpkgdir.go. But it is not necessary since goinstall will also do the compilation for you.

* Go back to the src directory and type the following command: goinstall pkgdirname.
* Now check the contents of the directory at GOPATH and you should now be seeing a new file there pkg/"system_arch"/pkgdirname.a. Since my system is a linux 64 bit machine, my directory is
$GOPATH/pkg/linux_amd64/pkgdirname.a.

To summarize how that happened, goinstall will
* automatically compile go files
* look for GOPATH
* find directory name following GOPATH/src, and retain that as the package directory structure for the target package
* create a folder GOPATH/pkg/"system_arch" if it is not already present
* and copy into it the compiled package with the name of the directory.

Notice here that the name used for the package is not the file name (fileinpkgdir.go), nor the name of the package that we put within the file (package pkgnameinfile). Instead it is the name of the directory that the file was in. I repeat that I named these items such only to distinguish for you which name is automatically taken by goinstall. Conventionally, you would have the same name for the directory as you would have in the package statement within the code.
*/
package main

import (
	"fmt"
	"pkgdirname" //use the file name without the ".a"
)

func main() {
	fmt.Println(pkgnameinfile.Add2Ints(3, 4)) //use the name of package in the file
}
