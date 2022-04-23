package service

import (
	"fmt"
	"testing"
	"time"

	"ArSearch/pkg/service/service_schema"
)

func TestPutToEs(t *testing.T) {
	t1 := time.Now()
	a1:=&service_schema.ArArticle{
		ID: "koehrOAeK5Lpc860JoLo-Gc6ODiCv0JCSwqtR8UWYGY",
		Title: "for all its parallels, ukraine war feels distant in taiwan",
		ArticleContext: "the russian attack on ukraine has put a spotlight on another place that could face an invasion by its larger neighbor, as some analysts draw quick comparisons to china's threats to assert its control over self-ruled taiwan.\nwhile similarities exist and taiwan is a democracy that has defied a more powerful authoritarian government, the differences are much greater. and for many on the island, the war in ukraine, and war in general, feels far away.\n\"i think our situation is not very similar to ukraine's, whether it's political or in terms of connections,\" said ethan lin, a 40-year-old who works in the service industry. \"china has many exchanges with taiwan in several areas, so i don't think it's that dangerous.\"\ntaiwan, an island of 23 million people about 160 kilometers (100 miles) off china's eastern coast, is self-ruled, but claimed by china. the decades-old issue has grown more intense since independence-leaning president tsai ing-wen took the helm in taiwan in 2016, and china stepped up military pressure on the island, sending ships into nearby waters and fighter jets in its direction.\non tuesday, china's people's liberation army's eastern command announced it had recently conducted landing drills in an undisclosed location in the east china sea.\nthe critical question for taiwan is whether the united states, which is not sending troops to defend ukraine, would intervene if china invaded. the u.s. has no official ties with taiwan but has historical relations and sells taiwan billions of dollars worth of weapons. it is also bound by its own law to ensure taiwan can defend itself.\ntaiwan is also a dominant player in the production of semiconductors that are used in everything from smartphones to cars.\n\"taiwan's economy and technology is important to the u.s., and perhaps the u.s. will value taiwan more, but we have to see how the conflict plays out,\" said kao-cheng wang, a professor at the graduate institute of international affairs and strategic studies at tamkang university in taiwan.\ntaiwan announced friday it would join global sanctions against russia, although it did not provide details on what those measures would be.\n\"we can't sit on the sidelines while a big power bullies a small neighbor,\" wang ting-yu, a lawmaker from tsai's ruling democratic progressive party, wrote on twitter.\nchina and taiwan split during a civil war in 1949. the u.s. cut formal diplomatic relations with taipei in 1979 in order to recognize beijing.\nwhile china's president xi jinping has stressed that \"peaceful reunification\" is in the best interests of both sides, china's cabinet-level taiwan affairs office routinely issues angry threats to crush moves by taiwanese politicians to continue pushing for an independent country, although it's been left with only 14 diplomatic allies.\nchina has not ruled out force if necessary to achieve reunification, but for now, military action remains unlikely and outside events will have relatively little effect on beijing's calculations, said li minjiang, a chinese international relations expert at singapore's nanyang technological university.\n\"external influences on china's decision over taiwan are minimal,\" li said, adding china would continue to use information campaigns and peaceful inducements to influence public opinion on taiwan.\nrussian president vladimir putin is different from xi and has previously used force against other countries, such as neighboring georgia, said wang, the tamkang university professor. \"xi jinping is rather strong, but he increased military activity, rather than starting a war.\"\nin taipei, the bustling capital, salesperson peter chiang doubted china would attack. \"i think even internally, they aren't that stable right now,\" he said.\nchinese communist party-owned global times newspaper has compared taiwan to ukraine's separatist eastern donetsk region, where the conflict first broke out in 2014. former u.s. president donald trump predicted taiwan would be attacked in an interview this week in which he praised putin's action.\nbut chinese officials are more careful. \"taiwan is indeed not ukraine,\" china's foreign ministry spokesperson hua chunying said this week, insisting that taiwan is an integral part of china.\non matsu, a group of outlying taiwanese islands whose closest point is just 10 kilometers (6 miles) from china, taiwanese politician wen lii is not dismissive of a possible invasion. but he rejected simplistic comparisons to ukraine.\n\"lazy comparisons often fuel an inevitable, triumphalist narrative for china, or weaken confidence in democracies, while ignoring different contexts for each region,\" wen, director of the matsu chapter of the democratic progressive party, wrote in an email.\n\"people in matsu always remain alert, but discussions about chinese threats are usually based on concrete observation instead of a foreign crisis,\" he said, saying there is no military buildup for now.",
	}
	PutToEs(a1)
	t2 := time.Now()

	fmt.Println(t2.Sub(t1))
}

func TestSearchInEs(t *testing.T) {
	es, err := SearchInEs("parallels")

	if err !=nil{
		fmt.Println(err)
	}

	fmt.Println(es)
}

func TestSearchMirrorData(t *testing.T) {
	data, err := SearchMirrorData("launch")

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(data)
}
