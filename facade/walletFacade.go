package facade

import "fmt"

type walletFacade struct {
	account      *account
	ledger       *ledger
	notification *notification
	securityCode *securityCode
	wallet       *wallet
}

func newWalletFacade(accountID string, code int) *walletFacade {
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}

	fmt.Println("Account created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting to add money to the wallet.")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
