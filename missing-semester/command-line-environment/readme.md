# 命令行环境

## 任务控制
- 优雅的退出
    ```
    SIGTERM
    kill -TERM <PID>
    ```
- 暂停
    ```
    SIGSTOP
    Ctrl-Z
    ```
- 前台恢复
    ```
    fg
    ```

- 后台恢复
    ```
    bg
    ```

## 终端多路复用

#### 工具
- tmux神器

## 配置文件（Dotfiles）
- 常见配置（建议维护自己的一份配置文件）
    ```
    bash - ~/.bashrc, ~/.bash_profile
    git - ~/.gitconfig
    vim - ~/.vimrc 和 ~/.vim 目录
    ssh - ~/.ssh/config
    tmux - ~/.tmux.conf
    ```
## shell远程登录
- 免密登陆配置
    ```
    密钥生成 ssh-keygen -o -a 100 -t ed25519 -f ~/.ssh/id_ed25519
    公钥上传 ssh-copy-id -i .ssh/id_ed25519.pub foobar@remote 或者 cat .ssh/id_ed25519.pub | ssh foobar@remote 'cat >> ~/.ssh/authorized_keys'
    ```

- SSH 配置

    使用~/.ssh/config
    ```
    Host vm
        User foobar
        HostName 172.16.174.141
        Port 2222
        IdentityFile ~/.ssh/id_ed25519
        LocalForward 9999 localhost:8888

    # 在配置文件中也可以使用通配符
    Host *.mit.edu
        User foobaz
    ```

- ssh复制文件

    scp ：当需要拷贝大量的文件或目录时，使用scp 命令则更加方便，因为它可以方便的遍历相关路径。语法如下：scp path/to/local_file remote_host:path/to/remote_file

    rsync 对 scp 进行了改进，它可以检测本地和远端的文件以防止`重复拷贝`。它还可以提供一些诸如符号连接、权限管理等精心打磨的功能。甚至还可以基于 --partial标记实现`断点续传`。rsync 的语法和scp类似

