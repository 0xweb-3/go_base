# 工作区模式
当私用包没有发布，但需要在同一项目中去使用时
```shell
# 目录结构
worksace
|____project
| |____go.mod // module github.com/0xweb-3/project
| |____main.go.
|____pkg1
| |____go.mod // module github.com/0xweb-3/pkg 
| |____pkg1.go
```


```shell
cd work

go work init pkg1 project
```