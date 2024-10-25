package main

import (
	"context"
	"fmt"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/imthaghost/goclone/pkg/crawler"
	"github.com/imthaghost/goclone/pkg/html"
)

func main() {
	domain := "https://learn.lianglianglee.com"
	projectPath := `/Users/xiaoxiang/GolandProjects/goclone/learn.lianglianglee.com`
	// finish `/专栏/10x程序员工作法`,
	zhuanlan := []string{`/专栏/12步通关求职面试-完`, `/专栏/22 讲通关 Go 语言-完`, `/专栏/24讲吃透分布式数据库-完`, `/专栏/300分钟吃透分布式缓存-完`, `/专栏/AB 测试从 0 到 1`, `/专栏/AI技术内参`, `/专栏/Android开发高手课`, `/专栏/CNCF X 阿里巴巴云原生技术公开课`, `/专栏/DDD 微服务落地实战`, `/专栏/DDD实战课`, `/专栏/DevOps实战笔记`, `/专栏/Dubbo源码解读与实战-完`, `/专栏/ElasticSearch知识体系详解`, `/专栏/Flutter入门教程`, `/专栏/Flutter核心技术与实战`, `/专栏/Go 语言项目开发实战`, `/专栏/Go语言核心36讲`, `/专栏/JVM 核心技术 32 讲（完）`, `/专栏/Java 业务开发常见错误 100 例`, `/专栏/Java 并发编程 78 讲-完`, `/专栏/Java 并发：JUC 入门与进阶`, `/专栏/Java 性能优化实战-完`, `/专栏/Java 核心技术面试精讲`, `/专栏/JavaScript 进阶实战课`, `/专栏/Java并发编程实战`, `/专栏/Java核心技术面试精讲`, `/专栏/Jenkins持续交付和持续部署`, `/专栏/Kafka核心技术与实战`, `/专栏/Kafka核心源码解读`, `/专栏/Kubernetes 从上手到实践`, `/专栏/Kubernetes 实践入门指南`, `/专栏/Kubernetes入门实战课`, `/专栏/Linux内核技术实战课`, `/专栏/Linux性能优化实战`, `/专栏/MySQL实战45讲`, `/专栏/MySQL实战宝典`, `/专栏/Netty 核心原理剖析与 RPC 实践-完`, `/专栏/OAuth2.0实战课`, `/专栏/OKR组织敏捷目标和绩效管理-完`, `/专栏/OpenResty从入门到实战`, `/专栏/PyTorch深度学习实战`, `/专栏/Python核心技术与实战`, `/专栏/Python自动化办公实战课`, `/专栏/RE实战手册`, `/专栏/RPC实战与核心原理`, `/专栏/Redis 核心原理与实战`, `/专栏/Redis 核心技术与实战`, `/专栏/Redis 源码剖析与实战`, `/专栏/RocketMQ 实战与进阶（完）`, `/专栏/Serverless 技术公开课（完）`, `/专栏/Serverless进阶实战课`, `/专栏/ShardingSphere 核心原理精讲-完`, `/专栏/Spark性能调优实战`, `/专栏/Spring Boot 实战开发`, `/专栏/Spring Security 详解与实操`, `/专栏/SpringCloud微服务实战（完）`, `/专栏/Spring编程常见错误50例`, `/专栏/To B市场品牌实战课`, `/专栏/Tony Bai · Go语言第一课`, `/专栏/Vim 实用技巧必知必会`, `/专栏/Web 3.0入局攻略`, `/专栏/WebAssembly入门课`, `/专栏/Web漏洞挖掘实战`, `/专栏/ZooKeeper源码分析与实战-完`, `/专栏/etcd实战课`, `/专栏/iOS开发高手课`, `/专栏/中间件核心技术与实战`, `/专栏/互联网消费金融高并发领域设计`, `/专栏/人工智能基础课`, `/专栏/从 0 开始学架构`, `/专栏/从0开始做增长`, `/专栏/从0开始学大数据`, `/专栏/从0开始学微服务`, `/专栏/从0开始学游戏开发`, `/专栏/代码之丑`, `/专栏/代码精进之路`, `/专栏/全解网络协议`, `/专栏/分布式中间件实践之路（完）`, `/专栏/分布式技术原理与实战45讲-完`, `/专栏/分布式技术原理与算法解析`, `/专栏/分布式金融架构课`, `/专栏/分布式链路追踪实战-完`, `/专栏/前端工程化精讲-完`, `/专栏/动态规划面试宝典`, `/专栏/即时消息技术剖析与实战`, `/专栏/反爬虫兵法演绎20讲`, `/专栏/后端技术面试38讲`, `/专栏/周志明的架构课`, `/专栏/大厂广告产品心法`, `/专栏/大厂设计进阶实战课`, `/专栏/大规模数据处理实战`, `/专栏/如何设计一个秒杀系统`, `/专栏/安全攻防技能30讲`, `/专栏/容器实战高手课`, `/专栏/容量保障核心技术与实战`, `/专栏/左耳听风`, `/专栏/微服务质量保障 20 讲-完`, `/专栏/成为AI产品经理`, `/专栏/打造爆款短视频`, `/专栏/技术与商业案例解读`, `/专栏/技术管理实战 36 讲`, `/专栏/技术领导力实战笔记`, `/专栏/持续交付36讲`, `/专栏/推荐系统三十六式`, `/专栏/操作系统实战45讲`, `/专栏/朱赟的技术管理课`, `/专栏/机器学习40讲`, `/专栏/李智慧 · 高并发架构实战课`, `/专栏/架构设计面试精讲`, `/专栏/案例上手 Spring Boot WebFlux（完）`, `/专栏/正则表达式入门课`, `/专栏/消息队列高手课`, `/专栏/深入剖析 MyBatis 核心原理-完`, `/专栏/深入剖析Java新特性`, `/专栏/深入剖析Kubernetes`, `/专栏/深入拆解Java虚拟机`, `/专栏/深入拆解Tomcat  Jetty`, `/专栏/深入浅出 Docker 技术栈实践课（完）`, `/专栏/深入浅出 Java 虚拟机-完`, `/专栏/深入浅出云计算`, `/专栏/深入浅出分布式技术原理`, `/专栏/深入浅出区块链`, `/专栏/深入浅出可观测性`, `/专栏/深入浅出计算机组成原理`, `/专栏/深入理解 Sentinel（完）`, `/专栏/由浅入深吃透 Docker-完`, `/专栏/白话法律42讲`, `/专栏/白话设计模式 28 讲（完）`, `/专栏/硅谷产品实战36讲`, `/专栏/程序员的个人财富课`, `/专栏/程序员的数学基础课`, `/专栏/程序员的数学课`, `/专栏/程序员的测试课`, `/专栏/程序员进阶攻略`, `/专栏/编译原理之美`, `/专栏/编译原理实战课`, `/专栏/计算机基础实战课`, `/专栏/许式伟的架构课`, `/专栏/说透低代码`, `/专栏/说透性能测试`, `/专栏/赵成的运维体系管理课`, `/专栏/超级访谈：对话张雪峰`, `/专栏/超级访谈：对话毕玄`, `/专栏/超级访谈：对话汤峥嵘`, `/专栏/超级访谈：对话玉伯`, `/专栏/趣谈网络协议`, `/专栏/跟着高手学复盘`, `/专栏/软件工程之美`, `/专栏/软件测试52讲`, `/专栏/透视HTTP协议`, `/专栏/重学操作系统-完`, `/专栏/重学数据结构与算法-完`, `/专栏/陈天 · Rust 编程第一课`, `/专栏/零基础入门Spark`, `/专栏/领域驱动设计实践（完）`, `/专栏/高并发系统实战课`, `/专栏/高并发系统设计40问`, `/专栏/高楼的性能工程实战课`}
	for _, item := range zhuanlan {
		oneZhuanlan := domain + item
		toVisitPage := []string{}

		c := colly.NewCollector()
		// Find and visit all links
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if strings.Contains(e.Attr("href"), `专栏`) {
				toVisitPage = append(toVisitPage, domain+e.Attr("href"))
			}
		})

		c.Visit(oneZhuanlan)

		for _, page := range toVisitPage {
			jar, _ := cookiejar.New(&cookiejar.Options{})

			var (
				newPagePath string
				err         error
			)

			if newPagePath, err = crawler.Collector(context.TODO(), page, projectPath, jar, "", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"); err != nil {
				panic(err)
			}
			if err = html.LinkRestructure(projectPath, newPagePath); err != nil {
				panic(err)
			}
			time.Sleep(15 * time.Second)
		}

		fmt.Println("finished:" + item)
	}
}