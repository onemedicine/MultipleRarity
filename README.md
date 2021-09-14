# MultipleRarity


#### 发呆时意识到一个严重问题，rarity是单纯基于角色构建的游戏模式，授权也意味着针对角色，那么其他任意账户也可以操作别人对合约授权过的角色，好处是：多账户的话方便太多了；坏处是：我曾加了个multiple_spend_xp，但spend_xp()玩法还不明，万一别人要是恶意消耗我的角色经验就完蛋了呀(::>_<::)。我已经改了，并且增加多角色claim gold和下地牢打材料，但是希望大家**务必立刻马上取消对三个旧合约的授权**！！！

### 更新说明

合约地址：[0x4a33076916289486cdc975ee28ef4041a5e5702a](https://ftmscan.com/address/0x4a33076916289486cdc975ee28ef4041a5e5702a#writeContract)

> * 所有接口增加授权判断
> 
> * 仅限角色的owner调用multiple_spend_xp()
> 
> * 增加批量claim gold 接口multiple_claim_gold()
> 
> * 增加批量打第一个地牢副本接口multiple_adventure_crafting_materials()
> 
> * 增加自毁接口destroy()，若发现批量接口被恶意利用，我将销毁合约，省去大家撤销授权的麻烦

------


### 使用说明

> * 1.按需调用对应的授权接口
> 
> * 2.批量冒险multiple_adventure()，批量升级multiple_level_up()，也可用自己的一个热账户去为其他冷钱包账户的角色冒险升级（急缺好心人啊）
> 
> * 3.批量领金币multiple_claim_gold()，rarity_gold合约原因，仅支持调用approve为角色授权
> 
> * 4.批量打地牢multiple_adventure_crafting_materials()，rarity_crafting_materials合约原因，仅支持调用approve为角色授权
> 
> * 5.批量消耗经验multiple_spend_xp(), 猜想将来AC肯定有其他合约调用该接口实现某种玩法，应该永远不会单独排上用场了，留着仅为把自己挖的坑给埋上。
> 
> * 6.单次传入太多角色id会超出block gaslimit，数量控制在130个左右较好。


**传参示例：**

![image](https://user-images.githubusercontent.com/20993492/133186363-239a8be6-e77b-46b5-ad68-616da0771e0e.png)

------


**rarity web**:https://rarity.game/

**rarity github**:https://github.com/andrecronje/rarity

[rarity: 0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#code)

------

## *重中之重，如何approve和撤销approve！！！*

***setApprovalForAll 和 approve并不相通，意味着使用approve的单个角色并不能通过setApprovalForAll来撤销，换句话说调哪个接口授权的就还调那个接口撤销***

### 如何撤销approve：
[在区块浏览器打开rarity](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#writeContract)

1.调用setApprovalForAll(撤销曾一键所有授权)
> * operator: 填写使用该接口授权过的合约地址
> * approved: 填写false

> 旧合约地址：
> 
> 0x8ACcaa4b940eaFC41b33159027cDBDb4A567d442
>
>0x8788f32939ff2a8eb014877fc734ff77aa8aa148
>
>0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3






2.调用approve(单个角色授权)
> * to: 填写零地址或其他自己的地址: 0x0000000000000000000000000000000000000000
> * tokenId: 填写你拥有的角色id

撤销过后可以查询确认一下，如下图分别为两种方式的成功撤销：

![image](https://user-images.githubusercontent.com/20993492/132992089-52225b4f-8109-4f4c-ab90-5bf10e78e9a7.png)



### 如何approve：
[在区块浏览器打开rarity](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#writeContract)

1.调用setApprovalForAll(一键授权所有)
> * operator: 填写MultipleRarity合约地址: 0x4a33076916289486cdc975ee28ef4041a5e5702a
> * approved: 填写true


2.调用approve(单个角色授权)
> * to: 填写MultipleRarity合约地址: 0x4a33076916289486cdc975ee28ef4041a5e5702a
> * tokenId: 填写你拥有的角色id


------

## 其他

### 如何快捷获取自己拥有的角色id

用到了下面两个网址：

> * 1.查看角色列表（当用脚本批量点属性时，可按职业筛选）：https://rarity-game.netlify.app/
> 
> 方法是：在该页面Ctrl+A，Ctrl+C 复制拿到所有文字信息，使用正则筛选出id
> 
![image](https://user-images.githubusercontent.com/20993492/132995223-5450f389-c563-4e27-9bac-740fb64aa6d5.png)


> * 2.在线正则：https://tool.lu/regex/
> 
> 用到的正则表达式：#[1-9]\d*

![image](https://user-images.githubusercontent.com/20993492/132995342-7be37782-d43b-4c23-aed0-96753d7f38d7.png)



### 角色属性加点如何分配

[medium原文](https://andrecronje.medium.com/rarity-attributes-19ff3cd457c8)

[rarity_attributes合约](https://ftmscan.com/address/0xb5f5af1087a8da62a23b08c00c6ec9af21f397a1)

所有职业加点推荐：(doc.google)[https://docs.google.com/spreadsheets/d/19GzfnCt9rofQPmA9GMUjvF3z5ObYfqCP3pCNX-7XePQ/edit#gid=619314779]

> * 力量(Strength)
> 
> 对力量值需求重要的职业有五个，分别是战士、野蛮人、圣骑士、游侠、僧人
> 
> * 敏捷(Dexterity)
> 
> 盗贼对敏捷的需求是重中之重，游戏中应该会有负重类的限制，类似重甲、轻甲的护甲装备，官方推荐游侠、野蛮人、巫师、术士、僧人都可以配置一些敏捷点数。
> 
> * 体质(Constitution)
> 
> 因为体质关系直接影响到人物的生命值和健康状态，所以对所有职业阵营都很重要，**都需要将其排在副属性里必点**。
> 
> * 智力(Intelligence)
> 
> 这个属性特别影响巫师，巫师类似游戏中的魔法师对施法的需求非常的迫切，如果你想你的角色有很多的法术技能搭配也可以分配一些智力给他。
> 
> * 感知(Wisdom)
> 
> 如果你想你的角色对冒险中发生的事情非常敏感的话，可以为他添加些许智慧。牧师、德鲁伊这两个职业对该属性要求较高，圣骑士和游侠也可以将该属性作为副属性分配。
> 
> * 魅力(Charisma)
> 
> 圣骑士、术士、诗人对魅力要求很高，说实话这个设定有些看不懂，牧师也可以添加一些，据说可以在冒险时驱散不死族的怪物。

属性默认每个8级，角色初始32点可在此基础上分配，角色每升4级获得1点可分配

9-14，每升1级消耗1点，该阶段完成消耗6点

15-16，每升1级消耗2点，该阶段完成消耗4

17-18，每升1级消耗3点，该阶段完成消耗6

19-20，每升1级消耗4点，该阶段完成消耗8
....

![image](https://user-images.githubusercontent.com/20993492/132832050-4f893437-89e8-47f7-8977-2cb44b3b2ab2.png)




### 加点模拟器
DFarm推出的模拟器（模拟加点及模拟探索）:https://dfarm.club/rarity.html

DFarm yyds!

![c1cb5c91cdf5d59843eba041ff1ee09](https://user-images.githubusercontent.com/20993492/132671153-4983bb11-aa17-4436-8bad-e549c133d5e4.png)

### 角色交易市场
**raritysea**:https://www.raritysea.io/

**summoner-market**:https://summoner-market.alphafinance.io/ 

**MarketplaceSalesStats**:https://cryptomarkettoday.com/rarity/

*靓号值钱，但高等级的角色后期一定是影响价格的主要因素，坚持冒险吧！*

![image](https://user-images.githubusercontent.com/20993492/132943146-62ddb89f-aeaf-451d-95b2-9e8e24942a27.png)


### 前端（按实用性排序，前端交互需注意账号风险）

  AC推特推荐：https://rarity.game/

  非常简洁（支持加点，探宝）：http://rarity-adventures.surge.sh/

  批量展示，升级方便：https://rarity-game.netlify.app/

  复古风格：https://adventure.major.tax/

  有更完整的角色UI：https://www.raritymanifested.com/

  红黑复古风：https://rarity-visualizer-ui.vercel.app/

  与第一个类似：https://rarity-xyz.surge.sh/


### 教程

  基础规则：https://weibo.com/ttarticle/p/show?id=2309404679169673134323

  新手入门：https://weibo.com/ttarticle/p/show?id=2309404678445002260494


### 获取FTM

#### 跨链（建议从BSC，Ploygon，xdai跨过来，整体手续费不高）

  https://cbridge.celer.network/#/
  
  https://stable.anyswap.exchange/#/swap
  
  https://spookyswap.finance/swap

#### 交易所

  币安不定时开放提币，还是建议不要去太小众的交易所

### claim token
当角色每升一级可以领金币

在区块浏览器打开[rarity_gold](https://ftmscan.com/address/0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2#writeContract)
![image](https://user-images.githubusercontent.com/20993492/132827633-68807f22-cbdc-473b-a0bc-7be80d76c61e.png)



### 感谢[DFarm](https://weibo.com/u/6112840709)曾指出合约批量升级错误, 非常抱歉给大家带来使用不便。
