# keyJSONFinder
find JSON files that contains specific keywords

## Output
Matched JSON files will be put in `output` folder

## How to run it

```shell script
./keyJSONFinder -input=FOLDER_NAME -keywords=KEYWORDS -top=TOP_K
```

- `-input=` accepts a folder name.
- `-keywords=` accepts a list of keywords, concatenated by `,`
- `-top=` accepts a number, indicating how many files to be matched. Default is `1`.

e.g.:


```shell script
./keyJSONFinder -input=/home/test -keywords="Ijk_Call,Ijk_Sys_syscall" -top=10
```

- Input folder is `/home/test`
- Keywords are "Ijk_Call", "Ijk_Sys_syscall"
- First 10 files will be returned.