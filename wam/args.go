package wam

import (
	"os"
	"github.com/taikedz/goargs/goargs"
)

type WamArgs struct {
	action string
	tokens []string
}

func ParseArgs() WamArgs {
	if len(goargs) <= 1 {
		fail(10, "Action required")
	}

	action := os.Args[1]
	args := os.Args[2:]
	switch action {
	case "prefix":
		return parserargsPrefix(args)
	case "label":
		return parseargsLabel(args)
	case "channel":
		return parseargsChannel(args)
	case "chan-del":
		return parseargsChandel(args)
	case "list":
		return parseargsList(args)
	case "ls":
		return parseargsLs(args)
	case "cleanup":
		return parseargsCleanup(args)
	case "retain":
		return parseargsRetain(args)
	case "unretain":
		return parseargsUnretain(args)
	case "get":
		return parseargsGet(args)
	default:
		fail(10, "Unknown action '%s'\n", action)
	}
}