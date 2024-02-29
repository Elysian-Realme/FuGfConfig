# 自动整合去广告规则

## 项目逻辑

以 [ios_rule_script Loon Advertising](https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list) 为基础，对其进行扩充

扩充内容：

1. [ios_rule_script QuantumultX Advertising](https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list)
2. [ios_rule_script Loon Advertising_Domain](https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/release/rule/Loon/Advertising/Advertising_Domain.list)
3. 自己搜集的
4. [CustomAdRules](https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomAdRules.conf)

## 程序不足

1. 无法处理 IPV6 规则
2. 对于 Qx 中的低优先级规则，直接予以丢弃
3. 对于 Qx 中的规则优先级，不会优先考虑
