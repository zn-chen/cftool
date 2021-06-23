# cftool
多种格式配置文件的编辑辅助工具

## 环境
unix,windows

## 安装

### 通过源码安装
```bash
git clone https://github.com/zn-chen/cftool
cd cftool
# linux
make && make install

# other
#go build -o cftool ./*.go
```

### 通过安装包安装
```bash
# redhat
yum install cftool-<version>-1.el7.x86_64.rpm

# debian
#apt-get install whisper_<version>_amd64.deb
```

### Example:
```bash
# read config
cftool r -f <config_file_path> <key1>.<key2>

# write config
cftool w -f <config_file_path> <key1>.<key2>=<value> > <config_file_path>
```
