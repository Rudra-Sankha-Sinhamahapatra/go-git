

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
decompressed size: 583
object type: blob
declared size: 574
actual content size: 574

```
 go run ./cmd/gogit cat-file -p 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c
 ```

error: cannot read compressed object 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c: open .gogit/objects/2c/f24dba5fb0a30e26e83b2ac5b9e29e1b161e5c: no such file or directory

# Write Tree

```
 go run ./cmd/gogit write-tree
 ```

object hash: 56294e459e92315a1f5fa0418e14e8394b54c04e
object hash: b6af5f1dc42e5a2882a2af36fc237f6b10280c56
object hash: a5c93a920b4ca10a8bd9ebdf359b57ec1600bfd3
object hash: ea5fe17da53fa122c7e3d22bbb0dc8223643d7d8
object hash: daf74f6896b5039cb2e481a5f3a136164aa2cd4c
object written to: .gogit/objects/da/f74f6896b5039cb2e481a5f3a136164aa2cd4c
tree written with hash: daf74f6896b5039cb2e481a5f3a136164aa2cd4c
object hash: ba7470c22d8fc4e1199d9d0b8d12f995ae7c3b97
object written to: .gogit/objects/ba/7470c22d8fc4e1199d9d0b8d12f995ae7c3b97
tree written with hash: ba7470c22d8fc4e1199d9d0b8d12f995ae7c3b97
object hash: 921a86cce097adb34953b4f3bf410055a4f0c52f
object hash: 20eaef3b03749979f0ea8b31a3a6857c208d347b
object hash: 988ae7f9f2161e306a2122bb5db7bebd8c0474b0
object hash: 615778b217873062420b3e83db391252579e21f1
object hash: 175c683e7201bc1c175d58938edbebcdedcb1bc7
object hash: 5ebd679a34db9307bbf1f7b3dadd5f83855c7873
object written to: .gogit/objects/5e/bd679a34db9307bbf1f7b3dadd5f83855c7873
tree written with hash: 5ebd679a34db9307bbf1f7b3dadd5f83855c7873
object hash: 8078de015ab9b22ef3768345b553857e285260f8
object written to: .gogit/objects/80/78de015ab9b22ef3768345b553857e285260f8
tree written with hash: 8078de015ab9b22ef3768345b553857e285260f8
object hash: e011ede034cd14ec0a2185feca7ad390a69366c4
object written to: .gogit/objects/e0/11ede034cd14ec0a2185feca7ad390a69366c4
tree written with hash: e011ede034cd14ec0a2185feca7ad390a69366c4
root tree hash: e011ede034cd14ec0a2185feca7ad390a69366c4


