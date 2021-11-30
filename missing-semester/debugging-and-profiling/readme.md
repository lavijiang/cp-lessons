# 调试及性能分析

## 静态分析
- 解决语法问题
    ```
    python: pyflakes\mypy\pylint\black
    shell: shellcheck
    go: gofmt
    rust: rustfmt
    javascript\html\css: prettier
    ```
## 性能分析
- 计时
- CPU
    ```
    追踪分析器（tracing）：记录程序的每一次函数调用
    采样分析器（sampling）：周期性的监测（通常为每毫秒）您的程序并记录程序堆栈
    ```
- 内存
- 可视化

    火焰图

## 资源监控
- 通用监控
    ```
    htop\top 显示当前运行进程的多种统计信息
    dstat 实时地计算不同子系统资源的度量数据，例如 I/O、网络、 CPU 利用率、上下文切换等
    ```
- I/O操作
    ```
    iotop 显示实时 I/O 占用信息而且可以非常方便地检查某个进程是否正在执行大量的磁盘读写操作
    ```
- 磁盘使用
    ```
    df 显示每个分区的信息
    du\ncdu 显示当前目录下每个文件的磁盘使用情况
    ```
- 内存使用
    ```
    free 显示系统当前空闲的内存
    htop 显示系统当前内存
    ```
- 打开文件 
    ```
    lsof 查看某个文件是被哪个进程打开
    ```
- 网络连接和配置
    ```
    ss\netstat 监控网络包的收发情况以及网络接口的显示信息
    ip\ifconfig 路由、网络设备和接口信息
    ```
