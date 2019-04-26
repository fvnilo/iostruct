# iostruct
> Write structs to files and read them back

## Why?
I needed a way to easily read/write structs from the filesystem for a project. Also: the written data had to be encrypted so it can't be accessed unless you have the right password.

## Installation

```
go get -u github.com/nylo-andry/iostruct
```

## Usage

### Simple Read/Write
```
type Foo struct {
	Bar string
}


f := Foo{"Bar"}
err := iostruct.Write(filename, f)
if err != nil {
	panic(err)
}

var f2 Foo
err = iostruct.Read(filename, &f2)
if err != nil {
  panic(err)	
}
```

### Encrypted Read/Write
```
type Foo struct {
	Bar string
}

passphrase := "some passphrase"
f := Foo{"Bar2"}
err := iostruct.WriteEncrypted(filename, passphrase, f)
if err != nil {
	panic(err)
}

var f2 Foo
err = iostruct.ReadEncrypted(filename, passphrase, &f2)
if err != nil {
	panic(err)	
}
```
