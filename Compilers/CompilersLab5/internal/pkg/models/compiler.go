package models

type MessageList []Message

type Compiler struct {
	MessageList []Message
}

func (c *Compiler) AddError(coord Position, text string) {
	c.MessageList = append(c.MessageList, Message{
		Value:    text,
		Position: coord,
		IsError:  true,
	})
}

func (c *Compiler) addWarning(coord Position, text string) {
	c.MessageList = append(c.MessageList, Message{
		Value:    text,
		Position: coord,
		IsError:  false,
	})
}
