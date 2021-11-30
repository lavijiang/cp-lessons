# Shell 工具和脚本

## Shell 脚本
#### [特殊变量](https://tldp.org/LDP/abs/html/special-chars.html)
```
$0 - 脚本名
$1 到 $9 - 脚本的参数。 $1 是第一个参数，依此类推。
$@ - 所有参数
$# - 参数个数
$? - 前一个命令的返回值
$$ - 当前脚本的进程识别码
!! - 完整的上一条命令，包括参数。常见应用：当你因为权限不足执行命令失败时，可以使用 sudo !!再尝试一次。
$_ - 上一条命令的最后一个参数。如果你正在使用的是交互式shell，你可以通过按下 Esc 之后键入 . 来获取这个值。
```

#### shell脚本校验工具

[shellcheck](https://github.com/koalaman/shellcheck)


#### 常用工具
```
find\fd\locate - 查找文件
grep\rg - 查找代码
history\Ctrl+R - 查找 shell 命令
fasd\autojump - 文件夹导航
```
