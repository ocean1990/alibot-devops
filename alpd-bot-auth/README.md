# alpd-bot-auth

用户认证服务

## 构建与运行

本地构建依赖docker环境，请提前安装好docker。

### 构建

```sh
make build
# docker images # 查看是否有生成alpd-bot-auth:latest的镜像
```

### 本地运行

```sh
docker run -p 9001:9001 alpd-bot-auth
```

### 单元测试(optional)

要求nodejs环境。

```sh
npm install
make test
```

### 部署到本地k8s

提前在本地安装好k8s环境（如minikube），并构建好容器镜像。

```sh
kubectl apply -f deployment-local.yml
# kubectl get svc
#   NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
#   alpd-bot-auth-svc     ClusterIP   10.108.106.62   <none>        9001/TCP   17m
# kubectl get deployments
#   NAME                         READY   UP-TO-DATE   AVAILABLE   AGE
#   alpd-bot-auth-deployment     1/1     1            1           16m
```

## TODO

- [ ] 添加服务单元测试
- [ ] 支持用户数据存储到MySQL数据库中

## Reference

* [docker安装](https://docs.docker.com/desktop/)
* [minikube安装](https://developer.aliyun.com/article/221687)
