cd ~
# 删除所有容器
docker rm -f $(docker ps -aq)
# 删除docker镜像

# 删除注册的管理员和用户
rm -f -r ./.hfc-key-store
# 下载代码仓库
# rm -f -r ./SCUT-Blockchain-Group2
# git clone https://github.com/h-hg/SCUT-Blockchain-Group2.git
# 设置权限
cd ./SCUT-Blockchain-Group2/
cd ./basic-network
chmod a+x ./start.sh
cd ../source-app
chmod a+x ./startFabric.sh
# 安装包
npm install
./startFabric.sh
# 更改端口
node registerAdmin.js
node registerUser.js
node server.js