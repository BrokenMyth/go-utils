# 简介

自己积累的一些 go 相关工具类



# 使用

```
go get github.com/BrokenMyth/go-utils 
```







# Linux 相关脚本



**安装 docker**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/docker/install_docker.sh | bash
```

**卸载 docker**

```
https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/docker/uninstall_docker.sh
```



**开启 rsa**

Key 填本地生成的秘钥。（TODO：如果配置里不存在一些 rsa 字段，需要新写入）

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/system/open_rsa.sh | bash -s "key"

随后使用本地机连接
ssh -i ~/.ssh/id_rsa root@ip  
ssh root@ip
```



**关闭 rsa**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/system/close_rsa.sh | bash
```

**安装 go**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/go/install_go.sh | bash
```

**卸载 go**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/go/uninstall_go.sh | bash
```

**安装 xray**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/xray/install_xray.sh | bash
```

**卸载 xray**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/xray/uninstall_xray.sh | bash
```

**安装 ffmpeg**  

注意：安装的是二进制文件，无需编译

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/ffmpeg/install_amd64_ffmpeg.sh | bash
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/ffmpeg/install_arm64_ffmpeg.sh | bash
```

(TODO: 根据系统架构来安装)

**卸载 ffmpeg**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/ffmpeg/uninstall_ffmpeg.sh | bash
```





**测试**

```
curl -sSL https://raw.githubusercontent.com/BrokenMyth/go-utils/main/shell/test.sh | bash -s key
```

