### goweb 网站开发脚手架项目

### 功能列表

[x] 多主题支持
[x] 多语言支持
[x] 缓存支持
[x] html/css/js 压缩
[x] 多站点支持
[x] 定时任务支持
[x] 谷歌广告支持
[x] mysql/sqlite 支持

### 使用

#### 依赖

```shell
go > 1.18 
make
supervisor
nginx / caddy 
mysql > 5.7 
```

#### 构建

```shell
make build 
```

#### supervisor 安装和配置

```shell
apt install supervisor
```

#### app安装

```shell
make install 
```

#### nginx || caddy 配置

1. nginx 参考 conf/nginx.conf
2. caddy 参考 conf/Caddyfile

#### nginx 重启

```shell
systemctl reload nginx 
```

#### caddy 重启

```shell
caddy reload
```

#### supervisor 重启

```shell
supervisorctl restart tools
```

#### 生产服务的运行目录

```shell
bin/ 
```

需要拷贝 `web` `conf/config.yaml` 到 bin 目录然后运行

