package cache

import (
	"log"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func StartRedis() error {
	log.Println("redis will be start in 3 seconds ...")
	cmd := exec.Command("nohup", "redis-server", "--port", "6379")
	f, err := os.Open("nohup.out")
	if os.IsNotExist(err) {
		f, err = os.Create("nohup.out")
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	cmd.Stdout = f
	cmd.Stderr = f
	if err := cmd.Start(); err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	return nil
}

func StopRedis() error {
	log.Println("redis shutdown...")
	return exec.Command("redis-cli", "shutdown", "nosave").Run()
}

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cache Suite")
}
