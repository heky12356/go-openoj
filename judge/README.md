部署(linux)：

1. 编译isolate
apt安装依赖(ubuntu)
```bash
sudo apt update
sudo apt install build-essential pkg-config libcap-dev
```

> 如果是内网服务器的话可以找一下这三个包
> 寻找libcap-dev、pkg-config、build-essential的本地dev包，或者yum包（centos）
> 然后使用本地安装的方式将这三个包安装到服务器上

编译(部分)
```bash
git clone https://github.com/ioi/isolate.git
cd ./isolate
make isolate
sudo cp isolate /usr/local/bin/
```

完整编译(可选)
```bash
# 新安装一个依赖
sudo apt-get install asciidoc
cd ./isolate
make
make install
```

2. 配置isolate
新建一个用到的文件夹
```bash
mkdir /var/local/lib/isolate
```
创建配置文件
```bash
cp default.cf.in /usr/local/etc/isolate
```

2. 运行
```bash
go run ./cmd/main.go
```

另，使用docker（不推荐在内网服务器上使用）
构建镜像：
```bash
docker build -t go-openoj-judge .
```
启动:
```bash
docker run -d -p 5050:5050 go-openoj-judge
```