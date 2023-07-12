# **项目说明**
## **字节青训营寒假专场后端项目 抖声 dousheng**
                                                                                                                                             
                                                                                                                                             
                                                                                                                                             
                                                                                                                                             
                                                                                                                                             
```                                                                                                                                    
            @@@^                                       =@@@                                                                            
            @@@^                                       =@@@                                                                            
            @@@^                                       =@@@                                  
      =@@@@@@@@@^  =@@@@@@@@@   @@@^   @@   @@@@@@@@@^  =@@@@@@@@@   @@@@@@@@@@^  @@@@@@@@@  =@@@@@@@@@@^                               
      =@^    @@@^  =@@@    @@   @@@^   @@   @@          =@@@    @@   @@@      @^  @@@^   @@  =@^     =@@^                                
      =@^    @@@^  =@@@    @@   @@@^   @@   @@@@@@@@@^  =@@@    @@   @@@@@@@@@@^  @@@^   @@  =@@     @@@^  
      =@^    @@@^  =@@@    @@   @@@^   @@          @@^   @@@    @@   @@@          @@@^   @@  =@@     @@@^
      =@@@@@@@@@^  =@@@@@@@@@   @@@@@@@@@   @@@@@@@@@^  =@@@    @@   @@@@@@@@@@`  @@@^   @@   =@@@@@@@@@^                               
                                                                                                    =@@@^
                                                                                              @@    =@@@^                                
                                                                                             ^@@@@@@@@@@^                                
                                                                                                                                                                
```                                                                                                                     
### **运行环境**：

>go.version = 1.19.5
>
>mysql.server = 8.0
>
>ffmpeg = 5.1.2

### **配置文件**
> confs
>  ├──dbmysql.json


### **运行方式 **
> go build ./main.go
>
> ./dousheng.exe

## **领域1：基础功能：**
- 视频Feed流，支持所有用户刷抖音，视频按照投稿时间倒序推出（完成）
- 视频投稿，支持登录用户自拍视频投稿（完成）
-  个人主页，支持查看用户的基本信息和投稿列表，注册用户流程简化（完成）
## **领域2：社交功能**
- 登录用户可以对视频点赞，在个人主页的喜欢Tab下能够查看点赞视频列表（完成）
- 支持未登录用户查看视频下的评论列表，登录用户能发表评论（完成）
- 登录用户可以关注其他用户，能够在个人主页查看本人的关注数量和粉丝数量，查看关注列表和粉丝列表（部分完成）
- 登录用户在消息页展示已关注的用户列表，点击用户头像进入聊天页面可以发送消息。（未完成）


**优化**
> 增加自定义日志模块
>
> 密码采用bcrypt模块加密
>
> jwt设置token
>
> ffmpeg抽帧设置视频封面


一、项目介绍
互动方向的极简版抖音后端项目
项目服务地址- http://123.60.180.147:8080/
Github  地址，权限设置为 public- https://github.com/vanyongqi/dousheng
项目演示bilibili地址：https://www.bilibili.com/video/BV1JT411i7v2/?vd_source=a96cf885606bcff2356c308aa6f73cbc
二、项目分工
好的团队协作可以酌情加分哟～请组长和组员做好项目分工与监督。
团队成员
主要贡献
范永奇
项目的设计、开发、测试、答辩等工作
三、项目实现
3.1 技术选型与相关开发文档
项目业务分析：
为视频浏览为核心机制的实时短视频App提供后端服务
后端业务分析：
  - 业务实体模型的模型设计与关系设计
  - 用户之间的实时通讯
  - 后端视频的存储与用户的推荐机制
存储分析：
- video大小：正常发布短视频：5-10M
- 数据库大小：用户单条记录200KB
技术选型： 
  - Http框架：Gin
  - ORM框架：GORM
  - 数据库：Mysql8（本地）、MariaDB（远程）
开发平台与工具：
  - Windows11
  - Goland（本地开发）
  - Vscode（远程部署）

后端项目地址：
- http://123.60.180.147:8080/
- http://www.fanyongqi.top:8080/
前端版本：
极简抖音APP (2023.2.6-15.55版本)
3.2 架构设计
单体类MVC架构
1. 控制层（Controller）：用于控制网络请求+业务实现。
2. 持久层（Dao）：用于进行数据库的操作。
[图片]
3.3 项目代码介绍
└──Configs配置文件
│   ├──msyql.json数据库 名称 密码
├──Controllers
│   ├──request  请求格式
│   ├──response 消息返回响应格式
│   ├──对应接口的controllers
├──Databases
│   ├──数据库初始化、
│   ├──DAO gorm数据库查询语句、
├──Middlewares
│   ├──Mylog.go自定义日志（根据logrus日志库实现）、
│   ├──Token的相关操作：生成与解析
├──Models
│   ├──gorm框架定义的user、video结构体以及对应的封装
├──Public 静态目录
│   ├──视频文件、视频封面、用户头像
├──Router
│   ├──router.go路由
├──Services
│   ├──chat用户聊天（未实现）
├──Test
│   ├── _test.go文件
├──Utils
│   ├──密码加密、利用content获取用户信息等函数

四、测试结果
功能测试：
- 视频Feed流，支持所有用户刷抖音，视频按照投稿时间倒序推出 （完成）
- 视频投稿，支持登录用户自拍视频投稿（完成）
-  个人主页，支持查看用户的基本信息和投稿列表，注册用户流程简化（完成）
- 登录用户可以对视频点赞，在个人主页的喜欢Tab下能够查看点赞视频列表（完成）
- 支持未登录用户查看视频下的评论列表，登录用户能发表评论（完成）
- 登录用户可以关注其他用户，能够在个人主页查看本人的关注数量和粉丝数量，查看关注列表和粉丝列表 （完成部分）
- 登录用户在消息页展示已关注的用户列表，点击用户头像进入聊天页面可以发送消息。（未完成）
性能测试：
测试平台：
  - CentOS 7系统
  - CPU1核|内存2G|磁盘40G|带宽1Mb/s
测试工具：wrk ，针对 Http 协议的基准测试工具，能够在单机多核 CPU 的条件下，使用系统自带的高性能 I/O 机制，如 epoll，kqueue 等，通过多线程和事件模式，对目标机器产生大量的负载。
测试结果：详细结果见图片
12线程  400连接 3秒 QPS:2021.48   TPS:292.17KB
30线程  400连接 30秒 QPS:1423.36 TPS：205.72KB
[图片]
五、Demo 演示视频 （必填）

小米8手机有部分bug，演示中无法模拟视频上传，录屏工具声音小但存在。
实际已使用postman+一加手机进行实机测试，均上传成功。
项目演示bilibili地址：https://www.bilibili.com/video/BV1JT411i7v2/?vd_source=a96cf885606bcff2356c308aa6f73cbc
暂时无法在{app_display_name}文档外展示此内容

六、项目总结与反思
1. 目前仍存在的问题
  1. 未完全开发完成，关注逻辑存在bug，消息功未能实现
  2. 自定义日志系统不够详细
  3. 测试不够完善
2. 已识别出的优化项
  1. 视频压缩减少磁盘占用
  2. 视频的推荐算法
  3. 微服务架构，将不同模块分成多个服务，最外层由统一的网关进行接收消息，然后转交给不同的端口号进行响应。甚至可以使不同机器的不同端口。
3. 架构演进的可能性
  1. 增加chat功能，因为前期的目录结构是按照完成所有功能完成的
  2. Redis进行热数据缓存如用户的关系表等
  3. 数据库表的二次分离
4. 项目过程中的反思与总结
  1. 项目前期的业务逻辑设计非常重要，关乎软件的生命周期
  2. 好的项目一定具备足够的复用性和拓展性
  3. 先去完成一个大的框架，再去进行补充修正，不要耽于开始的细节，否则后期的进度会很难赶上，甚至无法完成既定的任务。
  4. 即使不是领导者，也要多去联系争取志同道合的人去完成自己想完成的事情.

七、其他补充资料












## **接口**
#### **基础接口**
- >/douyin/feed/ - 视频流接口
- >/douyin/user/register/ - 用户注册接口
- > /douyin/user/login/ - 用户登录接口
- >/douyin/user/ - 用户信息
- >/douyin/publish/action/ - 视频投稿
- >/douyin/publish/list/ - 发布列表
- >/douyin/relation/action/ - 关注操作
- >/douyin/relation/follower/list/ - 用户粉丝列表
- >/douyin/relatioin/follow/list/ - 用户关注列表

###  **互动接口**
- >/douyin/favorite/action/ - 赞操作
- >/douyin/favorite/list/ - 赞列表
- >/douyin/comment/action/ - 评论操作
- >/douyin/comment/list/ - 视频评论列表
- >/douyin/message/chat/ - 聊天记录
- >douyin/message/action/ - 消息操作
- >/douyin/relation/friend/list/ - 用户好友列表
