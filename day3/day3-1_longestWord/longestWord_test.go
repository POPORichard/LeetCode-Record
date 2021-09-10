package day3_1_longestWord

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestWord(words []string) string {
	length := len(words)
	//以单词长度建立映射序列
	lengthSequence := make([]int, length, length)
	for i := range words {
		lengthSequence[i] = len(words[i])
	}
	//按单词长度重新排序
	for i:=0; i<length-1; i++{
		least := lengthSequence[i]
		target := -1
		for t:= i+1; t<length; t++{
			if lengthSequence[t] < least{
				least = lengthSequence[t]
				target = t
			}
		}
		if target != -1{
			tmp := lengthSequence[i]
			lengthSequence[i] = lengthSequence[target]
			lengthSequence[target] = tmp
			tmpS := words[i]
			words[i] = words[target]
			words[target] = tmpS
		}
	}

	//将满足条件的序列单词放入lines中
	lines := make([]string, 0,10)
	for i := range words{
		if lengthSequence[i] == 1{
			tmp := words[i]
			lines = append(lines,tmp)
			continue
		}
		for t := range lines{
			if len(lines[t]) != lengthSequence[i]-1{continue}
			if words[i][:lengthSequence[i]-1] == lines[t] {
				lines = append(lines, words[i])
			}
		}
	}

	//找到lines中最长的单词，放入res中
	if len(lines) == 0 {return ""}
	longest := len(lines[len(lines)-1])
	res := make([]string,0,10)
	for i:=len(lines)-1; i>=0; i--{
		if len(lines[i]) < longest {break}
		res = append(res, lines[i])
	}

	fmt.Println(res)

	//若res中只有一个单词则直接返回
	longestWord  := res[0]
	if len(res) == 1{return longestWord}
	//按字典序比较res中的单词
	for i:=1; i<len(res); i++{
		for t := range res[i]{
			if longestWord[t] > res[i][t]{
				longestWord = res[i]
				break
			}
			if longestWord[t] < res[i][t]{break}
		}
	}
	return longestWord
}

func Test_longestWord(t *testing.T){
	//case 1
	words := []string{"a", "banana","app","appl","ap","apply","apple"}
	res := longestWord(words)
	assert.Equal(t, "apple", res, "Error in case 1")

	//case 2
	words = []string{"m","mo","moc","moch","mocha","l","la","lat","latt","latte","c","ca","cat"}
	res = longestWord(words)
	assert.Equal(t, "latte", res, "Error in case 2")

	//case 3
	words = []string{"ovvnb","itcudluaycsspjrohf","cqwec","nz","mkomsvmrauphvm","mkomsvmrauphvmm","nndarhbeme","hdtvr","paqwiazmmvgu","zaydypcbfttmu","vkrgrzgrekblv","wwsaejhwqhjp","tveepjwpmghvi","hklbpj","ovl","wqmwehun","sqmh","yizslwlpndzjihon","xgjpvpjdiccwean","qfhggvatefvw","ixlfvbvdzwn","zaydypcbfttm","yfxsvg","nnda","uvdfwryvvugthazsf","keizhk","eurm","vkrgr","ncghfsccrgqkafhoa","omq","mnexyhlutc","ncghfsccrgqk","hklbpjmced","c","asglpcp","hklbpjmcedqcwnbuowt","asglpcpgct","hzzpjoos","yizslwlpndzj","zvkr","wwsae","momqxn","vkrgrzgrekbl","qdowgpmdjco","mnex","mnexyhlutclbwvvcg","lmwrglany","ctgu","lmwrglanyi","qdowgpmdjcowemkacwzoxu","it","dkwpdp","tiufhfhomricf","zfzp","itc","riwisyimddyguitvtdwg","aypotqdjlcevgvf","jap","ay","hzzpjoossvzgygxhgsiyd","y","nndarh","mnexyhlutclbwvvcgdzecsuvf","hklbpjmcedqcwnbuo","hzzpjoossvzgygxhg","tveepjwpmgh","hzzpjoossvzgyg","zvkroe","vkrgrzg","mnexyhlutclbwvvc","bzl","zaydypcbf","bqwkb","aqfmw","tpktncga","ovv","japawosnoio","ixlfvbv","japawosnoiobnr","fcndecv","cxto","xtwmnupe","uv","riwisyimddyguitvtdwgyzu","volzppbtszdnoocqylo","riwisyimddyguitvtdw","hklbpjmcedqcwnbuowtt","vgirdgjs","bzlrgwxnsqdzuxps","ixlfvbvdzwndvfi","riwis","yiz","uiuw","yizslwlpndzjihonftle","qlvtdmy","paqwiazm","kcudnbvfu","yizslwlpndzjihonftlel","fcnd","gslbk","omqn","gslbknct","uvdfwryvvugthazsfqukaikv","qlvtdmyeyfdxqamek","paqwiazmmv","bzlrgwxnsqdzuxp","ovvnbs","uvdfw","je","ov","wwsaejhwqhjpcr","vgirdg","iaydxhytj","zfzpuxghs","erkufypwnu","wqm","sqmht","wqmwehunbu","mnexyhlutclbwvvcgd","wqmwehunbuzfh","ncghfscc","qdowgpmdjcowemk","erkufyp","zaydypc","ixlfvbvdzwndvfir","nn","vg","hklbpjm","hzzpjoossvzgygx","mkomsvmrauph","ke","ixlfvbvd","asglpcpgctpdc","uiuwyorgme","lmwrglanyij","hk","omqndxpxevxoj","ncghfs","paqwiazmmvgudly","itcud","wvep","wws","nndarhbememc","mnexyh","ia","xgj","qdowgpmdjcowemkacwzo","utkd","japawosno","qfhggvatefvwmqb","xtwmnup","momqxnl","volzppbtszdnoocqy","tp","ixlfvbvdz","dcvdpdzoiveh","hklbpjmce","sqmhtpcgds","utkdk","eurmri","qdowgpmdjcowem","hzzpjoossv","omqndxp","xtwmnupebcuzb","kcudnbvfugbntkdgjst","zfzpuxghsf","ncgh","sqmhtpcgdsdf","tpktnc","kcudnbvfugbnt","eurmricmwbul","hdtvrl","sqmhtpcgdsdfqgxs","wwsaejhwqhjpcrgq","pgd","asglpc","wqmwehunbuzfhpia","asglpcpgctpd","itcudluaycsspjrohfr","wwsaejhwqhj","mnexyhlutclbwvvcgdz","ixlfv","momq","gslbknctqw","nzwwarlivffi","yf","cxtoslemgfy","riwisyimddyguitvtdwgyz","asglp","hzzp","dcvdpdz","qfhg","ui","riwisyimddyguitvtd","ncghfsccrg","qlvtd","wqmwehunbuz","kcudnbvfug","paqwiazmmvgudlyu","hklbpjmcedqcwnbu","mn","lemdqgqw","nzwwarl","dcvdpdzo","asglpcpgctpdcehw","zvkroeaj","qfh","qdow","zfzpuxghsfme","itcudluaycsspjrohfrniw","nndarhbememcvspya","paq","lem","lemdqgq","kcudnbvfugbntkdgjstx","omqndxpxevxojf","qdowgpmdjc","qfhggvatefvwm","wqmwehu","zfzpux","tveepjwpmghviay","uvdfwryvvugt","lemdqgqwum","qdowgpmdjcow","eurmricm","wve","tveepjwp","qdowgpmdjcowemkacw","japawosnoiobnrfjrsbrcuz","uiuwyorgmes","iay","uvdfwr","gkx","uvdfwryvvugthazsfq","itcudluaycsspj","dcvdp","cqwecm","ctguxt","hklb","vkrgrzgr","xt","wqmwe","ovvnbszsv","wqmweh","cx","momqx","ncghfsc","nzwwar","gnyrzd","xgjpvpjdiccw","bzlrgwxnsqdz","er","tpktncgausglce","zvkroeajidjenjf","hklbpjmc","bfbh","aypotqdjlc","fqda","h","m","sqmhtpcgdsdfqgxsjgqiz","qv","qfhggvatefv","sqmhtpcgdsdfqgxsjgqi","xa","pg","xgjpvpjd","riwisy","ixl","hdt","fqdauiilsnrwbw","hklbpjmcedqcwnbuowtti","vkrgrzgrek","wwsaejhwqhjpcrg","volzppbt","itcudluaycsspjrohfrn","mkom","xgjpvpjdiccwea","qd","riwisyimddy","ayp","erkufypw","mnexyhlutclbwvvcgdzec","uvdfwryvvugthaz","lemdqgqwu","qfhggvatefvwmqbxtd","k","fukw","volzpp","kcudnbvf","sqmhtp","eur","t","qf","fqdauiilsnrw","japawosnoiobnrf","tiufhfhomricfomkr","zvkroea","om","erkufypwnuhd","iaydxhytje","v","xgjpvp","ncghfsccrgqka","kcudnbv","gsl","aypo","mnexyhlutclbwvv","paqwia","volzppbtszdnoocq","nndarhbememcvsp","utkdkjorodpk","ovvn","hklbpjmcedq","paqwiazmm","mne","fcnde","riwisyimddyguitvtdwgyzui","uvdfwryvvugthazs","xtwmnu","nndarhb","iaydxhytjekusma","xtwmnupebcuzbgn","lemdqgqwumegvfexckxgiadwl","wqmwehunb","qdowgpmdj","aypotqdjlce","snv","paqwiazmmvgudlyuw","paqwi","yizslwlp","nzww","bzlrgwxnsqdzu","zvkroeajidj","hzzpjoo","lemdqgqwumegv","lmwrgla","cqwecmw","fqdauii","qdo","ksu","gnyr","qfhggvatefvwmqbxtdnb","ovvnbszsvkj","gslbknc","qdowgpmdjcowemka","omqndxpxevxojfw","hzzpjoossvzgygxhgsiydu","hzzpjoossvzgygxhgsi","qdowgpm","qdowgpmdjcowemkacwzoxukg","qdowg","tiufhfho","utkdkjorodpksiu","eurmricmw","hzzpjoossvzgygxhgsiyduod","cxtosl","zf","lemdqgqwumegvfexckxgiadwlc","em","jeqx","asglpcpgctpdce","kcudnbvfugbntk","z","xtwmnupebcuzbg","yizslw","bfbhnennr","vgir","vk","mko","ovvnbszsvkjn","tiufhfhomricfomk","lmwr","momqxnld","uiuwyorgmesm","zaydypcb","cxt","fqd","volzppb","tpktncgausglcelnpn","iaydxh","ctguxta","hzzpjoossvzgy","lemdq","omqndxpxe","mnexyhl","xtwmn","yizslwlpndzjihonftl","qlvt","bzlrgw","emap","dcvd","tveep","aypotqdj","o","tpktncgausglc","rxgjsers","qfhgg","bqwkbp","uvdfwryvvugthazsfqukaik","mkomsvmra","kcudn","nzwwarlivf","yizslwlpndz","dcvdpdzoive","iaydxhytjekusmae","lemdqgqwumegvfexc","yizslwlpndzjihonf","bfbhnennrn","wqmwehunbuzfhp","bzlr","zay","uvdfwryvvu","tpkt","uw","bqw","ovvnbsz","uiu","zayd","cxtos","g","fc","fcn","japa","riwisyimddygui","qfhggvat","hklbpjmcedqcwn","hklbp","mnexyhlut","lemdqgqwumegvfexckx","volzppbtszdnooc","xtwmnupeb","xtw","fqdauiilsnrwb","bzlrgwxns","keiz","kcudnbvfugb","bz","x","bzlrgwxnsqdzuxpsa","vol","mkomsvmrau","japawosnoiobnrfjrsbr","qlvtdmyey","ixlfvbvdzwndv","gny","vkrgrzgrekblvg","riwisyimd","itcudluaycs","gslbknctq","rj","yizslwl","mkoms","rxgjs","wv","tveepjwpmghviayy","yfxs","tiufhfhomricfom","q","qfhggv","tpktncg","yizslwlpndzjiho","tveepjwpmghviayyt","xgjpvpjdic","hzzpjo","tiufh","tpktncgausglcelnpno","hz","japawosnoiobnrfjrsbrcu","fqdauiilsn","japawosnoiob","volzppbtszdnoocqyl","dc","japawosn","ncg","hklbpjmcedqcwnbuowttik","gkxzjueyofw","dcv","utkdkjorodpksi","aq","japawosnoiobnrfjrsb","gkxz","lemdqgqwumegvfexckxg","u","volzppbtszdnoocqylol","fukws","tpktncgaus","bzlrgwxn","sqm","le","yfxsv","tveepjwpm","hzzpjooss","tveepjwpmg","volzppbtsz","bfb","omqndxpxevx","zaydypcbfttmuc","tiufhfh","uvdfwryvvugth","gkxzj","mnexyhlutclbwvvcgdzecs","ahx","uvdfwryv","cxtoslemgf","hzzpjoossvz","xg","ksui","ema","fqdauiil","ncghfsccr","cqw","qlvtdmyeyfdxqa","iaydxhy","xgjpvpj","fq","qdowgpmd","itcudluayc","japawosnoiobnrfjr","riwisyimddyguitvtdwgyzuim","asglpcpgc","ovly","sn","lemdqgqwumeg","hzzpjoossvzgygxhgsiyduodp","aypotqdjlcevg","japawosnoi","uvdfwryvvug","zvkroeajidjenjfn","tpktncgausgl","qdowgp","uvd","itcudlua","vkrg","xgjpvpjdiccweanvg","ixlfvb","omqndx","hklbpjmcedqcwnb","j","xtwmnupebc","erkufy","aypotq","uiuwyo","kcud","nzwwarlivffio","nndarhbemem","tveepjwpmghvia","mkomsvmr","utkdkjo","wwsaejhw","dd","aypotqdjlcev","vkrgrzgre","hkl","lemdqgqwumegvfe","nnd","tveepjwpmghviayytpa","ixlf","zfzpu","mnexyhlutcl","yizs","dcvdpd","tiufhfhomricfo","ah","jeq","kcudnbvfugbntkdgj","gslb","nc","rxgjsersc","uiuwyorg","aw","utkdkjorod","ctg","mnexyhlutclb","gslbknctqwk","lemdqgqwumegvf","vgird","tveepj","mnexyhlu","qlvtdmyeyfdxqam","emapkdooz","nndar","lemdqg","yizslwlpnd","zfzpuxghsfmeq","qfhggvatefvwmq","prc","ahxm","uvdfwryvvugthazsfquk","wwsaejhwqhjpc","omqndxpxevxojfws","rxgjserscv","uiuwy","fqdauiilsnrwbwxa","lemd","zfzpuxgh","tiufhfhomr","mkomsvmrauphv","utkdkjoro","ncghfsccrgq","mom","fqdau","bq","b","cxtoslemg","ncghfsccrgqkafho","ixlfvbvdzwndvf","cxtoslemgfyjfz","riwisyimddyguitvtdwgyzuimf","vgi","vkr","zv","tpktncgausglcelnp","ovvnbszsvkjnr","cxtoslemgfyj","dkwpdpvy","p","kcudnbvfugbn","utkdkjorodpksiuz","volzppbtszdnoo","utk","bqwkbpq","ct","mo","uvdfwry","rju","gkxzjueyo","omqndxpxevxojfwsm","itcudluay","tiuf","kei","ja","bzlrg","w","uiuwyor","zaydy","qdowgpmdjcowemkac","ks","lmw","uvdf","eurmric","hzz","xgjpvpjdiccweanv","riwisyimddygu","uvdfwryvvugtha","lemdqgqwumegvfexckxgia","japawosnoiobnrfj","momqxnldh","prci","vgirdgjsz","lmwrglan","volzp","r","bfbhnenn","paqwiazmmvgudlyuwa","volz","wwsaejhwq","nndarhbememcvspy","eurmricmwb","kcudnbvfugbntkdg","yizslwlpndzji","nzwwa","sqmhtpcgdsd","erkufypwnuh","erku","jeqxn","gn","gk","cxtoslemgfyjfzs","gkxzjuey","erk","bf","vkrgrzgrekb","asglpcpg","lmwrg","i","wqmwehunbuzfhpi","yi","wq","aypotqd","dk","japawosnoiobnrfjrsbrc","pa","wvepj","zvkroeajidjen","xaoq","hklbpjmcedqcwnbuowttikv","qfhggvatefvwmqbxtdnbo","tiufhfhomric","cxtoslemgfyjf","tv","cxtoslem","sqmhtpcgdsdfq","fu","cqwecmwk","itcudluaycsspjr","n","tpktncgausglceln","erkuf","yizslwlpndzjih","qlvtdmyeyfdxq","tve","hzzpjoossvzg","tpktn","kcu","paqwiaz","tpktncgau","wqmw","uiuwyorgm","ut","wwsaejhwqhjpcrgqk","ri","wwsa","hklbpjmcedqcw","eu","iaydxhytjek","fcndec","sqmhtpc","lemdqgqwumegvfexck","riwisyim","volzppbts","zvkro","ncghfsccrgqkaf","yy","fqdaui","tveepjw","vo","e","ncghf","qdowgpmdjcowe","uvdfwryvvugthazsfqu","sqmhtpcgd","zvk","a","ti","dcvdpdzoiv","qlvtdmye","xgjpvpjdi","hzzpjoossvzgygxhgsiy","itcudlu","xao","qlvtdm","aypotqdjlcevgvfr","ncghfsccrgqkafh","zfzpuxghsfmeqb","dkwpd","tiu","s","japawos","aqf","nzwwarli","asglpcpgctpdceh","tpktncgausglcel","mnexyhlutclbwvvcgdze","sqmhtpcg","japawosnoiobnrfjrs","ixlfvbvdzwnd","yfxsvgar","tiufhfhomri","ixlfvbvdzw","xgjpvpjdicc","rxgjser","nndarhbem","gkxzju","dcvdpdzoi","yizsl","cq","rxg","riwisyimddyg","tvee","asg","prcih","ovvnbszs","itcudluaycssp","fqdauiils","sqmhtpcgdsdfqgxsjg","yizslwlpn","omqnd","sqmhtpcgdsdfqgxsjgqizl","hdtv","gnyrz","jak","riwisyi","asgl","mkomsvm","qfhggvatef","paqw","as","qdowgpmdjcowemkacwz","paqwiazmmvgudl","sq","riwisyimddyguitvtdwgy","tpktncgausg","za","lemdqgqwumegvfexckxgiad","lemdqgqwume","japawo","l","utkdkjorodpksiuzc","fqdauiilsnrwbwx","fqdauiilsnr","sqmhtpcgdsdfqgxsjgq","omqndxpxevxo","uvdfwryvvugthazsfqukai","sqmhtpcgdsdfqgxsj","vkrgrz","qlvtdmyeyf","emapkdoo","qlvtdmyeyfd","fqdauiilsnrwbwxal","aypotqdjl","dkw","ww","tiufhf","japaw","tveepjwpmghv","eurmricmwbu","volzppbtszdno","volzppbtszd","d","volzppbtszdn","hklbpjmcedqcwnbuow","hzzpjoossvzgygxhgs","keizh","zaydypcbftt","aypot","xgjpvpjdiccwe","eurmr","iayd","lm","yfxsvgari","tiufhfhom","cqwe","xtwmnupebcuz","asglpcpgctp","qdowgpmdjcowemkacwzoxuk","mnexyhlutclbwv","uvdfwryvvugthazsfquka","fuk","paqwiazmmvgud","rxgjse","gkxzjue","riwisyimddyguit","uwp","xtwmnupebcu","zvkroeajid","qlvtdmyeyfdxqame","kc","nndarhbememcvs","zvkroeaji","bzlrgwxnsq","rxgj","emapk","utkdkjorodpks","ctguxtac","mkomsvmraup","itcudluaycsspjrohfrni","rx","utkdkjor","mkomsv","qfhggvatefvwmqbxtdn","mnexy","zfz","bfbhne","sqmhtpcgdsdfqg","yizslwlpndzjihonft","nzw","riwisyimddyguitv","lemdqgqwumegvfexckxgiadw","emapkdo","wwsaej","ql","bqwk","fukwsv","vgirdgj","ovvnbszsvk","f","iaydxhytjekusm","asglpcpgctpdcehwte","utkdkj","xgjp","paqwiazmmvg","lemdqgqwumegvfexckxgi","itcudluaycss","bzlrgwx","itcudluaycsspjro","hdtvrlw","yfx","cxtosle","mnexyhlutclbw","wwsaejhwqh","itcudl","xgjpv","bzlrgwxnsqd","sqmhtpcgdsdfqgx","itcudluaycsspjroh","hzzpjoossvzgygxh","dkwpdpv","qfhggvate","pr","nzwwarlivff","hzzpj","zvkroeajidje","uiuwyorgmesml","erkufypwn","nndarhbe","zvkroeajidjenj","omqndxpxev","wqmwehunbuzf","iaydxhytjekus","utkdkjorodp","riwisyimddyguitvt","kcudnbvfugbntkdgjs","dkwp","kcudnb","gslbkn","tveepjwpmghviayytp","yfxsvga","riw","omqndxpx","hd","asglpcpgctpdcehwt","uvdfwryvv","hzzpjoossvzgygxhgsiyduo","hklbpjmcedqc","zaydypcbft","nzwwarliv","emapkd","nndarhbememcv","ctgux","japawosnoiobn","xtwm","vgirdgjszo","iaydxhyt","zaydyp","qfhggvatefvwmqbxt","qdowgpmdjcowemkacwzox","mk","mnexyhlutclbwvvcgdzecsuv","iaydxhytjeku","bfbhnen","lmwrgl","iaydx","aypotqdjlcevgv","gs","bzlrgwxnsqdzux","qlv","tpk","riwi","qfhggvatefvwmqbx","iaydxhytjekusmaeg","zfzpuxghsfm","riwisyimdd","itcu","bfbhn","qlvtdmyeyfdx","wwsaejh","zfzpuxg","kcudnbvfugbntkd","mnexyhlutclbwvvcgdzecsu","gkxzjueyof","lemdqgqwumegvfex","zvkroeajidjenjfng","qfhggva","ix"}
	res = longestWord(words)
	assert.Equal(t, "lemdqgqwumegvfexckxgiadwlc", res, "Error in case 3")

	//case 4
	words = []string{"wo","wor","worl","world"}
	res = longestWord(words)
	assert.Equal(t, "", res, "Error in case 4")
}
