package common

import (
	"fmt"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli"
)

var NAME = "gitlab-ci-multi-runner"
var VERSION = "dev"
var REVISION = "HEAD"
var BRANCH = "HEAD"
var BUILT = "now"

var AppVersion AppVersionInfo

type AppVersionInfo struct {
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	Revision     string    `json:"revision"`
	Branch       string    `json:"branch"`
	GOVersion    string    `json:"go_version"`
	BuiltAt      time.Time `json:"built_at"`
	OS           string    `json:"os"`
	Architecture string    `json:"architecture"`
}

func (v *AppVersionInfo) Printer(c *cli.Context) {
	fmt.Print(v.Extended())
}

func (v *AppVersionInfo) Line() string {
	return fmt.Sprintf("%s %s (%s)", v.Name, v.Version, v.Revision)
}

func (v *AppVersionInfo) ShortLine() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.Revision)
}

func (v *AppVersionInfo) UserAgent() string {
	return fmt.Sprintf("%s %s (%s; %s; %s/%s)", v.Name, v.Version, v.Branch, v.GOVersion, v.OS, v.Architecture)
}

func (v *AppVersionInfo) Extended() string {
	version := fmt.Sprintf("Version:      %s\n", v.Version)
	version += fmt.Sprintf("Git revision: %s\n", v.Revision)
	version += fmt.Sprintf("Git branch:   %s\n", v.Branch)
	version += fmt.Sprintf("GO version:   %s\n", v.GOVersion)
	version += fmt.Sprintf("Built:        %s\n", v.BuiltAt.Format(time.RFC1123Z))
	version += fmt.Sprintf("OS/Arch:      %s/%s\n", v.OS, v.Architecture)

	return version
}

// NewMetricsCollector returns a prometheus.Collector which represents current build information.
func (v *AppVersionInfo) NewMetricsCollector() *prometheus.GaugeVec {
	labels := map[string]string{
		"name":         v.Name,
		"version":      v.Version,
		"revision":     v.Revision,
		"branch":       v.Branch,
		"go_version":   v.GOVersion,
		"built_at":     v.BuiltAt.String(),
		"os":           v.OS,
		"architecture": v.Architecture,
	}
	labelNames := make([]string, 0, len(labels))
	for n := range labels {
		labelNames = append(labelNames, n)
	}

	buildInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ci_runner_version_info",
			Help: "A metric with a constant '1' value labeled by different build stats fields.",
		},
		labelNames,
	)
	buildInfo.With(labels).Set(1)
	return buildInfo
}

func init() {
	builtAt := time.Now()
	if BUILT != "now" {
		builtAt, _ = time.Parse(time.RFC3339, BUILT)
	}

	AppVersion = AppVersionInfo{
		Name:         NAME,
		Version:      VERSION,
		Revision:     REVISION,
		Branch:       BRANCH,
		GOVersion:    runtime.Version(),
		BuiltAt:      builtAt,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
	}
}
