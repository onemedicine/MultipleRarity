# MultipleRarity

### 又又又更新了！
### MultipleRarity最新版:[0x8ACcaa4b940eaFC41b33159027cDBDb4A567d442](https://ftmscan.com/address/0x8accaa4b940eafc41b33159027cdbdb4a567d442#writeContract)
> * 注：角色冷却时间不统一时，可以不用管能不能冒险或升级，合约内部加了筛选，但消耗的gas增加了一点点，介意的可以使用常规修复版。

### MultipleRarity常规修复版:[0x8788f32939ff2a8eb014877fc734ff77aa8aa148](https://ftmscan.com/address/0x8788f32939ff2a8eb014877fc734ff77aa8aa148)

**[旧MultipleRarity合约](https://ftmscan.com/address/0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3)批量升级方法写错，请勿调用其批量升级！其他接口不受影响！**
### 感谢[DFarm](https://weibo.com/u/6112840709)及时指出合约错误, 非常抱歉给大家带来使用不便。


------

**rarity web**:https://rarity.game/

**rarity github**:https://github.com/andrecronje/rarity

[rarity: 0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#code)



[MultipleRarity: 0x8788f32939ff2a8eb014877fc734ff77aa8aa148](https://ftmscan.com/address/0x8788f32939ff2a8eb014877fc734ff77aa8aa148#code)


*第一次使用需要先approve*
### 如何approve：
[在区块浏览器打开rarity](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#writeContract)

1.调用setApprovalForAll(一键授权所有)
> * operator: 填写MultipleRarity合约地址: 0x8788f32939ff2a8eb014877fc734ff77aa8aa148
> * approved: 填写true

![image](https://user-images.githubusercontent.com/20993492/132890380-678de795-7be2-4299-a7e1-d26cd2870ce2.png)

2.调用approve(单个角色授权)
> * to: 填写MultipleRarity合约地址: 0x8788f32939ff2a8eb014877fc734ff77aa8aa148
> * tokenId: 填写你拥有的角色id

![image](https://user-images.githubusercontent.com/20993492/132890454-3faa4f68-b273-45f3-8882-babaaf1e3261.png)



### 如何使用该合约：
[在区块浏览器打开MultipleRarity](https://ftmscan.com/address/0x8788f32939ff2a8eb014877fc734ff77aa8aa148#writeContract)

1.调用multiple_adventure(多角色同时冒险)
> * _summoners: 填多个角色id，例[21,22,23,24,25,26]
> 
> 注: *须当角色冒险时间已到，并且approve过给该合约*
> 
> *FTM 矿工现在的block gaslimit太小了，太多角色容易超出block gaslimit，单次操作的角色数量控制在130以内最好; 同时矿工却还不断提高 minimum gas price，这就有点不厚道了*
  


![image](https://user-images.githubusercontent.com/20993492/132503821-be600618-4e33-453b-84bd-c7750465a85e.png)

其他功能multiple_spend_xp，multiple_level_up同样支持了多角色，自行使用。

可配合网站：https://rarity-game.netlify.app/, 观察自己所有角色的ID和冒险冷却时间



------

## 其他
### 角色属性加点如何分配

[medium原文](https://andrecronje.medium.com/rarity-attributes-19ff3cd457c8)

[rarity_attributes合约](https://ftmscan.com/address/0xb5f5af1087a8da62a23b08c00c6ec9af21f397a1)

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

