# mtk8
使用beego和client-go封装kubernetes的的接口层
## 目录规划
* client 目录存放公用客户端,如k8s client,mysql client......
* conf 目录存放beego配置文件
* controllers 目录存放各种控制器,针对k8s中每个object可设置一个文件
* models 暂时没用到
* routers 目录存放路由器,一般一个路由文件就好
## 导入包版本信息
* client-go: branch v2.0
* beego: tag v1.8.0
* 其他: branch master

