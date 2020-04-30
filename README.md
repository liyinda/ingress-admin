# ingress-admin

ingress-admin是基于kubernetes中ingress-controller的管理控制台，对于ingress-controller
推荐使用kubernetes官方默认的Nginx-ingress-controller。ingress-admin可以发布
到kubernetes集群中，作为Nginx-ingress-controller的ui管理控制台，通过它来管理
kubernetes中的ingress配置。

## 目录
* [环境](#环境)
* [下载](#下载)
* [结构](#结构)
* [编译](#编译)
  * [build binary](#build-binary)
  * [build docker image](#build-docker-image)
* [运行](#运行)
  * [运行在kubenetes内部](#run-binary)
  * [运行在kubenetes外部](#run-docker-image)
* [基本操作](#运行参数&google-authenticator手机端)
* [Annotations操作](#户后台与google-authenticator对接)
  * [限流操作](#redmine)
  * [Zabbix](#zabbix)



## 环境

* [Golang 1.13.4](https://golang.org/)
* [kubernetes 1.17.0](https://kubernetes.io/)
* [client-go 0.17.0](https://github.com/kubernetes/client-go)
* [k8s.io/api 0.17.3](https://github.com/kubernetes/api)

## 下载

Dockerfile can be downloaded from [Releases](https://github.com/liyinda/google-authenticator/releases) page.

## 结构

代码目录结构

``` shell
├── backend #后端代码
│   ├── api
│   ├── config
│   ├── dist
│   ├── example
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middleware
│   └── pkg
├── build #容器编译
│   ├── config
│   ├── dist
│   └── dockerfile
├── deploy #k8s部署
│   ├── namespace.yaml
│   ├── rbac.yaml
│   ├── svc.yaml
│   └── with-rbac.yaml
├── frontend #前端代码
│   └── vue-admin-template
├── LICENSE
├── Makefile
├── prometheus
└── README.md
``` 

## 编译

### build binary

``` shell
cd backend/
GOOS=linux CGO_ENABLED=0  go build -tags netgo -a -v -ldflags="-s -w" -o ingress-admin
```
### build docker image
``` shell
#编译前端
make build-frontend
#编译后端
make build-backend
#编译容器
make docker
#编译清空
make clean
```

## 运行

### 运行在kubenetes内部
``` shell
将deploy目录下载到kubernetes集群中，运行命令如下：

git clone https://github.com/liyinda/ingress-admin

cd deploy

kubectl apply  -f .

```
### 运行在kubenetes外部
```
如果不以ServiceAccount的方式，需要使用$HOME/.kube/config
同时需要更改backend/api/database/kube.go，建议使用内部部署方式
```

## 运行参数&google-authenticator手机端

### 可根据自身环境更改运行参数
``` shell
./verificationGoogleCode -h
Usage of ./verificationGoogleCode:
  -http.address string
        Address on HTTP Listen . (default ":8082")
  -log string
        Log file name (default "authenticator.log")
  -redis.address string
        Address on Redis Server . (default "127.0.0.1:6379")

```

### 手机下载google-authenticator客户端
iphone手机和android手机都有对应的客户端，请大家自行下载

![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/google-authenticator.jpg)


## 用户后台与google-authenticator对接

### Redmine
vi app/views/account/login.html.erb
``` shell 
添加
14 <tr>
15     <td style="text-align:right;"><label for="code">Google验证码:</label></td>
16     <td style="text-align:left;"><%= text_field_tag 'code', nil, :tabindex => '3' %></td>
17 </tr>

```

vi app/controllers/account_controller.rb
``` shell 
添加
1   require "open-uri"

192   def password_authentication
193 
194     uri = 'http://[google-authenticator服务端地址]/get?issuser=' + params[:username] + '&code=' + params[:code]
195     html_response = nil
196     open(uri) do |http|
197     html_response = http.read
198     end
199 
200     if html_response == 'ok'
201 
202     user = User.try_to_login(params[:username], params[:password], false)
203     if user.nil?
204       invalid_credentials
205     elsif user.new_record?
206       onthefly_creation_failed(user, {:login => user.login, :auth_source_id => user.auth_source_id })
207     else
208       # Valid user
209       if user.active?
210         successful_authentication(user)
211       else
212         handle_inactive_user(user)
213       end
214     end
215 
216     else
217         redirect_to(:action => 'login')
218     end
219 
220 
221   end

```

![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/redmine.jpg)

### Zabbix
vi include/views/general.login.php
``` shell 
添加
55         ->addItem([new CLabel(_('Password'), 'password'), (new CTextBox('password'))->setType('password')])
56         ->addItem([
57                 new CLabel(_('Google Code'), 'code'),
58                 (new CTextBox('code'))->setAttribute('', ''),
59                 $error
60         ])

```

vi index.php 
``` shell 
添加
65 if (isset($_REQUEST['enter']) && $_REQUEST['enter'] == _('Sign in')) {
66         // try to login
67         $autoLogin = getRequest('autologin', 0);
68         //print_r($_REQUEST);
69         $authflag=file_get_contents("http://[google-authenticator服务端地址]/get?issuser=".getRequest('name', '')."&code=".getRequest('code', ''));
70         //echo "http://[google-authenticator服务端地址]/get?issuser=".getRequest('name', '')."&code=".getRequest('code', '');
71         if ($authflag=='ok'){}else{
72             echo 'Google验证码错误'; header('Refresh: 2; url=http://zabbix.org/');exit;
73         }
74         //echo getRequest('code', '');

```


![image](https://github.com/liyinda/google-authenticator/blob/master/jpg/zabbix.jpg)


更多后台对接改造等您实现

