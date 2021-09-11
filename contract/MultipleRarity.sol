// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


interface rarity {

    function adventure(uint _summoner) external;

    function spend_xp(uint _summoner, uint _xp) external;

    function level_up(uint _summoner) external;

    function xp(uint _summoner) external pure returns (uint);

    function level(uint _summoner) external pure returns (uint);

    function xp_required(uint curent_level) external pure returns (uint);

    function adventurers_log(uint _summoner) external pure returns (uint);

}


contract MultipleRarity {
    rarity constant _rarity = rarity(0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb);

    function multiple_adventure(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (block.timestamp > _rarity.adventurers_log(_summoners[i])) {
                _rarity.adventure(_summoners[i]);
            }
        }
    }

    function multiple_spend_xp(uint[] calldata _summoners, uint[] calldata _xps) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_rarity.xp(_summoners[i]) >= _xps[i]) {
                _rarity.spend_xp(_summoners[i], _xps[i]);
            }
        }
    }

    function multiple_level_up(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_rarity.xp(_summoners[i]) >= _rarity.xp_required(_rarity.level(_summoners[i]))) {
                _rarity.level_up(_summoners[i]);
            }
        }
    }
}
