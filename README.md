# 关于秘书
秘书的工作非常重要，尤其在dev放弃owner权之后，秘书是全系列的唯一中心化实体，担负着维护全系列的价值的重任。秘书除了要负责、及时提交新闻候选列表之外，还有一些重要的要求：
* 秘书不要遗过真正的大新闻，需要秘书每天关注世界各处的事件。秘书需要为每天提交一次且至少三条新闻，不得有遗漏的日期。超过30天后，不能再提交该日期的新闻，一旦发生这将是秘书的重大失职。
* 描述只涉及事实，不加入感情色彩，不加入分析。
* 尽可能不记录非确定信息，所有的非确定的信息都需要加上“据称”，“可能”，“非可靠来源”之类的定语。
* 描述尽量简明扼要， 一是因为需要节约气费， 二是因为冗余词汇可能会带入立场，影响历史记述的客观性。
* 作为编年史，描述不带立场，不加评论，千秋功罪留与后人评说，我们这个项目只需要做好事实描述和记录。
* 希望秘书能在社交媒体上对本项目进行维护和推广。

后续在history合约中会加入秘书选举功能，一个NFT一票,希望这个机制可以监督秘书更好的为社区工作。

# 关于投票
* 用户需要首先使用 escrow deposit命令质押产生票权， 票权每天增加1/30，30天后票权为质押金额的100%。  
* 投票必须选择3条新闻的编号，编号从0开始递增，在history vote命令中用“-p 0,2,3”格式的参数提交投票。
* 需要认真投票
投票命中最终结果的第一名，获得的分红权重为 票权/3 * 8。  
投票命中最终结果的第二名，获得的分红权重为 票权/3 * 4。  
投票命中最终结果的第三名，获得的分红权重为 票权/3 * 2。  
其他情况获得的分红权重为 票权/3 * 1。  
这一设计是希望投票人认真的选择。
* 投票之后，可以用history settleVote结算分红，可以一次性把多次结算的分红通过history claim一次性提取到自己帐号。

# GoChainHis
ChainHis CLI utils

cd this project folder, then type "go run ." to execute. Also, it is OK to use 'go install .' to build and install GoChainHis to your $GOPATH/bin/ folder, then use GoChainHis directly.

## escrow deposit
- go run . escrow deposit --amount 0.1

## escrow withdraw
- go run . escrow withdraw --amount 0.1

## escrow dump
- go run . escrow dump

## history dump
"-n" param will dump continous date's history infomation.

e.g. :: 
- go run . history dump -d 2023-06-22 -n 2 
- go run . history dump -d 2023-06-23

## history submit [secretarty only]
submit command will submit candidate news for voters.

- go run . history submit -d 2023-06-23 -s subs/sub.0623.json
  
Note:: 
1). Date field of sub.0623.json must same as -d params e.g. 2023-06-23  
2). sub.0623.json should be as the following format:  
```
{
    "Date": "2023-06-23 UTC",
    "BigNews": [
        "BTC触及31456$创一年新高",
        "离岸人民币兑美元跌破7.22日内跌超250点",
        "美国深海潜水器泰坦内爆,5名乘员死亡",
        "Coinbase在最高院赢得与美国证交会的诉讼"
    ]
}
```
3). it is suggeust that the content should wrap/segment with <CN></CN> and <EN></EN> tags for internationalize, e.g.:
```
    "<CN><港媒:盒马鲜生最快11月IPO</CN><EN>Hong Kong media: Hema Xiansheng will IPO in November ASAP</EN>"
```

## history resolve [secretarty only]
resolve command will resolve a date's vote result once **VoteEndTm expired**.
- go run . history resolve -d 2023-06-22

## history mintAndAuc [secretarty only]
mintAndAuc command will mint NFT and issue its auction.
Note:: **don't miss festivals** if it is!
- go run . history mintAndAuc -d 2023-06-22 -e 端午节
- go run . history mintAndAuc -d 2023-10-01 -e 国庆节,中秋节
  
## history vote
go run . history vote -d 2023-06-23 -p 0,2,3

Note:: -p parameter must have 3 indexes which increased from 0 to indicate which news are the prefered big newes of the day, user can use **history dump** comamnd to see the details first.


