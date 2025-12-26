

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


# Commit Tree

```
 go run ./cmd/gogit commit-tree -m "first commit"
 ```

object hash: 56294e459e92315a1f5fa0418e14e8394b54c04e
object hash: 7c2fe49d1e08a313ac13925b2376ebfbd870d862
object hash: a5c93a920b4ca10a8bd9ebdf359b57ec1600bfd3
object hash: a050a32688507915ef79b7d85121885f6d6abc4b
object hash: d8188bca967cf236f7bf4546660317976c96896d
object written to: .gogit/objects/d8/188bca967cf236f7bf4546660317976c96896d
tree written with hash: d8188bca967cf236f7bf4546660317976c96896d
object hash: 96a093d1ff97c01d84a265db967929210f636f4b
object written to: .gogit/objects/96/a093d1ff97c01d84a265db967929210f636f4b
tree written with hash: 96a093d1ff97c01d84a265db967929210f636f4b
object hash: 921a86cce097adb34953b4f3bf410055a4f0c52f
object hash: 20eaef3b03749979f0ea8b31a3a6857c208d347b
object hash: b288c6993c230b52aa5984e4740cd5ac3a7d9762
object hash: 988ae7f9f2161e306a2122bb5db7bebd8c0474b0
object hash: 615778b217873062420b3e83db391252579e21f1
object hash: 175c683e7201bc1c175d58938edbebcdedcb1bc7
object hash: 1b3cdc6e810fb28480d2c7875007e318837e9787
object written to: .gogit/objects/1b/3cdc6e810fb28480d2c7875007e318837e9787
tree written with hash: 1b3cdc6e810fb28480d2c7875007e318837e9787
object hash: 3bb184ac61d0eba2c7b297f2be47b5926753dabb
object written to: .gogit/objects/3b/b184ac61d0eba2c7b297f2be47b5926753dabb
tree written with hash: 3bb184ac61d0eba2c7b297f2be47b5926753dabb
object hash: 89a1d6d8ea29f087437c15f0f1731d62e9548a28
object written to: .gogit/objects/89/a1d6d8ea29f087437c15f0f1731d62e9548a28
tree written with hash: 89a1d6d8ea29f087437c15f0f1731d62e9548a28
object hash: bda006dc3733322c9b6f68b1573c4933ce52fd90
object written to: .gogit/objects/bd/a006dc3733322c9b6f68b1573c4933ce52fd90
wrote commit object bda006dc3733322c9b6f68b1573c4933ce52fd90
commit hash: bda006dc3733322c9b6f68b1573c4933ce52fd90


## Reading commit message via cat file

``` 
go run ./cmd/gogit cat-file -p 5698904604f413d391d77fa13332587b2e10e00d
```

compressed object size: 142
object found at: .gogit/objects/56/98904604f413d391d77fa13332587b2e10e00d
decompressed size: 226
object type: commit
declared size: 215
actual content size: 215
tree 89a1d6d8ea29f087437c15f0f1731d62e9548a28
author Rudra Sankha Sinhamahapatra <rudra@rudrasankha.com> 1765865215 +0530
committer Rudra Sankha Sinhamahapatra <rudra@rudrasankha.com> 1765865215 +0530

first commit

# Head

```
 go run ./cmd/gogit head
 ```

HEAD points to: refs/heads/main
no commits yet

## To Try it more

```
 go run ./cmd/gogit commit-tree -m "first"
 ```
 
object hash: 56294e459e92315a1f5fa0418e14e8394b54c04e
object hash: ffcf7190d982cf54b2a2152cd65dd22c3db6a944
object hash: a5c93a920b4ca10a8bd9ebdf359b57ec1600bfd3
object hash: 87b52b4bed37c5a604cbefbd7f41d9052e901b86
object hash: fe1c96bcb0e9c85c8d62764aff185021d8d12db1
object written to: .gogit/objects/fe/1c96bcb0e9c85c8d62764aff185021d8d12db1
tree written with hash: fe1c96bcb0e9c85c8d62764aff185021d8d12db1
object hash: 0a3f1c80724bae5efcfe9a256cc19b0aa158c734
object written to: .gogit/objects/0a/3f1c80724bae5efcfe9a256cc19b0aa158c734
tree written with hash: 0a3f1c80724bae5efcfe9a256cc19b0aa158c734
object hash: 921a86cce097adb34953b4f3bf410055a4f0c52f
object hash: 20eaef3b03749979f0ea8b31a3a6857c208d347b
object hash: 69064a23fcd5b6cadd96795712d7a869549dd699
object hash: 988ae7f9f2161e306a2122bb5db7bebd8c0474b0
object hash: 615778b217873062420b3e83db391252579e21f1
object hash: 3fe12d4d38c57cce5a37253db96eb6af4326ee2c
object hash: 175c683e7201bc1c175d58938edbebcdedcb1bc7
object hash: 05bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
object written to: .gogit/objects/05/bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
tree written with hash: 05bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
object hash: cf0ba92794bcb937bb3f838cb3fff90a6dfd8241
object written to: .gogit/objects/cf/0ba92794bcb937bb3f838cb3fff90a6dfd8241
tree written with hash: cf0ba92794bcb937bb3f838cb3fff90a6dfd8241
object hash: 33a0c130da921cc466552ce0d064106cfefeb0c3
object written to: .gogit/objects/33/a0c130da921cc466552ce0d064106cfefeb0c3
tree written with hash: 33a0c130da921cc466552ce0d064106cfefeb0c3
object hash: e21b796e9c9735814596206baa947989e74c860d
object written to: .gogit/objects/e2/1b796e9c9735814596206baa947989e74c860d
wrote commit object e21b796e9c9735814596206baa947989e74c860d
commit hash: e21b796e9c9735814596206baa947989e74c860d


```
go run ./cmd/gogit head
```

HEAD points to: refs/heads/main
current commit: e21b796e9c9735814596206baa947989e74c860d

```
 go run ./cmd/gogit commit-tree -m "second"
 ```

object hash: 56294e459e92315a1f5fa0418e14e8394b54c04e
object hash: ffcf7190d982cf54b2a2152cd65dd22c3db6a944
object hash: a5c93a920b4ca10a8bd9ebdf359b57ec1600bfd3
object hash: 87b52b4bed37c5a604cbefbd7f41d9052e901b86
object hash: fe1c96bcb0e9c85c8d62764aff185021d8d12db1
object written to: .gogit/objects/fe/1c96bcb0e9c85c8d62764aff185021d8d12db1
tree written with hash: fe1c96bcb0e9c85c8d62764aff185021d8d12db1
object hash: 0a3f1c80724bae5efcfe9a256cc19b0aa158c734
object written to: .gogit/objects/0a/3f1c80724bae5efcfe9a256cc19b0aa158c734
tree written with hash: 0a3f1c80724bae5efcfe9a256cc19b0aa158c734
object hash: 921a86cce097adb34953b4f3bf410055a4f0c52f
object hash: 20eaef3b03749979f0ea8b31a3a6857c208d347b
object hash: 69064a23fcd5b6cadd96795712d7a869549dd699
object hash: 988ae7f9f2161e306a2122bb5db7bebd8c0474b0
object hash: 615778b217873062420b3e83db391252579e21f1
object hash: 3fe12d4d38c57cce5a37253db96eb6af4326ee2c
object hash: 175c683e7201bc1c175d58938edbebcdedcb1bc7
object hash: 05bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
object written to: .gogit/objects/05/bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
tree written with hash: 05bbf0c3c1e4b69cc93e01e3f274846d0a354cdd
object hash: cf0ba92794bcb937bb3f838cb3fff90a6dfd8241
object written to: .gogit/objects/cf/0ba92794bcb937bb3f838cb3fff90a6dfd8241
tree written with hash: cf0ba92794bcb937bb3f838cb3fff90a6dfd8241
object hash: 33a0c130da921cc466552ce0d064106cfefeb0c3
object written to: .gogit/objects/33/a0c130da921cc466552ce0d064106cfefeb0c3
tree written with hash: 33a0c130da921cc466552ce0d064106cfefeb0c3
object hash: 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147
object written to: .gogit/objects/16/8e2f3dfccc9dff3c6cc73c17e471a50bf5c147
wrote commit object 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147
commit hash: 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147

```
 go run ./cmd/gogit head        
 ```           

HEAD points to: refs/heads/main
current commit: 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147

```
 go run ./cmd/gogit cat-file -p 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147
 ```

compressed object size: 171
object found at: .gogit/objects/16/8e2f3dfccc9dff3c6cc73c17e471a50bf5c147
decompressed size: 268
object type: commit
declared size: 257
actual content size: 257
tree 33a0c130da921cc466552ce0d064106cfefeb0c3
parent e21b796e9c9735814596206baa947989e74c860d
author Rudra Sankha Sinhamahapatra <rudra@rudrasankha.com> 1766127821 +0530
committer Rudra Sankha Sinhamahapatra <rudra@rudrasankha.com> 1766127821 +0530

second


## Git Log

```
 go run ./cmd/gogit log
 ```

compressed object size: 171
object found at: .gogit/objects/16/8e2f3dfccc9dff3c6cc73c17e471a50bf5c147
decompressed size: 268
object type: commit
declared size: 257
actual content size: 257
commit 168e2f3dfccc9dff3c6cc73c17e471a50bf5c147
second


compressed object size: 140
object found at: .gogit/objects/e2/1b796e9c9735814596206baa947989e74c860d
decompressed size: 219
object type: commit
declared size: 208
actual content size: 208
commit e21b796e9c9735814596206baa947989e74c860d
first
