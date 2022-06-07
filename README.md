# simple-demo

## 功能实现

跳过登录验证

user.go中的usersLoginInfo里是测试的用户名密码，拿那个登录

## 投稿功能

token就是user+psd
    
    这个是否需要自己后续修改？（明文传输问题
    - 明文传输已经修改，后续应该改为jwt返回的token

数据库根据```struct video```的字段写了一下
    
    可以讨论一下有什么更改的

视频文件较大的话处理实际会很长，客户端接收不到response会提示失败

但实际服务端还在处理，最后数据库和对象存储中都会保存相关信息

TODO:videoList的不是只可以查询自己的，封装user需要用到user的数据，isFavorite未完成

TODO:有个bug不知道是我的bug还是客户端的，在publicList界面点视频会闪退

## 有个疑问

用户id和视频id，是int64吗？还是uuid？ 

- 用int64


## 关闭8080端口的进程

lsof -i:8080

kill pid

## test的问题

test时需要将config的路径改成../为什么我也不知道

正常pwd是mini-douyin，而test时路径是mini-douyin/test
- 因为test的时候config的路径用的是是绝对路径

## 注册

判断用户名是否重复好像失效
- 在上次提交中已经修改

注册数据库写了，但是客户端显示网络失败
- json返回类输出的判断条件写错，写成err!=nil了，已经修改

注册后自动登录

新注册的给用户id赋值 （可以改成按照时间戳的方式或者其他？