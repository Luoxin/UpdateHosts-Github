# UpdateHosts-Github

起源于 `update_hosts` 项目，换了个golang做了份重构

## 使用方式

- 下载二进制包，已管理员方式运行
### Windows下报错没有修改文件权限的解决方案
#### 方便但可能被恶意软件（比如我自己）利用的方案
- 在文件资源管理器打开目录 `C:\Windows\System32\drivers\etc`
- 选中 `hosts` 文件，选择 `属性`，将只读去掉
![image](https://user-images.githubusercontent.com/33619903/130169812-c2cf69b3-b45c-4353-9d73-a3dc81caad7b.png)
#### 不方便但是可用的方案
- 将工具生成的文件覆盖到`C:\Windows\System32\drivers\etc\hosts`
## TODO

- [x] 进行dns查询
- [x] 筛选最快的ip
- [ ] 写入到hosts文件
	- [x] Linux
	- [ ] Windows
- [ ] 以服务的方式运行
