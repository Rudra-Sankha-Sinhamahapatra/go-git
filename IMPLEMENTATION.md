

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

# Reads compressed object via cat file

``` go run ./cmd/gogit cat-file -p 56294e459e92315a1f5fa0418e14e8394b54c04e
```

compressed object size: 365
object found at: .gogit/objects/56/294e459e92315a1f5fa0418e14e8394b54c04e

```
 go run ./cmd/gogit cat-file -p 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c
 ```

error: cannot read compressed object 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c: open .gogit/objects/2c/f24dba5fb0a30e26e83b2ac5b9e29e1b161e5c: no such file or directory

