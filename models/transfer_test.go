package models

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/SmartMeshFoundation/SmartRaiden/utils"
	"github.com/stretchr/testify/assert"
)

func TestModelDB_NewReceivedTransfer(t *testing.T) {
	m := SetupTestDb(t)
	defer m.CloseDB()
	taddr := utils.NewRandomAddress()
	caddr := utils.NewRandomAddress()
	m.NewReceivedTransfer(2, caddr, caddr, taddr, 3, big.NewInt(10))
	key := fmt.Sprintf("%s-%d", caddr.String(), 3)
	r, err := m.GetReceivedTransfer(key)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, r.FromAddress, taddr)
	assert.Equal(t, r.ChannelAddress, caddr)
	assert.EqualValues(t, r.Nonce, 3)
	assert.EqualValues(t, r.Amount, big.NewInt(10))

	m.NewReceivedTransfer(3, caddr, caddr, taddr, 4, big.NewInt(10))
	m.NewReceivedTransfer(5, caddr, caddr, taddr, 6, big.NewInt(10))

	trs, err := m.GetReceivedTransferInBlockRange(0, 3)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, len(trs), 2)
	trs, err = m.GetReceivedTransferInBlockRange(0, 5)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, len(trs), 3)

	trs, err = m.GetReceivedTransferInBlockRange(0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, len(trs), 0)
}
