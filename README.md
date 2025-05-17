# My first Abstract Syntax Tree (AST)


## Hi !!, this is a simple (ast) with a recursive-descent expression compilation for parse my family generation it's just an excercie in my jurney for learn how compilers works and make my own some day. 


### How it's works

The scructure is like a ```json``` but you can add multiple names the first level it's the first generation

#### Ex:
```json
{
    "Odinn"
    "Zeus"
}
```

#### For add the Odin's generation:
```json
{
    "Odinn"
    {
        "Emiliano"
        "Atenea"
        "Mixra"
        {
            "Axtrox"
            "Loky"
            {
                "Zeus"
                "Hulk"
            }
        }
    }
    "Zeus"
}
```
#### Output:
```
Members of the 1 generation
        ["Odinn" "Zeus"]
Members of the 2 generation
        ["Emiliano" "Atenea" "Mixra"]
Members of the 3 generation
        ["Axtrox" "Loky"]
Members of the 4 generation
        ["Zeus" "Hulk"]
```

Under the hood the recursion tree it's saved into 
```go
type Node struct {
	data   string
	childs []*Node
}
```
