package main

import "net/http"

type MaskTransport struct {
	http.Transport
}

func (m *MaskTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", " Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003 Build/QQ3A.200605.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.62 XWEB/2469 MMWEBSDK/200701 Mobile Safari/537.36 MMWEBID/3234 MicroMessenger/7.0.17.1701(0x27001141) Process/toolsmp WeChat/arm64 GPVersion/1 NetType/WIFI Language/zh_CN ABI/arm64")
	request.Header.Set("Origin", "https://v-xxtb.zust.edu.cn")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-Requested-With", "com.tencent.mm")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Referer", "https://v-xxtb.zust.edu.cn/web/mobile47/")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	request.Header.Set("Cookie", "iPlanetDirectoryPro=AQIC5wM2LY4SfcyPOi8nV6eCJX6wOOsO2eZcn%2Bq33e4FGNE%3D%40AAJTSQACMDE%3D%23; Hm_lvt_0d261a1cc090e61ad0a6fc0eb2f4fada=1598807328,1600665366; Hm_lpvt_0d261a1cc090e61ad0a6fc0eb2f4fada=1600665366")

	return m.Transport.RoundTrip(request)
}
