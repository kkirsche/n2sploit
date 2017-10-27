package libnmap

// Address is the IP address that was scanned
type Address struct {
	AttrAddr     string `xml:" addr,attr"  json:",omitempty"`
	AttrAddrtype string `xml:" addrtype,attr"  json:",omitempty"`
}

// Cpe is common platform enumeration
type Cpe struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Debugging represents the debug level
type Debugging struct {
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
}

// Distance represents the number of hops to reach the target
type Distance struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type Elem struct {
	AttrKey string `xml:" key,attr"  json:",omitempty"`
	Text    string `xml:",chardata" json:",omitempty"`
}

type Extraports struct {
	AttrCount    string        `xml:" count,attr"  json:",omitempty"`
	AttrState    string        `xml:" state,attr"  json:",omitempty"`
	Extrareasons *Extrareasons `xml:" extrareasons,omitempty" json:"extrareasons,omitempty"`
}

type Extrareasons struct {
	AttrCount  string `xml:" count,attr"  json:",omitempty"`
	AttrReason string `xml:" reason,attr"  json:",omitempty"`
}

type Finished struct {
	AttrElapsed string `xml:" elapsed,attr"  json:",omitempty"`
	AttrExit    string `xml:" exit,attr"  json:",omitempty"`
	AttrSummary string `xml:" summary,attr"  json:",omitempty"`
	AttrTime    string `xml:" time,attr"  json:",omitempty"`
	AttrTimestr string `xml:" timestr,attr"  json:",omitempty"`
}

type Hop struct {
	AttrIpaddr string `xml:" ipaddr,attr"  json:",omitempty"`
	AttrRtt    string `xml:" rtt,attr"  json:",omitempty"`
	AttrTtl    string `xml:" ttl,attr"  json:",omitempty"`
}

type Host struct {
	AttrEndtime   string         `xml:" endtime,attr"  json:",omitempty"`
	AttrStarttime string         `xml:" starttime,attr"  json:",omitempty"`
	Address       *Address       `xml:" address,omitempty" json:"address,omitempty"`
	Distance      *Distance      `xml:" distance,omitempty" json:"distance,omitempty"`
	Hostnames     *Hostnames     `xml:" hostnames,omitempty" json:"hostnames,omitempty"`
	Ipidsequence  *Ipidsequence  `xml:" ipidsequence,omitempty" json:"ipidsequence,omitempty"`
	Os            *Os            `xml:" os,omitempty" json:"os,omitempty"`
	Ports         *Ports         `xml:" ports,omitempty" json:"ports,omitempty"`
	Status        *Status        `xml:" status,omitempty" json:"status,omitempty"`
	Tcpsequence   *Tcpsequence   `xml:" tcpsequence,omitempty" json:"tcpsequence,omitempty"`
	Tcptssequence *Tcptssequence `xml:" tcptssequence,omitempty" json:"tcptssequence,omitempty"`
	Times         *Times         `xml:" times,omitempty" json:"times,omitempty"`
	Trace         *Trace         `xml:" trace,omitempty" json:"trace,omitempty"`
	Uptime        *Uptime        `xml:" uptime,omitempty" json:"uptime,omitempty"`
}

type Hostnames struct {
}

type Hosts struct {
	AttrDown  string `xml:" down,attr"  json:",omitempty"`
	AttrTotal string `xml:" total,attr"  json:",omitempty"`
	AttrUp    string `xml:" up,attr"  json:",omitempty"`
}

type Ipidsequence struct {
	AttrClass  string `xml:" class,attr"  json:",omitempty"`
	AttrValues string `xml:" values,attr"  json:",omitempty"`
}

type Nmaprun struct {
	AttrArgs             string        `xml:" args,attr"  json:",omitempty"`
	AttrScanner          string        `xml:" scanner,attr"  json:",omitempty"`
	AttrStart            string        `xml:" start,attr"  json:",omitempty"`
	AttrStartstr         string        `xml:" startstr,attr"  json:",omitempty"`
	AttrVersion          string        `xml:" version,attr"  json:",omitempty"`
	AttrXmloutputversion string        `xml:" xmloutputversion,attr"  json:",omitempty"`
	Debugging            *Debugging    `xml:" debugging,omitempty" json:"debugging,omitempty"`
	Host                 *Host         `xml:" host,omitempty" json:"host,omitempty"`
	Runstats             *Runstats     `xml:" runstats,omitempty" json:"runstats,omitempty"`
	Scaninfo             *Scaninfo     `xml:" scaninfo,omitempty" json:"scaninfo,omitempty"`
	Taskprogress         *Taskprogress `xml:" taskprogress,omitempty" json:"taskprogress,omitempty"`
	Verbose              *Verbose      `xml:" verbose,omitempty" json:"verbose,omitempty"`
}

type Os struct {
	Osmatch  *Osmatch  `xml:" osmatch,omitempty" json:"osmatch,omitempty"`
	Portused *Portused `xml:" portused,omitempty" json:"portused,omitempty"`
}

type Osclass struct {
	AttrAccuracy string `xml:" accuracy,attr"  json:",omitempty"`
	AttrOsfamily string `xml:" osfamily,attr"  json:",omitempty"`
	AttrType     string `xml:" type,attr"  json:",omitempty"`
	AttrVendor   string `xml:" vendor,attr"  json:",omitempty"`
}

type Osmatch struct {
	AttrAccuracy string   `xml:" accuracy,attr"  json:",omitempty"`
	AttrLine     string   `xml:" line,attr"  json:",omitempty"`
	AttrName     string   `xml:" name,attr"  json:",omitempty"`
	Osclass      *Osclass `xml:" osclass,omitempty" json:"osclass,omitempty"`
}

type Port struct {
	AttrPortid   string    `xml:" portid,attr"  json:",omitempty"`
	AttrProtocol string    `xml:" protocol,attr"  json:",omitempty"`
	Script       []*Script `xml:" script,omitempty" json:"script,omitempty"`
	Service      *Service  `xml:" service,omitempty" json:"service,omitempty"`
	State        *State    `xml:" state,omitempty" json:"state,omitempty"`
}

type Ports struct {
	Extraports *Extraports `xml:" extraports,omitempty" json:"extraports,omitempty"`
	Port       []*Port     `xml:" port,omitempty" json:"port,omitempty"`
}

type Portused struct {
	AttrPortid string `xml:" portid,attr"  json:",omitempty"`
	AttrProto  string `xml:" proto,attr"  json:",omitempty"`
	AttrState  string `xml:" state,attr"  json:",omitempty"`
}

type Runstats struct {
	Finished *Finished `xml:" finished,omitempty" json:"finished,omitempty"`
	Hosts    *Hosts    `xml:" hosts,omitempty" json:"hosts,omitempty"`
}

type Scaninfo struct {
	AttrNumservices string `xml:" numservices,attr"  json:",omitempty"`
	AttrProtocol    string `xml:" protocol,attr"  json:",omitempty"`
	AttrServices    string `xml:" services,attr"  json:",omitempty"`
	AttrType        string `xml:" type,attr"  json:",omitempty"`
}

type Script struct {
	AttrId     string   `xml:" id,attr"  json:",omitempty"`
	AttrOutput string   `xml:" output,attr"  json:",omitempty"`
	Elem       []*Elem  `xml:" elem,omitempty" json:"elem,omitempty"`
	Table      []*Table `xml:" table,omitempty" json:"table,omitempty"`
}

type Service struct {
	AttrConf    string `xml:" conf,attr"  json:",omitempty"`
	AttrMethod  string `xml:" method,attr"  json:",omitempty"`
	AttrName    string `xml:" name,attr"  json:",omitempty"`
	AttrProduct string `xml:" product,attr"  json:",omitempty"`
	AttrTunnel  string `xml:" tunnel,attr"  json:",omitempty"`
	AttrVersion string `xml:" version,attr"  json:",omitempty"`
	Cpe         *Cpe   `xml:" cpe,omitempty" json:"cpe,omitempty"`
}

type State struct {
	AttrReason     string `xml:" reason,attr"  json:",omitempty"`
	AttrReason_ttl string `xml:" reason_ttl,attr"  json:",omitempty"`
	AttrState      string `xml:" state,attr"  json:",omitempty"`
}

type Status struct {
	AttrReason     string `xml:" reason,attr"  json:",omitempty"`
	AttrReason_ttl string `xml:" reason_ttl,attr"  json:",omitempty"`
	AttrState      string `xml:" state,attr"  json:",omitempty"`
}

type Table struct {
	AttrKey string   `xml:" key,attr"  json:",omitempty"`
	Elem    []*Elem  `xml:" elem,omitempty" json:"elem,omitempty"`
	Table   []*Table `xml:" table,omitempty" json:"table,omitempty"`
}

type Taskprogress struct {
	AttrEtc       string `xml:" etc,attr"  json:",omitempty"`
	AttrPercent   string `xml:" percent,attr"  json:",omitempty"`
	AttrRemaining string `xml:" remaining,attr"  json:",omitempty"`
	AttrTask      string `xml:" task,attr"  json:",omitempty"`
	AttrTime      string `xml:" time,attr"  json:",omitempty"`
}

type Tcpsequence struct {
	AttrDifficulty string `xml:" difficulty,attr"  json:",omitempty"`
	AttrIndex      string `xml:" index,attr"  json:",omitempty"`
	AttrValues     string `xml:" values,attr"  json:",omitempty"`
}

type Tcptssequence struct {
	AttrClass  string `xml:" class,attr"  json:",omitempty"`
	AttrValues string `xml:" values,attr"  json:",omitempty"`
}

type Times struct {
	AttrRttvar string `xml:" rttvar,attr"  json:",omitempty"`
	AttrSrtt   string `xml:" srtt,attr"  json:",omitempty"`
	AttrTo     string `xml:" to,attr"  json:",omitempty"`
}

type Trace struct {
	AttrPort  string `xml:" port,attr"  json:",omitempty"`
	AttrProto string `xml:" proto,attr"  json:",omitempty"`
	Hop       []*Hop `xml:" hop,omitempty" json:"hop,omitempty"`
}

type Uptime struct {
	AttrLastboot string `xml:" lastboot,attr"  json:",omitempty"`
	AttrSeconds  string `xml:" seconds,attr"  json:",omitempty"`
}

type Verbose struct {
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
}
