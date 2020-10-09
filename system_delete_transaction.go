package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type SystemDeleteTransaction struct {
	Transaction
	pb *proto.SystemDeleteTransactionBody
}

func NewSystemDeleteTransaction() *SystemDeleteTransaction {
	pb := &proto.SystemDeleteTransactionBody{}

	transaction := SystemDeleteTransaction{
		pb:          pb,
		Transaction: newTransaction(),
	}

	return &transaction
}

func (transaction *SystemDeleteTransaction) SetExpirationTime(expiration time.Time) *SystemDeleteTransaction {
	transaction.pb.ExpirationTime = &proto.TimestampSeconds{
		Seconds: expiration.Unix(),
	}
	return transaction
}

func (transaction *SystemDeleteTransaction) GetExpirationTime() int64 {
	return transaction.pb.GetExpirationTime().Seconds
}

func (transaction *SystemDeleteTransaction) SetContractID(contractID ContractID) *SystemDeleteTransaction {
	transaction.pb.Id = &proto.SystemDeleteTransactionBody_ContractID{ContractID: contractID.toProto()}
	return transaction
}

func (transaction *SystemDeleteTransaction) GetContract() ContractID {
	return contractIDFromProto(transaction.pb.GetContractID())
}

func (transaction *SystemDeleteTransaction) SetFileID(fileID FileID) *SystemDeleteTransaction {
	transaction.pb.Id = &proto.SystemDeleteTransactionBody_FileID{FileID: fileID.toProto()}
	return transaction
}

func (transaction *SystemDeleteTransaction) GetFileID() FileID {
	return fileIDFromProto(transaction.pb.GetFileID())
}

//
// The following methods must be copy-pasted/overriden at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func systemDeleteTransaction_getMethod(channel *channel) method {
	return method{
		transaction: channel.getCrypto().CreateAccount,
	}
}

func (transaction *SystemDeleteTransaction) IsFrozen() bool {
	return transaction.isFrozen()
}

// Sign uses the provided privateKey to sign the transaction.
func (transaction *SystemDeleteTransaction) Sign(
	privateKey PrivateKey,
) *SystemDeleteTransaction {
	return transaction.SignWith(privateKey.PublicKey(), privateKey.Sign)
}

func (transaction *SystemDeleteTransaction) SignWithOperator(
	client *Client,
) (*SystemDeleteTransaction, error) {
	// If the transaction is not signed by the operator, we need
	// to sign the transaction with the operator

	if client.operator == nil {
		return nil, errClientOperatorSigning
	}

	if !transaction.IsFrozen() {
		transaction.FreezeWith(client)
	}

	return transaction.SignWith(client.operator.publicKey, client.operator.signer), nil
}

// SignWith executes the TransactionSigner and adds the resulting signature data to the Transaction's signature map
// with the publicKey as the map key.
func (transaction *SystemDeleteTransaction) SignWith(
	publicKey PublicKey,
	signer TransactionSigner,
) *SystemDeleteTransaction {
	if !transaction.IsFrozen() {
		transaction.Freeze()
	}

	if transaction.keyAlreadySigned(publicKey) {
		return transaction
	}

	for index := 0; index < len(transaction.transactions); index++ {
		signature := signer(transaction.transactions[index].GetBodyBytes())

		transaction.signatures[index].SigPair = append(
			transaction.signatures[index].SigPair,
			publicKey.toSignaturePairProtobuf(signature),
		)
	}

	return transaction
}

// Execute executes the Transaction with the provided client
func (transaction *SystemDeleteTransaction) Execute(
	client *Client,
) (TransactionResponse, error) {
	if !transaction.IsFrozen() {
		transaction.FreezeWith(client)
	}

	transactionID := transaction.id

	if !client.GetOperatorID().isZero() && client.GetOperatorID().equals(transactionID.AccountID) {
		transaction.SignWith(
			client.GetOperatorKey(),
			client.operator.signer,
		)
	}

	_, err := execute(
		client,
		request{
			transaction: &transaction.Transaction,
		},
		transaction_shouldRetry,
		transaction_makeRequest,
		transaction_advanceRequest,
		transaction_getNodeId,
		systemDeleteTransaction_getMethod,
		transaction_mapResponseStatus,
		transaction_mapResponse,
	)

	if err != nil {
		return TransactionResponse{}, err
	}

	return TransactionResponse{TransactionID: transaction.id}, nil
}

func (transaction *SystemDeleteTransaction) onFreeze(
	pbBody *proto.TransactionBody,
) bool {
	pbBody.Data = &proto.TransactionBody_SystemDelete{
		SystemDelete: transaction.pb,
	}

	return true
}

func (transaction *SystemDeleteTransaction) Freeze() (*SystemDeleteTransaction, error) {
	return transaction.FreezeWith(nil)
}

func (transaction *SystemDeleteTransaction) FreezeWith(client *Client) (*SystemDeleteTransaction, error) {
	transaction.initFee(client)
	if err := transaction.initTransactionID(client); err != nil {
		return transaction, err
	}

	if !transaction.onFreeze(transaction.pbBody) {
		return transaction, nil
	}

	return transaction, transaction_freezeWith(&transaction.Transaction, client)
}

func (transaction *SystemDeleteTransaction) GetMaxTransactionFee() Hbar {
	return transaction.Transaction.GetMaxTransactionFee()
}

// SetMaxTransactionFee sets the max transaction fee for this SystemDeleteTransaction.
func (transaction *SystemDeleteTransaction) SetMaxTransactionFee(fee Hbar) *SystemDeleteTransaction {
	transaction.Transaction.SetMaxTransactionFee(fee)
	return transaction
}

func (transaction *SystemDeleteTransaction) GetTransactionMemo() string {
	return transaction.Transaction.GetTransactionMemo()
}

// SetTransactionMemo sets the memo for this SystemDeleteTransaction.
func (transaction *SystemDeleteTransaction) SetTransactionMemo(memo string) *SystemDeleteTransaction {
	transaction.Transaction.SetTransactionMemo(memo)
	return transaction
}

func (transaction *SystemDeleteTransaction) GetTransactionValidDuration() time.Duration {
	return transaction.Transaction.GetTransactionValidDuration()
}

// SetTransactionValidDuration sets the valid duration for this SystemDeleteTransaction.
func (transaction *SystemDeleteTransaction) SetTransactionValidDuration(duration time.Duration) *SystemDeleteTransaction {
	transaction.Transaction.SetTransactionValidDuration(duration)
	return transaction
}

func (transaction *SystemDeleteTransaction) GetTransactionID() TransactionID {
	return transaction.Transaction.GetTransactionID()
}

// SetTransactionID sets the TransactionID for this SystemDeleteTransaction.
func (transaction *SystemDeleteTransaction) SetTransactionID(transactionID TransactionID) *SystemDeleteTransaction {
	transaction.Transaction.SetTransactionID(transactionID)
	return transaction
}

func (transaction *SystemDeleteTransaction) GetNodeID() AccountID {
	return transaction.Transaction.GetNodeID()
}

// SetNodeID sets the node AccountID for this SystemDeleteTransaction.
func (transaction *SystemDeleteTransaction) SetNodeID(nodeID AccountID) *SystemDeleteTransaction {
	transaction.Transaction.SetNodeID(nodeID)
	return transaction
}