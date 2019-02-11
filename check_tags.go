package publiccode

import "github.com/thoas/go-funk"

// checkTag tells whether the tag is a valid tag or not and returns it.
func (p *Parser) checkTag(key, tag string) (string, error) {
	if funk.Contains(tagList, tag) {
		return tag, nil
	}

	return tag, newErrorInvalidValue(key, "unknown tag: %s", tag)
}

// A tagList represents a list of the possible tags.
var tagList = []string{
	// International tags.
	"3dgraphics",              // application for viewing, creating, or processing 3-d graphics
	"accessibility",           // accessibility
	"accounting",              // accounting software
	"amusement",               // a simple amusement
	"archiving",               // a tool to archive/backup data
	"art",                     // software to teach arts
	"artificial-intelligence", // artificial intelligence software
	"backend",                 // software not meant for end users
	"calculator",              // a calculator
	"calendar",                // calendar application
	"chat",                    // a chat client
	"classroom-management",    // classroom management software
	"clock",                   // a clock application/applet
	"content-management",      // a content management system (CMS)
	"compression",             // a tool to manage compressed data/archives
	"construction",            //
	"contact-management",      // e.g. an address book
	"database",                // application to manage a database
	"debugger",                // a tool to debug applications
	"dictionary",              // a dictionary
	"documentation",           // help or documentation
	"electronics",             // electronics software, e.g. a circuit designer
	"email",                   // email application
	"emulator",                // emulator of another platform, such as a dos emulator
	"engineering",             // engineering software, e.g. cad programs
	"file-manager",            // a file manager
	"file-transfer",           // tools like ftp or p2p programs
	"finance",                 // application to manage your finance
	"flowchart",               // a flowchart application
	"gui-designer",            // a gui designer application
	"identity",                // identity management
	"instant-messaging",       // an instant messaging client
	"library",                 // a library software
	"medical",                 // medical software
	"monitor",                 // monitor application/applet that monitors some resource or activity
	"museum",                  // museum software
	"music",                   // musical software
	"news",                    // software to manage and publish news
	"ocr",                     // optical character recognition application
	"parallel-computing",      // parallel computing software
	"photography",             // camera tools, etc.
	"presentation",            // presentation software
	"printing",                // a tool to manage printers
	"procurement",             // software for managing procurement
	"project-management",      // project management application
	"publishing",              // desktop publishing applications and color management tools
	"raster-graphics",         // application for viewing, creating, or processing raster (bitmap) graphics
	"remote-access",           // a tool to remotely manage your pc
	"revision-control",        // applications like git or subversion
	"robotics",                // robotics software
	"scanning",                // tool to scan a file/text
	"security",                // a security tool
	"sports",                  // sports software
	"spreadsheet",             // a spreadsheet
	"telephony",               // telephony via pc
	"terminal-emulator",       // a terminal emulator application
	"texteditor",              // a text editor
	"texttools",               // a text tool utility
	"translation",             // a translation tool
	"vector-graphics",         // application for viewing, creating, or processing vector graphics
	"video-conference",        // video conference software
	"viewer",                  // tool to view e.g. a graphic or pdf file
	"web-browser",             // a web browser
	"whistleblowing",          // software for whistleblowing / anticorruption
	"word-processor",          // a word processor
	"wordprocessor",           // a word processor
}
