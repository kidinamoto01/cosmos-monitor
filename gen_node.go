package main

import (
	crypto "github.com/tendermint/go-crypto"
	//"github.com/tendermint/tendermint/p2p"
	"fmt"
	"github.com/tendermint/go-amino"

	//"github.com/tendermint/tendermint/p2p"
)
type NodeKey struct {
	PrivKey crypto.PrivKey `json:"priv_key"` // our priv key
}
var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}

func GenNodeKey() (*NodeKey, error) {
	//nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519()}
	privKey := crypto.GenPrivKeyEd25519()
	nodeKey := &NodeKey{
		PrivKey: privKey,
	}

	jsonBytes, err := cdc.MarshalJSON(nodeKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(privKey.PubKey().Address())
	fmt.Println(jsonBytes)
	//err = ioutil.WriteFile(filePath, jsonBytes, 0600)
	//if err != nil {
	//	return nil, err
	//}
	return nodeKey, nil
}

func main(){
	for i:=0 ;i<10;i++{
		GenNodeKey()
	}


}