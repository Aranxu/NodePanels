# NodePanels

>一款使用简单功能全面的多服务器监控面板，一条指令即可对你的服务器了如指掌

<a href="https://t.me/nodepanels_group" target="_blank"><img src="https://s4.ax1x.com/2022/03/01/bQQs2t.png" alt="bQQs2t.png" border="0" /></a>
<a href="https://qm.qq.com/cgi-bin/qm/qr?k=rD7CokCfFoJ1OdHCr6CwaQGftWpCygFr&jump_from=webapi" target="_blank"><img src="https://s4.ax1x.com/2022/03/01/bQlaQ0.png" alt="bQlaQ0.png" border="0" /></a>
<a href="https://nodepanels-1256221051.file.myqcloud.com/web/static/pic/wechat_qrcode.jpg" target="_blank"><img src="https://s4.ax1x.com/2022/03/01/bQlWy6.png" alt="bQlWy6.png" border="0" /></a>
<a href="mailto:support@nodepanels.com"><img src="https://s4.ax1x.com/2022/03/01/bQ1AXV.png" alt="bQ1AXV.png" border="0" /></a>

## 网站： https://nodepanels.com

探针包：https://github.com/Aranxu/Nodepanels-probe

工具包：https://github.com/Aranxu/Nodepanels-tool

守护进程：https://github.com/Aranxu/Nodepanels-daemon

## 起因
因为手上的服务器挺多的，想要统一进行管理监控，尝试了一些已有的产品，最终使用了nodequery，简单明了。但是2020年9月开始发现网站经常无法打开、打开缓慢、签名过期等问题，而且一个号只能添加10台服务器，不支持windows。宝塔面板在2020年10月开始需要验证手机号，原因不深究。感觉市面上好像挺缺这种产品的，本人所在公司正好完成了一个云服务器监控系统，监控的是某机构某地市所有机房中的服务器，目前有几万台，主打安全。这不刚好对口了吗，于是此项目开起来了。

## 规划
**关于功能：** 一开始是想实现nodequery的功能就行了，监控服务器的几个指标信息。但觉得这样没什么亮点，后面计划向宝塔的功能靠拢，但难度确实摆在那里。目前先从简单监控开始，一步一步来。

**关于运营：** 此面板的目标是长久运营，会消耗大量的资源（服务器，开发，维护等），不可能一直用爱发电。后期会加入会员系统、投放广告维持运营，当然越早注册的用户会获得越大的优惠。

**关于隐私：** 本项目一个最重要的点就是用户隐私为第一位，绝不强制用户提供任何个人信息，账号在本系统产生的所有数据均可删除。（真的被某些产品恶心到了）

## 进展（每月更新）

详细更新记录：<a href="https://nodepanels.com/info/change" target="_blank">https://nodepanels.com/info/change</a>

<details>
  <summary>2022-11</summary>
  <br>
  <ul>
    <li>增加webhook告警通知方式</li>
  </ul>
</details>

<details>
  <summary>2022-09</summary>
  <br>
  <ul>
    <li>修改添加服务器方式，支持免配置添加</li>
  </ul>
</details>

<details>
  <summary>2022-06</summary>
  <br>
  <ul>
    <li>磁盘管理增加磁盘使用率历史数据</li>
    <li>支持磁盘告警</li>
    <li>修改告警规则操作方式，可批量设置</li>
    <li>移动端展示网络速率</li>
  </ul>
</details>

<details>
  <summary>2022-05</summary>
  <br>
  <ul>
    <li>支持获取实时数据，粒度为2秒</li>
    <li>探针每10分钟更新一次系统软硬件信息</li>
    <li>CPU、内存告警逻辑转为服务端实现</li>
    <li>探针数据上报增加备用域名，提高上报成功率</li>
    <li>探针升级至Go1.18，支持更多系统架构。更新windows端应用图标，优化程序逻辑，业务分离</li>
  </ul>
</details>

<details>
  <summary>2022-04</summary>
  <br>
  <ul>
    <li>服务器列表增加网络速率</li>
    <li>服务器列表增加详细信息弹框</li>
    <li>服务器列表新增网格显示方式</li>
    <li>增加删除账号功能，路径：用户中心</li>
    <li>新增“系统设置”页面</li>
    <li>增加traceroute功能（windows/linux），路径：服务器 -> 性能测试 -> 路由跟踪</li>
  </ul>
</details>

<details>
  <summary>2022-03</summary>
  <br>
  <ul>
    <li>前端框架重构完成</li>
    <li>后台适配新前端，系统前后分离</li>
    <li>新增SSH功能页</li>
    <li>新增”推荐有奖“</li>
    <li>更换稳定支付接口</li>
    <li>提高系统安全性</li>
  </ul>
</details>

<details>
  <summary>2022-02</summary>
  <br>
  <ul>
    <li>增加网络实时速率</li>
    <li>页面重构（耗时长）</li>
  </ul>
</details>

<details>
  <summary>2022-01</summary>
  <br>
  <ul>
    <li>网络测速增加数据分析，显示各地区最大最小平均值。</li>
    <li>文件管理器根据文件后缀展示对应icon。</li>
    <li>增加用户中心。</li>
    <li>增加费用中心，提供配额套餐。</li>
    <li>支持密码修改（电脑端）。</li>
    <li>支持密码修改。</li>
    <li>支持邮箱设置。</li>
    <li>创建QQ交流群，微信交流群，TG交流群、TG公告群。</li>
    <li>修复文件管理器无法预览文件。</li>
    <li>修复微信、QQ告警机器人</li>
    <li>支持cpu、内存、离线告警恢复通知。</li>
  </ul>
</details>

<details>
  <summary>2021-12</summary>
  <br>
  <ul>
    <li>上线文件管理（Linux）</li>
    <li>修复网络测速异常</li>
    <li>Linux网络测试无需sudo</li>
    <li>屏蔽微信告警、QQ告警，恢复显示telegram验证码</li>
    <li>页面增加埋点</li>
    <li>工具包支持Windows ARM</li>
    <li>提高探针通讯安全性</li>
  </ul>
</details>

<details>
  <summary>2021-11</summary>
  <br>
  <ul>
    <li>SSH页面设计，流程规划</li>
    <li>文件管理：文件列表、文件树、书签、新建文件/文件夹、复制、重命名、黏贴、移动、回收站、权限、属性。（未上线）</li>
  </ul>
</details>

<details>
  <summary>2021-10</summary>
  <br>
  <ul>
    <li>上线DNS设置功能（linux）</li>
    <li>上线主机名设置功能（linux/windows）</li>
    <li>上线YUM源配置功能（linux）</li>
    <li>上线时间管理功能（linux）</li>
    <li>上线环境变量列表功能（linux）</li>
    <li>上线服务列表功能（linux）</li>
  </ul>
</details>

<details>
  <summary>2021-09</summary>
  <br>
  <ul>
    <li>修复linux端网络测速功能（暂时需要系统带有sudo）</li>
    <li>修复探针虚假不在线异常</li>
    <li>DNS设置功能（未放出，待系统管理功能全部完成再上线）</li>
    <li>主机名设置功能（同上）</li>
  </ul>
</details>

<details>
  <summary>2021-08</summary>
  <br>
  <ul>
    <li>探针新增磁盘数据采集能力</li>
    <li>探针修改指令处理逻辑（后续将更方便更快捷的提供新功能）</li>
    <li>探针新增网络测速功能（已完成windows端，linux端将在近期上线）</li>
    <li>网站新增性能测试-网络测速页面</li>
    <li>升级存量探针版本至v1.0.2</li>
    <li>更新探针安装脚本</li>
  </ul>
</details>

<details>
  <summary>2021-07</summary>
  <br>
  <ul>
    <li>增加设置页面</li>
    <li>增加分享功能</li>
    <li>告警通知移至设置页面</li>
    <li>数据存储切换至时序数据库</li>
    <li>架构调整优化</li>
    <li>流量校正移至网络详情页</li>
  </ul>
</details>

<details>
  <summary>2021-06</summary>
  <br>
  <ul>
    <li>适配ARM服务器</li>
    <li>适配Windows服务器</li>
    <li>适配手机端页面</li>
    <li>升级邮箱系统</li>
  </ul>
</details>

<details>
  <summary>2021-05</summary>
  <br>
  <ul>
    <li>对已有功能查漏补缺</li>
    <li>上线正式环境，试运行</li>
    <li>目前上线功能为系统最基础的功能点</li>
  </ul>
</details>

<details>
  <summary>2021-04</summary>
  <br>
  <ul>
    <li>开发微信告警机器人</li>
    <li>开发QQ告警机器人</li>
    <li>开发Telegram告警机器人</li>
    <li>架构优化，提高系统可维护性，提高探针稳定性</li>
  </ul>
</details>

<details>
  <summary>2021-03</summary>
  <br>
  <ul>
    <li>完善各指标数据查询功能，并支持自定义时间查询，不同粒度查询。优化cpu、内存、swap、磁盘、流量的数据查询</li>
    <li>初步完成告警模块功能开发，暂未接入通知接口</li>
    <li>完善探针管理页面</li>
    <li>修复登录超时后跳转至首页bug</li>
  </ul>
</details>

<details>
  <summary>2021-02</summary>
  <br>
  <ul>
    <li>服务器列表页展示cpu、内存、swap、磁盘、流量等指标数据，丰富及美化列表显示内容，直观看出服务器租赁剩余时长和剩余流量</li>
    <li>探针优化，降低内存使用率</li>
  </ul>
</details>

<details>
  <summary>2021-01</summary>
  <br>
  <ul>
    <li>完成首页选型开发</li>
    <li>后台用户设计及开发</li>
    <li>数据采集模块划分优化</li>
  </ul>
</details>

<details>
  <summary>2020-12</summary>
  <br>
  <ul>
    <li>初步完成cpu，内存，磁盘，进程数据采集和页面设计及开发</li>
    <li>增加“压缩采集数据”程序，使用算法优化数据存储方式，减少数据库压力，提高数据处理能力</li>
  </ul>
</details>

<details>
  <summary>2020-11</summary>
  <br>
  <ul>
    <li>设计后台结构，前期没有那么多可用的机器和服务，暂时不以互联网项目去设计，后期看实际情况升级</li>
    <li>设计并测试探针可行方案</li>
    <li>确定前端模板（先决定自己做，后期第二版会找专业前端和UI改版）</li>
    <li>对探针和采集层进行压测，在本地网络较差情况下支持单机2000并发，在云服务器上测试支持单机10000+并发，大于实际应用场景，后期如规模增长会增加机器保证服务质量。目前计划采集机部署在亚太区、美区、欧洲区，后期视情况增加</li>
    <li>完成服务器总览页（列表）设计和页面开发</li>
    <li>完成服务器信息页设计和页面开发</li>
    <li>完成服务器、分组管理开发</li>
  </ul>
</details>

## 总进度
### 一阶段（2022-10-01更新）
类目|状态|进度
--|:--:|--:
首页|已完成|===================================100%
用户模块|已完成|===================================100%
探针开发|已完成|===================================100%
服务器总览页|已完成|===================================100%
服务器信息页|已完成|===================================100%
负载信息页|已完成|===================================100%
CPU信息页|已完成|===================================100%
内存信息页|已完成|===================================100%
磁盘信息页|已完成|===================================100%
进程信息页|已完成|===================================100%
网络信息页|已完成|===================================100%
SSH操作页|已完成|===================================100%
文件管理页|进行中|===================================100%
软件安装页|未确认|0%
安全管理|未确认|0%
系统信息|进行中|===　　　　　　　　　　　　　　　　　　　　　　10%
性能测试|进行中|=========　　　　　　　　　　　　　　　　　　30%
告警管理|已完成|===================================100%
探针管理|已完成|===================================100%
