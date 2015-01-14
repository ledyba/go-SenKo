#! bash
for file in data/*.csv
do
    iconv -f euc-jp -t utf8 -o "$file.utf8" "$file"
done

