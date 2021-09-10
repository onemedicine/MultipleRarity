// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


interface rarity {
    
    function adventure(uint _summoner) external ;
    
    function spend_xp(uint _summoner, uint _xp) external ;
    
    function level_up(uint _summoner) external ;
    
}


contract MultipleRarity {
    address constant RARITYADDRESS = 0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb;
    rarity _rarity = rarity(RARITYADDRESS);

    function multiple_adventure( uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++)
            _rarity.adventure(_summoners[i]);
    }
    
    function multiple_spend_xp( uint[] calldata _summoners,uint[] calldata _xps) external {
        for (uint256 i = 0; i < _summoners.length; i++)
            _rarity.spend_xp(_summoners[i],_xps[i]);
    }
    
    function multiple_level_up( uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++)
            _rarity.level_up(_summoners[i]);
    }
     
}
