go_format_recursively(){
	
	count=`ls -1 *.go 2>/dev/null | wc -l`
	if [ $count != 0 ]
	then 
		go fmt *.go
	fi 
	
	for d in "$@" ; do
		if [ -d "$d" ]; then
			cd $d
			go_format_recursively *
			cd ..
		fi
	done
}
go_format_recursively *
