package Redis

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"retail_shop/Config"
)

var (
	pool *redis.Pool
)

func init() {
	flag.Parse()
	pool = Config.NewPool()
}

const (
	lockScript = `
		return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2])
	`
	unlockScript = `
		if redis.call("get",KEYS[1]) == ARGV[1] then
		    return redis.call("del",KEYS[1])
		else
		    return 0
		end
	`
)

// Lock attempts to put a lock on the key for a specified duration (in milliseconds).
// If the lock was successfully acquired, true will be returned.
func Lock(key, value string, timeoutMs int) (bool, error) {
	r := pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, lockScript)
	fmt.Println(cmd)
	if res, err := cmd.Do(r, key, value, timeoutMs); err != nil {
		fmt.Print("heylllllyyyyyy")
		fmt.Println(res)
		fmt.Println("hellllo")
		fmt.Println(err)
		return false, err
	} else {
		fmt.Print("heyyyyyyy")
		return res == "OK", nil
	}
}

// Unlock attempts to remove the lock on a key so long as the value matches.
// If the lock cannot be removed, either because the key has already expired or
// because the value was incorrect, an error will be returned.
func Unlock(key, value string) error {
	r := pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, unlockScript)
	if res, err := redis.Int(cmd.Do(r, key, value)); err != nil {
		return err
	} else if res != 1 {
		return errors.New("ErrorMessage :Unlock failed, key or secret incorrect")
	}

	// Success
	return nil
}
