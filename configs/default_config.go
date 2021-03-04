package configs

var VolumeRegexp = [...]string{
	// `(?m)^\s*[第]*[0-9一二三四五六七八九十零〇百千两]+[\s]*[卷]+.*$`,
	`^[第]*\s*[0-9一二三四五六七八九十零百千两]+\s*[卷]+`,
	`^卷+\s*[0-9一二三四五六七八九十零百千两]+`,
}

var ChapterRegexp = [...]string{
	`^[第]*[0-9一二三四五六七八九十零百千两]+[章话]+`,
}

var BookInfoRegexp = [...]string{}
