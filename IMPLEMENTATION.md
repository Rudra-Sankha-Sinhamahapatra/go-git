

# For default gogit repo initialization

```
 go run ./cmd/gogit init

```
Initialized empty gogit repository in .gogit/
```
 go run ./cmd/gogit init
```
error: gogit repository already exists


# Reads Hash Object 

```
go run ./cmd/gogit hash-object .gitignore
```

read 574 bytes from .gitignore
header:  blob 574
raw object size:  583
object hash: 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c
object written to: .gogit/objects/2c/f24dba5fb0a30e26e83b2ac5b9e29e1b161e5c