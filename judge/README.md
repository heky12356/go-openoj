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
mkdir /var/lib/isolate 
#这是下面的box_root路径
#代码里面硬编码了，如果想改可以改judge/internal/service/service.go中
# boxPath := "/var/lib/isolate/" + boxID + "/box/"
# metaPath := "/var/lib/isolate/" + boxID + "/meta.txt"
```
创建配置文件
```bash
vim /usr/local/etc/isolate
# vim的使用可以参考：https://blog.heky.top/index.php/archives/38/
```

粘贴下面的内容：
```isolate
# cat /etc/isolate
# This is a configuration file for Isolate

# All sandboxes are created under this directory.
# To avoid symlink attacks, this directory and all its ancestors
# must be writeable only to root.
box_root = /var/lib/isolate

# Directory where lock files are created.
lock_root = /run/isolate/locks

# Control group under which we place our subgroups
# Either an explicit path to a subdirectory in cgroupfs, or "auto:file" to read
# the path from "file", where it is put by isolate-cg-helper.
# cg_root = /sys/fs/cgroup/isolate.slice/isolate.service
cg_root = auto:/run/isolate/cgroup

# Block of UIDs and GIDs reserved for sandboxes
first_uid = 60000
first_gid = 60000
num_boxes = 1000

# Only root can create new sandboxes (default: 0=everybody can)
#restricted_init = 1

# Per-box settings of the set of allowed CPUs and NUMA nodes
# (see linux/Documentation/cgroups/cpusets.txt for precise syntax)

#box0.cpus = 4-7
#box0.mems = 1
```
2. 创建judge需要用到的文件夹
```bash
# 在go-openoj根目录下创建
mkdir ./judge/tmp
mkdir ./judge/submissions
```
3. 运行
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