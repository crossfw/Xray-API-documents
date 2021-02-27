# Xray-API-documents
## 写在开头
使用API是一个技术活, 请在编写相关程序前对 Xray 配置有所了解<br>
[点我了解](https://xtls.github.io/about/) <br>
以及你至少会一门支持grpc的语言, 这里以 Golang 为例, ~~因为他比较方便~~.

## 文档
说明文档在 docs 目录下<br>
相关示例在 example 目录下, 每一个方法都对应一个test文件演示完整的过程<br>

## 其他说明
- 请尽量不要并发执行， 经过测试，大约在第10个线程后，Xray内核会丢弃多余的请求

相关问题请提issue.