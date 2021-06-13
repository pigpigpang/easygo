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
根据 [golang-standards](https://github.com/golang-standards/project-layout) 改编的单应用web快速开发目录


├─cmd <br>
├─config <br>
├─docs <br>
├─internal <br>
│  ├─common <br>
│  │  └─config <br>
│  ├─controller <br>
│  ├─dao <br>
│  ├─entity <br>
│  ├─middleware <br>
│  ├─model <br>
│  ├─router <br>
│  └─schema <br>
├─log <br>
├─pkg <br>
│  ├─cryptox <br>
│  ├─jwt <br>
│  ├─logger <br>
│  ├─redis <br>
│  ├─request <br>
│  ├─trace <br>
│  └─wrapper <br>
└─ go.mod

### /cmd
本项目的主干。

每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，/cmd/myapp)。

不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。你会惊讶于别人会怎么做，所以要明确你的意图!

通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。


### /internal
私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。注意，你并不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录。

我们的单应用程序的基本函数层，就在internal里。这意味着，当我们需要给我们的应用添加功能时，我们需要向internal的一级子目录的所有文件进行更改（包括controller、model等）  

我们的应用不需要经常更改的层在 /internal/common ，这里是一些不对外暴露的内部的初始化等层

需要格外注意的是，当我们要更改目录结构时，我们可能需要对common的wire做一些修改

你可以选择向 internal 包中添加一些额外的结构，以分隔共享和非共享的内部代码。这不是必需的(特别是对于较小的项目)，但是最好有有可视化的线索来显示预期的包的用途。

### /pkg
外部应用程序可以使用的库代码。其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思:-)

注意，internal 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。

在我们的单应用目录结构中，pkg存放我们可能会复用的方法

往往根据代码重构思想，当一个方法使用超过3次时，我们就需要考虑封装方法，而不是大量的复制粘贴等操作。这时候，pkg库是一个比较良好的方法。

### /test

这里需要注意的是，我在 internal 的 controller 层预留了一个test函数，用来测试基本功能(主要是添加新的common以后的测试)，比如 redis ， jwt ，测试基本功能等

这里的test目录，是我们可以自己mock一些数据，自己构造请求，测试结果是否与预期一致。(我个人认为 类似于我们手动的postman测试等)


### /docs

docs一般放置我们的接口文档，供交接所用  

如果有需求，也可以放置用户手册等文档

---

**v0.0.1**在工作室开源发布<br>

[0.0.1](http://gitlab.qnxg.net/qgo) <br>
[MIT](http://gitlab.qnxg.net/qgo/qgo/-/blob/master/LICENSE)

**v0.0.2**<br>

- 增加了redis和jwt的默认支持
- 修改部分冗余代码
- 增加测试层


**v0.0.3**<br>
- 根据[golang-standards](https://github.com/golang-standards/project-layout) 修改目录结构


---

### 预告

- 完善测试层，规范代码结构
- 删除多余函数
- 完善升级上下文和错误处理，方便进行debug时的堆栈追踪
- json解码修改为proto解码