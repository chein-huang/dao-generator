/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-03-21 16:20:15
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-05-06 14:21:34
 * @FilePath: /dao-generator/cmd/root.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dao-generator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	encoderConf := zap.NewProductionEncoderConfig()

	consoleEncoder := zapcore.NewJSONEncoder(encoderConf)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zap.InfoLevel)

	// 创建zap core对象
	core := zapcore.NewTee(consoleCore)

	// 创建zap logger对象，同时添加两个option：日志打印行号、error级别的日志打印堆栈信息
	newLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer func() {
		_ = newLogger.Sync()
	}()

	zap.ReplaceGlobals(newLogger)
}

type CustomJSONEncoder struct {
	zapcore.Encoder
}

func (e *CustomJSONEncoder) Clone() zapcore.Encoder {
	return &CustomJSONEncoder{e.Encoder.Clone()}
}

// nolint
func (e *CustomJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	fields = append(fields, zap.String("dt", strconv.FormatInt(entry.Time.UnixNano(), 10)))
	return e.Encoder.EncodeEntry(entry, fields)
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// log.Info("Using config file:", zap.String("config file", viper.ConfigFileUsed()))
	} else {
		cobra.CheckErr(err)
	}
}
