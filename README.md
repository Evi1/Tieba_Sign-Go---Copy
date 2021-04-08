# Tieba Sign in Go

**WARNING: This project is currently under BETA, use it on your risk**

A faster && lighter sign robot for http://tieba.baidu.com which can help you get more credit at tieba.  

From https://github.com/kookxiang/Tieba_Sign-Go

Demo:
https://sign.bilibili.network/

使用方法  
1、安装go  
2、go get github.com/rikaaa0928/Tieba_Sign-Go---Copy  
完成后 $GOPATH/bin （GOPATH默认为：~/go） 目录下因该有可执行程序 Tieba_Sign-Go---Copy ，如果没有请手动build  
3、将需要签到的贴吧账号的cookies（需要BDUSS和STOKEN两项）放到txt后缀的文件里，每个账号一个单独文件，文件名随意。  
4、把准备好的文件放到 $GOPATH/src/github.com/rikaaa0928/Tieba_Sign-Go---Copy/cookies/ 目录下  
5、直接运行 Tieba_Sign-Go---Copy 或使用systemd启动，使用浏览器访问ip+:60080端口即可查看签到情况，需要修改ip或端口请直接修改global/global.go中的Server = ":60080"部分，重新build即可  

附：简单的tbsign.service（个人使用，请酌情修改）  
[Unit]  
Description=Tieba_Sign-Go---Copy  
After=network.target  
Wants=network.target  

[Service]  
Type=simple  
Environment=GOPATH="/usr/share/go/contrib"  
PIDFile=/var/run/tbsign.pid  
ExecStart=Tieba_Sign-Go---Copy  
Restart=on-failure  

[Install]
WantedBy=multi-user.target  
