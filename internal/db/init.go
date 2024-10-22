package db

import (
	"fmt"
)

func Init() error {
	if err := DBInit(); err != nil {
		return fmt.Errorf("DBInit() : %v", err)
	}
	// if err := RedisInit(); err != nil {
	// 	return fmt.Errorf("RedisInit() : %v", err)
	// }
	// if err := KVStoreInit(); err != nil {
	// 	return fmt.Errorf("KVStoreInit() : %v", err)
	// }
	return nil
}

func Close() error {
	var err error
	if err1 := DBClose(); err1 != nil {
		err = fmt.Errorf("DBClose(): %v", err1)
	}
	if err1 := RDBClose(); err1 != nil {
		err1 = fmt.Errorf("RDBClose(): %v", err1)
		if err1 != nil {
			err = fmt.Errorf("%v %v", err, err1)
		} else {
			err = err1
		}
	}
	if err1 := KvStoreClose(); err1 != nil {
		err1 := fmt.Errorf("KvStoreClose(): %v", err1)
		if err1 != nil {
			err = fmt.Errorf("%v %v", err, err1)
		} else {
			err = err1
		}
	}
	return err
}
