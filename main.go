package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var hashtagGroupPerNiche = map[string]map[string]string{}

	hashtagGroupPerNiche["mindset"] = map[string]string{}
	hashtagGroupPerNiche["mindset"]["high"] = `#positivemindset #determination #thinkbig #mindsetmatters #selfconfidence #selfdevelopment #growthmindset #mindsetshift #mindsetofgreatness #changeyourmindset #successfulmindset #healthymindset`
	hashtagGroupPerNiche["mindset"]["medium"] = `#mindsetiskey #mindsetreset #mindsetmastery #mindsetchange #mindsetmotivation #changeyourmindsetchangeyourlife #strongmindset #motivationalmindset #growthmindset🌱 #mindsetgrowth #mindsetgoals #mindsetforgreatness #postivemindset #shiftyourmindset #gratefulmindset #manifestationmindset #successisamindset #positivemindset💯`
	hashtagGroupPerNiche["mindset"]["low"] = `#mindsetpositif #mindsetshifting #changemindset #itsamindset #mindsetovereverything #powerfulmindset #mindsetmatter #fixmindset #goodmindset #growthmindsets #successfulmindsets #happymindset #mindsetfirst #clearmindset #sucessfulmindset #mindsetpower #mindsetmaster #mindsetpositive #healthymindsets #growmindset #mindset101 #changingmindsets #evolveyourmindset #motivasimindset #powerofmindset`

	hashtagGroupPerNiche["motivation"] = map[string]string{}
	hashtagGroupPerNiche["motivation"]["high"] = `#motivationforlife #motivationdaily #motivations #motivateyourself #motivationalwords #motivationoftheday #motivationalpost #motivation101 #inspiringquotes #selfmotivation #instamotivation`
	hashtagGroupPerNiche["motivation"]["medium"] = `#lifemotivation #motivationlife #motivates #motivationeveryday #fearlessmotivation #getmotivated #strongmindset #motivationalmindset #successmind #inspirationalthoughts #motivationmindset #motivateothers #achieveyourgoals #lessonsoflife #actiontaker #motivation100 #beleiveinyourself #motivation💪 #lifechangingquotes #bestmotivation`
	hashtagGroupPerNiche["motivation"]["low"] = `#motivationalquotesaboutlife #motivationquotesforlife #lifestylemotivation #motivationoflife #lifetimemotivation #lifechangingmotivation #motivationinlife #motivation4life #motivationalforlife #motivationallineaboutlife #lifesgoalsmotivation #lifegoalsmotivation #motivationlifestyle #motivatedminds #motivatedlife #motivationmilestone #motivationwords #successthoughts #mindthinking #thinkforchange #thinkforchanges #motivation_thoughts #power_of_positivity #motivation_of_the_day #motivation_line #instagrammotivation #inspirationmindset`

	hashtagGroupPerNiche["habits"] = map[string]string{}
	hashtagGroupPerNiche["habits"]["high"] = `#betterlife #discipline #starttoday #consistency #routine #decisions #bebetter #willpower #choices #habits #healthyhabits #youcandoit #habit`
	hashtagGroupPerNiche["habits"]["medium"] = `#habitandhome #habitacion #successhabits #goodhabits #dailyhabits #habitchange #highperformancehabits #changeyourhabits #starting #becomebetter #forwardthinking #selfdiscipline #workonyourself #longtermgoals #betteryourself #helpyourself #changelife`
	hashtagGroupPerNiche["habits"]["low"] = `#changinghabits #successfulhabits #habitstacking #oldhabits #createhealthyhabits #betterhabits #habitsforhappiness #creatinghealthyhabits #habitsforlife #goodhabitsforlife #breakinghabits #buildinghabits #healtyhabits #createnewhabits #creatinghabits #breakbadhabits #habitsofmind #creategoodhabits #smallhabits #newhabitsnewresults #habitsandlifestyles #delinquenthabits #successfulpeoplehabits #buildhabits #habitsforwellness #mindfulhabits #goodhabitsstartyoung #selfcarehabits #greathabits #habitsofhappiness`

	hashtagGroupPerNiche["productivity"] = map[string]string{}
	hashtagGroupPerNiche["productivity"]["high"] = `#productivity #timemanagement #productive #careergoals #planning #worklifebalance #organized #workmode #studytime`
	hashtagGroupPerNiche["productivity"]["medium"] = `#timemanagementtips #mindsetreset #productivityhacks #productivityhabits #goalplanning #gettingthingsdone #getthingsdone #decisionmaking #achievemore #planforsuccess #juststart #workgoals #makeaplan #scheduling #procrastinate #selfmanagement #organizedlife #selfknowledge #productivitytips #productiveday #beproductive #growthhacking #onmydesk #worksmart`
	hashtagGroupPerNiche["productivity"]["low"] = `#timemanagementskills #timemanagementhacks #timemanagement101 #timemanagementgoals #productivityplanner #productivitytip #productivityhack #lifeplanning #dailyhabitsforsuccess #taskmanagement #productivemindset #takeactiontoday #actionplan #worklifeintegration #developyourself #planyourday #stopprocrastinating #managementskills #timeblocking #lifeplan #getstuffdone #timemanagment #managementwaktu`

	hashtagGroupPerNiche["self reminder"] = map[string]string{}
	hashtagGroupPerNiche["self reminder"]["high"] = `#selfreminder #selfreflection #reminder #notetoself #journaling #dailyreminder #selfawareness #selfhealing #selfworth #janganlupabahagia`
	hashtagGroupPerNiche["self reminder"]["medium"] = `#remindertoself #selfreminders #selfreflect #jadilebihbaik #selftalk #katasemangat #introspeksidiri #simplereminders #intropeksidiri #notetomyself #selfnote #pengingatdiri #dailyreminders #lebihbaik #perbaikidiri #reminders #journalingcommunity #nasehatdiri #journalwithme #salingmengingatkan #selfcareisntselfish #lifeadvice #affirmationsoftheday #anxietyfighter`
	hashtagGroupPerNiche["self reminder"]["low"] = `#reminderformyself #remindertomyself #reminderself #remindermyself #reminder4myself #selfcarereminder #selflovereminder #reminderstoself #reminderforself #reminderstomyself #selfireminder #reminder2self #selfreminder❤ #selfreminder😊 #thoughtfortoday #bepatientwithyourself #behopeful #renungandiri #reminderoftheday #mindfulllife #semangatberjuang #emotionalgrowth #soulpeace #ketenangan #jangantakut #selalusemangat #bodoamatlah`

	hashtagGroupPerNiche["self improvement"] = map[string]string{}
	hashtagGroupPerNiche["self improvement"]["high"] = `#selfconfidence #selfempowerment #personalgrowth #selfgrowth #selfworth #personaldevelopment #selfimprovement #selfhelp #selfdiscovery #changeyourlife #innerstrength #selfrespect #selfesteem #bethebestyou #beyourbestself #youmatter`
	hashtagGroupPerNiche["self improvement"]["medium"] = `#selfimprovementdaily #selfmastery #becomebetter #strongmind #mindpower #positiveselftalk #findyourpurpose #healthyboundaries #personaldevelopmentjourney #selfdevelopmentjourney #selfdiscipline #selfeducation #selfmotivated #improveyourself #betteryourself #personalgrowthanddevelopment #selfgrowthjourney #personaljourney #smallsteps #yourjourney #lifejourney #selfhelptips #overcomingobstacles #selfaffirmation #valueyourself #selfvalue #overcomefear #selfactualization #thejourney #pengembangandiri`
	hashtagGroupPerNiche["self improvement"]["low"] = `#improvementofself #selfimprovements #selfimprovementquotes #selfimprovementtips #selfimprovementjourney #selfimprovementeveryday #selfimprovementsdaily #selfimprovementmovement #selfimprovementgoals #selfdevelopement #controlyourmind #lifemastery #selfimprove #selﬁmprovement #berubahlebihbaik #confidencefromwithin #lifeimprovement`

	// ctx, _ := json.MarshalIndent(hashtagGroupPerNiche["self improvement"], "", "\t")
	// fmt.Println("json:", string(ctx))

	makeGroup(5, 20, hashtagGroupPerNiche["self reminder"])
}

func makeGroup(nGroupResult int, nHashtag int, hashtagGroupPerNiche map[string]string){
	var highRatio float64 = 1
	var mediumRatio float64 = 3
	var lowRatio float64 = 5
	var total float64 = highRatio + mediumRatio + lowRatio

	// hitung jumlah hashtag dalam grup sesuai ratio dari masing-masing tingkat reachment
	countHighHashtagGroup := (highRatio/total)  * float64(nHashtag)
	countMediumHashtagGroup := (mediumRatio/total) * float64(nHashtag)
	countLowHashtagGroup := (lowRatio/total) * float64(nHashtag)

	// jadikan string menjadi slice of string
	highReachmentHashtags := strings.Split(hashtagGroupPerNiche["high"]," ")
	mediumReachmentHashtags := strings.Split(hashtagGroupPerNiche["medium"]," ")
	lowReachmentHashtags := strings.Split(hashtagGroupPerNiche["low"]," ")

	for i := 0; i < int(countHighHashtagGroup); i++ {
		randIdx := rand.Intn(len(highReachmentHashtags))
		fmt.Println("len high:",len(highReachmentHashtags));
		fmt.Println("high:",randIdx);
	}

	for i := 0; i < int(countMediumHashtagGroup); i++ {
		randIdx := rand.Intn(len(mediumReachmentHashtags))
		fmt.Println("len medium:",len(mediumReachmentHashtags));
		fmt.Println("medium:",randIdx);
	}

	for i := 0; i < int(countLowHashtagGroup); i++ {
		randIdx := rand.Intn(len(lowReachmentHashtags))
		fmt.Println("len low:",len(lowReachmentHashtags));
		fmt.Println("low:",randIdx);
	}
	

	// fmt.Println(math.Round(countHighHashtagGroup));
	// fmt.Println(math.Round(countMediumHashtagGroup));
	// fmt.Println(math.Round(countLowHashtagGroup));
}