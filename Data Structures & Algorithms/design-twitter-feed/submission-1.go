type Tweet struct {
	TweetId int
	UserId  int
}

type Twitter struct {
	Users  map[int]map[int]struct{} // this maps a user to the list of people they follow (including themselves)
	Tweets []*Tweet                 // a slice that holds pointers to tweets
}

func Constructor() Twitter {
	return Twitter{map[int]map[int]struct{}{}, []*Tweet{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	_, succ := this.Users[userId]
	if !succ {
		this.Users[userId] = map[int]struct{}{userId: struct{}{}}
	}

	this.Tweets = append(this.Tweets, &Tweet{tweetId, userId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followers := this.Users[userId]
	output := make([]int, 0, 10)

	for i := 1; i <= len(this.Tweets); i++ {
		tweet := this.Tweets[len(this.Tweets)-i]
		tweetUser, tweetId := tweet.UserId, tweet.TweetId
		if _, succ := followers[tweetUser]; succ {
			output = append(output, tweetId)
		}

		if len(output) >= 10 {
			break
		}
	}
	return output
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, succ := this.Users[followerId]
	if !succ {
		this.Users[followerId] = map[int]struct{}{followerId: struct{}{}}
	}

	this.Users[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	_, succ := this.Users[followerId]
	if !succ {
		return
	}

	delete(this.Users[followerId], followeeId)
}