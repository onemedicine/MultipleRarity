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
    
    function ownerOf(uint) external view returns (address);
    
    function getApproved(uint) external view returns (address);
    
    function isApprovedForAll(address owner, address operator) external pure returns (bool) ;

}

interface rarity_gold {
    function claim(uint summoner) external;
    
    function claimed(uint summoner) external pure returns (uint);
}

interface rarity_crafting_materials {
    function adventure(uint _summoner) external returns (uint);
    
    function adventurers_log (uint _summoner) external pure returns (uint);
}

contract MultipleRarity {
    rarity constant _rarity = rarity(0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb);
    
    rarity_gold constant _rarity_gold = rarity_gold(0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2);
    
    rarity_crafting_materials constant _rarity_crafting_materials = rarity_crafting_materials(0x2A0F1cB17680161cF255348dDFDeE94ea8Ca196A);
    
    address payable  owner;
 
    constructor() { owner = payable(msg.sender); }
    
    function multiple_adventure(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_isApprovedForAll(_summoners[i]) &&  block.timestamp > _rarity.adventurers_log(_summoners[i])) {
                _rarity.adventure(_summoners[i]);
            }
        }
    }

    function multiple_spend_xp(uint[] calldata _summoners, uint[] calldata _xps) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_isOwner(_summoners[i]) && _rarity.xp(_summoners[i]) >= _xps[i]) {
                _rarity.spend_xp(_summoners[i], _xps[i]);
            }
        }
    }

    function multiple_level_up(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_isApprovedForAll(_summoners[i]) && _rarity.xp(_summoners[i]) >= _rarity.xp_required(_rarity.level(_summoners[i]))) {
                _rarity.level_up(_summoners[i]);
            }
        }
    }
    
    function multiple_claim_gold(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_isApproved(_summoners[i]) && _rarity.level(_summoners[i]) >= 2 && _rarity.level(_summoners[i]) > _rarity_gold.claimed(_summoners[i])) {
                _rarity_gold.claim(_summoners[i]);
            }
        }
    }
    
    function multiple_adventure_crafting_materials(uint[] calldata _summoners) external {
        for (uint256 i = 0; i < _summoners.length; i++) {
            if (_isApproved(_summoners[i]) && block.timestamp > _rarity_crafting_materials.adventurers_log(_summoners[i])) {
                _rarity_crafting_materials.adventure(_summoners[i]);
            }
        }
    }
    
    function destroy() external {
        require(msg.sender == owner, "Only owner can call this function.");
        selfdestruct(owner);
    }
    
    function _isOwner(uint _summoner) internal view returns (bool) {
        return (_rarity.ownerOf(_summoner) == msg.sender);
    }
    
    function _isApproved(uint _summoner) internal view returns (bool) {
        return (_rarity.getApproved(_summoner) == address(this));
    }
    
    function _isApprovedForAll(uint _summoner) internal view virtual returns (bool) {
        return (_rarity.getApproved(_summoner) == address(this) || _rarity.isApprovedForAll(_rarity.ownerOf(_summoner), address(this)));
    }

}
