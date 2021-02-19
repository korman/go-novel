package configs

var VolumeRegexp = [...]string{
	// `(?m)^\s*[第]*[0-9一二三四五六七八九十零〇百千两]+[\s]*[卷]+.*$`,
	`^[第]*[0-9一二三四五六七八九十零百千两]+[卷]+`,
	`^卷+[0-9一二三四五六七八九十零百千两]+`,
}

var ChapterRegexp = [...]string{}

var BookInfoRegexp = [...]string{}
