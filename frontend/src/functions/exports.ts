import { createSignal } from "solid-js";
import { createStore } from "solid-js/store";

export interface Host {
	ID:    number;
	Name:  string;
	DNS:   string;
	Iface: string;
	IP:    string;
	Mac:   string;
	Hw:    string;
	Date:  string;
	Known: number;
	Now:   number;
};

export interface Conf {
	Host:	   string;
	Port:	   string;
	Theme:	   string;
	Color:     string;
	DirPath:   string;
	Timeout:   number;
	NodePath:  string;
	LogLevel:  string;
	Ifaces:	   string;
	ArpArgs:   string;
	ArpStrs:   string[];
	TrimHist:  number;
	ShoutURL:  string;
	UseDB:     string;
	PGConnect: string;
	// InfluxDB
	InfluxEnable:  boolean;
	InfluxAddr:    string;
	InfluxToken:   string;
	InfluxOrg:     string;
	InfluxBucket:  string;
	InfluxSkipTLS: boolean;
	// Prometheus
	PrometheusEnable: boolean;
};

export const emptyHost:Host = {
	ID:    0,
	Name:  "",
	DNS:   "",
	Iface: "",
	IP:    "",
	Mac:   "",
	Hw:    "",
	Date:  "",
	Known: 0,
	Now:   0,
};

export const emptyConf:Conf = {
	Host:	 "",
	Port:	 "",
	Theme:	 "",
	Color:   "",
	DirPath: "",
	Timeout: 120,
	NodePath: "",
	LogLevel: "",
	Ifaces:	 "",
	ArpArgs: "",
	ArpStrs: [],
	TrimHist: 48,
	ShoutURL: "",
	UseDB: "",
	PGConnect: "",
	InfluxEnable:  false,
	InfluxAddr:    "",
	InfluxToken:   "",
	InfluxOrg:     "",
	InfluxBucket:  "",
	InfluxSkipTLS: false,
	PrometheusEnable: false,
};

export const [allHosts, setAllHosts] = createStore<Host[]>([]);
export const [bkpHosts, setBkpHosts] = createSignal<Host[]>([]);

export const [ifaces, setIfaces] = createSignal<string[]>([]);

export const [appConfig, setAppConfig] = createSignal<Conf>(emptyConf);

export const [editNames, setEditNames] = createSignal(false);

export const [show, setShow] = createSignal<number>(200);

export const [histUpdOnFilter, setHistUpdOnFilter] = createSignal(false);

export const [selectedIDs, setSelectedIDs] = createSignal<number[]>([]);