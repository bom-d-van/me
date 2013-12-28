#!/usr/bin/env bash

me=/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/me

mv -f $me/thoughts_creation_file.txt{,.backup}
for f in $(find /Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/me/thoughts)
do
	echo "$f://$(GetFileInfo -d $f)" >> $me/thoughts_creation_file.txt
done

sed -i "" "s/\/Users\/bom_d_van\/Code\/go\/workspace\/src\/github.com\/bom-d-van\/me\/thoughts//g" $me/thoughts_creation_file.txt
