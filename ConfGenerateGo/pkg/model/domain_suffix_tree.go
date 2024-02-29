// 域名后缀树
// 自己尝试实现一下
// 生产环境请使用 这个大佬的项目 https://github.com/golang-infrastructure/go-domain-suffix-trie

// 我要计划实现的东西：
// 给定一个 基础域名字符串集合，通过这个集合构建出一个域名后缀树
// 给定一个待匹配的域名字符串集合，跟后缀树比对，不在后缀树里的可以选择插入或者是加入一个 list 最后返回
// 本质上先给本项目服务，上面的描述也是本项目的需求

package model

// 定义 后缀树的一个节点
type Node struct {
	Data         string
	ChildrenNode []*Node
}

// 定义一个后缀树
type DomianSuffixTree struct {
	Data string
}

// 初始化函数

// 插入函数
// 查找函数
// 输出函数
