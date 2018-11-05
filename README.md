# rft - Registrate new file type in Windows

	download: go get -u github.com/softlandia/rft

using

	>rft txt txtFile "text document" "notepad++ %1" "c:\windows\ShellNew\std text doc.txt"  
or

	>rft (i) (infArc) "описание архива" "notepad++ %1" "c:\windows\ShellNew\(arcInfo).(i)"  

parameters 1-4 required  
parameter  5   not required  

(c) softlandia@gmail.com