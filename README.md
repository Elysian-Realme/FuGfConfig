# FuGfConfig

## 特别声明

本项目中所有内容只供学习和研究使用，不得将本项目中任何内容用于违反任何国家/地区/组织等的法律法规或相关规定的其他用途。

## 支持

本项目对 Loon、QuantumultX、AdGuard Home、Surge 提供完全支持

优先级：Loon = QuantumultX = AdGuard Home = Surge

> 经过鄙人的重构，现在提供 Host, DomainSet, Loon, QuantumultX 的支持，各位可以按需引用

> 被迫修改了一些规则的路径和名字，我尽量争取此类事件不再发生

经过我个人的测试，FuckRogueSoftware 规则集误杀应该 ~~（大概 可能）~~ 没有那么严重了~~吧~~，计划在未来本项目全面向 FuckRogueSoftware 转移，淘汰掉 以前的 AdRules

> 仅个人测试，请添加时准备好 direct 规则 (

经过测试的软件：todo

计划测试的软件：暂无

## 使用方法

### QuantumultX

建议配合 ios_rule_script 项目中的 [QuantumultX 去广告](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/QuantumultX/Advertising) 一起使用（已去重

别配合使用了，我感觉误杀有点高（~~虽然好像 FuckRogueSoftware 规则集的误杀也不低~~

优先级从高到低：

```txt
FuckRogueSoftware

AppleRules

自定义的策略组，走代理的或者走不同代理路线的

ChinaASN
DirectRules
BasicRules
```

```
# 隐私保护
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/FuckRogueSoftwareRules.conf, force-policy=FuckRogueSoftware, tag=FuckRogueSoftware, enabled=true

# 去广告
https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list, force-policy=Advertising, tag=CustomAd, enabled=true

# Telegram 代理
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/TelegramRules.conf, force-policy=Telegram, tag=TelegramRules, enabled=true

# Youtube
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/YouTubeRules.conf, force-policy=YouTube, tag=YouTubeRules, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/ChinaASN.conf, force-policy=direct, tag=ChinaASN, enabled=true

# 自定义的直连
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/DirectRules.conf, force-policy=direct, tag=CustomDirect, enabled=true
```

### Loon

#### Loon 分流规则

优先级从高到低：

```
FuckRogueSoftware
ProxyRules
TelegramRules
AppleNoChinaCDNRules
AppleRules
AppleAPIRules
AppleCDNRules
GFWRules
DirectRules
BaseRules
```

若有精准去广告的需求，请~~在未来~~使用 FuckRogueSoftware，现在大概依然有误杀

```
# 去广告
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/FuckRogueSoftwareRules.conf, policy=Advertising, tag=FuckRogueSoftware, enabled=true

# 自定义的代理
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/ProxyRules.conf, policy=PROXY, tag=CustomProxy, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/TelegramRules.conf, policy=PROXY, tag=TelegramRules, enabled=true

# Apple 规则
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Apple/AppleNoChinaCDNRules.conf, policy=AppleNoChinaCDN, tag=AppleNoChinaCDN, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Apple/AppleRules.conf, policy=Apple, tag=Apple, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Apple/AppleAPIRules.conf, policy=AppleAPI, tag=AppleAPI, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Apple/AppleCDNRules.conf, policy=AppleCDN, tag=AppleCDN, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/GFWRules.conf, policy=PROXY, tag=FuckGFW, enabled=true

# 自定义的直连
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/DirectRules.conf, policy=DIRECT, tag=CustomDirect, enabled=true

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/BaseRules.conf, policy=DIRECT, tag=BaseRules, enabled=true

```

### 对于 FuckRogueSoftware 规则的说明

此规则对某些国内软件强屏蔽，包括但不限于广告，跟踪，数据分析

在 [FuckRogueSoftware.txt](https://github.com/Elysian-Realme/FuGfConfig/blob/main/ConfigFile/DataFile/FuckRogueSoftware/FuckRogueSoftware.txt) 中可以看到部分屏蔽说明

```
# qx
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/FuckRogueSoftwareRules.conf, force-policy=FuckRogueSoftware, tag=FuckRogueSoftware, enabled=true

# loon
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/FuckRogueSoftwareRules.conf, policy=Advertising, tag=FuckRogueSoftware, enabled=true
```

### Apple 系统更新

另外，若有屏蔽 Apple 系统更新的需求，可以引用 `AppleUpdateRules` 规则集

```
# loon
AppleUpdate = select,REJECT-DROP,DIRECT,img-url = https://raw.githubusercontent.com/Koolson/Qure/master/IconSet/Color/Apple_Update.png

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Apple/AppleUpdateRules.conf, policy=REJECT ,tag=AppleUpdate, enable=true

# qx
static= AppleUpdate, reject, direct, img-url = https://raw.githubusercontent.com/Koolson/Qure/master/IconSet/Color/Apple_Update.png

https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/QuantumultX/Apple/AppleUpdateRules.conf, force-policy=AppleUpdate, tag=AppleUpdateRules, enabled=true
```

### 对于 Apple 解锁功能

请查看 [iRingo 解锁完整的 Apple 功能和集成服务](https://github.com/VirgilClyne/iRingo) 仓库

**建议优先采用上文仓库的规则**

### 对于 Apple 分流规则

请参考 [这篇文章](https://blog.royli.dev/2019/better-proxy-rules-for-apple-services)

对于本项目

AppleRules 是与本地化息息相关的规则，比如地图、天气、查找、Facetime、Apple Pay
( iCloud 上传与下载也归于此规则集

AppleCDNRules 是苹果的全球 CDN

AppleNoChinaCDNRules 是大陆没有的 CDN 节点

AppleAPIRules 是苹果的 API 域名

请把 NoChinaCDN 和 APIRules 放在最前面

#### 使用中国区账号（App Store + iCloud）

AppleRules 直连

AppleCDNRules 直连

AppleNoChinaCDNRules 代理

AppleAPIRules 直连

#### 使用美国区账号（App Store + iCloud）

AppleRules 直连

AppleCDNRules 直连

AppleNoChinaCDNRules 代理

AppleAPIRules 代理

建议 AppleAPIRules 依然直连，上文是根据上述文章给出的建议，请结合自身情况使用

### Loon 插件

#### DNSMap

对常用的国内域名进行 DNS 解析分流，国内走国内的各大 Doh

```
# DNS 映射
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonPlugin/DNSMap.plugin, tag=DNS Map, enabled=true
```

#### DNSMapAd

对一些广告域名 DNS 解析重定向至 `127.0.0.1` 或 `0.0.0.0`

有用没用的诸君自己看着用吧

```
# DNS 去广告映射
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonPlugin/DNSMapAd.plugin, tag=DNS Map Ad, enabled=true
```

#### 抖音屏蔽

部分规则来自 <https://gist.github.com/JamesHopbourn/b5cf9219bdacfa8b6dbb3414276c149b>

在此表示感谢

```
# FuckDouyin
https://raw.githubusercontent.com/Elysian-Realme/FuGfConfig/main/ConfigFile/Loon/LoonPlugin/FuckDouyin.plugin, proxy=Advertising, tag=Fuck Douyin, enabled=true
```

### 光明的未来

可以预见的未来越来越光明啦

特此提前准备了白名单模式，随时准备敬献。

## 项目路线图

- [x] 自动化根据规则集生成适配不同客户端的规则文件
- [x] 去重，根据域名后缀去重
- [] 构造域名后缀树，从 HashMap 切换到域名后缀树
- [] 对大佬们的 AdGuard 规则进行去重合并

## 感谢

本项目的数据大部分来自一下项目，在此对他们表示感谢（以下排名不分先后

[GetSomeFries](https://github.com/VirgilClyne/GetSomeFries)

[ios_rule_script](https://github.com/blackmatrix7/ios_rule_script)

[NextDNS 维护的系统级跟踪列表](https://github.com/nextdns/metadata/tree/master/privacy/native)

[Shadowrocket-ADBlock-Rules](https://github.com/h2y/Shadowrocket-ADBlock-Rules)

[neohosts](https://github.com/neoFelhz/neohosts)

[gfwlist](https://github.com/gfwlist/gfwlist)

[SS-Rule-Snippet](https://github.com/Hackl0us/SS-Rule-Snippet#%E5%85%B3%E4%BA%8E%E9%A1%B9%E7%9B%AE)

[ACL4SSR](https://github.com/ACL4SSR/ACL4SSR/tree/master)

[anudeepND-blacklist](https://github.com/anudeepND/blacklist)

[neodevhost](https://github.com/neodevpro/neodevhost)

[BlueSkyXN-AdGuardHomeRules](https://github.com/BlueSkyXN/AdGuardHomeRules)

[WindowsSpyBlocker](https://github.com/crazy-max/WindowsSpyBlocker)

[hoshsadiq/adblock-nocoin-list](https://github.com/hoshsadiq/adblock-nocoin-list/)

## Fuck you Gitcode

Fuck you, Gitcode.

皇后大道西又皇后大道東

皇后大道東轉皇后大道中

皇后大道東上為何無皇宮

皇后大道中人民如潮湧

有個貴族朋友在硬幣背後

青春不變名字叫做皇后

每次買賣隨我到處去奔走

面上沒有表情卻匯聚成就

知己一聲拜拜遠去這都市

要靠偉大同志搞搞新意思

照買照賣樓花處處有單位

但是旺角可能要換換名字

皇后大道西又皇后大道東

皇后大道東轉皇后大道中

皇后大道東上為何無皇宮

皇后大道中人民如潮湧

這個正義朋友面善又友善

因此批準馬匹一週跑兩天

百姓也自然要鬥快過終點

若做大國公民祇須身有錢

知己一聲拜拜遠去這都市

要靠偉大同志搞搞新意思

冷暖氣候同樣影響這都市

但是換季可能靠特異人士

8964
