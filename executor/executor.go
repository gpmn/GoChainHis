package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"GoChainHis/Card"
	"GoChainHis/Escrow"
	"GoChainHis/History"
	"GoChainHis/util"
)

type Configure struct {
	Provider        string
	ChainID         int64
	MyAddr          string
	MyPrivKey       string
	HistoryContract string
	CardContract    string
	EscrowContract  string
}

type Executor struct {
	Configure
	client          *ethclient.Client
	myAddr          common.Address
	myPrivKey       *ecdsa.PrivateKey
	historyContract common.Address
	cardContract    common.Address
	escrowContract  common.Address
	escrow          *Escrow.Escrow
	history         *History.History
	card            *Card.Card
	signer          bind.SignerFn
}

var executor Executor

func GetExecutor() *Executor {
	return &executor
}

func CheckAck(hint string, extraWaitSec int) (passed bool) {
	fmt.Printf(hint)
	var ack string
	cnt, err := fmt.Scanf("%s", &ack)
	if nil != err {
		log.Printf("CheckAck - fmt.Scanf failed : %s", err.Error())
		return false
	}
	if cnt != 1 {
		log.Printf("CheckAck - fmt.Scanf got %d resp, invalid", cnt)
		return false
	}
	if ack != "Y" && ack != "y" {
		log.Printf("CheckAck - fmt.Scanf got ack '%s', not 'Y' or 'y', skip withdraw", ack)
		return false
	}
	if extraWaitSec > 0 {
		fmt.Printf("Please DOUBLE CHECK, press CTRL-C to break\n")
		for extraWaitSec > 0 {
			fmt.Printf("    wait %d second ...\n", extraWaitSec)
			time.Sleep(time.Second)
			extraWaitSec--
		}
	}
	return true
}

func DaySlotFromStr(daySlotStr string, offset int) (tm time.Time, err error) {
	tm, err1 := time.Parse("2006-01-02 UTC", daySlotStr+" UTC")
	if nil != err {
		sec, err2 := strconv.ParseInt(daySlotStr, 10, 64)
		if err2 != nil {
			log.Printf("DaySlotFromStr - '%s' can not interpret as date:%s. also can not interpret as integer:%s.",
				daySlotStr, err1.Error(), err2.Error())
			err = errors.New("bad day slot")
			return
		}
		tm = time.Unix(sec, 0)
	}

	if tm.Unix()%(60*60*24) != 0 {
		log.Printf("DaySlotFromStr - daySlotStr:'%s' => %s, not align to UTC day", daySlotStr, tm.Format(util.FavoredTimeFormat))
		err = errors.New("not align to a UTC day")
		return
	}

	if offset != 0 {
		tm = tm.Add(time.Duration(offset*24) * time.Hour)
	}

	return
}

func (exe *Executor) LoadConfigure(confPath string) (err error) {
	var content []byte
	content, err = ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("exe.LoadConfigure - read file %s failed : %s", confPath, err.Error())
		return err
	}

	if err = json.Unmarshal(content, &exe.Configure); nil != err {
		log.Printf("exe.LoadConfigure - json.Unmarshal for file %s failed : %s", confPath, err.Error())
		return err
	}

	if exe.MyPrivKey != "" { // compare private key and addr
		privateKey, err := crypto.HexToECDSA(exe.MyPrivKey[2:])
		if err != nil {
			log.Printf("exe.LoadConfigure - private key in %s invalid, HexToECDSA failed  : %s", confPath, err.Error())
			return err
		}

		pubKey := crypto.PubkeyToAddress(privateKey.PublicKey)
		if pubKey.String() != exe.MyAddr {
			log.Printf("exe.LoadConfigure - in %s private key and address not match", confPath)
			return errors.New("key and address not match")
		}
		exe.myPrivKey = privateKey
	}

	// log.Printf("LoadConfigure - Provider        : %s", exe.Provider)
	// log.Printf("                MyAddr          : %s", exe.MyAddr)
	// log.Printf("                HistoryContract : %s", exe.HistoryContract)
	// log.Printf("                CardContract    : %s", exe.CardContract)
	// log.Printf("                EscrowContract  : %s", exe.EscrowContract)
	exe.myAddr = common.HexToAddress(exe.MyAddr)
	exe.historyContract = common.HexToAddress(exe.HistoryContract)
	exe.cardContract = common.HexToAddress(exe.CardContract)
	exe.escrowContract = common.HexToAddress(exe.EscrowContract)

	exe.client, err = ethclient.Dial(exe.Provider)
	if err != nil {
		log.Printf("LoadConfigure - Dial to %s failed:%s", exe.Provider, err.Error())
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(exe.myPrivKey, big.NewInt(exe.ChainID))
	if nil != err {
		log.Printf("LoadConfigure - NewKeyedTransactorWithChainID failed : %s", err.Error())
		return err
	}
	exe.signer = auth.Signer

	exe.escrow, err = Escrow.NewEscrow(exe.escrowContract, exe.client)
	if nil != err {
		log.Printf("LoadConfigure - Escrow.NewEscrow failed: %s", err.Error())
		return err
	}

	exe.card, err = Card.NewCard(exe.cardContract, exe.client)
	if nil != err {
		log.Printf("LoadConfigure - Card.NewCard failed: %s", err.Error())
		return err
	}

	exe.history, err = History.NewHistory(exe.historyContract, exe.client)
	if nil != err {
		log.Printf("LoadConfigure - History.NewHistory failed: %s", err.Error())
	}

	return
}

func (exe *Executor) EscrowDump(addrStr string) (err error) {
	var addr common.Address
	if addrStr == "" {
		addr = exe.myAddr
	} else {
		addr = common.HexToAddress(addrStr)
	}
	res, err := exe.escrow.GetDepositInfoTo0(&bind.CallOpts{}, addr)
	if nil != err {
		log.Printf("EscrowDump - GetDepositInfoTo0 failed : %s", err.Error())
		return err
	}
	balance, err := exe.client.BalanceAt(context.Background(), addr, nil)
	fmt.Println("******************************************************************************")
	fmt.Printf("* Escrow Dump for addr : %-53s*\n", addr.String())
	fmt.Println("******************************************************************************")
	fmt.Printf("    Deposit Time    : %-20s   (may be twick if portion withdrawn)\n", time.Unix(res.DepositTm.Int64(), 0).UTC().Format(util.FavoredTimeFormat))
	fmt.Printf("    My Balance(ETH) : %-20s   (balance of my wallet : %s)\n", util.ToDecimal(balance, 18), addr.String())
	fmt.Printf("    My Escrow(ETH)  : %-20s   (record of contract   : %s)\n", util.ToDecimal(res.Amount, 18), exe.EscrowContract)
	fmt.Printf("    My Vote Shares  : %-20s   (compared with ETH)\n", util.ToDecimal(res.VoteAmt, 18))

	return nil
}

func (exe *Executor) EscrowDeposit(amount float64) (err error) {
	hint := fmt.Sprintf(`This will deposit %f eth from your wallet:%s to contract:%s.
Are Your Sure?(y/N)`,
		amount, exe.MyAddr, exe.escrowContract)
	if !CheckAck(hint, 3) {
		log.Printf("EscrowDeposit - Canceled by user")
		return errors.New("canceled by user")
	}

	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		Value:     util.ToWei(amount, 18),
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.escrow.Deposit(opts)
	if nil != err {
		log.Printf("EscrowDeposit - failed, err:%s", err.Error())
		return err
	}

	log.Printf("EscrowDeposit - Deposit %f ETH successfully, tx hash : %s", amount, tx.Hash())

	return err
}

func (exe *Executor) EscrowWithdraw(amount float64) (err error) {
	hint := fmt.Sprintf(`This will withdraw %f eth from contract:%s to your wallet:%s (if affordable, s.a. output of : GoChainHis escrow dump).
Are Your Sure?(y/N)`,
		amount, exe.MyAddr, exe.escrowContract)

	if !CheckAck(hint, 3) {
		log.Printf("EscrowWithdraw - Canceled by user")
		return errors.New("canceled by user")
	}
	var ack string
	cnt, err := fmt.Scanf("%s", &ack)
	if nil != err {
		log.Printf("EscrowWithdraw - fmt.Scanf failed : %s", err.Error())
		return err
	}
	if cnt != 1 {
		log.Printf("EscrowWithdraw - fmt.Scanf got %d resp, invalid", cnt)
		return errors.New("bad input")
	}
	if ack != "Y" && ack != "y" {
		log.Printf("EscrowWithdraw - fmt.Scanf got ack '%s', not 'Y' or 'y', skip withdraw", ack)
		return errors.New("canceled by user")
	}

	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.escrow.Withdraw(opts, util.ToWei(amount, 18))
	if nil != err {
		log.Printf("EscrowWithdraw - failed, err:%s", err.Error())
		return err
	}

	log.Printf("EscrowWithdraw - Withdraw %f ETH successfully, tx hash : %s", amount, tx.Hash())

	return err
}

func (exe *Executor) HistoryDump(daySlotStr string, offset int) (err error) {
	tmpTm, err := DaySlotFromStr(daySlotStr, offset)
	if nil != err {
		log.Printf("HistoryDump - DaySlotFromStr(%s, %d) failed : %s", daySlotStr, offset, err.Error())
		return err
	}
	unixSec := tmpTm.Unix()

	maPrice, err := exe.history.MAAuctionPrice(&bind.CallOpts{})

	storyCnt, err := exe.history.GetHisRecStoriesCnt(&bind.CallOpts{}, uint64(unixSec))
	if nil != err {
		log.Printf("HistoryDump - daySlotStr:'%s' => %s GetHisRecStoriesCnt failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
			err.Error())
	}
	record, err := exe.history.HistoryMap(&bind.CallOpts{}, uint64(unixSec))
	if nil != err {
		log.Printf("HistoryDump - daySlotStr:'%s' => %s HistoryDump failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
			err.Error())
	}
	bigStories, err := exe.history.GetBigStories(&bind.CallOpts{}, uint64(unixSec))
	if nil != err {
		log.Printf("HistoryDump - daySlotStr:'%s' => %s GetBigStories failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
			err.Error())
	}

	exists, err := exe.card.Exists(&bind.CallOpts{}, uint64(unixSec))
	if nil != err {
		log.Printf("HistoryDump - daySlotStr:'%s' => %s Exists failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
			err.Error())
	}
	var bidInfo string
	var tokenURI string
	var nftOwner common.Address
	if exists {
		curPrice, err := exe.history.QueryCurPrice(&bind.CallOpts{}, uint64(unixSec))
		if nil != err {
			log.Printf("HistoryDump - daySlotStr:'%s' => %s QueryCurPrice failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
				err.Error())
		}
		if record.Status == 3 {
			bidInfo = fmt.Sprintf("最新价:%s(eth) Status:%d 已拍卖", util.ToDecimal(curPrice, 18), record.Status)
		} else {
			bidInfo = fmt.Sprintf("最新价:%s(eth) Status:%d", util.ToDecimal(curPrice, 18), record.Status)
		}
		tokenURI, err = exe.card.TokenURI(&bind.CallOpts{}, big.NewInt(unixSec))
		if nil != err {
			log.Printf("HistoryDump - daySlotStr:'%s' => %s TokenURI failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
				err.Error())
		}
		nftOwner, err = exe.card.OwnerOf(&bind.CallOpts{}, big.NewInt(unixSec))
		if nil != err {
			log.Printf("HistoryDump - daySlotStr:'%s' => %s OwnerOf failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
				err.Error())
		}
	} else {
		bidInfo = "Not mint yet"
	}

	qcp, err := exe.history.QueryCurPrice(&bind.CallOpts{}, uint64(unixSec))
	var curPrice *big.Int
	if nil != err {
		curPrice = big.NewInt(0)
		if record.Status >= 2 { // _HistoryStatusMinted = 2
			log.Printf("HistoryDump - QueryCurPrice for %s failed:%s", daySlotStr, err.Error())
		}
	} else {
		curPrice = qcp.CurPrice
	}

	vi, err := exe.history.GetVoteInfo(&bind.CallOpts{}, uint64(unixSec), exe.myAddr)
	if nil != err {
		log.Printf("HistoryDump - GetVoteInfo for daySlot:'%s' addr:%s failed : %s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
			exe.myAddr, err.Error())
	}
	statusDesc := ""
	switch record.Status {
	case 0:
		statusDesc = fmt.Sprintf("%d - Prepare", record.Status)
	case 1:
		statusDesc = fmt.Sprintf("%d - Solved", record.Status)
	case 2:
		statusDesc = fmt.Sprintf("%d - Minted", record.Status)
	case 3:
		statusDesc = fmt.Sprintf("%d - Bidden", record.Status)
	default:
		statusDesc = fmt.Sprintf("%d - Invalid", record.Status)

	}
	fmt.Printf("%sHistory Dump :: %d => %s StoryCount:%d MAAuctionPrice:%s(ETH) %s\n",
		util.ColorYellow, unixSec, tmpTm.UTC().Format(util.FavoredTimeFormat), storyCnt,
		util.ToDecimal(maPrice, 18).String(),
		util.ColorSuffix)
	fmt.Printf("    VoteEndTm       : %s (As Local Time)\n", time.Unix(int64(record.VoteEndTm), 0).Local().Format(util.FavoredTimeFormat))
	fmt.Printf("    VoterCnt        : %d\n", record.VoterCnt)
	fmt.Printf("    VoteSum         : %s\n", record.VoteSum.String())
	fmt.Printf("    WeightedVoteSum : %s\n", record.WeightedVoteSum.String())
	fmt.Printf("    Status          : %s\n", statusDesc)
	fmt.Printf("    AucInitTm       : %s\n", time.Unix(int64(record.AucInitTm), 0).Local().Format(util.FavoredTimeFormat))
	fmt.Printf("    FinalPrice      : %s(ETH)\n", util.ToDecimal(record.FinalPrice, 18))
	fmt.Printf("    AucInitPrice    : %s(ETH)\n", util.ToDecimal(record.AucInitPrice, 18))
	fmt.Printf("    CurPrice        : %s(ETH)\n", util.ToDecimal(curPrice, 18))
	fmt.Printf("    Winner          : %s\n", record.Winner)
	fmt.Printf("    NFT Owner       : %s\n", nftOwner)
	fmt.Printf("    BigStoriesIdx   : [%d, %d, %d]\n", bigStories.Bis[0], bigStories.Bis[1], bigStories.Bis[2])
	fmt.Printf("    BidInfo         : %s\n", bidInfo)
	fmt.Printf("    Shares(D/S/C/V) : %d %d %d %d\n", record.ShareOfDev, record.ShareOfSecretary, record.ShareOfCards, record.ShareOfVoters)
	fmt.Printf("    NFT TokenURI    : %s\n", tokenURI)
	fmt.Printf("%s    Stories:%s\n", util.ColorYellow, util.ColorSuffix)

	for si := uint8(0); si < uint8(storyCnt); si++ {
		story, err := exe.history.GetHisRecStoryAt(&bind.CallOpts{}, uint64(unixSec), si)
		if nil != err {
			log.Printf("HistoryDump - daySlotStr:'%s' => %s GetHisRecStoryAt failed:%s", daySlotStr, tmpTm.Format(util.FavoredTimeFormat),
				err.Error())
		}
		mark := util.ColorWhite + " [ ] " + util.ColorSuffix
		if record.Status >= 1 {
			if bigStories.Bis[0] == si {
				mark = util.ColorRed + " [0] " + util.ColorSuffix
			} else if bigStories.Bis[1] == si {
				mark = util.ColorYellow + " [1] " + util.ColorSuffix
			} else if bigStories.Bis[2] == si {
				mark = util.ColorLightblue + " [2] " + util.ColorSuffix
			}
		}

		fmt.Printf("    Seq:%d Mark:%s RsvWeight:%d  VoteSum:%-20s(as eth)  Content:%s\n",
			si, mark,
			story.RsvWeight,
			util.ToDecimal(story.VoteSum, 18).String(),
			story.Content)
	}
	statusDesc = ""
	if vi.Status == 0 {
		statusDesc = "0 - NotVoted"
	} else if vi.Status == 1 {
		statusDesc = "1 - Voted"
	} else if vi.Status == 2 {
		statusDesc = "2 - Settled"
	} else {
		statusDesc = fmt.Sprintf("%d - Invalid", vi.Status)
	}
	fmt.Printf("%s    Vote Info for %s%s\n", util.ColorYellow, exe.myAddr, util.ColorSuffix)
	fmt.Printf("    MyVoteStatus    : %s\n", statusDesc)
	fmt.Printf("    MyVotePrefers   : [%d %d %d]\n", vi.Prefer0, vi.Prefer1, vi.Prefer2)
	fmt.Printf("    MyVoteAmt       : %s(as eth)\n", util.ToDecimal(vi.VoteAmt, 18).String())
	fmt.Printf("    MyVoteReward    : %s(as eth)\n", util.ToDecimal(vi.Reward, 18).String())
	fmt.Println()
	return nil
}

func (exe *Executor) SubmitCandidates(daySlot uint64, candidates []string) (err error) {
	sec, err := exe.history.Secretary(&bind.CallOpts{})
	if nil != err {
		log.Printf("SubmitCandidates - history.Secretary() failed: %s", err.Error())
		return err
	}
	if sec != exe.myAddr {
		log.Printf("SubmitCandidates - MyAddr is:%s, secretary is:%s, operation not allowed.", exe.myAddr, sec)
		return err
	}

	storyCnt, err := exe.history.GetHisRecStoriesCnt(&bind.CallOpts{}, daySlot)
	if nil != err {
		log.Printf("SubmitCandidates - GetHisRecStoriesCnt for daySlot:%d failed:%s", daySlot, err.Error())
		return err
	}
	if storyCnt > 0 {
		log.Printf("SubmitCandidates - daySlot:%d have submitted", daySlot)
		return errors.New("have submitted")
	}

	if !CheckAck("This will Submit candidate to history contract, only secretary allowed.Are Your Sure?(y/N)\n", 10) {
		log.Printf("SubmitCandidates - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.history.SubmitCandidates(opts, daySlot, candidates)
	if nil != err {
		log.Printf("SubmitCandidates - history.SubmitCandidates failed : %s", err.Error())
		return err
	}
	log.Printf("SubmitCandidates - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("SubmitCandidates - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("SubmitCandidates - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("SubmitCandidates - operation successfully, tx hash : %s", tx.Hash())
	return nil
}

func (exe *Executor) HistoryVote(daySlotStr string, prefers [3]uint8) (err error) {
	tm, err := DaySlotFromStr(daySlotStr, 0)
	if nil != err {
		log.Printf("HistoryVote - DaySlotFromStr(%s,0) failed : %s", daySlotStr, err.Error())
		return err
	}

	if !CheckAck("This will vote big news in history contract. Are Your Sure?(y/N)\n", 3) {
		log.Printf("HistoryVote - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}
	tx, err := exe.history.Vote(opts, uint64(tm.Unix()), prefers)

	if nil != err {
		log.Printf("HistoryVote - history.HistoryVote failed : %s", err.Error())
		return err
	}

	log.Printf("HistoryVote - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("HistoryVote - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("HistoryVote - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("HistoryVote - operation successfully, tx hash : %s", tx.Hash())
	return nil
}

func (exe *Executor) HistorySettleVoterReward(daySlots []string) (err error) {
	var tms []uint64
	for idx := range daySlots {
		tm, err := DaySlotFromStr(daySlots[idx], 0)
		if nil != err {
			log.Printf("HistorySettleVoterReward - bad dayslot '%s', err:%s", daySlots[idx], err.Error())
			return err
		}
		tms = append(tms, uint64(tm.Unix()))
	}

	hint := fmt.Sprintf("This will settle voters rewards for %+v. Are Your Sure?(y/N)\n", daySlots)
	if !CheckAck(hint, 3) {
		log.Printf("HistorySettleVoterReward - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}
	tx, err := exe.history.SettleVoteReward(opts, tms)

	if nil != err {
		log.Printf("HistorySettleVoterReward - history.SettleVoteReward failed : %s", err.Error())
		return err
	}

	log.Printf("HistorySettleVoterReward - operation successfully, tx hash : %s", tx.Hash())
	return nil
}

func (exe *Executor) HistorySettleOpsReward(daySlots []string) (err error) {
	var slots []uint64
	for idx := range daySlots {
		tm, err := DaySlotFromStr(daySlots[idx], 0)
		if nil != err {
			log.Printf("HistorySettleOpsReward - bad dayslot '%s', err:%s", daySlots[idx], err.Error())
			return err
		}
		slots = append(slots, uint64(tm.Unix()))
	}

	hint := fmt.Sprintf("This will settle ops rewards for %+v. Are Your Sure?(y/N)\n", daySlots)
	if !CheckAck(hint, 3) {
		log.Printf("HistorySettleOpsReward - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}
	tx, err := exe.history.SettleOpsReward(opts, slots)

	if nil != err {
		log.Printf("HistorySettleOpsReward - history.SettleOpsReward failed : %s", err.Error())
		return err
	}

	log.Printf("HistorySettleOpsReward - operation successfully, tx hash : %s", tx.Hash())
	return nil
}

func (exe *Executor) HistorySettleCardReward(daySlot string, fromSlots []string) (err error) {
	var froms []uint64
	for idx := range fromSlots {
		tm, err := DaySlotFromStr(fromSlots[idx], 0)
		if nil != err {
			log.Printf("HistorySettleCardReward - bad fromSlots '%s', err:%s", fromSlots[idx], err.Error())
			return err
		}
		froms = append(froms, uint64(tm.Unix()))
	}

	to, err := DaySlotFromStr(daySlot, 0)
	if nil != err {
		log.Printf("HistorySettleCardReward - bad dayslot '%s', err:%s", daySlot, err.Error())
		return err
	}

	hint := fmt.Sprintf("This will settle card rewards for %+v. Are Your Sure?(y/N)\n", fromSlots)
	if !CheckAck(hint, 3) {
		log.Printf("HistorySettleCardReward - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}
	tx, err := exe.history.SettleCardReward(opts, uint64(to.Unix()), froms)

	if nil != err {
		log.Printf("HistorySettleCardReward - history.SettleCardReward failed : %s", err.Error())
		return err
	}

	log.Printf("HistorySettleCardReward - operation successfully, tx hash : %s", tx.Hash())
	return nil
}

func (exe *Executor) HistoryResolve(daySlotStr string) (err error) {
	var slot uint64

	tm, err := DaySlotFromStr(daySlotStr, 0)
	if nil != err {
		log.Printf("HistoryResolve - bad dayslot '%s', err:%s", daySlotStr, err.Error())
		return err
	}
	slot = uint64(tm.Unix())
	record, err := exe.history.HistoryMap(&bind.CallOpts{}, slot)
	if nil != err {
		log.Printf("HistoryResolve - HistoryMap for %s failed : %s", daySlotStr, err.Error())
		return err
	}
	if record.Status != 0 { // must be in _HistoryStatusPrepare=0 status
		log.Printf("HistoryResolve - HistoryMap for %s status is :%d, invalid", daySlotStr, record.Status)
		return errors.New("bad record status")
	}

	hint := fmt.Sprintf("This will resolve for %+v. Are Your Sure?(y/N)\n", daySlotStr)
	if !CheckAck(hint, 3) {
		log.Printf("HistoryResolve - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.history.Resolve(opts, slot)
	if nil != err {
		log.Printf("HistoryResolve - failed, err:%s", err.Error())
		return err
	}
	log.Printf("HistoryResolve - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("HistoryResolve - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("HistoryResolve - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("HistoryResolve - Resolve for %s successfully, tx hash : %s", daySlotStr, tx.Hash())
	return nil
}

func (exe *Executor) HistoryBid(daySlotStr string) (err error) {
	var slot uint64

	tm, err := DaySlotFromStr(daySlotStr, 0)
	if nil != err {
		log.Printf("HistoryBid - bad dayslot '%s', err:%s", daySlotStr, err.Error())
		return err
	}
	slot = uint64(tm.Unix())

	qcp, err := exe.history.QueryCurPrice(&bind.CallOpts{}, slot)
	if nil != err {
		log.Printf("HistoryBid - QueryCurPrice for %s failed:%s", daySlotStr, err.Error())
		return err
	}

	if qcp.HisStatus != 2 { // must be _HistoryStatusMinted = 2 status
		log.Printf("HistoryBid - HistoryMap for %s status is :%d, invalid", daySlotStr, qcp.HisStatus)
		return errors.New("bad record status")
	}

	hint := fmt.Sprintf("This will bid for %+v at price : %s eth. Are Your Sure?(y/N)\n", daySlotStr, util.ToDecimal(qcp.CurPrice, 18))
	if !CheckAck(hint, 3) {
		log.Printf("HistoryBid - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
		Value:     qcp.CurPrice,
	}

	tx, err := exe.history.Bid(opts, slot)
	if nil != err {
		log.Printf("HistoryBid - failed, err:%s", err.Error())
		return err
	}
	log.Printf("HistoryBid - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("HistoryBid - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("HistoryBid - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("HistoryBid - Bid for %s successfully, tx hash : %s", daySlotStr, tx.Hash())
	return nil
}

func (exe *Executor) HistoryMintAndAuction(daySlotStr string, festivals []string) (err error) {
	var slot uint64
	tm, err := DaySlotFromStr(daySlotStr, 0)
	if nil != err {
		log.Printf("HistoryMintAndAuction - bad dayslot '%s', err:%s", daySlotStr, err.Error())
		return err
	}
	slot = uint64(tm.Unix())

	record, err := exe.history.HistoryMap(&bind.CallOpts{}, slot)
	if nil != err {
		log.Printf("HistoryMintAndAuction - HistoryMap for %s failed : %s", daySlotStr, err.Error())
		return err
	}
	if record.Status != 1 { // must be in _HistoryStatusSolved = 1 status
		log.Printf("HistoryMintAndAuction - HistoryMap for %s status is :%d, invalid", daySlotStr, record.Status)
		return errors.New("bad record status")
	}

	hint := fmt.Sprintf("This will Mint and Auction for %+v, with Festivals:%+v. Are Your Sure?(y/N)\n", daySlotStr, festivals)
	if !CheckAck(hint, 3) {
		log.Printf("HistoryMintAndAuction - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.history.MintAndAuction(opts, slot, festivals)
	if nil != err {
		log.Printf("HistoryMintAndAuction - MintAndAuction failed, err:%s", err.Error())
		return err
	}
	log.Printf("HistoryMintAndAuction - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("HistoryMintAndAuction - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("HistoryMintAndAuction - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("HistoryMintAndAuction - MintAndAuction for %s successfully, tx hash : %s", daySlotStr, tx.Hash())
	return nil
}

func (exe *Executor) CardClaim(daySlotStr string) (err error) {
	var slot uint64
	tm, err := DaySlotFromStr(daySlotStr, 0)
	if nil != err {
		log.Printf("CardClaim - bad dayslot '%s', err:%s", daySlotStr, err.Error())
		return err
	}
	slot = uint64(tm.Unix())

	record, err := exe.history.HistoryMap(&bind.CallOpts{}, slot)
	if nil != err {
		log.Printf("CardClaim - HistoryMap for %s failed : %s", daySlotStr, err.Error())
		return err
	}
	if record.Status != 3 { // must be in _HistoryStatusBidden = 3 status
		log.Printf("CardClaim - HistoryMap for %s status is :%d, invalid", daySlotStr, record.Status)
		return errors.New("bad record status")
	}

	hint := fmt.Sprintf("This will Claim card for %+v, only winner of auction allowed. Are Your Sure?(y/N)\n", daySlotStr)
	if !CheckAck(hint, 3) {
		log.Printf("CardClaim - canceled by user")
		return errors.New("canceled by user")
	}
	opts := &bind.TransactOpts{
		From:      exe.myAddr,
		GasTipCap: big.NewInt(GAS_TIP_IN_GWEI),
		GasLimit:  GAS_LIMIT,
		Signer:    exe.signer,
	}

	tx, err := exe.card.Claim(opts, slot)
	if nil != err {
		log.Printf("CardClaim - MintAndAuction failed, err:%s", err.Error())
		return err
	}
	log.Printf("CardClaim - TX %s executed, wait 5 seconds for the receipt ...", tx.Hash())
	time.Sleep(time.Second * 5)
	receipt, err := exe.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("CardClaim - TransactionReceipt %s failed:%s", tx.Hash(), err.Error())
		return err
	}

	if receipt.Status != 1 {
		log.Printf("CardClaim - tx %s operation status %d, error log:%v, failed!", tx.Hash(), receipt.Status, receipt.Logs)
		return errors.New("failed status")
	}
	log.Printf("CardClaim - MintAndAuction for %s successfully, tx hash : %s", daySlotStr, tx.Hash())
	return nil
}
