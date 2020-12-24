## 概述

自用，自动为主流平台自动编译Github上的go程序的程序。

只能在Linux运行，适用于制作api。

## 运行

```shell
./main -repo=https://github.com/... -backend=https://... -name=... 
```

repo为仓库git地址，backend为回调地址，name为编译结果命名