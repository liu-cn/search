# search

全局文本搜索，可以搜索出指定文件内的指定文本，支持跳转到文本所在的行。

如果有go环境的话可以直接go install 安装，

```go
    go install github.com/liu-cn/search
```

没有go环境，可自行下载windows macos linux 指定操作系统文件夹下可执行文件，下载后添加到环境变量即可使用



使用方式

```shell
search -f "文件" -t "文本" -p "目录(不指定默认从当前目录搜索)" -e "排除搜索的文本"
```

例如：我想要搜索该目录下所有go文件中包含 "跑路" 的文本，应该这样的搜索

```shell
search -f "*.go" -t "跑路"
```

goland/idea 显示效果 只要是jetbrains的编辑器都是支持直接点击后跳转到该行的

![输出效果](http://cdn.motianli.com/cdn/%E6%88%AA%E5%B1%8F2022-04-15%2000.21.53.png)

vscode显示效果，也支持点击跳转到文本位置

![](http://cdn.motianli.com/cdn/%E6%88%AA%E5%B1%8F2022-04-15%2000.39.34.png)



可以看到搜索到原代码是这样的，

![](http://cdn.motianli.com/cdn/%E6%88%AA%E5%B1%8F2022-04-15%2000.41.40.png)

可以看到第八行是注释，我们不想搜索到第八行，我们可以排除 注释 // 的行

接下来命令可以改一下，改成这样

```shell
search -f "*.go" -t "跑路" -e "//"
```

显示效果，注释行已经被排除掉了

![](http://cdn.motianli.com/cdn/%E6%88%AA%E5%B1%8F2022-04-15%2000.46.02.png)