/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"GoChainHis/executor"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var subFile string

func (ut *UTCTime) UnmarshalJSON(b []byte) (err error) {
	if len(b) < 2 {
		log.Printf("UTCTime.UnmarshalJSON - content is empty")
		return errors.New("bad format")
	}

	if b[0] != '"' || b[len(b)-1] != '"' {
		log.Printf("UTCTime.UnmarshalJSON - parse '%s' failed : need a pair of '\"' surrounded", string(b))
		return errors.New("bad format")
	}

	str := string(b[1 : len(b)-1])
	if !strings.HasSuffix(str, " UTC") {
		log.Printf("UTCTime.UnmarshalJSON - parse '%s' failed : must has ' UTC' suffix, such as '2023-06-06 UTC'", string(b))
		return errors.New("bad format")
	}

	tmpTm, err := time.Parse("2006-01-02 UTC", str)
	if nil != err {
		log.Printf("UTCTime.UnmarshalJSON - parse '%s' failed : %s", str, err.Error())
		return err
	}
	*ut = UTCTime(tmpTm)
	return
}

type UTCTime time.Time

type CandidateDef struct {
	Date    UTCTime // Must has a suffix of 'UTC'
	BigNews []string
}

func checkNews(news string) (valid bool) {
	valid, err := regexp.MatchString(`^<CN>.+</CN><EN>.+</EN>$`, news)
	if err != nil {
		log.Printf("checkNews - MatchString failed : %s", err.Error())
		return false
	}
	return valid
}

// submitCmd represents the submit command
var SubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: "candidates file of big the days's news.",
	Long: `candidates file of big the days's news. Should follow json format of this structure:
type CandidateDef struct {
	Date    time.Time /* such as '2023-06-06 UTC', MUST has UTC suffix */
	BigNews []string /* at least 3 news */
}
'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("subFile : %+v\n", subFile)
		var content []byte
		content, err := ioutil.ReadFile(subFile)
		if err != nil {
			log.Printf("SubmitCmd - read file %s failed : %s", subFile, err.Error())
			os.Exit(1)
		}

		var cdf CandidateDef
		if err = json.Unmarshal(content, &cdf); nil != err {
			log.Printf("SubmitCmd - json.Unmarshal for file '%s' failed : %s", subFile, err.Error())
			os.Exit(1)
		}

		cdfUnix := time.Time(cdf.Date).Unix()
		if cdfUnix%(60*60*24) != 0 {
			log.Printf("SubmitCmd - bad time format, in file '%s', Date must align to day", subFile)
			os.Exit(1)
		}
		todaySec := (24 * 60 * 60) * (time.Now().Unix() / (24 * 60 * 60))
		if cdfUnix < todaySec-60*60*24*30 {
			log.Printf("SubmitCmd - bad time, in file '%s',time too early for now, should within 30 days before now.", subFile)
			os.Exit(1)
		}
		if cdfUnix >= todaySec {
			log.Printf("SubmitCmd - bad time, in file '%s', Date should before today.", subFile)
			os.Exit(1)
		}
		if len(cdf.BigNews) < 3 {
			log.Printf("SubmitCmd - at least 3 big news in file '%s'.", subFile)
			os.Exit(1)
		}
		for idx, news := range cdf.BigNews {
			if !checkNews(news) {
				log.Printf("SubmitCmd - news[%d]:%s is not a valid news item :", idx, news)
				os.Exit(1)
			}
		}
		if len(argDaySlots) != 1 {
			log.Printf("SubmitCmd - need one -d/--dayslot args in command line.")
			os.Exit(1)
		}
		var utm UTCTime
		parseStr := argDaySlots[0]
		if !strings.HasSuffix(parseStr, " UTC") {
			parseStr += " UTC"
		}
		if err = utm.UnmarshalJSON([]byte("\"" + parseStr + "\"")); nil != err {
			log.Printf("SubmitCmd - cmd line's -d/--dayslot (%s) not match with date item in %s", argDaySlots[0], subFile)
			os.Exit(1)
		}
		if (time.Time(utm)).Unix() != (time.Time(cdf.Date)).Unix() {
			log.Printf("SubmitCmd - cmd line's -d/--dayslot (%s) not match with date item in %s", argDaySlots[0], subFile)
			os.Exit(1)
		}
		log.Printf("SubmitCmd - will submit content : %s", string(content))
		executor.GetExecutor().SubmitCandidates(uint64(cdfUnix), cdf.BigNews)
	},
}

func init() {
	SubmitCmd.Flags().StringVarP(&subFile, "subFile", "s", "", "candidate file path which hold the content of the days's big news in json format")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
