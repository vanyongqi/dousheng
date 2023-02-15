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


### **运行方式**
> go build
>
> ./dousheng.exe

## **领域1：基础功能：**
- 视频Feed流，支持所有用户刷抖音，视频按照投稿时间倒序推出
- 视频投稿，支持登录用户自拍视频投稿
-  个人主页，支持查看用户的基本信息和投稿列表，注册用户流程简化
## **领域2：社交功能**
- 登录用户可以对视频点赞，在个人主页的喜欢Tab下能够查看点赞视频列表
- 支持未登录用户查看视频下的评论列表，登录用户能发表评论
- 登录用户可以关注其他用户，能够在个人主页查看本人的关注数量和粉丝数量，查看关注列表和粉丝列表
- 登录用户在消息页展示已关注的用户列表，点击用户头像进入聊天页面可以发送消息。

## **评分项目**：
- **功能实现60分**，服务能够正常运行，接口实现完整性，边界情况处理等
- **代码质量10分**，项目结构清晰，代码符合编码规范
- **服务性能10分**，数据表是否设置了合理的索引，处理了常见的性能问题
- **安全可靠20分**，是否考虑过 SQL 注入，越权等安全问题的防御方式

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
### **未实现**
- >/douyin/message/chat/ - 聊天记录
- >douyin/message/action/ - 消息操作
- >/douyin/relation/friend/list/ - 用户好友列表


**优化：**
> 日志目前是直接写入控制台信息,os.Stdout，可以自定义日志输出，如 fatal error  warning print
> message 模块没有加 好友列表没有加