# 数据整理

## regex
```
. 除换行符之外的”任意单个字符”
* 匹配前面字符零次或多次
+ 匹配前面字符一次或多次
[abc] 匹配 a, b 和 c 中的任意一个
(RX1|RX2) 任何能够匹配RX1 或 RX2的结果
^ 行首
$ 行尾
```

- 贪婪模式 *、+
- 非贪婪模式 在*、+后面加？

## 工具
- sed
```
基于文本编辑器ed构建的”流编辑器” 
常用指令: s - 替换
s/REGEX(正则表达式)/SUBSTITUTION(替换文本)/
```
- awk

