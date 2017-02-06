package telegrams

// description of structures is just copy-paste from https://core.telegram.org/bots/api

// This object represents a Telegram user or bot.
type user struct {
	id        int    `json:"id"`         // Unique identifier for this user or bot
	firstName string `json:"first_name"` // User‘s or bot’s first name
	lastName  string `json:"last_name"`  // Optional. User‘s or bot’s last name
	username  string `json:"username"`   // Optional. User‘s or bot’s username
}

// This object represents a chat.
type chat struct {
	id                          int    `json:"id"`                             // Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	chtype                      string `json:"type"`                           // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	title                       string `json:"title"`                          // Optional. Title, for supergroups, channels and group chats
	username                    string `json:"username"`                       // Optional. Username, for private chats, supergroups and channels if available
	firstName                   string `json:"first_name"`                     // Optional. First name of the other party in a private chat
	lastName                    string `json:"last_name"`                      // Optional. Last name of the other party in a private chat
	allMembersAreAdministrators bool   `json:"all_members_are_administrators"` // Optional. True if a group has ‘All Members Are Admins’ enabled.
}

// This object represents a message.
type message struct {
	messageId             int             `json:"message_id"`              // Unique message identifier inside this chat
	from                  user            `json:"from"`                    // Optional. Sender, can be empty for messages sent to channels
	date                  int             `json:"date"`                    // Date the message was sent in Unix time
	chat                  chat            `json:"chat"`                    // Conversation the message belongs to
	forwardFrom           user            `json:"forward_from"`            // Optional. For forwarded messages, sender of the original message
	forwardFromChat       chat            `json:"forward_from_chat"`       // Optional. For messages forwarded from a channel, information about the original channel
	forwardFromMessage_id int             `json:"forward_from_message_id"` // Optional. For forwarded channel posts, identifier of the original message in the channel
	forwardDate           int             `json:"forward_date"`            // Optional. For forwarded messages, date the original message was sent in Unix time
	replyToMessage        message         `json:"reply_to_message"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	editDate              int             `json:"edit_date"`               // Optional. Date the message was last edited in Unix time
	text                  string          `json:"text"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	entities              []messageEntity `json:"entities"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	audio                 audio           `json:"audio"`                   // Optional. Message is an audio file, information about the file
	document              document        `json:"document"`                // Optional. Message is a general file, information about the file
	game                  game            `json:"game"`                    // Optional. Message is a game, information about the game. More about games »
	photo                 []photoSize     `json:"photo"`                   // Optional. Message is a photo, available sizes of the photo
	sticker               sticker         `json:"sticker"`                 // Optional. Message is a sticker, information about the sticker
	video                 video           `json:"video"`                   // Optional. Message is a video, information about the video
	voice                 voice           `json:"voice"`                   // Optional. Message is a voice message, information about the file
	caption               string          `json:"caption"`                 // Optional. Caption for the document, photo or video, 0-200 characters
	contact               contact         `json:"contact"`                 // Optional. Message is a shared contact, information about the contact
	location              location        `json:"location"`                // Optional. Message is a shared location, information about the location
	venue                 venue           `json:"venue"`                   // Optional. Message is a venue, information about the venue
	newChatMember         user            `json:"new_chat_member"`         // Optional. A new member was added to the group, information about them (this member may be the bot itself)
	leftChatMember        user            `json:"left_chat_member"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	newChatTitle          string          `json:"new_chat_title"`          // Optional. A chat title was changed to this value
	newChatPhoto          []photoSize     `json:"new_chat_photo"`          // Optional. A chat photo was change to this value
	deleteChatPhoto       bool            `json:"delete_chat_photo"`       // Optional. Service message: the chat photo was deleted
	groupChatCreated      bool            `json:"group_chat_created"`      // Optional. Service message: the group has been created
	supergroupChatCreated bool            `json:"supergroup_chat_created"` // Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	channelChatCreated    bool            `json:"channel_chat_created"`    // Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	migrateToChatId       int             `json:"migrate_to_chat_id"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	migrateFromChatId     int             `json:"migrate_from_chat_id"`    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	pinnedMessage         message         `json:"pinned_message"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
}

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type messageEntity struct {
	messageEntitytype string `json:"type"`   // Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	offset            int    `json:"offset"` // Offset in UTF-16 code units to the start of the entity
	length            int    `json:"length"` // Length of the entity in UTF-16 code units
	url               string `json:"url"`    // Optional. For “text_link” only, url that will be opened after user taps on the text
	user              user   `json:"user"`   // Optional. For “text_mention” only, the mentioned user
}

// This object represents one size of a photo or a file / sticker thumbnail.
type photoSize struct {
	fileId   string `json:"file_id"`   // Unique identifier for this file
	width    int    `json:"width"`     // Photo width
	height   int    `json:"height"`    // Photo height
	fileSize int    `json:"file_size"` // Optional. File size
}

// This object represents an audio file to be treated as music by the Telegram clients.
type audio struct {
	fileId    string `json:"file_id"`   // Unique identifier for this file
	duration  int    `json:"duration"`  // Duration of the audio in seconds as defined by sender
	performer string `json:"performer"` // Optional. Performer of the audio as defined by sender or by audio tags
	title     string `json:"title"`     // Optional. Title of the audio as defined by sender or by audio tags
	mimeType  string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	fileSize  int    `json:"file_size"` // Optional. File size
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
type document struct {
	fileId   string    `json:"file_id"`   // Unique file identifier
	thumb    photoSize `json:"thumb"`     // Optional. Document thumbnail as defined by sender
	fileName string    `json:"file_name"` // Optional. Original filename as defined by sender
	mimeType string    `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	fileSize int       `json:"file_size"` // Optional. File size
}

// This object represents a sticker.
type sticker struct {
	fileId   string    `json:"file_id"`   // Unique identifier for this file
	width    int       `json:"width"`     // Sticker width
	height   int       `json:"height"`    // Sticker height
	thumb    photoSize `json:"thumb"`     // Optional. Sticker thumbnail in .webp or .jpg format
	emoji    string    `json:"emoji"`     // Optional. Emoji associated with the sticker
	fileSize int       `json:"file_size"` // Optional. File size
}

// This object represents a video file.
type video struct {
	fileId   string    `json:"file_id"`   // Unique identifier for this file
	width    int       `json:"width"`     // Video width as defined by sender
	height   int       `json:"height"`    // Video height as defined by sender
	duration int       `json:"duration"`  // Duration of the video in seconds as defined by sender
	thumb    photoSize `json:"thumb"`     // Optional. Video thumbnail
	mimeType string    `json:"mime_type"` // Optional. Mime type of a file as defined by sender
	fileSize int       `json:"file_size"` // Optional. File size
}

// This object represents a voice note.
type voice struct {
	fileId    string `json:"file_id"`   // Unique identifier for this file
	duration  int    `json:"duration"`  // Duration of the audio in seconds as defined by sender
	mime_type string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	file_size int    `json:"file_size"` // Optional. File size
}

// This object represents a phone contact.
type contact struct {
	phoneNumber string `json:"phone_number"` // Contact's phone number
	firstName   string `json:"first_name"`   // Contact's first name
	lastName    string `json:"last_name"`    // Optional. Contact's last name
	userId      int    `json:"user_id"`      // Optional. Contact's user identifier in Telegram
}

// This object represents a point on the map.
type location struct {
	longitude float32 `json:"longitude"` // Longitude as defined by sender
	latitude  float32 `json:"latitude"`  // Latitude as defined by sender
}

// This object represents a venue.
type venue struct {
	location     location `json:"location"`      // Venue location
	title        string   `json:"title"`         // Name of the venue
	address      string   `json:"address"`       // Address of the venue
	foursquareId string   `json:"foursquare_id"` // Optional. Foursquare identifier of the venue
}

// UserProfilePhotos
type userProfilePhotos struct {
	totalCount int         `json:"total_count"` // Total number of profile pictures the target user has
	photos     []photoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type file struct {
	fileId   string `json:"file_id"`   // Unique identifier for this file
	fileSize int    `json:"file_size"` // Optional. File size, if known
	filePath string `json:"file_path"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	keyboard        [][]keyboardButton `json:"keyboard"`          // Array of button rows, each represented by an Array of KeyboardButton objects
	resizeKeyboard  bool               `json:"resize_keyboard"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	oneTimeKeyboard bool               `json:"one_time_keyboard"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.

	selective bool `json:"selective"` // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
}

// This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type keyboardButton struct {
	text            string `json:"text"`             // Text of the button. If none of the optional fields are used, it will be sent to the bot as a message when the button is pressed
	requestContact  bool   `json:"request_contact"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	requestLocation bool   `json:"request_location"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type replyKeyboardRemove struct {
	removeKeyboard bool `json:"remove_keyboard"` // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)

	selective bool `json:"selective"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type inlineKeyboardMarkup struct {
	inlineKeyboard [][]inlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type inlineKeyboardButton struct {
	text              string `json:"text"`                // Label text on the button
	url               string `json:"url"`                 // Optional. HTTP url to be opened when button is pressed
	callback_data     string `json:"callback_data"`       // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	switchInlineQuery string `json:"switch_inline_query"` // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.

	switchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"` // Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.
	//This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.

	callbackGame callbackGame `json:"callback_game"` // Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
}

// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type callbackQuery struct {
	id              string  `json:"id"`                // Unique identifier for this query
	from            user    `json:"from"`              // Sender
	message         message `json:"message"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	inlineMessageId string  `json:"inline_message_id"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	chatInstance    string  `json:"chat_instance"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	data            string  `json:"data"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	gameShortName   string  `json:"game_short_name"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type forceReply struct {
	forceReply bool `json:"force_reply"` // Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	selective  bool `json:"selective"`   // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// This object contains information about one member of the chat.
type chatMember struct {
	user   user   `json:"user"`   // Information about the user
	status string `json:"status"` // The member's status in the chat. Can be “creator”, “administrator”, “member”, “left” or “kicked”
}

// Contains information about why a request was unsuccessfull.
type ResponseParameters struct {
	migrateToChatId int `json:"migrate_to_chat_id"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	retryAfter      int `json:"retry_after"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}
