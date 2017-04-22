#bin/bash
for d in */ ; do
    cd $d
    go test -v
    cd ..
done