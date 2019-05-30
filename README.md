# 简介

这是一个基于Hyperledger fabric的快递溯源应用。传统物流公司一般采用中心化系统，使得系统存在信息不完全透明，数据可以自行修改等缺陷。同时也存在如下的痛点：

1. 我们的快递包在运送过程可能被某个接触到这个快递包的人员串改或调换，但我们无法确定是谁

2. 快递包在运送的过程中损坏却难以确定在那个环节被损坏

3. 电商行业与物流公司之间存在某些虚包合作，不利于人们对购买的商品的了解。

# 使用
## 如何部署
1. 确保当前主机已经部署了如下环境
```
1. Hyperledger Fabric (本项目基于Hyperledger Fabric v1.1)
2. npm ≥ 5.6.0
3. node.js ≥ 8.11.3
4. Angularjs 1.4.3
5. git
6. docker
```
2. 下载代码仓库
```shell
git clone https://github.com/h-hg/SCUT-Blockchain-Group2.git
```
3. 进入`SCUT-Blockchain-Group2`目录，并赋予执行脚本权限
```shell
cd ./SCUT-Blockchain-Group2
chmod a+x ./deployment.sh
```
4. 执行部署脚本
```shell
./deployment.sh
```
## 其他说明
### 第二次部署
第二次部署时，需要删除`~/.hfc-key-store`这个文件夹，脚本里已经写好，不然会出现执行`node registerUser.js `错误

### 关于修改智能合约`source-app.go`
	修改智能合约后，重新部署需要先删除docker镜像名以`dev-mycc`开头的镜像文件，此镜像储存了上一次的智能合约文件。

### 部署在服务器上
修改`server.js`文件的代码
```js
var port = process.env.PORT || 3389; //端口修改为服务器的安全组端口
app.listen(port, '0.0.0.0', function(){
 console.log("Live on port: " + port);
});
```
# 参考资料

1. [Hyperledger中文文档](https://hyperledgercn.github.io/hyperledgerDocs/)
2. [Writing Your First Application](https://hyperledger-fabric.readthedocs.io/en/release-1.1/write_first_app.html)
3. [Hyperledger英文文档](https://hyperledger-fabric.readthedocs.io/)