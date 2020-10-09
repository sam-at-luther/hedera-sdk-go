package hedera

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestSerializeFreezeTransaction(t *testing.T) {
	startTime := time.Date(
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		12, 30, 0, time.Now().Nanosecond(), time.Now().Location(),
	)

	endTime := time.Date(
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		14, 30, 0, time.Now().Nanosecond(), time.Now().Location(),
	)

	mockClient, err := newMockClient()
	assert.NoError(t, err)

	privateKey, err := PrivateKeyFromString(mockPrivateKey)
	assert.NoError(t, err)

	tx, err := NewFreezeTransaction().
		SetTransactionID(testTransactionID).
		SetStartTime(startTime).
		SetEndTime(endTime).
		SetMaxTransactionFee(HbarFromTinybar(1e6)).
		SetTransactionID(testTransactionID).
		FreezeWith(mockClient)

	assert.NoError(t, err)

	tx.Sign(privateKey)

	assert.Equal(t, `bodyBytes:"\n\016\n\010\010\334\311\007\020\333\237\t\022\002\030\003\022\002\030\003\030\300\204=\"\002\010x\272\001\010\010\014\020\036\030\016\036"sigMap:<sigPair:<pubKeyPrefix:"\344\361\300\353L}\315\303\347\353\021p\263\010\212=\022\242\227\364\243\353\342\362\205\003\375g5F\355\216"ed25519:"\236\2040\201\314\033\272\346\250Zi$\240\016\322!\351\211\342\000[z\036\346\345\314,\220\314F0\377\356\370k*\020\270I\320\273<zp\260\241~\255\204L\274\272\2328\240\004\256\356\3603+\200\r\005">>transactionID:<transactionValidStart:<seconds:124124nanos:151515>accountID:<accountNum:3>>nodeAccountID:<accountNum:3>transactionFee:1000000transactionValidDuration:<seconds:120>freeze:<startHour:12startMin:30endHour:14endMin:30>`, strings.ReplaceAll(strings.ReplaceAll(tx.String(), " ", ""), "\n", ""))
}
