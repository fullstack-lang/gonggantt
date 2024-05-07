package gonggantt

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gonggantt/dist/ng-github.com-fullstack-lang-gonggantt
var NgDistNg embed.FS
