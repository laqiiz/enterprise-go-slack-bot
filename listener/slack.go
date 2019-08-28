package listener

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

type SlackListener struct {
	Client *slack.Client
}

// ListenAndResponse listens slack events and response
func (s *SlackListener) ListenAndResponse() {
	rtm := s.Client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.FileSharedEvent:
			if err := s.handleFileShareEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message event: %s\n", err)
			}
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message event: %s\n", err)
			}
		}
	}
}

// handleMessageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {

	if len(ev.Files) > 0 {
		if _, _, err := s.Client.PostMessage(ev.Channel, slack.MsgOptionText("🚨ファイルUploadを検出しました👮", false), ); err != nil {
			return fmt.Errorf("failed to post message: %s\n", err)
		}
		return nil
	}

	fmt.Printf("[INFO] Receive post message: %s\n", ev.Msg.Text)
	return nil
}

func (s *SlackListener) handleFileShareEvent(ev *slack.FileSharedEvent) error {
	//if _, _, err := s.Client.PostMessage(ev.File.Channels[0], slack.MsgOptionText("🚨ファイルUploadを検出しました👮", false), ); err != nil {
	//	return fmt.Errorf("failed to post message: %s\n", err)
	//}
	return nil
}

//// handleFileCreatedEvent handles FileCreated events.
//func (s *SlackListener) handleFileCreatedEvent(ev *slack.FileCreatedEvent) error {
//	fmt.Printf("[INFO] Receive file upload event: %s\n", ev.File.Name)
//	if _, _, err := s.Client.PostMessage(ev.File.Channels[0], slack.MsgOptionText("🚨ファイルUploadを検出しました👮", false), ); err != nil {
//		return fmt.Errorf("failed to post message: %s\n", err)
//	}
//	return nil
//}
