
# Create Bank Account Request

Request for creating a bank account

## Structure

`CreateBankAccountRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HolderName` | `string` | Required | Bank account holder name |
| `HolderType` | `string` | Required | Bank account holder type |
| `HolderDocument` | `string` | Required | Bank account holder document |
| `Bank` | `string` | Required | Bank |
| `BranchNumber` | `string` | Required | Branch number |
| `BranchCheckDigit` | `Optional[string]` | Optional | Branch check digit |
| `AccountNumber` | `string` | Required | Account number |
| `AccountCheckDigit` | `string` | Required | Account check digit |
| `Type` | `string` | Required | Bank account type |
| `Metadata` | `map[string]string` | Required | Metadata |
| `PixKey` | `Optional[string]` | Optional | Pix key |

## Example (as JSON)

```json
{
  "holder_name": "holder_name4",
  "holder_type": "holder_type0",
  "holder_document": "holder_document8",
  "bank": "bank6",
  "branch_number": "branch_number4",
  "branch_check_digit": "branch_check_digit4",
  "account_number": "account_number8",
  "account_check_digit": "account_check_digit4",
  "type": "type2",
  "metadata": {
    "key0": "metadata5",
    "key1": "metadata6",
    "key2": "metadata7"
  },
  "pix_key": "pix_key8"
}
```

