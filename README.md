Tree utility.

Prints tree of catalogs and files (the last one if is used option -f).

```
go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
	
go run main.go .
└───testdata
	├───project
	├───static
	│	├───css
	│	├───html
	│	└───js
	└───zline
```