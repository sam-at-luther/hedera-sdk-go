package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/v2/proto"
)

// ContractCallQuery calls a function of the given smart contract instance, giving it ContractFunctionParameters as its
// inputs. It will consume the entire given amount of gas.
//
// This is performed locally on the particular node that the client is communicating with. It cannot change the state of
// the contract instance (and so, cannot spend anything from the instance's Hedera account). It will not have a
// consensus timestamp. It cannot generate a record or a receipt. This is useful for calling getter functions, which
// purely read the state and don't change it. It is faster and cheaper than a ContractExecuteTransaction, because it is
// purely local to a single  node.
type ContractCallQuery struct {
	Query
	pb *proto.ContractCallLocalQuery
}

// NewContractCallQuery creates a ContractCallQuery query which can be used to construct and execute a
// Contract Call Local Query.
func NewContractCallQuery() *ContractCallQuery {
	header := proto.QueryHeader{}
	query := newQuery(true, &header)
	pb := proto.ContractCallLocalQuery{Header: &header}
	query.pb.Query = &proto.Query_ContractCallLocal{
		ContractCallLocal: &pb,
	}

	query.SetMaxQueryPayment(NewHbar(2))

	return &ContractCallQuery{
		Query: query,
		pb:    &pb,
	}
}

// SetContractID sets the contract instance to call
func (query *ContractCallQuery) SetContractID(id ContractID) *ContractCallQuery {
	query.pb.ContractID = id.toProtobuf()
	return query
}

func (query *ContractCallQuery) GetContractID() ContractID {
	return contractIDFromProtobuf(query.pb.ContractID)
}

// SetGas sets the amount of gas to use for the call. All of the gas offered will be charged for.
func (query *ContractCallQuery) SetGas(gas uint64) *ContractCallQuery {
	query.pb.Gas = int64(gas)
	return query
}

func (query *ContractCallQuery) GetGas() uint64 {
	return uint64(query.pb.Gas)
}

// SetMaxResultSize sets the max number of bytes that the result might include. The run will fail if it would have
// returned more than this number of bytes.
func (query *ContractCallQuery) SetMaxResultSize(size uint64) *ContractCallQuery {
	query.pb.MaxResultSize = int64(size)
	return query
}

// SetFunction sets which function to call, and the ContractFunctionParams to pass to the function
func (query *ContractCallQuery) SetFunction(name string, params *ContractFunctionParameters) *ContractCallQuery {
	if params == nil {
		params = NewContractFunctionParameters()
	}

	query.pb.FunctionParameters = params.build(&name)
	return query
}

func (query *ContractCallQuery) SetFunctionParameters(byteArray []byte) *ContractCallQuery {
	query.pb.FunctionParameters = byteArray
	return query
}

func (query *ContractCallQuery) GetFunctionParameters() []byte {
	return query.pb.FunctionParameters
}

func (query *ContractCallQuery) GetCost(client *Client) (Hbar, error) {
	if client == nil || client.operator == nil {
		return Hbar{}, errNoClientProvided
	}

	paymentTransaction, err := query_makePaymentTransaction(TransactionID{}, AccountID{}, client.operator, Hbar{})
	if err != nil {
		return Hbar{}, err
	}

	query.pbHeader.Payment = paymentTransaction
	query.pbHeader.ResponseType = proto.ResponseType_COST_ANSWER
	query.nodeIDs = client.network.getNodeAccountIDsForExecute()

	resp, err := execute(
		client,
		request{
			query: &query.Query,
		},
		query_shouldRetry,
		costQuery_makeRequest,
		costQuery_advanceRequest,
		costQuery_getNodeAccountID,
		contractCallQuery_getMethod,
		contractCallQuery_mapResponseStatus,
		query_mapResponse,
	)

	if err != nil {
		return Hbar{}, err
	}

	cost := int64(resp.query.GetContractCallLocal().Header.Cost)
	return HbarFromTinybar(cost), nil
}

func contractCallQuery_mapResponseStatus(_ request, response response) Status {
	return Status(response.query.GetContractCallLocal().Header.NodeTransactionPrecheckCode)
}

func contractCallQuery_getMethod(_ request, channel *channel) method {
	return method{
		query: channel.getContract().ContractCallLocalMethod,
	}
}

func (query *ContractCallQuery) Execute(client *Client) (ContractFunctionResult, error) {
	if client == nil || client.operator == nil {
		return ContractFunctionResult{}, errNoClientProvided
	}

	if len(query.Query.GetNodeAccountIDs()) == 0 {
		query.SetNodeAccountIDs(client.network.getNodeAccountIDsForExecute())
	}

	query.queryPayment = NewHbar(2)
	query.paymentTransactionID = TransactionIDGenerate(client.operator.accountID)

	var cost Hbar
	if query.queryPayment.tinybar != 0 {
		cost = query.queryPayment
	} else {
		cost = client.maxQueryPayment

		actualCost, err := query.GetCost(client)
		if err != nil {
			return ContractFunctionResult{}, err
		}

		if cost.tinybar > actualCost.tinybar {
			return ContractFunctionResult{}, ErrMaxQueryPaymentExceeded{}
		}

		cost = actualCost
	}

	err := query_generatePayments(&query.Query, client, cost)
	if err != nil {
		return ContractFunctionResult{}, err
	}

	resp, err := execute(
		client,
		request{
			query: &query.Query,
		},
		query_shouldRetry,
		query_makeRequest,
		query_advanceRequest,
		query_getNodeAccountID,
		contractCallQuery_getMethod,
		contractCallQuery_mapResponseStatus,
		query_mapResponse,
	)

	if err != nil {
		return ContractFunctionResult{}, err
	}

	return contractFunctionResultFromProtobuf(resp.query.GetContractCallLocal().FunctionResult), nil
}

// SetMaxQueryPayment sets the maximum payment allowed for this Query.
func (query *ContractCallQuery) SetMaxQueryPayment(maxPayment Hbar) *ContractCallQuery {
	query.Query.SetMaxQueryPayment(maxPayment)
	return query
}

// SetQueryPayment sets the payment amount for this Query.
func (query *ContractCallQuery) SetQueryPayment(paymentAmount Hbar) *ContractCallQuery {
	query.Query.SetQueryPayment(paymentAmount)
	return query
}

// SetNodeAccountIDs sets the node AccountID for this ContractCallQuery.
func (query *ContractCallQuery) SetNodeAccountIDs(accountID []AccountID) *ContractCallQuery {
	query.Query.SetNodeAccountIDs(accountID)
	return query
}

func (query *ContractCallQuery) SetMaxRetry(count int) *ContractCallQuery {
	query.Query.SetMaxRetry(count)
	return query
}
