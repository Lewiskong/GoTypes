package main

type ReqPayPageInfo struct {
	PayPageId     int
	VipBid        int
	Qua           string
	KtBossChannel string
	UserType      string
}

type ReqCheckUserPayValid struct {
	PayPageId int

	UserType string
	UserId   string

	ClientIp    string
	Appid       string
	Openid      string
	AccessToken string

	VuserId   string
	Vusession string
}

type ReqGetValidPay struct {
	PayPageId     int
	VipBid        int
	Qua           string
	KtBossChannel string
	UserType      string

	UserId string

	ClientIp    string
	Appid       string
	Openid      string
	AccessToken string

	VuserId   string
	Vusession string
}

var req = ReqGetValidPay{
	20, 90, "QV=1&PR=VIDEO&PT=TVMORE&CHID=10009&RL=1920*1080&VN=2.6.0&VN_CODE=1604&SV=5.1.1&DV=JurassicPark&VN_BUILD=1007&MD=MiBOX3&BD=JurassicPark&TVKPlatform=2280603",
	"tx_tvmore", "qq", "", "", "101414262", "17070F5256E262DD8A0E011F03B8AE50", "B007AFFE847BE5A0BCBFB724ABAD1446", "", "",
}
