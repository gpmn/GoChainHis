# 关于ChainHis
ChainHis出发点是创建一个去中心化的历史书，可以认为是一个偏文化的区块链项目。  
为了让项目长期生存下去，为了在项目内形成内循环激励机制，确保项目长期生存，所以设置了一些分红和拍卖的机制。但是根本出发点不是为了割韭菜或者圈钱，希望大家理性参与。  
项目是业余时间开发，开发和参与的人都很少。前期只有一些朋友参与，希望以后能够得到社区的喜爱。  
合约分为三部分，escrow质押模块，history核心逻辑模块，card是NFT模块。目前escrow合约已经提交源码到[arbiscan](https://arbiscan.io/address/0xaff56ee7bdd72354e5b664ab5921f564df66bf5c#contracts)。  
history和card还有部分开发工作，暂未开源，预计在三个月内实现全部功能后会全部提交源代码。  
相关网站刚启动，还需要一段时间才能提供完整功能（合并图片和新闻文字呈现）。  

# 关于秘书
秘书的工作非常重要，尤其在dev放弃owner权之后，秘书是全社区唯一中心化实体，担负着维护全社区价值的重任。秘书除了负责及时提交新闻候选列表之外，还有一些重要的要求：
* 秘书不能遗漏真正的大新闻，需要秘书每天关注世界各处的事件。秘书需要为每天提交一次且至少三条新闻，不得有遗漏的日期。超过30天后，不能再提交该日期的新闻，一旦发生这将是秘书的重大失职。
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
* 需要认真投票。  
投票命中最终结果的第一名，获得的分红权重为 票权/3 * 8。  
投票命中最终结果的第二名，获得的分红权重为 票权/3 * 4。  
投票命中最终结果的第三名，获得的分红权重为 票权/3 * 2。  
其他情况获得的分红权重为 票权/3 * 1。  
这一设计是希望投票人认真的选择。
* 投票之后，可以用history settleVote结算分红，可以一次性把多次结算的分红通过history claim一次性提取到自己帐号(提取的奖励包括投票奖励以及NFT持有人奖励)。

# 关于拍卖
* 在秘书通过mintAndAuc创建并发起拍卖之后,NFT在合约内启动了拍卖，拍卖前的NFT所有者是Card合约。  
* 拍卖方式采用荷兰拍，先到先得。价格从初始拍卖价格每隔1小时降价1%，最低降价到0.01ETH。初始拍卖价格是最近十次成交价的EMA平均价的1.5倍，且不低于0.02ETH(曾经是0.04ETH，已修改)。  
* 可以在history dump命令中看到初始拍卖价AucInitPrice，最新拍卖价CurPrice，最终成交价FinalPrice。  
* 通过 history bid命令即可以最新拍卖价竞拍。  
* 拍卖成功之后，通过card claim命令把NFT从Card合约取回自己帐户。

# 关于奖励
现在有三部分奖励。
* 运维奖励。开发团队和秘书分别获得拍卖金额5%的奖励，合计10%。 秘书或者开发团队可以通过history settleOps结算这部分奖励。
* 投票奖励。凡是参与了投票的帐户，根据投票命中率和票权计算其分红权重。投票人根据投票权重分享拍卖金额的80%。投票人可以通过history settleVote结算这部分奖励，可以在history dump最后打印中看到这部分奖励。
* NFT持有人的奖励。早期的NFT持有人，可以平分后来的NFT拍卖金额的10%（包括自己这个NFT的奖励）。NFT持有人可以通过命令如"history settleCard -l 2023-06-23 -r 2023-06-24"结算这部分奖励，可以通过比如"hisotry dumpReward -l 2023-06-22 -r 2023-06-23" 命令查看是否已经结算。

奖励汇总后通过history claim提取。奖励的结算和提取记录通过history dumpReward命令查看。  

需要强调的是，大部分分红都给到投票人是为了保证系统的长期运行。NFT是全社区的长期利益同路人，但是更广大的投票人是NFT持有人参与社区的原动力，所以必须向投票人倾斜。  

按照目前的设想，在代码全部稳定后，开发团队会放弃owner权限，其5%的奖励作为社区的公积金留存备用。

# NFT的定制和呈现
* 受开发资源限制，全部NFT都用的同一个默认主图显示。预期在chainhis.net建设之后，可以叠加NFT的新闻/节日/投票信息/自定的附图/指定的祝福辞，其呈现将会出现差异。
* 秘书可以用"card setBaseImage"命令定制某个日期的NFT的背景图，NFT持有人可以向秘书申请定制其背景图。但是定制的主图必须和NFT基础图的风格一致，需要由秘书确保这一点。
* NFT持有人可以通过card customize定制nft的祝福辞和附图，以及选择主图呈现方式（最后这点需等待chainhis.net网站建设）。

# 后续工作
1. 实现投票选举秘书功能。
2. chainhis.net网站建设，支持网页+metamask的主流操作方式。
3. chainhis.net支持融合图片，把主图/日期/新闻/节日/附图/祝福辞混合在一起呈现，为不同的NFT显示不同的主图，且自动推送到IPFS。nft持有人可以向秘书申请更新主图为推送到IPFS的图片。
4. escrow和history中的存量资金，可以考虑挂结defi，在安全的前提下实现增值。

# GoChainHis

## 下载和安装
1. 获取源码: git clone https://github.com/gpmn/GoChainHis
2. cd GoChainHis
3. go install .

执行go install .之后，会生成GoChainHis可执行文件，后续步骤的"go run ."可以改为GoChainHis。

## 编辑自己的conf.json
需要先参考conf.json.demo,编辑自己的conf.json文件，只需要修改其中的MyAddr和MyPrivKey。注意::需要确保MyAddr地址**大小写正确**。  
自己的conf.json要么放在默认的$Home/.GoChainHis/目录下,要么通过--conf参数指定其全路径。  

## 命令列表
### escrow deposit
- go run . escrow deposit --amount 0.1
向抵押合约抵押指定数量的ETH。

### escrow withdraw
- go run . escrow withdraw --amount 0.1
从抵押合约提取指定数量的ETH。

### escrow dump
- go run . escrow dump
查看本帐号的抵押信息。只有抵押后才能获得投票的票权。  
票权前30天每天递增1/30，30天后加满。

### history dump
- go run . history dump -d 2023-06-22
查看06-22的新闻记录。
- go run . history dump -d 2023-06-22 -n 2 
查看06-22和06-23连续两天的新闻记录。

### history submit [secretarty only]
- go run . history submit -d 2023-06-23 -s subs/sub.0623.json
秘书提交某日新闻供社区投票，提交后24小时截止投票。
  
注意:: 
1). -d参数和sub文件中的日期必须一致
2). 建议使用<CN></CN> 和 <EN></EN> 标签实现国际化.
3). sub.0623.json u必须符合如下格式:  
```
{
    "Date": "2023-06-23 UTC",
    "BigNews": [
        "<CN><港媒:盒马鲜生最快11月IPO</CN><EN>Hong Kong media: Hema Xiansheng will IPO in November ASAP</EN>",
        "<CN>离岸人民币兑美元跌破7.22日内跌超250点</CN>",
    ]
}
```

### history resolve [secretarty only]
- go run . history resolve -d 2023-06-22
在投票时间结束（超过VoteEndTm）后，秘书结算该日期的新闻和投票信息。

### history mintAndAuc [secretarty only]
- go run . history mintAndAuc -d 2023-06-22
- go run . history mintAndAuc -d 2023-06-22 -e 端午节
- go run . history mintAndAuc -d 2023-10-01 -e 国庆节,中秋节

只有秘书可以执行 mintAndAuc, 用以生成NFT并发起拍卖，需要注意: **不要遗漏-e参数**以指定节日。
  
### history vote
* go run . history vote -d 2023-06-23 -p 0,2,3

为2023-06-23的新闻投票。  

Note:: 
1. 投票-p参数必须给三个不同的编号，编号针对的是history dump命令中Stories列表的编号 
2. VoteEndTm截止之后不可投票
3. 投票命中群选结果第一名权重*8; 第二名权重*4; 第三名权重*2; 未命中前三权重*1
4. 可以用history dump查看待投票日期信息。

### history dumpReward
* go run . history dumpReward  
查看汇总的已结算和已提取的分红，此命令查询不到尚未结算的分红。在history dump尾部，如果MyVoteStatus不为0就有资格领取分红，是否领取分红/是否可以领取分红可以在MyVoteReward字段查询到。
* go run . history dumpReward -l 2023-06-22 -r 2023-06-26
查看06-22的持卡人从06-26卡拍卖获得的分红

### history settleVote
* go run . history settleVote -d 2023-06-23  
投票人结算投票分红，奖励结算后可以在history dumpReward中查询到，也可以在history dump尾部查询到。

### history settleCard
* go run . history settleCard -l 2023-06-22 -r 2023-06-26 
持卡人（持有06-22的卡）去结算06-26的卡拍卖分红，结算后可以在go run . history dumpReward -l 2023-06-22 -r 2023-06-26 查看。

### history settleOps
* go run . history settleOps -d 2023-06-23  
运维团队结算分红，只有开发团队和秘书可以执行。

### history claimReward
* go run . history claimReward -d 2023-06-23    
提取分红到自己钱包,参见 history dumpReward

### card claim [bid winner only]
* go run . card claim  -d 2023-06-23  
竞拍成功后(hitory dump的Winner字段)， 买家通过card claim命令，将NFT提取回自己钱包。

### card customize [NFT owner only]
* go run . card customize -d 2023-06-22 -r 0 -g "庆祝链史项目启动！希望大家理性参与！希望链史项目能长期运作！希望链史项目能够为人类真实记录下每一天！"  -i "https://ipfs.io/ipfs/Qmd5yq6Qtkr6xshHzWXXRSPaboq1LounKxogVDMPMCKRAY"
定制NFT主图呈现方式，自定的祝福辞和附图。







