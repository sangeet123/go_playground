#bin/bash
for d in */ ; do
    cd $d
    for file in *; do
    	go test -v $file
    done
    cd ..
done