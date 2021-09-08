# MultipleRarity
After approve this contract, you can use the contract to adventure with multiple characters at the same time

[rarity: 0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#code)

[MultipleRarity: 0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3](https://ftmscan.com/address/0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3#code)



第一次使用需要先approve
### 如何approve：
[在浏览器打开rarity](https://ftmscan.com/address/0xce761d788df608bd21bdd59d6f4b54b2e27f25bb#writeContract)
1.调用approve(授权)
> * to: 填写MultipleRarity合约地址: 0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3
> * tokenId: 填写你拥有的角色id

![image](https://user-images.githubusercontent.com/20993492/132502933-8f3f048c-8500-4bea-96bc-3353f22ef8ad.png)

### 如何使用该合约：
[在浏览器打开MultipleRarity](https://ftmscan.com/address/0xb3e2dea302f43df164758f1a8ded7ac6c87741b3#writeContract)
1.调用multiple_adventure(多角色同时冒险)
> * _summoners: 填多个角色id，例[21,22,23,24,25,26]
注: *须当角色冒险时间已到，并且approve过给该合约*


![image](https://user-images.githubusercontent.com/20993492/132503821-be600618-4e33-453b-84bd-c7750465a85e.png)

其他功能multiple_spend_xp，multiple_level_up同样支持了多角色，自行使用。

可配合网站：https://rarity-game.netlify.app/, 观察自己所有角色的ID和冒险冷却时间


点属性，探宝：https://rarity-adventures.surge.sh/

