# MultipleRarity
After approve this contract, you can use the contract to adventure with multiple characters at the same time

**rarity web**:https://rarity.game/

**rarity github**:https://github.com/andrecronje/rarity

[rarity: 0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#code)



[MultipleRarity: 0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3](https://ftmscan.com/address/0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3#code)



*第一次使用需要先approve*
### 如何approve：
[在浏览器打开rarity](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#writeContract)

1.调用setApprovalForAll(一键授权所有)
> * operator: 填写MultipleRarity合约地址: 0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3
> * approved: 填写true

![image](https://user-images.githubusercontent.com/20993492/132617874-214bb754-c480-4779-9f38-27930f9fddc6.png)

2.调用approve(单个角色授权)
> * to: 填写MultipleRarity合约地址: 0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3
> * tokenId: 填写你拥有的角色id

![image](https://user-images.githubusercontent.com/20993492/132502933-8f3f048c-8500-4bea-96bc-3353f22ef8ad.png)



### 如何使用该合约：
[在浏览器打开MultipleRarity](https://ftmscan.com/address/0xb3e2dea302f43df164758f1a8ded7ac6c87741b3#writeContract)

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
### 角色加点如何分配
> * 力量
> 
> 对力量值需求重要的职业有五个，分别是战士、野蛮人、圣骑士、游侠、僧人
> 
> * 敏捷
> 
> 盗贼对敏捷的需求是重中之重，游戏中应该会有负重类的限制，类似重甲、轻甲的护甲装备，官方推荐游侠、野蛮人、巫师、术士、僧人都可以配置一些敏捷点数。
> 
> * 体质
> 
> 因为体质关系直接影响到人物的生命值和健康状态，所以对所有职业阵营都很重要。
> 
> * 魔力
> 
> 这个属性特别影响巫师，巫师类似游戏中的魔法师对施法的需求非常的迫切，如果你想你的角色有很多的法术技能搭配也可以分配一些魔力给他。
> 
> * 智慧
> 
> 和生活中的智慧是同样重要的，如果你想你的角色对冒险中发生的事情非常敏感的话，可以为他添加些许智慧。牧师、德鲁伊这两个职业对该属性要求较高，圣骑士和游侠也可以将该属性作为副属性分配。
> 
> * 魅力
> 
> 圣骑士、术士、诗人对魅力要求很高，说实话这个设定有些看不懂，牧师也可以添加一些，据说可以在冒险时驱散不死族的怪物。

![7987b24183449a48bbddafa41bb6431](https://user-images.githubusercontent.com/20993492/132791736-a1e52cd5-9996-4fc6-bc78-f3d0d75c6d15.png)



### 加点模拟器
DFarm推出的模拟器（模拟加点及模拟探索）:https://dfarm.club/rarity.html

yyds!
![c1cb5c91cdf5d59843eba041ff1ee09](https://user-images.githubusercontent.com/20993492/132671153-4983bb11-aa17-4436-8bad-e549c133d5e4.png)



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
当角色达到2级可以领金币
