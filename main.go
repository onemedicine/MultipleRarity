package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_common "go-rarity/common"
	"go-rarity/config"
	"go-rarity/multiplerarity"
	"go-rarity/rarity"
	"log"
	"math/big"
)

var (
	Client           *ethclient.Client
	Rarity           *rarity.Rarity
	MultipleRarity   *multiplerarity.Multiplerarity
	RarityAttributes *rarity.RarityAttributes
	RarityGold       *rarity.RarityGold
	ChainID *big.Int

	FromAddress common.Address
	PrivateKey *ecdsa.PrivateKey
)

var CustomAttributes = map[int64]_common.Attributes{
	// Barbarian 野蛮人
	1:  {20, 14, 10, 8, 8, 8},
	// Bard 吟游诗人
	2:  {12, 10, 10, 8, 8, 20},
	// Cleric 牧师
	3:  {8, 10, 10, 10, 19, 14},
	// Druid 德鲁伊
	4:  {10, 10, 10, 10, 20, 8},
 	// Fighter 战士
	5:  {20, 10, 14, 8, 8, 8},
	// Monk 武僧
	6:  {20, 14, 10, 8, 8, 8},
	// Paladin 圣骑士
	7:  {15, 8, 10, 8, 14, 18},
	// Ranger 游侠
	8:  {18, 14, 10, 8, 15, 8},
	// Rogue 盗贼
	9:  {10, 20, 10, 12, 8, 8},
	// Sorcerer 巫师
	10: {8, 10, 10, 20, 12, 8},
	// Wizard 男巫师
	11: {2, 12, 10, 20, 10, 8},
}

func init() {
	var err error
	Client, err = ethclient.Dial("https://rpc.ftm.tools")
	if err != nil {
		log.Fatalf("Failed to connect to the RPC client: %v", err)
	}
	ChainID,err = Client.ChainID(nil)
	if err != nil {
		log.Fatalf("Failed to connect to the RPC client: %v", err)
	}

	Rarity, err = rarity.NewRarity(common.HexToAddress("0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb"), Client)
	if err != nil {
		log.Fatalf("Failed to instantiate Rarity contract: %v", err)
	}

	MultipleRarity, err = multiplerarity.NewMultiplerarity(common.HexToAddress("0xB3e2dEa302f43Df164758f1A8Ded7Ac6C87741b3"), Client)
	if err != nil {
		log.Fatalf("Failed to instantiate MultipleRarity contract: %v", err)
	}

	RarityAttributes, err = rarity.NewRarityAttributes(common.HexToAddress("0xB5F5AF1087A8DA62A23b08C00C6ec9af21F397a1"), Client)
	if err != nil {
		log.Fatalf("Failed to instantiate RarityAttributes contract: %v", err)
	}

	RarityGold, err = rarity.NewRarityGold(common.HexToAddress("0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2"), Client)
	if err != nil {
		log.Fatalf("Failed to instantiate RarityGold contract: %v", err)
	}

	PrivateKey, err = crypto.HexToECDSA(config.Conf.Rarity.Private)
	if err != nil {
		log.Fatal(err)
	}


	pk, ok := PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	FromAddress = crypto.PubkeyToAddress(*pk)

}

func main() {



}


func summon(class *big.Int) {
	nonce, err := Client.PendingNonceAt(context.Background(), FromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, ChainID)
	if err != nil {
		log.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(800000)
	auth.GasPrice = gasPrice

	tx, err := Rarity.Summon(auth, class)
	if err != nil {
		log.Println(err)
	}

	log.Println("tx sent:", tx.Hash().Hex())
}

func summoner(summoner *big.Int) (*_common.Summoner,error) {
	info, err := Rarity.Summoner(nil,summoner)
	if err != nil {
		log.Fatalf("Failed to retrieve summoner: %v", err)
		return nil,err
	}
	fmt.Printf("%v\n",info)
	return &_common.Summoner{
		Xp: info.Xp,
		Log: info.Log,
		Class: info.Class,
		Level: info.Level,
	},nil
}


func pointBuy(summoner *big.Int, class int64) error {
	nonce, err := Client.PendingNonceAt(context.Background(), FromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, ChainID)
	if err != nil {
		log.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(800000)
	auth.GasPrice = gasPrice



	tx, err := RarityAttributes.PointBuy(auth,
		summoner,
		CustomAttributes[class].Strength,
		CustomAttributes[class].Dexterity,
		CustomAttributes[class].Constitution,
		CustomAttributes[class].Intelligence,
		CustomAttributes[class].Wisdom,
		CustomAttributes[class].Charisma)
	if err != nil {
		log.Println(err)
	}

	log.Println("tx sent:", tx.Hash().Hex())

	return nil

}
