# easygo

**简易的 开源的 面向开发的 go单应用web脚手架**



---



框架采用 **wire**依赖注入

**gin**做web路由

**gorm**做数据库映射

配置文件采用**yaml文件**读取

配置文件集成了redis、jwt、mysql等常用功能

默认需要打开redis:6379，和mysql :3306






---

规定 <br>

controller层和model层通过schema交互 <br>
model层和dao层交互通过dao层定义struct交互

---

目录结构

├─cmd &nbsp;&nbsp;&nbsp; //文件的启动目录 <br>
├─internal // 对内函数 <br>
│ ├─config      
│ ├─controller  
│ ├─dao  
│ ├─entity  
│ ├─middleware  
│ ├─model  
│ ├─router  
│ └─schema  
├─log  
└─pkg //代码仓库 ，可引用 复用  <br>
&nbsp;&nbsp;&nbsp;├─cryptox  
&nbsp;&nbsp;&nbsp;├─jwt  
&nbsp;&nbsp;&nbsp;├─logger  
&nbsp;&nbsp;&nbsp;├─redis  
&nbsp;&nbsp;&nbsp;├─request  
&nbsp;&nbsp;&nbsp;├─trace   
&nbsp;&nbsp;&nbsp;└─wrapper


---

v0.0.1在工作室开源发布<br>

[0.0.1](http://gitlab.qnxg.net/qgo) <br>
[MIT](http://gitlab.qnxg.net/qgo/qgo/-/blob/master/LICENSE)

v0.0.2<br>

- 增加了redis和jwt的默认支持
- 修改部分冗余代码
- 增加测试层

---

### 预告

- 完善测试层，规范代码结构
- 删除多余函数
- 完善升级上下文和错误处理，方便进行debug时的堆栈追踪
- json解码修改为proto解码