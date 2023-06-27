# 关于秘书
秘书的工作非常重要，尤其在dev放弃owner权之后，秘书是全系列的唯一中心化实体，担负着维护全系列的价值的重任。秘书除了要负责、及时提交新闻候选列表之外，还有一些重要的要求：
* 秘书不要遗过真正的大新闻，需要秘书每天关注世界各处的事件。秘书需要为每天提交一次且至少三条新闻，不得有遗漏的日期。超过30天后，不能再提交该日期的新闻，一旦发生这将是秘书的重大失职。
* 描述只涉及事实，不加入感情色彩，不加入分析。
* 尽可能不记录非确定信息，所有的非确定的信息都需要加上“据称”，“可能”，“非可靠来源”之类的定语。
* 描述尽量简明扼要， 一是因为需要节约气费， 二是因为冗余词汇可能会带入立场，影响历史记述的客观性。
* 作为编年史，描述不带立场，不加评论，千秋功罪留与后人评说，我们这个项目只需要做好事实描述和记录。
* 希望秘书能在社交媒体上对本项目进行维护和推广。
* 秘书需要确保定制的NFT主图必须和NFT基础图的风格一致。

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

# 关于拍卖
* 在秘书通过mintAndAuc创建并发起拍卖之后,NFT在合约内启动了拍卖，拍卖前的NFT所有者是Card合约。  
* 拍卖方式采用荷兰拍，先到先得。价格从初始拍卖价格每隔1小时降价1%，最低降价到0.01ETH。初始拍卖价格是最近十次成交价的EMA平均价的1.5倍，且不低于0.02ETH(曾经是0.04ETH，已修改)。  
* 可以在history dump命令中看到初始拍卖价AucInitPrice，最新拍卖价CurPrice，最终成交价FinalPrice。  
* 通过 history bid命令即可以最新拍卖价竞拍。

# 关于奖励
现在有三部分奖励。
* 运维奖励。开发团队和秘书分别获得拍卖金额5%的奖励，合计10%。 秘书或者开发团队可以通过history settleOps结算这部分奖励。
* 投票奖励。凡是参与了投票的帐户，根据投票命中率和票权计算其分红权重。投票人根据投票权重分享拍卖金额的80%。投票人可以通过history settleVote结算这部分奖励，可以在history dump最后打印中看到这部分奖励。
* NFT持有人的奖励。早期的NFT持有人，可以平分后来的NFT拍卖金额的10%（包括自己这个NFT的奖励）。NFT持有人可以通过history settleCard结算这部分奖励，可以通过比如"hisotry dumpReward -l 2023-06-22 -r 2023-06-23" 命令查看是否已经结算。

需要强调的是，大部分分红都给到投票人是为了保证系统的长期运行。NFT是全社区的长期利益同路人，但是更广大的投票人是NFT持有人参与社区的原动力，所以必须向投票人倾斜。  

按照目前的设想，在代码全部稳定后，开发团队会放弃owner权限，其5%的奖励作为社区的公积金留存备用。

# NFT的定制和呈现
* 受开发资源限制，全系列现在都用的同一个默认主图显示。预期在chainhis.net建设之后，可以叠加NFT的新闻/节日/投票信息/自定的附图/指定的祝福辞。
* 秘书可以用"card setBaseImage"命令定制某个日期的NFT的背景图，NFT持有人可以向秘书申请定制其背景图。但是定制的主图必须和NFT基础图的风格一致，需要由秘书确保这一点。
* NFT持有人可以通过card customize定制nft的祝福辞和附图，以及选择主图呈现方式（最后这点需等待chainhis.net网站建设）。

# GoChainHis
ChainHis CLI utils  

cd this project folder, then type "go run ." to execute. Also, it is OK to use 'go install .' to build and install GoChainHis to your $GOPATH/bin/ folder, then use GoChainHis directly.  

需要先参考conf.json.demo,编辑自己的conf.json文件，只需要修改其中的MyAddr和MyPrivKey。注意::需要确保MyAddr地址**大小写正确**。  
自己的conf.json要么放在默认的$Home/.GoChainHis/目录下,要么通过--conf参数指定其全路径。  

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


