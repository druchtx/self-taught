# Install Hands-on

## Linux packages

- [Gitlab Docs](https://docs.gitlab.com/install/package/debian/)
  
官方文档需要设置 外部访问域名，配置dns，ssl等，在本地测试时，外部访问地址使用 "EXTERNAL_URL=http://gitlab.local"或者虚拟机IP

修改外部访问地址使用如下命令
```shell
# 将gitlab.rb中的相关字段修改如下
# external_url "http://gitlab.local"
# letsencrypt['enable'] = false
vim /etc/gitlab/gitlab.rb  

# 修改完成后，执行如下命令
sudo gitlab-ctl reconfigure
sudo gitlab-ctl restart
```

查看 gitlab 状态
```shell
druchtx@debian:~$ sudo gitlab-ctl status
run: alertmanager: (pid 25279) 2045s; run: log: (pid 25085) 2096s
run: gitaly: (pid 23649) 2661s; run: log: (pid 22877) 2815s
run: gitlab-exporter: (pid 25246) 2047s; run: log: (pid 25015) 2115s
run: gitlab-kas: (pid 25208) 2049s; run: log: (pid 23170) 2795s
run: gitlab-workhorse: (pid 23612) 2661s; run: log: (pid 23335) 2704s
run: logrotate: (pid 22785) 2830s; run: log: (pid 22793) 2829s
run: nginx: (pid 25222) 2048s; run: log: (pid 23373) 2698s
run: node-exporter: (pid 25238) 2048s; run: log: (pid 25004) 2121s
run: postgres-exporter: (pid 25293) 2044s; run: log: (pid 25112) 2090s
run: postgresql: (pid 22917) 2807s; run: log: (pid 22939) 2804s
run: prometheus: (pid 25258) 2046s; run: log: (pid 25058) 2104s
run: puma: (pid 25152) 2080s; run: log: (pid 23248) 2716s
run: redis: (pid 22820) 2824s; run: log: (pid 22830) 2823s
run: redis-exporter: (pid 25249) 2047s; run: log: (pid 25035) 2108s
run: sidekiq: (pid 25128) 2085s; run: log: (pid 23263) 2712s
```

