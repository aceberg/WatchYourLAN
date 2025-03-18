import { createSignal } from "solid-js";

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

export interface Conf {
	Color:   string;
	Timeout: number;
};


export const [allHosts, setAllHosts] = createSignal<Host[]>([]);
export const [bkpHosts, setBkpHosts] = createSignal<Host[]>([]);

export const [ifaces, setIfaces] = createSignal<string[]>([]);

export const [appConfig, setAppConfig] = createSignal<Conf>();

export const [editNames, setEditNames] = createSignal(false);

export const [currentHost, setCurrentHost] = createSignal<Host>(emptyHost);