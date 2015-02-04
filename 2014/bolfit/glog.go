glog.Info("Always printed")

glog.Infof("Printed on %v", time.Now())

glog.Debug("Printed only if at least on DEBUG level")

if glog.DebugEnabled() {
    glog.Debugf("Printed only if at least on %s level", "DEBUG")
}

glog.Trace("Printed only is at least on TRACE level")

if glog.TraceEnabled() {
    glog.Tracef("Printed only if at least on %s level", "TRACE")
} 