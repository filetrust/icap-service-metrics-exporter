package icap

import (
	"regexp"
	"strconv"
)

type RunningServerStatistics struct {
	children int
	free     int
	used     int
	started  int
	closed   int
	crashed  int
	closing  int
}

type GWRebuildStatistics struct {
	reqmods          int
	respmods         int
	options          int
	allow204         int
	requestsScanned  int
	rebuildFailures  int
	rebuildErrors    int
	scanRebuilt      int
	unprocessed      int
	unprocessable    int
	bytesIn          int
	bytesOut         int
	httpBytesIn      int
	httpBytesOut     int
	bodyBytesIn      int
	bodyBytesOut     int
	bodyBytesScanned int
}

var (
	icapRunningStatsChildrenRegexp = regexp.MustCompile(`Children number: ([0-9]*)`)
	icapRunningStatsFreeRegexp     = regexp.MustCompile(`Free Servers: ([0-9]*)`)
	icapRunningStatsUsedRegexp     = regexp.MustCompile(`Used Servers: ([0-9]*)`)
	icapRunningStatsStartedRegexp  = regexp.MustCompile(`Started Processes: ([0-9]*)`)
	icapRunningStatsClosedRegexp   = regexp.MustCompile(`Closed Processes: ([0-9]*)`)
	icapRunningStatsCrashedRegexp  = regexp.MustCompile(`Crashed Processes: ([0-9]*)`)
	icapRunningStatsClosingRegexp  = regexp.MustCompile(`Closing Processes: ([0-9]*)`)

	gwRebuildStatsReqmodsRegexp          = regexp.MustCompile(`Service gw_rebuild REQMODS : ([0-9]*)`)
	gwRebuildStatsRespmodsRegexp         = regexp.MustCompile(`Service gw_rebuild RESPMODS : ([0-9]*)`)
	gwRebuildStatsOptionsRegexp          = regexp.MustCompile(`Service gw_rebuild OPTIONS : ([0-9]*)`)
	gwRebuildStatsAllow204Regexp         = regexp.MustCompile(`Service gw_rebuild ALLOW 204 : ([0-9]*)`)
	gwRebuildStatsRequestsScannedRegexp  = regexp.MustCompile(`Service gw_rebuild REQUESTS SCANNED : ([0-9]*)`)
	gwRebuildStatsRebuildFailuresRegexp  = regexp.MustCompile(`Service gw_rebuild REBUILD FAILURES : ([0-9]*)`)
	gwRebuildStatsRebuildErrorsRegexp    = regexp.MustCompile(`Service gw_rebuild REBUILD ERRORS : ([0-9]*)`)
	gwRebuildStatsScanRebuiltRegexp      = regexp.MustCompile(`Service gw_rebuild SCAN REBUILT : ([0-9]*)`)
	gwRebuildStatsUnprocessedRegexp      = regexp.MustCompile(`Service gw_rebuild UNPROCESSED : ([0-9]*)`)
	gwRebuildStatsUnprocessableRegexp    = regexp.MustCompile(`Service gw_rebuild UNPROCESSABLE : ([0-9]*)`)
	gwRebuildStatsBytesInRegexp          = regexp.MustCompile(`Service gw_rebuild BYTES IN : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsBytesOutRegexp         = regexp.MustCompile(`Service gw_rebuild BYTES OUT : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsHTTPBytesInRegexp      = regexp.MustCompile(`Service gw_rebuild HTTP BYTES IN : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsHTTPBytesOutRegexp     = regexp.MustCompile(`Service gw_rebuild HTTP BYTES OUT : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsBodyBytesInRegexp      = regexp.MustCompile(`Service gw_rebuild BODY BYTES IN : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsBodyBytesOutRegexp     = regexp.MustCompile(`Service gw_rebuild BODY BYTES OUT : ([0-9]*) Kbs ([0-9]*) bytes`)
	gwRebuildStatsBodyBytesScannedRegexp = regexp.MustCompile(`Service gw_rebuild BODY BYTES SCANNED : ([0-9]*) Kbs ([0-9]*) bytes`)

	icapRespCodeRegexp = regexp.MustCompile(`ICAP/1\.0 (\d+)`)
)

func parseRunningServerStatistics(icapRes []byte) (stats RunningServerStatistics) {
	childrenRegex := icapRunningStatsChildrenRegexp.FindSubmatch(icapRes)
	if len(childrenRegex) == 2 {
		stats.children, _ = strconv.Atoi(string(childrenRegex[1]))
	}
	freeRegex := icapRunningStatsFreeRegexp.FindSubmatch(icapRes)
	if len(freeRegex) == 2 {
		stats.free, _ = strconv.Atoi(string(freeRegex[1]))
	}
	usedRegex := icapRunningStatsUsedRegexp.FindSubmatch(icapRes)
	if len(usedRegex) == 2 {
		stats.used, _ = strconv.Atoi(string(usedRegex[1]))
	}
	startedRegex := icapRunningStatsStartedRegexp.FindSubmatch(icapRes)
	if len(startedRegex) == 2 {
		stats.started, _ = strconv.Atoi(string(startedRegex[1]))
	}
	closedRegex := icapRunningStatsClosedRegexp.FindSubmatch(icapRes)
	if len(closedRegex) == 2 {
		stats.closed, _ = strconv.Atoi(string(closedRegex[1]))
	}
	crashedRegex := icapRunningStatsCrashedRegexp.FindSubmatch(icapRes)
	if len(crashedRegex) == 2 {
		stats.crashed, _ = strconv.Atoi(string(crashedRegex[1]))
	}
	closingRegex := icapRunningStatsClosingRegexp.FindSubmatch(icapRes)
	if len(closingRegex) == 2 {
		stats.closing, _ = strconv.Atoi(string(closingRegex[1]))
	}
	return
}

func parseGWRebuildStatistics(icapRes []byte) (stats GWRebuildStatistics) {
	reqmodsRegex := gwRebuildStatsReqmodsRegexp.FindSubmatch(icapRes)
	if len(reqmodsRegex) == 2 {
		stats.reqmods, _ = strconv.Atoi(string(reqmodsRegex[1]))
	}
	respmodsRegex := gwRebuildStatsRespmodsRegexp.FindSubmatch(icapRes)
	if len(respmodsRegex) == 2 {
		stats.respmods, _ = strconv.Atoi(string(respmodsRegex[1]))
	}
	optionsRegex := gwRebuildStatsOptionsRegexp.FindSubmatch(icapRes)
	if len(optionsRegex) == 2 {
		stats.options, _ = strconv.Atoi(string(optionsRegex[1]))
	}
	allow204Regex := gwRebuildStatsAllow204Regexp.FindSubmatch(icapRes)
	if len(allow204Regex) == 2 {
		stats.allow204, _ = strconv.Atoi(string(allow204Regex[1]))
	}
	requestScannedRegex := gwRebuildStatsRequestsScannedRegexp.FindSubmatch(icapRes)
	if len(requestScannedRegex) == 2 {
		stats.requestsScanned, _ = strconv.Atoi(string(requestScannedRegex[1]))
	}
	rebuildFailuresRegex := gwRebuildStatsRebuildFailuresRegexp.FindSubmatch(icapRes)
	if len(rebuildFailuresRegex) == 2 {
		stats.rebuildFailures, _ = strconv.Atoi(string(rebuildFailuresRegex[1]))
	}
	rebuildErrorsRegex := gwRebuildStatsRebuildErrorsRegexp.FindSubmatch(icapRes)
	if len(rebuildErrorsRegex) == 2 {
		stats.rebuildErrors, _ = strconv.Atoi(string(rebuildErrorsRegex[1]))
	}
	scanRebuiltRegex := gwRebuildStatsScanRebuiltRegexp.FindSubmatch(icapRes)
	if len(scanRebuiltRegex) == 2 {
		stats.scanRebuilt, _ = strconv.Atoi(string(scanRebuiltRegex[1]))
	}
	unprocessRegex := gwRebuildStatsUnprocessedRegexp.FindSubmatch(icapRes)
	if len(unprocessRegex) == 2 {
		stats.unprocessed, _ = strconv.Atoi(string(unprocessRegex[1]))
	}
	unprocessableRegex := gwRebuildStatsUnprocessableRegexp.FindSubmatch(icapRes)
	if len(unprocessableRegex) == 2 {
		stats.unprocessable, _ = strconv.Atoi(string(unprocessableRegex[1]))
	}
	bytesInRegex := gwRebuildStatsBytesInRegexp.FindSubmatch(icapRes)
	if len(bytesInRegex) == 3 {
		kbs, _ := strconv.Atoi(string(bytesInRegex[1])) //kbs
		b, _ := strconv.Atoi(string(bytesInRegex[2]))   // bytes
		stats.bytesIn = (kbs * 1024) + b
	}
	bytesOutRegex := gwRebuildStatsBytesOutRegexp.FindSubmatch(icapRes)
	if len(bytesOutRegex) == 3 {
		kbs, _ := strconv.Atoi(string(bytesOutRegex[1])) //kbs
		b, _ := strconv.Atoi(string(bytesOutRegex[2]))   // bytes
		stats.bytesOut = (kbs * 1024) + b
	}
	httpBytesInRegex := gwRebuildStatsHTTPBytesInRegexp.FindSubmatch(icapRes)
	if len(httpBytesInRegex) == 3 {
		kbs, _ := strconv.Atoi(string(httpBytesInRegex[1])) //kbs
		b, _ := strconv.Atoi(string(httpBytesInRegex[2]))   // bytes
		stats.httpBytesIn = (kbs * 1024) + b
	}
	httpBytesOutRegex := gwRebuildStatsHTTPBytesOutRegexp.FindSubmatch(icapRes)
	if len(httpBytesOutRegex) == 3 {
		kbs, _ := strconv.Atoi(string(httpBytesOutRegex[1])) //kbs
		b, _ := strconv.Atoi(string(httpBytesOutRegex[2]))   // bytes
		stats.httpBytesOut = (kbs * 1024) + b
	}
	bodyBytesInRegex := gwRebuildStatsBodyBytesInRegexp.FindSubmatch(icapRes)
	if len(bodyBytesInRegex) == 3 {
		kbs, _ := strconv.Atoi(string(bodyBytesInRegex[1])) //kbs
		b, _ := strconv.Atoi(string(bodyBytesInRegex[2]))   // bytes
		stats.bodyBytesIn = (kbs * 1024) + b
	}
	bodyBytesOutRegex := gwRebuildStatsBodyBytesOutRegexp.FindSubmatch(icapRes)
	if len(bodyBytesOutRegex) == 3 {
		kbs, _ := strconv.Atoi(string(bodyBytesOutRegex[1])) //kbs
		b, _ := strconv.Atoi(string(bodyBytesOutRegex[2]))   // bytes
		stats.bodyBytesOut = (kbs * 1024) + b
	}
	bodyBytesScannedRegex := gwRebuildStatsBodyBytesScannedRegexp.FindSubmatch(icapRes)
	if len(bodyBytesScannedRegex) == 3 {
		kbs, _ := strconv.Atoi(string(bodyBytesScannedRegex[1])) //kbs
		b, _ := strconv.Atoi(string(bodyBytesScannedRegex[2]))   // bytes
		stats.bodyBytesScanned = (kbs * 1024) + b
	}
	return
}

func ParseIcapHeader(icapRes []byte) (code int) {
	code = -1

	c := icapRespCodeRegexp.FindSubmatch(icapRes)
	if len(c) == 2 {
		code, _ = strconv.Atoi(string(c[1]))
	}
	return
}
