package models

import (
	"time"
)

// Response object for listing transactions
type ListTransactionsResponseInterface interface {
	GetData() Optional[[]GetTransactionResponseInterface]
	GetPaging() Optional[PagingResponse]
}

// Response object for getting a charge
type GetChargeResponseInterface interface {
	GetId() Optional[string]
	GetCode() Optional[string]
	GetGatewayId() Optional[string]
	GetAmount() Optional[int]
	GetStatus() Optional[string]
	GetCurrency() Optional[string]
	GetPaymentMethod() Optional[string]
	GetDueAt() Optional[time.Time]
	GetCreatedAt() Optional[time.Time]
	GetUpdatedAt() Optional[time.Time]
	GetLastTransaction() Optional[GetTransactionResponseInterface]
	GetInvoice() Optional[GetInvoiceResponse]
	GetOrder() Optional[GetOrderResponse]
	GetCustomer() Optional[GetCustomerResponse]
	GetMetadata() Optional[map[string]string]
	GetPaidAt() Optional[time.Time]
	GetCanceledAt() Optional[time.Time]
	GetCanceledAmount() Optional[int]
	GetPaidAmount() Optional[int]
	GetInterestAndFinePaid() Optional[int]
	GetRecurrencyCycle() Optional[string]
}

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() *string
	GetBankTid() *string
	GetBank() *string
	GetPaidAt() *time.Time
	GetPaidAmount() *int
}

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() Optional[string]
	GetBankTid() Optional[string]
	GetPaidAt() Optional[time.Time]
	GetPaidAmount() Optional[int]
}

// Response for voucher transactions
type GetVoucherTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() Optional[string]
	GetAcquirerName() Optional[string]
	GetAcquirerAffiliationCode() Optional[string]
	GetAcquirerTid() Optional[string]
	GetAcquirerNsu() Optional[string]
	GetAcquirerAuthCode() Optional[string]
	GetAcquirerMessage() Optional[string]
	GetAcquirerReturnCode() Optional[string]
	GetOperationType() Optional[string]
	GetCard() Optional[GetCardResponse]
}

// Response object for getting a boleto transaction
type GetBoletoTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() Optional[string]
	GetBarcode() Optional[string]
	GetNossoNumero() Optional[string]
	GetBank() Optional[string]
	GetDocumentNumber() Optional[string]
	GetInstructions() Optional[string]
	GetBillingAddress() Optional[GetBillingAddressResponse]
	GetDueAt() Optional[time.Time]
	GetQrCode() Optional[string]
	GetLine() Optional[string]
	GetPdfPassword() Optional[string]
	GetPdf() Optional[string]
	GetPaidAt() Optional[time.Time]
	GetPaidAmount() Optional[string]
	GetType() Optional[string]
	GetCreditAt() Optional[time.Time]
	GetStatementDescriptor() Optional[string]
}

// Response object for getting a debit card transaction
type GetDebitCardTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() Optional[string]
	GetAcquirerName() Optional[string]
	GetAcquirerAffiliationCode() Optional[string]
	GetAcquirerTid() Optional[string]
	GetAcquirerNsu() Optional[string]
	GetAcquirerAuthCode() Optional[string]
	GetOperationType() Optional[string]
	GetCard() Optional[GetCardResponse]
	GetAcquirerMessage() Optional[string]
	GetAcquirerReturnCode() Optional[string]
	GetMpi() Optional[string]
	GetEci() Optional[string]
	GetAuthenticationType() Optional[string]
	GetThreedAuthenticationUrl() Optional[string]
	GetFundingSource() Optional[string]
}

// Generic response object for getting a transaction.
type GetTransactionResponseInterface interface {
	GetGatewayId() Optional[string]
	GetAmount() Optional[int]
	GetStatus() Optional[string]
	GetSuccess() Optional[bool]
	GetCreatedAt() Optional[time.Time]
	GetUpdatedAt() Optional[time.Time]
	GetAttemptCount() Optional[int]
	GetMaxAttempts() Optional[int]
	GetSplits() Optional[[]GetSplitResponse]
	GetNextAttempt() Optional[time.Time]
	GetTransactionType() *string
	GetId() Optional[string]
	GetGatewayResponse() Optional[GetGatewayResponseResponse]
	GetAntifraudResponse() Optional[GetAntifraudResponse]
	GetMetadata() Optional[map[string]string]
	GetSplit() Optional[[]GetSplitResponse]
	GetInterest() Optional[GetInterestResponse]
	GetFine() Optional[GetFineResponse]
	GetMaxDaysToPayPastDue() Optional[int]
}

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() Optional[string]
	GetAcquirerName() Optional[string]
	GetAcquirerAffiliationCode() Optional[string]
	GetAcquirerTid() Optional[string]
	GetAcquirerNsu() Optional[string]
	GetAcquirerAuthCode() Optional[string]
	GetOperationType() Optional[string]
	GetCard() Optional[GetCardResponse]
	GetAcquirerMessage() Optional[string]
	GetAcquirerReturnCode() Optional[string]
	GetInstallments() Optional[int]
}

// Response object for getting a cash transaction
type GetCashTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetDescription() Optional[string]
}

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() Optional[string]
	GetAcquirerName() *string
	GetAcquirerAffiliationCode() Optional[string]
	GetAcquirerTid() *string
	GetAcquirerNsu() *string
	GetAcquirerAuthCode() Optional[string]
	GetOperationType() Optional[string]
	GetCard() Optional[GetCardResponse]
	GetAcquirerMessage() Optional[string]
	GetAcquirerReturnCode() Optional[string]
	GetInstallments() Optional[int]
	GetThreedAuthenticationUrl() Optional[string]
	GetFundingSource() Optional[string]
}

// Response object for listing charge transactions
type ListChargeTransactionsResponseInterface interface {
	GetData() Optional[[]GetTransactionResponseInterface]
	GetPaging() Optional[PagingResponse]
}

// Response object when getting a pix transaction
type GetPixTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetQrCode() Optional[string]
	GetQrCodeUrl() Optional[string]
	GetExpiresAt() Optional[time.Time]
	GetAdditionalInformation() Optional[[]PixAdditionalInformation]
	GetEndToEndId() Optional[string]
	GetPayer() Optional[GetPixPayerResponse]
	GetPixProviderTid() Optional[string]
}

// Generic response object for getting a BalanceOperation.
type GetBalanceOperationResponseInterface interface {
	GetId() Optional[string]
	GetStatus() Optional[string]
	GetBalanceAmount() Optional[string]
	GetBalanceOldAmount() Optional[string]
	GetType() Optional[string]
	GetAmount() Optional[string]
	GetFee() Optional[string]
	GetCreatedAt() Optional[string]
	GetMovementObject() GetMovementObjectBaseResponseInterface
}

// Generic response object for getting a MovementObjectBase.
type GetMovementObjectBaseResponseInterface interface {
	GetObject() *string
	GetId() Optional[string]
	GetStatus() Optional[string]
	GetAmount() Optional[string]
	GetCreatedAt() Optional[string]
	GetType() Optional[string]
	GetChargeId() Optional[string]
	GetGatewayId() Optional[string]
}

// Generic response object for getting a MovementObjectRefund.
type GetMovementObjectRefundResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetFraudCoverageFee() Optional[string]
	GetChargeFeeRecipientId() Optional[string]
	GetBankAccountId() Optional[string]
	GetLocalTransactionId() Optional[string]
	GetUpdatedAt() Optional[string]
}

// Generic response object for getting a MovementObjectFeeCollection.
type GetMovementObjectFeeCollectionResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetDescription() Optional[string]
	GetPaymentDate() Optional[string]
	GetRecipientId() Optional[string]
}

type GetMovementObjectPayableResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetFee() Optional[string]
	GetAnticipationFee() string
	GetFraudCoverageFee() string
	GetInstallment() string
	GetSplitId() string
	GetBulkAnticipationId() string
	GetAnticipationId() string
	GetRecipientId() string
	GetOriginatorModel() string
	GetOriginatorModelId() string
	GetPaymentDate() string
	GetOriginalPaymentDate() string
	GetPaymentMethod() string
	GetAccrualAt() string
	GetLiquidationArrangementId() string
}

type GetMovementObjectTransferResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetSourceType() Optional[string]
	GetSourceId() Optional[string]
	GetTargetType() Optional[string]
	GetTargetId() Optional[string]
	GetFee() Optional[string]
	GetFundingDate() Optional[string]
	GetFundingEstimatedDate() Optional[string]
	GetBankAccount() Optional[string]
}
