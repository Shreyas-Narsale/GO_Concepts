package main

import (
	"fmt"
)

/* Observer Pattern:( Behavioral Design Pattern)
	a one-to-many relationship
	where one object (the Subject) automatically notifies multiple other objects (Observers) about any state changes.
	ex. One Youtube video post atuomatically alert to all user , video is live.
	Subject → list of observers(users)
	Observers → listen for changes; interface which contain methods with needs to implenetd for all user which get call on notify
	When subject updates → like video posted, all observers are notified automatically

ex.
Think of:
YouTube channel (Subject)
Subscribers (Observers)
When a new video is uploaded → all subscribers get notified.
*/

// 1. Observer Interface
type Observer interface {
	Update(videoTitle string)
}

// 2. Subject: (Publisher)
type Subject interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	NotifyAll()
}

// 3.Concrete Subject
type YouTubeChannel struct {
	ChannelName string
	subscribers []Observer
}

func (yt *YouTubeChannel) Subscribe(o Observer) {
	yt.subscribers = append(yt.subscribers, o)
}

func (yt *YouTubeChannel) Unsubscribe(o Observer) {

	for index, obs := range yt.subscribers {
		if o == obs {
			yt.subscribers = append(yt.subscribers[:index], yt.subscribers[index+1:]...)
			break
		}
	}
}

func (yt *YouTubeChannel) UploadVideo(title string) {
	fmt.Println(yt.ChannelName, " is uploading a video ", title)
	yt.NotifyAll(title)
}

func (yt *YouTubeChannel) NotifyAll(title string) {
	for _, subscriber := range yt.subscribers {
		subscriber.Update(title)
	}
}

// 3. Concrete Observers
type User struct {
	Name  string
	Email string
}

func (u *User) Update(title string) {
	fmt.Printf("Hey %s, new video alert: %s\n", u.Name, title)
}

func main() {

	CarryMinatiChannel := YouTubeChannel{ChannelName: "CarryMinati"}
	FlyingBeastChannel := YouTubeChannel{ChannelName: "CarryMinati"}

	User1 := User{Name: "Shreyas", Email: "xyz@gmaail.com"}
	User2 := User{Name: "Nishu", Email: "mno@gmaail.com"}
	User3 := User{Name: "RamBhau", Email: "pqr@gmaail.com"}

	CarryMinatiChannel.Subscribe(&User1)
	CarryMinatiChannel.Subscribe(&User2)
	CarryMinatiChannel.Subscribe(&User3)

	FlyingBeastChannel.Subscribe(&User1)
	FlyingBeastChannel.Subscribe(&User3)

	CarryMinatiChannel.Unsubscribe(&User3)

	CarryMinatiChannel.UploadVideo("Yalgaar")
	FlyingBeastChannel.UploadVideo("Rashi")

}
