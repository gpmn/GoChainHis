# GoChainHis
ChainHis CLI utils

cd this project folder, then type "go run ." to execute.

# escrow deposit
- go run . escrow deposit --amount 0.1

# escrow withdraw
- go run . escrow withdraw --amount 0.1

# escrow dump
- go run . escrow dump

# history dump
"-n" param will dump continous date's history infomation.

e.g. :: 
- go run . history dump -d 2023-06-22 -n 2 
- go run . history dump -d 2023-06-23

# history submit [secretarty only]
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

# history resolve [secretarty only]
resolve command will resolve a date's vote result once **VoteEndTm expired**.
- go run . history resolve -d 2023-06-22

# history mintAndAuc [secretarty only]
mintAndAuc command will mint NFT and issue its auction.
- go run . history mintAndAuc -d 2023-06-22
  
# history vote
go run . history vote -d 2023-06-23 -p 0,2,3

Note:: -p parameter must have 3 indexes which increased from 0 to indicate which news are the prefered big newes of the day, user can use **history dump** comamnd to see the details first.


