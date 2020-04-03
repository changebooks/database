package database

import "fmt"

func IsWrite(mode int) bool {
	return mode == Write
}

func IsRead(mode int) bool {
	return mode == Read
}

func IsBackup(mode int) bool {
	return mode == Backup
}

func IsMode(num int) bool {
	return num == Write || num == Read || num == Backup
}

func CheckMode(num int) error {
	if IsMode(num) {
		return nil
	} else {
		return fmt.Errorf("unsupported mode %d, must be %d or %d or %d", num, Write, Read, Backup)
	}
}
