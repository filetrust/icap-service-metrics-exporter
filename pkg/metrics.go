package icap

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type IcapOptions struct {
	Host    string
	Port    string
	Service string
}

type IcapChecker struct {
	opts                 IcapOptions
	promChildrenNumber   *prometheus.Desc
	promFreeServers      *prometheus.Desc
	promUsedServers      *prometheus.Desc
	promStartedProcesses *prometheus.Desc
	promClosedProcesses  *prometheus.Desc
	promCrashedProcesses *prometheus.Desc
	promClosingProcesses *prometheus.Desc
	promReqmods          *prometheus.Desc
	promRespmods         *prometheus.Desc
	promOptions          *prometheus.Desc
	promAllow204         *prometheus.Desc
	promRequestsScanned  *prometheus.Desc
	promRebuildFailures  *prometheus.Desc
	promRebuildErrors    *prometheus.Desc
	promScanRebuilt      *prometheus.Desc
	promUnprocessed      *prometheus.Desc
	promUnprocessable    *prometheus.Desc
	promBytesIn          *prometheus.Desc
	promBytesOut         *prometheus.Desc
	promHTTPBytesIn      *prometheus.Desc
	promHTTPBytesOuts    *prometheus.Desc
	promBodyBytesIn      *prometheus.Desc
	promBodyBytesOut     *prometheus.Desc
	promBodyBytesScanned *prometheus.Desc
}

func NewIcapChecker(host, port, service string) *IcapChecker {
	return &IcapChecker{
		opts:                 IcapOptions{Host: host, Port: port, Service: service},
		promChildrenNumber:   prometheus.NewDesc("gw_icap_server_children_number", "ICAP Server Children Number", []string{}, nil),
		promFreeServers:      prometheus.NewDesc("gw_icap_server_free_servers", "ICAP Server Free Servers Number", []string{}, nil),
		promUsedServers:      prometheus.NewDesc("gw_icap_server_used_servers", "ICAP Server Used Servers Number", []string{}, nil),
		promStartedProcesses: prometheus.NewDesc("gw_icap_server_started_processes", "ICAP Server Started Processes Number", []string{}, nil),
		promClosedProcesses:  prometheus.NewDesc("gw_icap_server_closed_processes", "ICAP Server Closed Processes Number", []string{}, nil),
		promCrashedProcesses: prometheus.NewDesc("gw_icap_server_crashed_processes", "ICAP Server Crashed Processes Number", []string{}, nil),
		promClosingProcesses: prometheus.NewDesc("gw_icap_server_closing_processes", "ICAP Server Closing Processes Number", []string{}, nil),
		promReqmods:          prometheus.NewDesc("gw_rebuild_reqmods", "GW Rebuild REQMODS Number", []string{}, nil),
		promRespmods:         prometheus.NewDesc("gw_rebuild_respmods", "GW Rebuild RESPMOD Number", []string{}, nil),
		promOptions:          prometheus.NewDesc("gw_rebuild_options", "GW Rebuild OPTIONS Number", []string{}, nil),
		promAllow204:         prometheus.NewDesc("gw_rebuild_allow204", "GW Rebuild ALLOW204 Number", []string{}, nil),
		promRequestsScanned:  prometheus.NewDesc("gw_rebuild_requests_scanned", "GW Rebuild Requests Scanned Number", []string{}, nil),
		promRebuildFailures:  prometheus.NewDesc("gw_rebuild_rebuild_failures", "GW Rebuild Rebuilt Failures Number", []string{}, nil),
		promRebuildErrors:    prometheus.NewDesc("gw_rebuild_rebuild_errors", "GW Rebuild Rebuild Errors Number", []string{}, nil),
		promScanRebuilt:      prometheus.NewDesc("gw_rebuild_scan_rebuilt", "GW Rebuild Scan Rebuilt Number", []string{}, nil),
		promUnprocessed:      prometheus.NewDesc("gw_rebuild_unprocessed", "GW Rebuild Unprocessed Number", []string{}, nil),
		promUnprocessable:    prometheus.NewDesc("gw_rebuild_unprocessable", "GW Rebuild Unprocessable Number", []string{}, nil),
		promBytesIn:          prometheus.NewDesc("gw_rebuild_bytes_in", "GW Rebuild Bytes In Number", []string{}, nil),
		promBytesOut:         prometheus.NewDesc("gw_rebuild_bytes_out", "GW Rebuild Bytes Out Number", []string{}, nil),
		promHTTPBytesIn:      prometheus.NewDesc("gw_rebuild_http_bytes_in", "GW Rebuild HTTP Bytes In Number", []string{}, nil),
		promHTTPBytesOuts:    prometheus.NewDesc("gw_rebuild_http_bytes_out", "GW Rebuild HTTP Bytes Out Number", []string{}, nil),
		promBodyBytesIn:      prometheus.NewDesc("gw_rebuild_body_bytes_in", "GW Rebuild Body Bytes In Number", []string{}, nil),
		promBodyBytesOut:     prometheus.NewDesc("gw_rebuild_body_bytes_out", "GW Rebuild Body Bytes Out Number", []string{}, nil),
		promBodyBytesScanned: prometheus.NewDesc("gw_rebuild_body_bytes_scanned", "GW Rebuild Body Bytes Scanned Number", []string{}, nil),
	}
}

func (c *IcapChecker) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.promChildrenNumber
	ch <- c.promFreeServers
	ch <- c.promUsedServers
	ch <- c.promStartedProcesses
	ch <- c.promClosedProcesses
	ch <- c.promCrashedProcesses
	ch <- c.promClosingProcesses
	ch <- c.promReqmods
	ch <- c.promRespmods
	ch <- c.promOptions
	ch <- c.promAllow204
	ch <- c.promRequestsScanned
	ch <- c.promRebuildFailures
	ch <- c.promRebuildErrors
	ch <- c.promScanRebuilt
	ch <- c.promUnprocessed
	ch <- c.promUnprocessable
	ch <- c.promBytesIn
	ch <- c.promBytesOut
	ch <- c.promHTTPBytesIn
	ch <- c.promHTTPBytesOuts
	ch <- c.promBodyBytesIn
	ch <- c.promBodyBytesOut
	ch <- c.promBodyBytesScanned
}

func (c *IcapChecker) Collect(ch chan<- prometheus.Metric) {
	fmt.Println("Collecting Statistics from ICAP Service")
	res, err := collectStatistics(c.opts.Host, c.opts.Port, c.opts.Service)
	fmt.Println("Collected statistics from ICAP Service")

	fmt.Println("Parsing Server Statistics")
	serverStats := parseRunningServerStatistics(res)
	fmt.Println("Parsed Server Statistics")

	fmt.Println("Parsing GW Rebuild Statistics")
	rebuildStats := parseGWRebuildStatistics(res)
	fmt.Println("Parsed GW Rebuild Statistics")

	if err != nil {
		return
	}

	ch <- prometheus.MustNewConstMetric(c.promChildrenNumber, prometheus.GaugeValue, float64(serverStats.children))
	ch <- prometheus.MustNewConstMetric(c.promFreeServers, prometheus.GaugeValue, float64(serverStats.free))
	ch <- prometheus.MustNewConstMetric(c.promUsedServers, prometheus.GaugeValue, float64(serverStats.used))
	ch <- prometheus.MustNewConstMetric(c.promStartedProcesses, prometheus.GaugeValue, float64(serverStats.started))
	ch <- prometheus.MustNewConstMetric(c.promClosedProcesses, prometheus.GaugeValue, float64(serverStats.closed))
	ch <- prometheus.MustNewConstMetric(c.promCrashedProcesses, prometheus.GaugeValue, float64(serverStats.crashed))
	ch <- prometheus.MustNewConstMetric(c.promClosingProcesses, prometheus.GaugeValue, float64(serverStats.closing))
	ch <- prometheus.MustNewConstMetric(c.promReqmods, prometheus.GaugeValue, float64(rebuildStats.reqmods))
	ch <- prometheus.MustNewConstMetric(c.promRespmods, prometheus.GaugeValue, float64(rebuildStats.respmods))
	ch <- prometheus.MustNewConstMetric(c.promOptions, prometheus.GaugeValue, float64(rebuildStats.options))
	ch <- prometheus.MustNewConstMetric(c.promAllow204, prometheus.GaugeValue, float64(rebuildStats.allow204))
	ch <- prometheus.MustNewConstMetric(c.promRequestsScanned, prometheus.GaugeValue, float64(rebuildStats.requestsScanned))
	ch <- prometheus.MustNewConstMetric(c.promRebuildFailures, prometheus.GaugeValue, float64(rebuildStats.rebuildFailures))
	ch <- prometheus.MustNewConstMetric(c.promRebuildErrors, prometheus.GaugeValue, float64(rebuildStats.rebuildErrors))
	ch <- prometheus.MustNewConstMetric(c.promScanRebuilt, prometheus.GaugeValue, float64(rebuildStats.scanRebuilt))
	ch <- prometheus.MustNewConstMetric(c.promUnprocessed, prometheus.GaugeValue, float64(rebuildStats.unprocessed))
	ch <- prometheus.MustNewConstMetric(c.promUnprocessable, prometheus.GaugeValue, float64(rebuildStats.unprocessable))
	ch <- prometheus.MustNewConstMetric(c.promBytesIn, prometheus.GaugeValue, float64(rebuildStats.bytesIn))
	ch <- prometheus.MustNewConstMetric(c.promBytesOut, prometheus.GaugeValue, float64(rebuildStats.bytesOut))
	ch <- prometheus.MustNewConstMetric(c.promHTTPBytesIn, prometheus.GaugeValue, float64(rebuildStats.httpBytesIn))
	ch <- prometheus.MustNewConstMetric(c.promHTTPBytesOuts, prometheus.GaugeValue, float64(rebuildStats.httpBytesOut))
	ch <- prometheus.MustNewConstMetric(c.promBodyBytesIn, prometheus.GaugeValue, float64(rebuildStats.bodyBytesIn))
	ch <- prometheus.MustNewConstMetric(c.promBodyBytesOut, prometheus.GaugeValue, float64(rebuildStats.bodyBytesOut))
	ch <- prometheus.MustNewConstMetric(c.promBodyBytesScanned, prometheus.GaugeValue, float64(rebuildStats.bodyBytesScanned))
}
