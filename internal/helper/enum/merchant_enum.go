package enum

type MerchantStatus uint16

const (
	New MerchantStatus = iota
	Open
	Requested
	Pending
	Approved
	Rejected
	Disabled
)

func (m MerchantStatus) String() string {
	switch m {
	case New:
		return "New"
	case Approved:
		return "Approved"
	case Open:
		return "Open"
	case Disabled:
		return "Disabled"
	case Requested:
		return "Requested"
	case Pending:
		return "Pending"
	case Rejected:
		return "Rejected"
	}

	return "unknown"
}
