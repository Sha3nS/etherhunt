/*
 Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"fmt"
	"github.com/shawncles/etherhunt/main/cmd"
	"github.com/spf13/cobra"
)

func main() {
	mainCmd := &cobra.Command{Use: "eth-client"}
	mainCmd.AddCommand(cmd.InvokeCMD())
	err := mainCmd.Execute()
	if err != nil {
		fmt.Printf("eth-client server start error, %v", err)
	}

	return
}