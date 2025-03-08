package wam

func Channel_set(label string, version string, channels ...string)

func Channel_del(label string, channel string)

func Channel_list(label string) []KvPair
