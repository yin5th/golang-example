### 怎么使用
###### 项目文件夹 goroutine_example
```
    cd goroutine_example
    go build
    ./goroutine_example http://www.baidu.com http://www.qq.com http://mail.qq.com
```
###### 结果展示
```
    url http://www.baidu.com cost 0.11s, content length: 153392
    url http://mail.qq.com cost 0.83s, content length: 12240
    url http://www.qq.com cost 0.95s, content length: 239731
    total cost 0.97 s

```
