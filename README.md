# Resizer Core

这是[Reizer](https://github.com/Zhoucheng133/Resizer)核心模块

你可以使用下面的命令生成动态库

```bash
# macOS
go build -buildmode=c-shared -o build/core.dylib
# Windows
go build -buildmode=c-shared -o build/core.dll

# 如果你使用比较新版本的golang，使用下面的命令生成动态库
#  macOS
go build -buildmode=c-shared -ldflags="-s -w" -o build/core.dylib
# Windows
go build -buildmode=c-shared -ldflags="-s -w" -o build/core.dll
```
