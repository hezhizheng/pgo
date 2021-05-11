# pgo 使用golang实现一些常用的内置函数(包括但不限于php一些常用的内置函数)

收集一些常用的操作函数(包括但不限于php一些常用的内置函数)，帮助更快的完成开发工作，并减少重复代码，
主要来源于 [https://www.php2golang.com](https://www.php2golang.com/)

## 安装
`go get -u github.com/hezhizheng/php2golang`

## 使用
```go
pgo.InArray("hello",[]string{"hello","world"}) // true
pgo.ArrayColumn()
pgo.ArrayPush()
pgo.ArrayPop()
pgo.ArrayUnshift()
pgo.ArrayShift()

pgo.Md5("123465") // e10adc3949ba59abbe56e057f20f883e
pgo.Uniqid("") // 608a594ee0624
pgo.MbStrlen("中文 1") // 4
pgo.Explode(",","hello,world")
pgo.Strpos("+1s","s") // 2
// 将字符串以指定长度进行截断
pgo.StrLimit("测试2q文字超出，符号补充 1a",10,"...") // 测试2q文字超出，符...

pgo.Blank(0) // false
pgo.IsNumeric("000") // true

pgo.FileExists("path") // bool
// 用法可参考test , 更多函数完善中...
```

## License
[MIT](./LICENSE.txt)