/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 10:16:01
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-28 11:02:20
 * @FilePath: /dao-generator/cmd/gen.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cmd

import (
	"github.com/chein-huang/dao-generator/pkg/generator"
	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	// genCmd represents the commit command
	var genCmd = &cobra.Command{
		Use: "gen",
	}
	var input string
	var mode model.ORMType
	var output string
	genCmd.Flags().StringVarP(&input, "input", "i", "", "input path")
	genCmd.Flags().IntVar((*int)(&mode), "orm", 1, "1: gorm")
	genCmd.Flags().StringVarP(&output, "output", "o", "", "output path")

	genCmd.Run = func(cmd *cobra.Command, args []string) {
		if input == "" {
			zap.L().Info("file invalid", zap.String("file", input))
			return
		}
		err := generator.Generate(input, output, mode)
		CheckErr(err)
	}
	rootCmd.AddCommand(genCmd)
}
