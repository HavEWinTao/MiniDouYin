# simple-demo

## 功能实现

跳过登录验证

user.go中的usersLoginInfo里是测试的用户名密码，拿那个登录

## 投稿功能

data是一个视频文件,客户端已经处理好了,会保存到public文件夹
    
    后续是否需要用对象存储等？

token就是user+psd
    
    这个是否需要自己后续修改？（明文传输问题

title是在app里输入的视频描述

数据库根据```struct video```的字段写了一下
    
    可以讨论一下有什么更改的

实测用户查看视频时下载速度有点慢，该怎么解决
    
    时慢时快，不知道客户端是否有缓存，也不清楚反正确实有时候很慢

每个视频是有个url的,封面也是有url的
    
    客户端不支持客户选择封面
    
    TODO:在视频中截取封面，并保存，得到封面url

## 有个疑问

用户id和视频id，是int64吗？还是uuid？

## 关闭8080端口的进程

lsof -i:8080

kill pid

## test的问题

test时需要将config的路径改成../为什么我也不知道

正常pwd是mini-douyin，而test时路径是mini-douyin/test
