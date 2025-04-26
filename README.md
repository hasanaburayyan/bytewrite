# bytewrite

**Write real bytes, not text.**

`bytewrite` is a simple command-line tool for creating **binary files** from **human-readable binary or hexadecimal input**.

When you type `00000001` into a text file, you're writing eight separate characters ('0', '0', ..., '1') — **not** a single byte.  
`bytewrite` lets you **speak in bits or hex** and generate **true binary output**, exactly how the machine sees it.

Perfect for:
- Learning how bytes, memory, and filesystems *really* work
- Building custom file formats
- Bootstrapping tiny virtual machines or games
- Hacking, CTFs, systems programming, retro computing
- Finally understanding the difference between *text* and *bytes*.

---

## 🚀 Install

```bash
go install github.com/hasanaburayyan/bytewrite/cmd/bytewrite@v0.1.1
```


## 📖 Usage

```bash
# write a single byte
$ bytewrite -b 00000001 > out.bin

# verify the output
$ xxd out.bin

00000000: 01                                       .
```

```bash
# write a byte array
$ bytewrite -b 00000001 00000010 00000011 > out.bin

# verify the output
$ xxd out.bin
00000000: 01 02 03                                 ...
```
