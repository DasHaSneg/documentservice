package types

import "strconv"

var NumberOfSigners = 2

var (
	PendingSupplementCreation     = "Pending application creation"
	WaitingForSupplementSignature = "Waiting for application signature"
	WaitingForSignature           = getWaitingForSignature
	Signed                        = "Signed"
	Completed                     = "Completed"
	Cancelled                     = "Cancelled"
)

func getWaitingForSignature(num int) string {
	return "Waiting for signature " + strconv.Itoa(num) + " of " + strconv.Itoa(NumberOfSigners)
}
