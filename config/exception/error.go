package exception

import (
	"fmt"
	"os"
)

func CustomError(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func ZookeeperConnectError() {
	fmt.Println("Zookeeper container not found, please run container first")
	os.Exit(11)
}
