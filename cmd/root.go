package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "neko",
	Short: "A large model-based ACM-ICPC algorithm problem automatic generation system",
	Long:  "A large model-based ACM-ICPC algorithm problem automatic generation system that can automatically generate algorithm problems, testcases, and problem solutions. ",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println(" -------- NekoACM -------- ")
			fmt.Println(" problem: 生成题目")
			fmt.Println(" testcase: 生成测试用例")
			fmt.Println(" solution: 生成题解")
			fmt.Println(" translate: 翻译题目")
			fmt.Println(" joke: 生成算法冷笑话")
			fmt.Println(" server: 启动服务器模式")
			fmt.Println(" exit: 退出")
			fmt.Println(" ------------------------- ")
			if err := clearBuffer(reader); err != nil {
				return err
			}
			fmt.Print("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			input = strings.ToLower(input)
			//fmt.Println(input)
			switch input {
			case "exit":
				os.Exit(0)
			case "server":
				if err := ServerCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "problem":
				if err := ProblemCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "testcase":
				if err := TestcaseCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "solution":
				if err := SolutionCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "translate":
				if err := TranslateCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "joke":
				if err := JokeCmd.RunE(nil, nil); err != nil {
					return err
				}
			default:
				fmt.Println("未知命令！")
			}
		}
	},
}

// 初始化
func init() {
	rootCmd.AddCommand(ServerCmd)
	rootCmd.AddCommand(ProblemCmd)
	rootCmd.AddCommand(TestcaseCmd)
	rootCmd.AddCommand(SolutionCmd)
	rootCmd.AddCommand(TranslateCmd)
	rootCmd.AddCommand(JokeCmd)
}

// 执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
