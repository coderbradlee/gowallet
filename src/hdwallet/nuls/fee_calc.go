package nuls

import "errors"

const (
	MIN_PRECE_PRE_1024_BYTES   = 100000
	OTHER_PRECE_PRE_1024_BYTES = 1000000
	KB                         = 1024
)

func GetTransferFee(size int) Na {
	fee := MIN_PRECE_PRE_1024_BYTES * size / KB
	if size%KB > 0 {
		fee += MIN_PRECE_PRE_1024_BYTES
	}
	return Na(fee)
}

func GetMaxFee(size int) Na {
	fee := OTHER_PRECE_PRE_1024_BYTES * size / KB
	if size%KB > 0 {
		fee += OTHER_PRECE_PRE_1024_BYTES
	}
	return Na(fee)
}


func getFee(size int, price Na) (Na,error) {
	if price<MIN_PRECE_PRE_1024_BYTES {
		return 0,errors.New("The price is too low!")
	}
	if price>OTHER_PRECE_PRE_1024_BYTES {
		return 0,errors.New("The price is too high!")
	}
	fee := int(price) * int(size / KB)
	if size%KB > 0 {
		fee += int(price)
	}
	return Na(fee),nil
}