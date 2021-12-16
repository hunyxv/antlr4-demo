# ANTLR4 简单四则运算demo

ANTLR4 官网：[antlr.org](https://www.antlr.org/)

生成 go 代码：

```shell
// 同时生成 观察者模式 和 访问者模式 代码
antlr4 -Dlanguage=Go -visitor -o parser Hello.g4
// 只生成 访问者模式 代码
antlr4 -Dlanguage=Go -no-listener -visitor -o parser Hello.g4
// 只生成 观察者模式 代码
antlr4 -Dlanguage=Go -o parser Hello.g4
```

1. 简单四则运算
2. go 代码嵌入