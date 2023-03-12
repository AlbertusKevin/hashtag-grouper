package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/AlbertusKevin/hashtag-grouper/constant"
	"github.com/gorilla/mux"
)

func createMapHashtagGroup() map[string]map[string]string{
	hashtagGroupPerNiche := map[string]map[string]string{}

	hashtagGroupPerNiche[constant.MINDSET] = map[string]string{}
	hashtagGroupPerNiche[constant.MINDSET][constant.HIGH_REACHMENT] = `#positivemindset #determination #thinkbig #mindsetmatters #selfconfidence #selfdevelopment #growthmindset #mindsetshift #mindsetofgreatness #changeyourmindset #successfulmindset #healthymindset`
	hashtagGroupPerNiche[constant.MINDSET][constant.MED_REACHMENT] = `#mindsetiskey #mindsetreset #mindsetmastery #mindsetchange #mindsetmotivation #changeyourmindsetchangeyourlife #strongmindset #motivationalmindset #growthmindset🌱 #mindsetgrowth #mindsetgoals #mindsetforgreatness #postivemindset #shiftyourmindset #gratefulmindset #manifestationmindset #successisamindset #positivemindset💯`
	hashtagGroupPerNiche[constant.MINDSET][constant.LOW_REACHMENT] = `#mindsetpositif #mindsetshifting #changemindset #itsamindset #mindsetovereverything #powerfulmindset #mindsetmatter #fixmindset #goodmindset #growthmindsets #successfulmindsets #happymindset #mindsetfirst #clearmindset #sucessfulmindset #mindsetpower #mindsetmaster #mindsetpositive #healthymindsets #growmindset #mindset101 #changingmindsets #evolveyourmindset #motivasimindset #powerofmindset`

	hashtagGroupPerNiche[constant.MOTIVATION] = map[string]string{}
	hashtagGroupPerNiche[constant.MOTIVATION][constant.HIGH_REACHMENT] = `#motivationforlife #motivationdaily #motivations #motivateyourself #motivationalwords #motivationoftheday #motivationalpost #motivation101 #inspiringquotes #selfmotivation #instamotivation`
	hashtagGroupPerNiche[constant.MOTIVATION][constant.MED_REACHMENT] = `#lifemotivation #motivationlife #motivates #motivationeveryday #fearlessmotivation #getmotivated #strongmindset #motivationalmindset #successmind #inspirationalthoughts #motivationmindset #motivateothers #achieveyourgoals #lessonsoflife #actiontaker #motivation100 #beleiveinyourself #motivation💪 #lifechangingquotes #bestmotivation`
	hashtagGroupPerNiche[constant.MOTIVATION][constant.LOW_REACHMENT] = `#motivationalquotesaboutlife #motivationquotesforlife #lifestylemotivation #motivationoflife #lifetimemotivation #lifechangingmotivation #motivationinlife #motivation4life #motivationalforlife #motivationallineaboutlife #lifesgoalsmotivation #lifegoalsmotivation #motivationlifestyle #motivatedminds #motivatedlife #motivationmilestone #motivationwords #successthoughts #mindthinking #thinkforchange #thinkforchanges #motivation_thoughts #power_of_positivity #motivation_of_the_day #motivation_line #instagrammotivation #inspirationmindset`

	hashtagGroupPerNiche[constant.HABITS] = map[string]string{}
	hashtagGroupPerNiche[constant.HABITS][constant.HIGH_REACHMENT] = `#betterlife #discipline #starttoday #consistency #routine #decisions #bebetter #willpower #choices #habits #healthyhabits #youcandoit #habit`
	hashtagGroupPerNiche[constant.HABITS][constant.MED_REACHMENT] = `#habitandhome #habitacion #successhabits #goodhabits #dailyhabits #habitchange #highperformancehabits #changeyourhabits #starting #becomebetter #forwardthinking #selfdiscipline #workonyourself #longtermgoals #betteryourself #helpyourself #changelife`
	hashtagGroupPerNiche[constant.HABITS][constant.LOW_REACHMENT] = `#changinghabits #successfulhabits #habitstacking #oldhabits #createhealthyhabits #betterhabits #habitsforhappiness #creatinghealthyhabits #habitsforlife #goodhabitsforlife #breakinghabits #buildinghabits #healtyhabits #createnewhabits #creatinghabits #breakbadhabits #habitsofmind #creategoodhabits #smallhabits #newhabitsnewresults #habitsandlifestyles #delinquenthabits #successfulpeoplehabits #buildhabits #habitsforwellness #mindfulhabits #goodhabitsstartyoung #selfcarehabits #greathabits #habitsofhappiness`

	hashtagGroupPerNiche[constant.PRODUCTIVITY] = map[string]string{}
	hashtagGroupPerNiche[constant.PRODUCTIVITY][constant.HIGH_REACHMENT] = `#productivity #timemanagement #productive #careergoals #planning #worklifebalance #organized #workmode #studytime`
	hashtagGroupPerNiche[constant.PRODUCTIVITY][constant.MED_REACHMENT] = `#timemanagementtips #mindsetreset #productivityhacks #productivityhabits #goalplanning #gettingthingsdone #getthingsdone #decisionmaking #achievemore #planforsuccess #juststart #workgoals #makeaplan #scheduling #procrastinate #selfmanagement #organizedlife #selfknowledge #productivitytips #productiveday #beproductive #growthhacking #onmydesk #worksmart`
	hashtagGroupPerNiche[constant.PRODUCTIVITY][constant.LOW_REACHMENT] = `#timemanagementskills #timemanagementhacks #timemanagement101 #timemanagementgoals #productivityplanner #productivitytip #productivityhack #lifeplanning #dailyhabitsforsuccess #taskmanagement #productivemindset #takeactiontoday #actionplan #worklifeintegration #developyourself #planyourday #stopprocrastinating #managementskills #timeblocking #lifeplan #getstuffdone #timemanagment #managementwaktu`

	hashtagGroupPerNiche[constant.SELF_REMINDER] = map[string]string{}
	hashtagGroupPerNiche[constant.SELF_REMINDER][constant.HIGH_REACHMENT] = `#selfreminder #selfreflection #reminder #notetoself #journaling #dailyreminder #selfawareness #selfhealing #selfworth #janganlupabahagia`
	hashtagGroupPerNiche[constant.SELF_REMINDER][constant.MED_REACHMENT] = `#remindertoself #selfreminders #selfreflect #jadilebihbaik #selftalk #katasemangat #introspeksidiri #simplereminders #intropeksidiri #notetomyself #selfnote #pengingatdiri #dailyreminders #lebihbaik #perbaikidiri #reminders #journalingcommunity #nasehatdiri #journalwithme #salingmengingatkan #selfcareisntselfish #lifeadvice #affirmationsoftheday #anxietyfighter`
	hashtagGroupPerNiche[constant.SELF_REMINDER][constant.LOW_REACHMENT] = `#reminderformyself #remindertomyself #reminderself #remindermyself #reminder4myself #selfcarereminder #selflovereminder #reminderstoself #reminderforself #reminderstomyself #selfireminder #reminder2self #selfreminder❤ #selfreminder😊 #thoughtfortoday #bepatientwithyourself #behopeful #renungandiri #reminderoftheday #mindfulllife #semangatberjuang #emotionalgrowth #soulpeace #ketenangan #jangantakut #selalusemangat #bodoamatlah`

	hashtagGroupPerNiche[constant.SELF_IMPROVEMENT] = map[string]string{}
	hashtagGroupPerNiche[constant.SELF_IMPROVEMENT][constant.HIGH_REACHMENT] = `#selfconfidence #selfempowerment #personalgrowth #selfgrowth #selfworth #personaldevelopment #selfimprovement #selfhelp #selfdiscovery #changeyourlife #innerstrength #selfrespect #selfesteem #bethebestyou #beyourbestself #youmatter`
	hashtagGroupPerNiche[constant.SELF_IMPROVEMENT][constant.MED_REACHMENT] = `#selfimprovementdaily #selfmastery #becomebetter #strongmind #mindpower #positiveselftalk #findyourpurpose #healthyboundaries #personaldevelopmentjourney #selfdevelopmentjourney #selfdiscipline #selfeducation #selfmotivated #improveyourself #betteryourself #personalgrowthanddevelopment #selfgrowthjourney #personaljourney #smallsteps #yourjourney #lifejourney #selfhelptips #overcomingobstacles #selfaffirmation #valueyourself #selfvalue #overcomefear #selfactualization #thejourney #pengembangandiri`
	hashtagGroupPerNiche[constant.SELF_IMPROVEMENT][constant.LOW_REACHMENT] = `#improvementofself #selfimprovements #selfimprovementquotes #selfimprovementtips #selfimprovementjourney #selfimprovementeveryday #selfimprovementsdaily #selfimprovementmovement #selfimprovementgoals #selfdevelopement #controlyourmind #lifemastery #selfimprove #selﬁmprovement #berubahlebihbaik #confidencefromwithin #lifeimprovement`

	return hashtagGroupPerNiche
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello");
	}).Methods("GET")

	hashtagGroupPerNiche := createMapHashtagGroup()
	makeHashtagNicheGroup(5, 14, hashtagGroupPerNiche[constant.MOTIVATION])
	
	fmt.Println("Connected to port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func makeHashtagNicheGroup(nGroupResult int, nHashtag int, hashtagGroupPerNiche map[string]string){
	// hasil grup akhir
	finalGroup := map[int][]string{}

	// map untuk jumlah group per reachment
	highGroup := map[int][]string{}
	mediumGroup := map[int][]string{}
	lowGroup := map[int][]string{}

	// rasio komposisi hashtag dari setiap grup
	var highRatio float64 = 1
	var mediumRatio float64 = 3
	var lowRatio float64 = 5

	var total float64 = highRatio + mediumRatio + lowRatio

	// hitung jumlah hashtag dalam grup sesuai ratio dari masing-masing tingkat reachment
	countHighHashtagGroup := (highRatio/total) * float64(nHashtag)
	countMediumHashtagGroup := (mediumRatio/total) * float64(nHashtag)
	countLowHashtagGroup := (lowRatio/total) * float64(nHashtag)

	// jadikan string menjadi slice of string
	highReachmentHashtags := strings.Split(hashtagGroupPerNiche[constant.HIGH_REACHMENT]," ")
	mediumReachmentHashtags := strings.Split(hashtagGroupPerNiche[constant.MED_REACHMENT]," ")
	lowReachmentHashtags := strings.Split(hashtagGroupPerNiche[constant.LOW_REACHMENT]," ")

	// buat n grup per masing-masing level
	for i := 0; i < nGroupResult; i++{
		highGroup[i] = makeReachmentGroup(int(countHighHashtagGroup), highReachmentHashtags)
		mediumGroup[i] = makeReachmentGroup(int(countMediumHashtagGroup), mediumReachmentHashtags)
		lowGroup[i] = makeReachmentGroup(int(countLowHashtagGroup), lowReachmentHashtags)
	}

	// tiap level digabungkan sehingga menjadi n grup yang sudah mix
	for i := 0; i < nGroupResult; i++{
		finalGroup[i] = append(finalGroup[i], highGroup[i]...)
		finalGroup[i] = append(finalGroup[i], mediumGroup[i]...)
		finalGroup[i] = append(finalGroup[i], lowGroup[i]...)
	}

	ctx, _ := json.MarshalIndent(highGroup, "", "\t")
	fmt.Println("json high group:", string(ctx))

	ctx, _ = json.MarshalIndent(mediumGroup, "", "\t")
	fmt.Println("json medium group:", string(ctx))

	ctx, _ = json.MarshalIndent(lowGroup, "", "\t")
	fmt.Println("json low group:", string(ctx))

	ctx, _ = json.MarshalIndent(finalGroup, "", "\t")
	fmt.Println("json final group:", string(ctx))

	ctx, _ = json.MarshalIndent(lowGroup, "", "\t")
	fmt.Println("json buang:", string(ctx))
}

func makeReachmentGroup(countHashtagGroup int, reachmentHashtagsDB []string) (reachmentGroup []string){
	var idxIncluded []int

	for len(reachmentGroup) != countHashtagGroup {
		randIdx := rand.Intn(len(reachmentHashtagsDB))
		if !indexContains(idxIncluded, randIdx){
			idxIncluded = append(idxIncluded, randIdx)
			reachmentGroup = append(reachmentGroup, reachmentHashtagsDB[randIdx])
		}
	}

	return reachmentGroup
}