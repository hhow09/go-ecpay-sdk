package ecpay_test

import (
	"strconv"
	"time"

	"github.com/hhow09/go-ecpay-sdk"
	ecpayBase "github.com/hhow09/go-ecpay-sdk/base"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"github.com/hhow09/go-ecpay-sdk"
)

var _ = Describe("Client", func() {
	Context("CheckMac", func() {
		It("mac should calculate correct", func() {
			client := ecpay.NewStageClient()
			values := map[string]string{"test": "test"}
			mac := client.GenerateCheckMacValue(values)
			Expect(mac).To(Equal("B5947417B0FB8C1B8EBCA08E2FF3D66B7FB2329C8529C1331FE2268E9430903A"))
		})
	})

	Context("AutoPostTmpl", func() {
		It("shouldn't fail to generate auto post form html", func() {
			client := ecpay.NewStageClient()
			values := map[string]string{
				"test": "test",
				"Test": "C",
				"D":    "E",
				"F":    "G",
			}
			html := client.GenerateAutoSubmitHtmlForm(values, "http://test.test.test/test")
			//print(html)
			Expect(html).To(Equal(`<form id="order_form" action="http://test.test.test/test" method="post">
    <input type="hidden" name="D" id="D" value="E" />
    <input type="hidden" name="F" id="F" value="G" />
    <input type="hidden" name="Test" id="Test" value="C" />
    <input type="hidden" name="test" id="test" value="test" />
</form>
<script>document.querySelector("#order_form").submit();</script>`))
		})
	})

	Context("Order", func() {
		It("should success to create a normal order request", func() {
			client := ecpay.NewStageClient(ecpay.WithReturnURL("https://example.com/ecpay/return"), ecpay.WithPeriodReturnURL("https://example.com/ecpay/period"))
			now := time.Now()
			tradeNo := "testLuck" + strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
			request := client.CreateOrder(tradeNo, now, 400, "無論如何都不太會獲得獎勵的服務，你錢太多了\n不要錢可以送我\n(每分鐘檢測一次)\nspecial character -_.!)(*", []string{"你好"}).
				SetCreditPayment().
				WithCreditOptional(ecpay.AioCheckOutCreditOptional{
					BindingCard:      ecpayBase.BINDINGCARDENUM_BINDING,
					MerchantMemberID: client.MerchantID() + "_test_member",
				}).
				WithCreditPeriodOptional(ecpayBase.CREDITPERIODTYPEENUM_DAY, 2, 4)
			//
			//testdata,err := request.GenerateUrlQuery()
			//Expect(err).To(Succeed())
			//fmt.Println(testdata)
			//fmt.Println(ecpay.FormUrlEncode(testdata))
			//
			mac, err := request.GenerateCheckMac()
			Expect(err).To(Succeed())
			html, err := request.GenerateRequestHtml()
			Expect(err).To(Succeed())
			//fmt.Println(html)
			expected := `<form id="order_form" action="https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5" method="post">
    <input type="hidden" name="BindingCard" id="BindingCard" value="1" />
    <input type="hidden" name="CheckMacValue" id="CheckMacValue" value="` + mac + `" />
    <input type="hidden" name="ChoosePayment" id="ChoosePayment" value="Credit" />
    <input type="hidden" name="EncryptType" id="EncryptType" value="1" />
    <input type="hidden" name="ExecTimes" id="ExecTimes" value="4" />
    <input type="hidden" name="Frequency" id="Frequency" value="2" />
    <input type="hidden" name="ItemName" id="ItemName" value="你好" />
    <input type="hidden" name="MerchantID" id="MerchantID" value="2000132" />
    <input type="hidden" name="MerchantMemberID" id="MerchantMemberID" value="2000132_test_member" />
    <input type="hidden" name="MerchantTradeDate" id="MerchantTradeDate" value="` + ecpayBase.ECPayDateTime(now).String() + `" />
    <input type="hidden" name="MerchantTradeNo" id="MerchantTradeNo" value="` + tradeNo + `" />
    <input type="hidden" name="PaymentType" id="PaymentType" value="aio" />
    <input type="hidden" name="PeriodAmount" id="PeriodAmount" value="400" />
    <input type="hidden" name="PeriodReturnURL" id="PeriodReturnURL" value="https://example.com/ecpay/period" />
    <input type="hidden" name="PeriodType" id="PeriodType" value="D" />
    <input type="hidden" name="ReturnURL" id="ReturnURL" value="https://example.com/ecpay/return" />
    <input type="hidden" name="TotalAmount" id="TotalAmount" value="400" />
    <input type="hidden" name="TradeDesc" id="TradeDesc" value="無論如何都不太會獲得獎勵的服務，你錢太多了 不要錢可以送我 (每分鐘檢測一次) special character -_.!)(*" />
</form>
<script>document.querySelector("#order_form").submit();</script>`
			//fmt.Println(html)
			Expect(html).To(Equal(expected))
		})
	})
})
