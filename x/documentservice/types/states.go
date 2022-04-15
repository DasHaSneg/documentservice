package types

import "strconv"

var NumberOfSigners = 2

var (
	PendingAnnexCreation     = "Pending annex creation"
	WaitingForAnnexSignature = getWaitingForAnnexSignature
	WaitingForSignature      = getWaitingForSignature
	Signed                   = "Signed"
	Completed                = "Completed"
	Cancelled                = "Cancelled"
)

func getWaitingForSignature(num int) string {
	return "Waiting for signature " + strconv.Itoa(num) + " of " + strconv.Itoa(NumberOfSigners)
}

func getWaitingForAnnexSignature(num int) string {
	return "Waiting for annex signature " + strconv.Itoa(num) + " of " + strconv.Itoa(NumberOfSigners)
}
