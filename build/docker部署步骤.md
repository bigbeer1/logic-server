### 1、预装Gcc docker环境  还有docker-compose


### 2、启动项目所依赖的环境  根据实际情况配置是否需要国内代理

```
docker engine 配置 

{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": false,
  "features": {
    "buildkit": true
  },
  "registry-mirrors": [
    "https://docker.registry.cyou",
    "https://docker-cf.registry.cyou",
    "https://dockercf.jsdelivr.fyi",
    "https://docker.jsdelivr.fyi",
    "https://dockertest.jsdelivr.fyi",
    "https://mirror.aliyuncs.com",
    "https://dockerproxy.com",
    "https://mirror.baidubce.com",
    "https://docker.m.daocloud.io",
    "https://docker.nju.edu.cn",
    "https://docker.mirrors.sjtug.sjtu.edu.cn",
    "https://docker.mirrors.ustc.edu.cn",
    "https://mirror.iscas.ac.cn",
    "https://docker.rainbond.cc"
  ]
}


#### 步骤3：部署etcd
$ docker-compose -f docker-compose-etcd.yml up -d

#### 步骤12：如果需要本地语音合成人工只能功能
$ docker-compose -f docker-compose-paddlespeech.yml up -d 




$ docker exec -it mysql mysql -uroot -p
##输入密码：PXDNA999999
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;